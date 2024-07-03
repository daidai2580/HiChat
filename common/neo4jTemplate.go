package common

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"golang.org/x/net/context"
	"strings"
)

type GenericStruct[T any] struct {
	Id         int64    `json:"Id"`
	ElementId  string   `json:"ElementId"`
	ElementId2 string   `json:"elementId"`
	Labels     []string `json:"Labels"`
	Props      T        `json:"Props"`
}

func Neo4jContext() (context.Context, neo4j.DriverWithContext) {
	// scheme://host(:port) (default port is 7687)
	//driver, err := neo4j.NewDriverWithContext(fmt.Sprintf("%s:%d", global.ServiceConfig.Neo4jDB.Host, global.ServiceConfig.Neo4jDB.Port), neo4j.BasicAuth(global.ServiceConfig.Neo4jDB.User, global.ServiceConfig.Neo4jDB.Password, ""))
	driver, err := neo4j.NewDriverWithContext(fmt.Sprintf("%s:%d", "bolt://neo4j.comm.gov.jwis.cn", 47687), neo4j.BasicAuth("neo4j", "neo4j123", ""))
	if err != nil {
		panic(err)
	}
	// 自5.0开始，您可以控制大多数driver API的执行
	// 为了简化，我们这里创建一个永不取消的上下文
	// 了解更多关于上下文的信息，请访问https://pkg.go.dev/context
	ctx := context.Background()
	// 根据你的应用程序生命周期要求来管理driver的生命周期。
	// driver的生命周期通常绑定到应用程序生命周期，通常意味着一个应用程序对应一个driver实例
	return ctx, driver
}

func SelectItem(sql string, param map[string]any) (*[]byte, error) {
	ctx, driver := Neo4jContext()
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"match (n:User { oid: $oid}) RETURN n LIMIT 1",
		map[string]any{
			"oid": "b13ab27f-9ba6-4edf-89cf-798bd1c2aaae",
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "n")
	if err != nil {
		return nil, fmt.Errorf("无法找到节点n")
	}
	jsonData, err := json.Marshal(itemNode.GetProperties())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &jsonData, nil
}

func SelectT[T any](sql string, param map[string]any, flag string) (T, error) {
	ctx, driver := Neo4jContext()
	result, err := neo4j.ExecuteQuery(ctx, driver,
		sql,
		param, neo4j.EagerResultTransformer)
	if err != nil {
		return *new(T), err
	}

	itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], flag)
	if err != nil {
		return *new(T), fmt.Errorf("无法找到节点n")
	}
	jsonData, err := json.Marshal(itemNode.GetProperties())
	// 反序列化 JSON 数据到 GenericStruct[T]
	var v T
	err = json.Unmarshal(jsonData, &v)
	if err != nil {
		return *new(T), err
	}
	return v, err

}

func SelectJson(sql string, param map[string]any, flag string) (map[string]interface{}, error) {
	ctx, driver := Neo4jContext()
	var v map[string]interface{}
	result, err := neo4j.ExecuteQuery(ctx, driver,
		sql,
		param, neo4j.EagerResultTransformer)
	if err != nil {
		return v, err
	}

	itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], flag)
	if err != nil {
		return v, fmt.Errorf("无法找到节点n")
	}
	jsonData, err := json.Marshal(itemNode.GetProperties())
	// 反序列化 JSON 数据到 GenericStruct[T]

	err = json.Unmarshal(jsonData, &v)
	if err != nil {
		return v, err
	}
	return v, err

}

func SelectCount(sql string, param map[string]any, flag string) (int64, error) {
	ctx, driver := Neo4jContext()
	session := driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(ctx)
	var count int64
	result, err := session.Run(ctx,
		sql,
		param)
	if err != nil {
		return count, err
	}

	for result.Next(ctx) {
		record := result.Record()
		value, exist := record.Get(flag)
		if !exist {
			return count, fmt.Errorf("无法找到对应返回值")
		}
		count = value.(int64)
		return count, nil
	}

	return count, err
}

func SelectList[T any](sql string, param map[string]any, flag string) ([]T, error) {
	ctx, driver := Neo4jContext()
	session := driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(ctx)
	result, err := session.Run(ctx,
		sql,
		param)
	list := make([]T, 0)
	if err != nil {
		return list, err
	}
	for result.Next(ctx) {
		record := result.Record()

		node, err := json.Marshal(record.AsMap()[flag])
		if err != nil {
			return list, fmt.Errorf("无法找到节点n")
		}
		var v GenericStruct[T]
		err = json.Unmarshal(node, &v)
		list = append(list, v.Props)
	}

	return list, err

}

func UpdateT[T any](v T, oid string, modelType string) (int64, error) {
	ctx, driver := Neo4jContext()
	session := driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(ctx)
	var mapData map[string]interface{}

	jsonData, err := json.Marshal(&v)
	err = json.Unmarshal(jsonData, &mapData)
	size := len(mapData)
	index := 0
	param := make(map[string]any, size)
	if err != nil {
		return 0, err
	}
	var str strings.Builder
	str.WriteString("match (n:")
	str.WriteString(modelType)
	str.WriteString(")  WHERE n.oid=$oid  set ")
	param["oid"] = oid

	for key, value := range mapData {
		str.WriteString("n.")
		str.WriteString(key)
		param[key] = value
		str.WriteString("=")
		str.WriteString("$")
		str.WriteString(key)
		index = index + 1
		if index < size {
			str.WriteString(",")
		}
	}
	str.WriteString(" return count(n) as n ")
	fmt.Println(str.String())
	result, err := session.Run(ctx,
		str.String(),
		param)
	if err != nil {
		return 0, err
	}
	for result.Next(ctx) {
		record := result.Record()
		value, exist := record.Get("n")
		if !exist {
			return 0, fmt.Errorf("无法找到对应返回值")
		}
		return value.(int64), nil
	}

	return 0, err
}

func InsertT[T any](v T, modelType string) (int64, error) {
	ctx, driver := Neo4jContext()
	session := driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close(ctx)

	var mapData map[string]interface{}

	jsonData, err := json.Marshal(&v)
	err = json.Unmarshal(jsonData, &mapData)
	size := len(mapData)
	index := 0
	param := make(map[string]any, size)
	if err != nil {
		return 0, err
	}
	var str strings.Builder
	str.WriteString("create (n:")
	str.WriteString(modelType)
	str.WriteString("{ ")

	for key, value := range mapData {
		str.WriteString(key)
		param[key] = value
		str.WriteString(":")
		str.WriteString("$")
		str.WriteString(key)
		str.WriteString(" ")
		index = index + 1
		if index < size {
			str.WriteString(",")
		}
	}
	str.WriteString("} ")
	str.WriteString(") ")
	str.WriteString(" return count(n) as n ")
	fmt.Println(str.String())
	result, err := session.Run(ctx,
		str.String(),
		param)
	if err != nil {
		return 0, err
	}
	for result.Next(ctx) {
		record := result.Record()
		value, exist := record.Get("n")
		if !exist {
			return 0, fmt.Errorf("无法找到对应返回值")
		}
		return value.(int64), nil
	}

	return 0, err
}

/*func GetGenericValue[T any]() (T, error) {
	params := map[string]interface{}{
		"oid": "oid",
		"id":  "id",
		"Id":  "Id",
	}
	jsonData, err := json.Marshal(params)
	var v T
	err = json.Unmarshal(jsonData, &v)
	return v, err
}*/

func GetGenericValue[T any]() (T, error) {
	params := map[string]interface{}{
		"oid": "oid",
		"id":  "id",
		"Id":  "Id",
	}
	jsonData, err := json.Marshal(params)
	if err != nil {
		return *new(T), err
	}

	var v T
	err = json.Unmarshal(jsonData, &v)
	if err != nil {
		return *new(T), err
	}

	return v, nil
}

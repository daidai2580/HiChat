package main

import (
	"encoding/json"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main2() {
	/*	f := excelize.NewFile()
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		//创建一个工作表
		sheet, err := f.NewSheet("表格2")
		if err != nil {
			fmt.Println(err)
			return
		}
		f.SetCellValue("表格2", "A2", "name")
		f.SetCellValue("Sheet1", "b2", "你好")
		f.SetActiveSheet(sheet)
		// 根据指定路径保存文件
		if err := f.SaveAs("Book1.xlsx"); err != nil {
			fmt.Println(err)
		}
	*/
	/*	f, err := excelize.OpenFile("odm.xlsx")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		rows, err := f.GetRows("Sheet1")
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, r := range rows {
			for _, l := range r {
				fmt.Print(l, "\t")
			}
			fmt.Println()
		}*/
	/*	f := excelize.NewFile()
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		_, err := f.NewSheet("表格1")
		for idx, row := range [][]interface{}{
			{nil, "Apple", "Orange", "Pear"},
			{"Small", 2, 3, 3},
			{"Normal", 5, 2, 4},
			{"big", 10, 6, 8},
		} {
			cell, err := excelize.CoordinatesToCellName(1, idx+1)
			if err != nil {
				fmt.Println(err)
				return
			}
			f.SetSheetRow("表格1", cell, &row)
		}
		err = f.AddChart("表格1", "E1", &excelize.Chart{
			Type: excelize.Col3DClustered,
			Series: []excelize.ChartSeries{
				{
					Name:       "表格1!$A$2",
					Categories: "表格1!$B$1:$D$1",
					Values:     "表格1!$B$2:$D$2",
				},
				{
					Name:       "表格1!$A$3",
					Categories: "表格1!$B$1:$D$1",
					Values:     "表格1!$B$3:$D$3",
				},
				{
					Name:       "表格1!$A$4",
					Categories: "表格1!$B$1:$D$1",
					Values:     "表格1!$B$4:$D$4",
				},
			},
			Title: []excelize.RichTextRun{
				{
					Text: "Fruit 3D Clustered Column Chart",
				},
			},
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		err = f.SaveAs("Book2.xlsx")
		if err != nil {
			fmt.Println(err)
		}*/

	/*f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	err := f.AddPicture("Sheet1", "A2", "42580805.jpg", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = f.AddPicture("Sheet1", "D2", "42580805.jpg", &excelize.GraphicOptions{ScaleX: 0.5, ScaleY: 0.5})
	if err != nil {
		fmt.Println(err)
		return
	}
	enable, disable := true, false
	err = f.AddPicture("Sheet1", "E2", "42580805.jpg", &excelize.GraphicOptions{
		PrintObject:     &enable,
		LockAspectRatio: false,
		OffsetX:         15,
		OffsetY:         10,
		Locked:          &disable,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = f.SaveAs("Bool3.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}*/

	/*dbUri := "neo4j://neo4j.comm.gov.jwis.cn:57687" // scheme://host(:port) (default port is 7687)
	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth("neo4j", "neo4j123", ""))
	if err != nil {
		panic(err)
	}
	// 自5.0开始，您可以控制大多数driver API的执行
	// 为了简化，我们这里创建一个永不取消的上下文
	// 了解更多关于上下文的信息，请访问https://pkg.go.dev/context
	ctx := context.Background()
	// 根据你的应用程序生命周期要求来管理driver的生命周期。
	// driver的生命周期通常绑定到应用程序生命周期，通常意味着一个应用程序对应一个driver实例

	defer driver.Close(ctx) // 确保处理延迟调用中的错误*/
	/*	item, err := dao.SelectT[models.Item]("match (n:User) RETURN n LIMIT 3", map[string]any{
			"oid": "b13ab27f-9ba6-4edf-89cf-798bd1c2aaae",
		}, "n")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", item)*/
	/*	value, err := dao.SelectList[models.Item]("match (n:User) RETURN n LIMIT 3", map[string]any{
			"oid": "b13ab27f-9ba6-4edf-89cf-798bd1c2aaae",
		}, "n")
		if err != nil {
			panic(err)
		}
		for _, p := range value {
			fmt.Printf("name: %s,id:%s,account:%s,email:%s", p.Name, p.Id, p.Account, p.Email)
		}
		fmt.Printf("%v\n", value)*/
	/*	t, err := dao.SelectJson("match (n:User) RETURN   n limit 1 ", map[string]any{}, "n")
		if err != nil {
			panic(err)
		}
		fmt.Printf("count: %d", t["result"])*/
	/*count, err := dao.SelectCount("match (n:User) RETURN count(n) as n", map[string]any{}, "n")
	if err != nil {
		panic(err)
	}
	print(count)*/

	//count, err := dao.InsertT(models.News{Title: "标题", Content: "内容", AuthorId: 980}, "news")
	/*count, err := dao.UpdateT(models.News{Title: "标题2", Content: "内容32", AuthorId: 9803}, "5567", "news")
	if err != nil {
		panic(err)
	}
	print(count)*/
	jsonStr := `{"name":"Alice","age":30,"gender":"Female","skills":["Go", "Python", "JavaScript"],"user":{"id":"43"}}`
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("name: %v", data["name"])
	//m := data["user"].(map[string]interface{})
	n := data["skills"].([]interface{})
	for i, skill := range n {
		fmt.Printf("i=%d,skill=%s\n", i, skill)

	}
}

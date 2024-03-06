package main

import (
	"HiChat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
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
	dsn := "root:root@tcp(120.24.168.49:3306)/hichat?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.News{})
	if err != nil {
		panic(err)
	}
}

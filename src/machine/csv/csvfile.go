package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main1() {
	path := "/Users/apple/test/GoFamily/src/machine/data/iris.csv"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	//-1代表从头扫描到结束
	reader.FieldsPerRecord = -1
	// 一口气读完
	csvdata, err := reader.ReadAll()
	fmt.Println(len(csvdata))
	for _, lien := range csvdata {
		fmt.Println(lien)
	}
}
func main() {
	path := "/Users/apple/test/GoFamily/src/machine/data/iris.csv"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	//-1代表从头扫描到结束
	reader.FieldsPerRecord = -1
	var data [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		//插入记录
		data = append(data, record)
	}
	for _, lien := range data {
		fmt.Println(lien)
	}
}

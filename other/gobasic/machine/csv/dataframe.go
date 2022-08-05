package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"os"
)

func main2() {
	path := "/Users/apple/test/GoFamily/src/machine/data/iris.csv"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	irisDF := dataframe.ReadCSV(file)
	fmt.Println(irisDF)
}
func main() {
	path := "/Users/apple/test/GoFamily/src/machine/data/iris.csv"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	irisDF := dataframe.ReadCSV(file)
	fmt.Println(irisDF)
	filter := dataframe.F{
		Colname:    "species",
		Comparator: "==",
		Comparando: "Iris-setosa",
	}
	fmt.Println(filter)
	//筛选功能
	mydf := irisDF.Filter(filter)
	fmt.Println(mydf)
	newdf := irisDF.Filter(filter).Select([]string{"Petal.Length", "Species"})
	fmt.Println(newdf)
	fmt.Println("=========")
	//Subset([]int{0,1,2,3})取前四个
	newdfx := irisDF.Filter(filter).Select([]string{"Petal.Length", "Species"}).Subset([]int{0, 1, 2, 3})
	fmt.Println(newdfx)

}

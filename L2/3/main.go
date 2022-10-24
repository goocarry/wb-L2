package main

import (
	"flag"
	"fmt"

	"github.com/goocarry/wb-L2/L2/3/sort"
)

func main() {
	filePath := flag.String("f", "./file.txt", "path to file to read")
	columnNumber := flag.Int("k", 1, "path to file to read")
	isNumeric := flag.Bool("n", false, "sort column by numeric value")
	isReverse := flag.Bool("r", false, "sort column in reversed order")
	
	flag.Parse()
	
	file := sort.File{}
	file.ReadByLine(*filePath)

	fmt.Println(*isReverse)

	if *isNumeric {
		file.SortNumericColumn(*columnNumber-1, *isReverse)
	} else {
		file.SortStringColumn(*columnNumber-1, *isReverse)
	}

	fmt.Println(file.Lines)

	file.WriteByLine("./sorted.txt")

}
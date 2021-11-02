package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var prefix string
	fmt.Println("Add prefix?")
	fmt.Scanln(&prefix)
	data, _ := ioutil.ReadFile("./input.txt")
	contents := string(data)
	var result string = ""
	result += "Name	Path	Type	Format	Length	Precision	Currency	Decimal	Group	Trim type	Repeat\n"

	var namePrefix string
	if strings.Trim(prefix, " ") == "" {
		namePrefix = ""
	} else {
		namePrefix = fmt.Sprintf("%v_", prefix)
	}
	contentArr := strings.Split(contents, "\r\n")
	for _, v := range contentArr {
		result += fmt.Sprintf("%v%v\t%v\n", namePrefix, v, fmt.Sprintf("$.*.%v", v))
	}

	result = result[:strings.LastIndex(result, "\n")]

	ioutil.WriteFile("./output.txt", []byte(result), 0644)
}

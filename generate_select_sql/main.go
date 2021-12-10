package main

import (
	"ctbc_util/common"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("Please provide target table name!")
	}
	tableName := args[1]
	mainColumn, err := ioutil.ReadFile("./main_column.txt")
	common.CheckError(err)

	mainColumnStr := string(mainColumn)
	fmt.Println(mainColumnStr)

	targetColumn, err := ioutil.ReadFile("./target_column.txt")
	common.CheckError(err)

	targetColumnStr := string(targetColumn)
	fmt.Println(targetColumnStr)

	mainArr := strings.Split(mainColumnStr, "\r\n")
	targetArr := strings.Split(targetColumnStr, "\r\n")

	var result string = "select \n"

	for i, v := range mainArr {
		name := targetArr[i]
		if len(name) == 0 {
			name = "NULL"
		}
		name = strings.ToUpper(name)
		name = strings.ReplaceAll(name, " ", "_")
		name = strings.ReplaceAll(name, "/", "_")
		name = strings.ReplaceAll(name, "(", "")
		name = strings.ReplaceAll(name, ")", "")
		name = strings.ReplaceAll(name, ":", "")
		v = strings.ToUpper(v)
		v = strings.ReplaceAll(v, " ", "_")
		v = strings.ReplaceAll(v, "/", "_")
		v = strings.ReplaceAll(v, "(", "")
		v = strings.ReplaceAll(v, ")", "")
		v = strings.ReplaceAll(v, ":", "")
		result += fmt.Sprintf("%s as %s,\n", name, v)
	}

	result += fmt.Sprintf("from %v", tableName)

	fmt.Println(result)

	common.OutputResults("./result", "output.txt", result)
}

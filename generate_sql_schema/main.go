package main

import (
	c "ctbc_util/common"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func checkForOutputDirectories() {
	check, _ := c.Exists("./result_sql")
	if !check {
		os.Mkdir("./result_sql", 0755)
	}
	check, _ = c.Exists("./result_column_name")
	if !check {
		os.Mkdir("./result_column_name", 0755)
	}
	check, _ = c.Exists("./result_column_type")
	if !check {
		os.Mkdir("./result_column_type", 0755)
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("Please provide a table name")
	}
	tableName := strings.ToUpper(strings.Join(args[1:], "_"))

	inputFile, _ := ioutil.ReadFile("./columns.txt")
	data := string(inputFile)
	dataArr := strings.Split(data, "\r\n")
	var outputFileString string = ""
	var outputColumns string = ""
	var outputColumnName string = ""
	outputFileString = fmt.Sprintf("CREATE TABLE %v (\n", tableName)
	for _, v := range dataArr {
		fmt.Println(v)
		splitData := strings.Split(v, "\t")
		fmt.Println(splitData)
		for i := 0; i < len(splitData); i++ {
			splitData[i] = strings.ToUpper(splitData[i])
			if i == 0 {
				splitData[i] = strings.ReplaceAll(splitData[i], " ", "_")
				splitData[i] = strings.ReplaceAll(splitData[i], "/", "_")
				splitData[i] = strings.ReplaceAll(splitData[i], "(", "_")
				splitData[i] = strings.ReplaceAll(splitData[i], ")", "_")
				splitData[i] = strings.ReplaceAll(splitData[i], "’", "")
			}
			if i == 1 {
				// splitData[i] = strings.ReplaceAll(splitData[i], "CHAR", "VARCHAR(200)")
				splitData[i] = strings.ReplaceAll(splitData[i], "DATE", "DATETIME")
				if splitData[i] == "VARCHAR" || splitData[i] == "CHAR" {
					splitData[i] = "VARCHAR(200)"
				}
				splitData[i] = strings.ReplaceAll(splitData[i], "NUMBER", "DECIMAL(25, 10)")
				splitData[i] = strings.ReplaceAll(splitData[i], "NUMERIC", "DECIMAL(25, 10)")
			}
		}
		for i := 0; i < len(splitData); i++ {
			if splitData[0] == strings.ToUpper("AS_OF_DATE") {
				splitData[1] = strings.ReplaceAll(splitData[1], "VARCHAR(200)", "VARCHAR(8)")
			}
		}
		for i := 0; i < len(splitData); i++ {
			if i == 0 {
				outputColumnName += strings.ToUpper(splitData[0]) + "\n"
			}
			if i == 1 {
				outputColumns += strings.ToLower(splitData[1]) + "\n"
			}
		}
		outputFileString += "\t" + strings.Join(splitData, "\t") + " NULL,\n"
	}
	outputFileString += ");"
	checkForOutputDirectories()
	ioutil.WriteFile(fmt.Sprintf("./result_sql/%v.sql", tableName), []byte(outputFileString), 0644)
	ioutil.WriteFile(fmt.Sprintf("./result_column_name/%v_column_name.txt", tableName), []byte(outputColumnName), 0644)
	ioutil.WriteFile(fmt.Sprintf("./result_column_type/%v_column_type.txt", tableName), []byte(outputColumns), 0644)
}

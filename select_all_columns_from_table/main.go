package main

import (
	"ctbc_util/common"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	configBytes, err := ioutil.ReadFile("./config.yaml")
	common.CheckError(err)
	config := common.SelectAllColumnsFromTable{}
	yaml.Unmarshal(configBytes, &config)

	db := common.GetDataBaseConnection()

	// allColumns := make([]string, 0)
	var column string
	var allColumns []string = make([]string, 0)
	result := ""
	query := "SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '$1';"
	query = strings.Replace(query, "$1", config.TableName, 1)
	rows, err := db.Query(query)
	common.CheckError(err)
	result += "SELECT \n"

	for rows.Next() {
		err = rows.Scan(&column)
		common.CheckError(err)
		allColumns = append(allColumns, column)
	}

	for i, v := range allColumns {
		if i == len(allColumns)-1 {
			result += fmt.Sprintf("\t%s.%s\n", config.Alias, v)
		} else {
			result += fmt.Sprintf("\t%s.%s,\n", config.Alias, v)
		}
	}

	result += fmt.Sprintf("FROM %s %s", config.TableName, config.Alias)
	fmt.Println(result)

	common.OutputResults("./result", "output.txt", result)
}

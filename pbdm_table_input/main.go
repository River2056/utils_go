package main

import (
	"ctbc_util/common"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type PbdmTableInput struct {
	PbdmTableName string `yaml:"pbdm-table"`
}

func main() {
	db := common.GetDataBaseConnection()
	allColumns := make([]string, 0)

	configBytes, err := ioutil.ReadFile("./config.yaml")
	common.CheckError(err)
	config := PbdmTableInput{}
	yaml.Unmarshal(configBytes, &config)

	query := "SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = @p1;"
	stmt, err := db.Prepare(query)
	common.CheckError(err)
	rows, err := stmt.Query(config.PbdmTableName)
	common.CheckError(err)

	var column string
	for rows.Next() {
		err = rows.Scan(&column)
		common.CheckError(err)
		allColumns = append(allColumns, column)
	}

	fmt.Println(allColumns)
}

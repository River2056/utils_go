package main

import (
	"ctbc_util/common"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	db := common.GetDataBaseConnection()

	bytes, err := ioutil.ReadFile("./config.yaml")
	common.CheckError(err)
	config := common.AlterTableTool{}
	yaml.Unmarshal(bytes, &config)

	for _, column := range config.AlterColumns {
		sqlStr := "ALTER TABLE $NAME $COMMAND $COLUMN_NAME;"
		sqlStr = strings.Replace(sqlStr, "$NAME", config.TableName, 1)
		sqlStr = strings.Replace(sqlStr, "$COMMAND", config.AlterCmd, 1)
		sqlStr = strings.Replace(sqlStr, "$COLUMN_NAME", column, 1)

		fmt.Printf("sqlStr: %v\n", sqlStr)

		result, err := db.Exec(sqlStr)
		common.CheckError(err)
		affected, err := result.RowsAffected()
		common.CheckError(err)
		fmt.Printf("rows affected: %v\n", affected)
	}

	fmt.Println("alter table done!")
}

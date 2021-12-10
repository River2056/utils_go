package main

import (
	"ctbc_util/common"
	"flag"
	"fmt"
	"strings"
)

func oldALterTool() {
	db := common.GetDataBaseConnection()
	// sqlStr := "ALTER TABLE $1 $COMMAND $COLUMN_NAME $DATA_TYPE"

	// args := os.Args
	tableName := flag.String("name", "", "table name, e.g. AVQ_DWH_OBJ_ASSET_SG")
	cmd := flag.String("cmd", "", "alter command to execute, e.g. ADD")
	columnName := flag.String("col", "", "table column name, e.g. ID")
	dataType := flag.String("type", "VARCHAR(200)", "data type, default: VARCHAR(200)")
	flag.Parse()

	sqlStr := "ALTER TABLE $NAME $COMMAND $COLUMN_NAME $DATA_TYPE;"
	sqlStr = strings.Replace(sqlStr, "$NAME", *tableName, 1)
	sqlStr = strings.Replace(sqlStr, "$COMMAND", *cmd, 1)
	sqlStr = strings.Replace(sqlStr, "$COLUMN_NAME", *columnName, 1)
	sqlStr = strings.Replace(sqlStr, "$DATA_TYPE", *dataType, 1)

	fmt.Printf("sqlStr: %v\n", sqlStr)

	result, err := db.Exec(sqlStr)
	common.CheckError(err)
	affected, err := result.RowsAffected()
	common.CheckError(err)
	fmt.Printf("rows affected: %v\n", affected)
}

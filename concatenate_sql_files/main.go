package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"ctbc_util/common"
)

func main() {
	var result string
	var fileDir string
	args := os.Args
	if len(args) < 2 {
		panic("Please provide file directory!")
	}
	fileDir = args[1]

	files, err := ioutil.ReadDir(fileDir)
	common.CheckError(err)

	for _, file := range files {
		fileBytes, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", fileDir, file.Name()))
		common.CheckError(err)

		fmt.Println(string(fileBytes))
		result += fmt.Sprintf("%v\n", string(fileBytes))
	}

	common.CheckDirectoryExistsAndMkdir("./result")
	ioutil.WriteFile("./result/concatenate_sql.sql", []byte(result), 0644)
}

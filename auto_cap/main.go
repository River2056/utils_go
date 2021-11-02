package main

import (
	"ctbc_util/common"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileBytes, err := ioutil.ReadFile("./columns.txt")
	common.CheckError(err)
	fileContent := string(fileBytes)

	fileArr := strings.Split(fileContent, "\r\n")
	result := ""
	for _, word := range fileArr {
		result += fmt.Sprintf("%v\n", strings.ToUpper(word))
	}

	result = result[:strings.LastIndex(result, "\n")]

	common.OutputResults("./result", "output.txt", result)
}

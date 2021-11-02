package main

import (
	"ctbc_util/common"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileBytes, err := ioutil.ReadFile("file_names.txt")
	common.CheckError(err)
	fileName := string(fileBytes)

	fileArr := strings.Split(fileName, "\r\n")
	common.CheckDirectoryExistsAndMkdir("./result")
	for _, file := range fileArr {
		os.Create(fmt.Sprintf("./result/PBDM %v_測試報告.doc", file))
	}
}

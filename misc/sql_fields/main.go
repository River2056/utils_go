package misc

import (
	"ctbc_util/common"
	"fmt"
	"io/ioutil"
	"strings"
)

func GenerateNULLasField() {
	bytes, err := ioutil.ReadFile("./input.txt")
	common.CheckError(err)
	content := string(bytes)

	arr := strings.Split(content, "\r\n")
	for _, v := range arr {
		v = strings.ToUpper(v)
		v = strings.ReplaceAll(v, " ", "_")
		v = strings.ReplaceAll(v, "/", "_")
		v = strings.ReplaceAll(v, "(", "_")
		v = strings.ReplaceAll(v, ")", "_")
		fmt.Printf("NULL as %s,\n", v)
	}

}

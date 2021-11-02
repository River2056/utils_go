package misc

import (
	"ctbc_util/common"
	"fmt"
	"regexp"
	"strings"
)

func SetMethodCaps(line string) {
	var result string = ""
	arr := strings.Split(line, "\t")
	m := regexp.MustCompile("set[a-z]{1}")
	fmt.Println(len(arr))
	for _, v := range arr {
		v = strings.TrimSpace(v)
		if len(v) >= 4 {
			subset := m.Find([]byte(v))
			converted := m.ReplaceAllString(v, "set"+strings.ToUpper(string(subset[3])))
			fmt.Println(converted)
			result += fmt.Sprintf("%v\n", converted)
		}
	}
	common.OutputResults("./results", "output.txt", result)
}

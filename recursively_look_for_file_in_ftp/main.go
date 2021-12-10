package main

import (
	"ctbc_util/common"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	ftpConnectionConfig := common.FtpConnectionConfig{}
	common.InitializeFromGeneralConfig(&ftpConnectionConfig)
	conn := common.GetFtpConnection(ftpConnectionConfig)
	config := common.RecursiveLookForFileInFtpModule{}
	common.InitializeFromGeneralConfig(&config)

	walker := conn.Walk(config.RecursiveLookForFileInFtp.Path)
	common.RecursiveExecuteCustomFunctionOnFtp(walker, func(fileName string) {
		for _, reg := range config.RecursiveLookForFileInFtp.File {
			m := regexp.MustCompile("(?i)" + reg)
			fileSubString := strings.Trim(m.FindString(fileName), " ")
			if len(fileSubString) > 0 && strings.Contains(fileName, fileSubString) {
				stat := walker.Stat()
				fmt.Printf("%10s\t%10s\n", stat.Time, fileName)
			}
		}
	})
}

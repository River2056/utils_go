package main

import (
	"ctbc_util/common"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	ftpConnectionConfig := common.FtpConnectionConfig{}
	common.InitializeFromGeneralConfig(&ftpConnectionConfig)
	conn := common.GetFtpConnection(ftpConnectionConfig)
	functionConfig := common.RetriveFileFromFtpModule{}
	common.InitializeFromGeneralConfig(&functionConfig)

	walker := conn.Walk(functionConfig.RetriveFileFromFtp.Path)
	common.RecursiveExecuteCustomFunctionOnFtp(walker, func(fileName string) {
		for _, reg := range functionConfig.RetriveFileFromFtp.FilesReg {
			matcher := regexp.MustCompile("(?i)" + reg)
			fileSubName := strings.Trim(matcher.FindString(fileName), " ")
			if len(fileSubName) > 0 && strings.Contains(fileName, fileSubName) {
				fileSubNameArrTmp := strings.Split(strings.Replace(fileSubName, functionConfig.RetriveFileFromFtp.Path, "", 1), "/")
				var fileSubNameArr []string = make([]string, 0)
				for _, v := range fileSubNameArrTmp {
					if len(v) > 0 {
						fileSubNameArr = append(fileSubNameArr, v)
					}
				}
				destArr := strings.Split(functionConfig.RetriveFileFromFtp.Dest, "\\")
				var tempFilePathArr []string = append(destArr, fileSubNameArr...)
				tempFilePath := strings.Join(tempFilePathArr, string(os.PathSeparator))
				fileFromFtpRes, err := conn.Retr(fileName)
				common.CheckError(err)

				tempDirPath := string(tempFilePath[:strings.LastIndex(tempFilePath, "\\")])
				dirExists, err := common.Exists(tempDirPath)
				common.CheckError(err)
				if !dirExists {
					err = os.MkdirAll(tempDirPath, 0755)
					common.CheckError(err)
				}

				bytes, _ := ioutil.ReadAll(fileFromFtpRes)
				err = ioutil.WriteFile(tempFilePath, bytes, 0755)
				common.CheckError(err)

				fmt.Printf("Done Copying files: from %s, to %s\n", fileName, tempFilePath)
				fileFromFtpRes.Close()
			}
		}
	})
}

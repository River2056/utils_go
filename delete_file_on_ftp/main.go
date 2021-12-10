package main

import (
	"ctbc_util/common"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	ftpConnectionConfig := common.FtpConnectionConfig{}
	common.InitializeFromGeneralConfig(&ftpConnectionConfig)
	conn := common.GetFtpConnection(ftpConnectionConfig)
	functionConfig := common.DeleteFileOnFtpModule{}
	common.InitializeFromGeneralConfig(&functionConfig)

	walker := conn.Walk(functionConfig.DeleteFileOnFtp.Path)

	common.RecursiveExecuteCustomFunctionOnFtp(walker, func(fileName string) {
		for _, filesToDeleteReg := range functionConfig.DeleteFileOnFtp.FilesToDelete {
			matcher := regexp.MustCompile("(?i)" + filesToDeleteReg)
			fileSubName := strings.Trim(matcher.FindString(fileName), " ")
			if len(fileSubName) > 0 && strings.Contains(fileName, fileSubName) {
				// keep file in recycle bin
				fileFromFtpRes, err := conn.Retr(fileName)
				common.CheckError(err)
				if check, err := common.Exists("./recycle_bin"); !check {
					common.CheckError(err)
					err = os.Mkdir("./recycle_bin", 0755)
					common.CheckError(err)
				}

				recycleBinPath := filepath.Join("./recycle_bin", fileSubName)
				recycleBinDirPath := recycleBinPath[:strings.LastIndex(recycleBinPath, string(os.PathSeparator))]
				if exists, _ := common.Exists(recycleBinDirPath); !exists {
					err = os.MkdirAll(recycleBinDirPath, 0755)
					common.CheckError(err)
				}

				fileBytes, err := ioutil.ReadAll(fileFromFtpRes)
				common.CheckError(err)
				err = ioutil.WriteFile(recycleBinPath, fileBytes, 0755)
				common.CheckError(err)
				fmt.Printf("backup to recycle bin: %s\n", recycleBinPath)
				fileFromFtpRes.Close()

				// remove the file from the ftp
				err = conn.Delete(fileName)
				common.CheckError(err)
				fmt.Printf("file deleted: %v\n", fileName)
			}
		}
	})
	fmt.Println("delete complete!")
}

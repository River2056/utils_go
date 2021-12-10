package main

import (
	"ctbc_util/common"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	ftpConnectionConfig := common.FtpConnectionConfig{}
	common.InitializeFromGeneralConfig(&ftpConnectionConfig)

	// initializing connection to ftp
	conn := common.GetFtpConnection(ftpConnectionConfig)

	// function configs
	functionConfig := common.LookForFileAndUploadToDestinationModule{}
	common.InitializeFromGeneralConfig(&functionConfig)

	filesToStore := make(map[string]string)

	walker := conn.Walk(functionConfig.LookForFileAndUploadToDestination.Dest)

	// files to store
	filepath.Walk(functionConfig.LookForFileAndUploadToDestination.Src, func(path string, info os.FileInfo, err error) error {
		common.CheckError(err)
		if !info.IsDir() {
			filesToStore[info.Name()] = path

		}
		return nil
	})

	for walker.Next() {
		fileName := walker.Path()
		for k, v := range filesToStore {
			if strings.Contains(fileName, k) {
				file, err := os.Open(v)
				common.CheckError(err)
				defer file.Close()

				err = conn.Stor(fileName, file)
				common.CheckError(err)
				fmt.Printf("update complete: from %s to %s\n", k, fileName)
				delete(filesToStore, k)
			}
		}
	}

	for k, v := range filesToStore {
		// ftp does not contain file, upload as new file
		// navigate to parent folder first
		err := conn.ChangeDir("/")
		common.CheckError(err)
		file, err := os.Open(v)
		common.CheckError(err)
		defer file.Close()

		ftpDestPath := strings.Replace(file.Name(), functionConfig.LookForFileAndUploadToDestination.Src, functionConfig.LookForFileAndUploadToDestination.Dest, -1)
		ftpDestPath = strings.ReplaceAll(ftpDestPath, "\\", "/")
		ftpDestDirPath := string(ftpDestPath[:strings.LastIndex(ftpDestPath, "/")])
		ftpArr := strings.Split(ftpDestDirPath, "/")
		for _, ftpPath := range ftpArr {
			if len(ftpPath) > 0 {
				err = conn.MakeDir(ftpPath)
				err = conn.ChangeDir(ftpPath)
			}
		}

		err = conn.Stor(ftpDestPath, file)
		common.CheckError(err)
		fmt.Printf("upload new file: from %s to %s\n", file.Name(), ftpDestPath)
		delete(filesToStore, k)
	}

	fmt.Println("Done Updating!")
}

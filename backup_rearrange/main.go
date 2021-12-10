package main

import (
	"ctbc_util/common"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	backupFolderPath := filepath.Join("C:\\", "", "Users", "user", "Desktop", "backup_files")
	backupTempFolders := []string{
		filepath.Join(backupFolderPath, "AVQ_DWH_KTR"),
		filepath.Join(backupFolderPath, "AVQ_DWH_schema"),
		filepath.Join(backupFolderPath, "PBDM_KTR"),
		filepath.Join(backupFolderPath, "PBDM_schema"),
		filepath.Join(backupFolderPath, "misc"),
	}
	if check, _ := common.Exists(backupTempFolders[0]); !check {
		err := os.MkdirAll(backupTempFolders[0], 0755)
		common.CheckError(err)
	}
	if check, _ := common.Exists(backupTempFolders[1]); !check {
		err := os.MkdirAll(backupTempFolders[1], 0755)
		common.CheckError(err)
	}
	if check, _ := common.Exists(backupTempFolders[2]); !check {
		err := os.MkdirAll(backupTempFolders[2], 0755)
		common.CheckError(err)
	}
	if check, _ := common.Exists(backupTempFolders[3]); !check {
		err := os.MkdirAll(backupTempFolders[3], 0755)
		common.CheckError(err)
	}

	filepath.Walk(backupFolderPath, func(path string, file os.FileInfo, err error) error {
		tempPath := strings.Replace(path, file.Name(), "", 1)
		tempPath = tempPath[:strings.LastIndex(tempPath, string(os.PathSeparator))]
		if hasItem := common.Contains(backupTempFolders, tempPath); !hasItem && !file.IsDir() {
			subDir := tempPath[strings.LastIndex(tempPath, string(os.PathSeparator)):]
			subDir = strings.Replace(subDir, string(os.PathSeparator), "", 1)
			arrangePath := filepath.Join(backupFolderPath, subDir, file.Name())
			err = os.Rename(path, arrangePath)
			common.CheckError(err)
			fmt.Println(subDir)
		}
		return nil
	})

	dirsToRemove := make([]string, 0)
	filepath.Walk(backupFolderPath, func(path string, file os.FileInfo, err error) error {
		if backupFolderPath != path {
			if hasItem := common.Contains(backupTempFolders, path); !hasItem && file.IsDir() {
				dirsToRemove = append(dirsToRemove, path)
			}
		}
		return nil
	})
}

package main

import (
	"ctbc_util/common"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("Please provide root folder path to purge!")
	}
	folderPath := args[1]

	err := filepath.Walk(folderPath, func(path string, file os.FileInfo, e error) error {
		common.CheckError(e)

		if !file.IsDir() {
			// delete file
			if e := os.Remove(path); e != nil {
				common.CheckError(e)
			}
			fmt.Printf("done purging: %v\n", path)
		}
		return nil
	})

	common.CheckError(err)

	fmt.Printf("done purging folder: %v\n", folderPath)
}

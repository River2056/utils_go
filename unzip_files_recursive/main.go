package main

import (
	"archive/zip"
	"ctbc_util/common"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide root folder path!")
		os.Exit(0)
	}
	rootFolderPath := args[1]

	filepath.Walk(rootFolderPath, func(path string, file os.FileInfo, e error) error {
		common.CheckError(e)

		if !file.IsDir() && strings.HasSuffix(file.Name(), ".zip") {
			r, e := zip.OpenReader(path)
			common.CheckError(e)
			defer r.Close()

			for _, f := range r.File {
				rc, e := f.Open()
				common.CheckError(e)
				defer rc.Close()

				dPath := path[:strings.LastIndex(path, string(os.PathSeparator))]
				fPath := filepath.Join(dPath, f.Name)

				if !f.FileInfo().IsDir() {
					var fDir string
					if lastIndex := strings.LastIndex(fPath, string(os.PathSeparator)); lastIndex != -1 {
						fDir = fPath[:lastIndex]
					}

					e = os.MkdirAll(fDir, f.Mode())
					common.CheckError(e)
					openedFile, e := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
					common.CheckError(e)
					defer openedFile.Close()

					_, e = io.Copy(openedFile, rc)
					common.CheckError(e)
					fmt.Printf("done extracting %v\n", fPath)
				}
			}
		}
		return nil
	})

	fmt.Println("extraction complete!")
}

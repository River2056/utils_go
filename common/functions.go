package common

import (
	"io/ioutil"
	"os"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckDirectoryExistsAndMkdir(fileDirPath string) {
	if ok, _ := Exists(fileDirPath); !ok {
		os.Mkdir(fileDirPath, 0755)
	}
}

func OutputResults(resultDir, resultFilePath, outputData string) {
	CheckDirectoryExistsAndMkdir(resultDir)

	if len(outputData) != 0 {
		ioutil.WriteFile(resultDir+"/"+resultFilePath, []byte(outputData), 0644)
	}
}

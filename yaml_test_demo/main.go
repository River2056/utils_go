package main

import (
	"ctbc_util/common"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name     string   `yaml:"name"`
	Elements []string `yaml:"elements"`
}

func main() {
	fileBytes, err := ioutil.ReadFile("./config.yaml")
	common.CheckError(err)

	config := Config{}
	yaml.Unmarshal(fileBytes, &config)

	fmt.Println(config)
	fmt.Println(config.Name)
	fmt.Println(config.Elements)
	fmt.Println(config.Elements[0])

	fileInfo, err := ioutil.ReadDir(config.Name)
	common.CheckError(err)

	fmt.Println(fileInfo)
	for _, v := range fileInfo {
		fmt.Println(v)
	}
}

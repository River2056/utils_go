package main

import (
	"ctbc_util/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type Configuration struct {
	Path  string   `json:"json-file-path"`
	Attrs []string `json:"attr-to-search"`
}

var config Configuration
var result string

func lookUpForAttribute(jsonObj map[string]interface{}, attr string) {
	for k, v := range jsonObj {
		valueType := reflect.TypeOf(v).Kind()
		if valueType.String() != "slice" && k == attr {
			strV := fmt.Sprintf("%v", v)
			if len(strV) != 0 {
				fmt.Println(fmt.Sprintf("%v: %v", k, v))
				result += fmt.Sprintf("%v: %v\n", k, v)
			}
		} else if valueType.String() == "slice" {
			if k == attr {
				// check if slice has info
				if len(v.([]interface{})) != 0 {
					for _, element := range v.([]interface{}) {
						switch element.(type) {
						case map[string]interface{}:
							fmt.Println(k)
							result += fmt.Sprintf("%v\n", k)
							for elemK, elemV := range element.(map[string]interface{}) {
								fmt.Println(fmt.Sprintf("*\t%v: ", elemK))
								result += fmt.Sprintf("*\t%v: \n", elemK)
								fmt.Println(fmt.Sprintf("\t\t%v", elemV))
								result += fmt.Sprintf("\t\t%v\n", elemV)
							}
						default:
							fmt.Println(fmt.Sprintf("%v: %v", k, v))
							result += fmt.Sprintf("%v: %v\n", k, v)
						}
					}
				}
				return
			}
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				value := s.Index(i).Interface()
				switch item := value.(type) {
				case map[string]interface{}:
					lookUpForAttribute(item, attr)
				default:
					fmt.Println(value)
					result += fmt.Sprintf("%v\n", value)
				}
			}
		} else if valueType.String() == "map" {
			if k == attr {
				strV := fmt.Sprintf("%v", v)
				if len(strV) != 0 {
					fmt.Println(v)
					result += fmt.Sprintf("%v\n", v)
					return
				}
			}
			m := reflect.ValueOf(v).Interface()
			lookUpForAttribute(m.(map[string]interface{}), attr)
		}
	}
	if ok, _ := common.Exists("./result"); !ok {
		os.Mkdir("./result", 0755)
	}

	ioutil.WriteFile("./result/result.txt", []byte(result), 0644)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	configFile, err := ioutil.ReadFile("./config.json")
	checkError(err)
	json.Unmarshal(configFile, &config)
}

func main() {
	filePath := config.Path
	attrs := config.Attrs
	jsonFile, err := os.Open(filePath)
	if err != nil {
		panic("json file not found!")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var jsonObj map[string]interface{}
	json.Unmarshal(byteValue, &jsonObj)

	for _, attr := range attrs {
		lookUpForAttribute(jsonObj, attr)
	}
}

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
	args := os.Args
	if len(args) < 2 {
		panic("PLease provide java model name!")
	}
	javaModelName := args[1]
	file, err := ioutil.ReadFile("./fields.txt")
	common.CheckError(err)

	types, err := ioutil.ReadFile("./types.txt")
	common.CheckError(err)

	fieldsStr := string(file)
	typesStr := string(types)
	fieldArr := strings.Split(fieldsStr, "\r\n")
	typeArr := strings.Split(typesStr, "\r\n")

	if len(fieldArr) != len(typeArr) {
		panic("Please provide enough fields/types in both inputs!")
	}

	var result string = fmt.Sprintf("public class %v { \n", javaModelName)

	for index, field := range fieldArr {
		javaField := strings.ToLower(field)
		matcher := regexp.MustCompile(`_[a-zA-Z]{1}`)
		convertedString := matcher.ReplaceAllStringFunc(javaField, strings.ToUpper)
		convertedString = strings.ReplaceAll(convertedString, "_", "")
		typeAttribute := typeArr[index]
		typeAttribute = strings.ToLower(typeAttribute)
		switch typeAttribute {
		case "varchar":
			typeAttribute = "String"
		case "numberic":
			typeAttribute = "BigDecimal"
		case "numeric":
			typeAttribute = "BigDecimal"
		case "date":
			typeAttribute = "Date"
		case "tinyint":
			typeAttribute = "Long"
		case "integer":
			typeAttribute = "Integer"
		}

		result += fmt.Sprintf("\tprivate %v %v;\n", typeAttribute, convertedString)
	}

	result += fmt.Sprintf("}")

	common.OutputResults("./result", fmt.Sprintf("%v.java", javaModelName), result)
}

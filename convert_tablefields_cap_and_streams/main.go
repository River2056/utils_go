package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	content := string(file)
	data := strings.Split(content, "\r\n")

	for i := 1; i < len(data); i++ {
		row := strings.Split(data[i], "\t")
		row[0] = strings.ToUpper(row[0])
		data[i] = row[0] + "\t" + row[1]
	}

	for _, v := range data {
		fmt.Println(v)
	}

	outputData := strings.Join(data, "\n")

	// outputFile, _ := os.OpenFile("./output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// dataWriter := bufio.NewWriter(outputFile)
	// for _, d := range data {
	// 	_, _ = dataWriter.WriteString(d + "\n")
	// }
	// dataWriter.Flush()
	// outputFile.Close()
	ioutil.WriteFile("./output.txt", []byte(outputData), 0644)
}

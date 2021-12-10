package common

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jlaffaye/ftp"
	"gopkg.in/yaml.v2"
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

func ArrStrContains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
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

func InitializeValuesFromConfigForFtp() FtpConfig {
	config := FtpConfig{}
	configFile, err := ioutil.ReadFile("./config.yaml")
	yaml.Unmarshal(configFile, &config)
	CheckError(err)
	return config
}

func InitializeParamsFromConfigForFtp() FtpConnectionConfig {
	config := FtpConnectionConfig{}
	configFile, err := ioutil.ReadFile("./config.yaml")
	yaml.Unmarshal(configFile, &config)
	CheckError(err)
	return config
}

func InitializeFromGeneralConfig(customType interface{}) {
	if generalConfigExists, err := Exists(GENERAL_CONFIG_FILE_PATH); generalConfigExists {
		CheckError(err)
		bytes, err := ioutil.ReadFile(GENERAL_CONFIG_FILE_PATH)
		CheckError(err)
		yaml.Unmarshal(bytes, customType)
		return
	}

	configFile, err := ioutil.ReadFile("./config.yaml")
	yaml.Unmarshal(configFile, customType)
	CheckError(err)
	return
}

func GetFtpConnection(config FtpConnectionConfig) *ftp.ServerConn {
	options := ftp.DialWithDisabledEPSV(true)
	conn, err := ftp.Dial(config.Ftp, ftp.DialWithTimeout(5*time.Second), options)
	CheckError(err)
	err = conn.Login(config.User, config.Password)
	CheckError(err)
	return conn
}

func RecursiveExecuteCustomFunctionOnFtp(walker *ftp.Walker, fn func(string)) {
	for walker.Next() {
		fileName := walker.Path()
		fn(fileName)
	}
}

func GetDataBaseConnection() *sql.DB {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", SERVER, USER, PASSWORD, PORT, DATABASE)

	db, err := sql.Open("sqlserver", connString)
	CheckError(err)
	fmt.Printf("Connection established: %s\n", db)
	return db
}

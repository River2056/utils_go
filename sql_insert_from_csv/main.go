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
		panic("Please provide csv file path!")
	}

	csvFilePath := args[1]
	bytes, err := ioutil.ReadFile(csvFilePath)
	common.CheckError(err)

	contents := string(bytes)
	contentArr := strings.Split(contents, "\r\n")
	var resultSql string = ""

	for _, row := range contentArr {
		cellValues := strings.Split(row, ",")
		for i := 0; i < len(cellValues); i++ {
			m := regexp.MustCompile("^\\d+?(\\.?)\\d+?$")
			if m.Match([]byte(cellValues[i])) {
				cellValues[i] = fmt.Sprintf("%v", cellValues[i])
			} else {
				cellValues[i] = fmt.Sprintf("'%v'", cellValues[i])
			}
			if len(cellValues[i]) == 2 {
				cellValues[i] = fmt.Sprintf("%v", "null")
			}
			if i == 2 || i == 25 || i == 28 {
				cellValues[i] = fmt.Sprint("''")
			}
		}
		values := strings.Join(cellValues, ",")
		resultSql += fmt.Sprintf("insert into PBRCLendingPortfolioLimit (CR_Name, CR_Root_Id, Mainline_Id, Facility_Id, Line_Name, Limit_CCY_Code, Limit_Amt, Available_Amt, Outstanding_Amt, Start_Date, Expiry_Date, Avail_Flag, Annual_Review_Date, Portfolio_Code, Process_Date, Business_Entity, Facility_Limit_Amt, Facility_Available_Amt, Facility_Collateral_Amt, Freeze_Date, Freeze_Reason_Code, Interest_Margin, Std_Interest_Margin, Internal_Remark, New_Loan_Industrial_Code, Loan_Purpose, Activation_Date, ORR_Adjust, Status, Active_Flow, Flow_Id) values (%v)\n", values)
		// fmt.Println(fmt.Sprintf("index: %v, len: %v", index, len(cellValues)))
	}

	common.OutputResults("./result", "output.sql", resultSql)
}

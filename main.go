package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type planData struct {
    PartnerCode string
    PartnerName string
    AddressInfo string
	Fullname string
	TemplateName string
	RouteDate string
	CheckDate string
	VisitDate string
	CheckOutAt string
	SurveyDate string
	Uid string
	TemplateId string
	SurveyId string
	VisitID string
}
func main() {

    csvFile, err := os.Open("SurveyId.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
    
    csvLines, err := csv.NewReader(csvFile).ReadAll()
	
    if err != nil {
        fmt.Println(err)
    }    
    for _, line := range csvLines {
        plan := planData{
            PartnerCode: line[0],
            PartnerName: line[1],
			AddressInfo: line[2],
			Fullname: line[3],
			TemplateName: line[4],
			RouteDate: line[5],
			CheckDate: line[6],
			VisitDate: line[7],
			CheckOutAt: line[8],
			SurveyDate: line[9],
			Uid: line[10],
			TemplateId: line[11],
			SurveyId: line[12],
			VisitID: line[13],
        }
        fmt.Println(plan.PartnerCode + " - " + plan.PartnerName + " - " + plan.AddressInfo+ " - " + plan.Fullname+" - "+ plan.TemplateName + " - " + plan.RouteDate + " - " + plan.CheckDate+ " - " + plan.VisitDate+" - "+ plan.CheckOutAt+ " - " + plan.SurveyDate + " - " + plan.Uid+ " - " + plan.TemplateId+" - "+ plan.SurveyId+" - "+plan.VisitID)
		fmt.Println("Go lang Insert")

		db, err := sql.Open("mysql", "root:irisgroup@tcp(127.0.0.1:3306)/demo")
		defer db.Close()
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
	
		// insForm, err := db.Prepare("INSERT INTO plan(PartnerCode, PartnerName, AddressInfo, Fullname, TemplateName, RouteDate, CheckDate, VisitDate, CheckOutAt, SurveyDate, Uid, TemplateId, SurveyId, VisitID) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
        // if err != nil {
        //     panic(err.Error())
        // }
		// defer db.Close()
        // res, err := insForm.Exec(plan.PartnerCode, plan.PartnerName, plan.AddressInfo, plan.Fullname, plan.TemplateName, plan.RouteDate, plan.CheckDate, plan.VisitDate, plan.CheckOutAt, plan.SurveyDate, plan.Uid, plan.TemplateId, plan.SurveyId, plan.VisitID)
		// if err != nil {
		// 	fmt.Println(err)
        // }
		
		// id, err := res.LastInsertId()
		
		// fmt.Println("Insert id", id)
		// defer db.Close()
    }
	totalDataRows := len(csvLines)
    fmt.Println("Total CVS: of rows:", totalDataRows)

}
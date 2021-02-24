package main

import (
	"fmt"
	"vtx.io/repsub"
	"bufio"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to RepSub! A Report Submission Board")
	fmt.Print("Title:")
    reportName, _ := reader.ReadString('\n')
	fmt.Print("Description:")
	reportDesc, _ := reader.ReadString('\n')
	fmt.Print("Severity (1-10):")
	reportSvrty, _ := reader.ReadString('\n')

	response := repsub.SubmitReport(reportName,  reportDesc, reportSvrty, "Ray")
	fmt.Println(response)

	fmt.Println("Do you want to view your reports? (Y/N)")
	viewReports, _ := reader.ReadString('\n')
	if viewReports === "Y" {
		reports := repsub.ViewReports()
		fmt.Println(reports)
	}
}
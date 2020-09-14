package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

//Details details
type Details struct {
	ID   string `csv:"ID"`
	Name string `csv:"Name"`
	Delta  string `csv:"Delta"`
}

//Log logs
type Log struct {
	ID    int  `csv:"ID"`
	Found bool `csv:"Found"`
	Count int  `csv:"Count"`
}

func main() {

	AuditDetails := []Details{}
	A := Details{
		ID:   "User",
		Delta:  "92",
		Name: "ReferenceName",
	}
	AuditDetails = append(AuditDetails, A)

	AuditLogs := []Log{}
	for i := 0; i < 5; i++ {
		AuditLog := Log{}
		AuditLog.ID = i
		AuditLog.Found = false
		AuditLog.Count = i * 5
		AuditLogs = append(AuditLogs, AuditLog)
	}

	FinalCOntent := []byte{}
	FileContentBytes, err := gocsv.MarshalBytes(AuditDetails)
	if err != nil {
		fmt.Print("AERROR: Could not marshal file:", err)
		os.Exit(1)
	}

	LogContent, err := gocsv.MarshalBytes(AuditLogs)
	if err != nil {
		fmt.Print("BERROR: Could not marshal file:", err)
		os.Exit(1)
	}

	FinalCOntent = append(FinalCOntent, FileContentBytes...)
	FinalCOntent = append(FinalCOntent, []byte("\n")...)
	FinalCOntent = append(FinalCOntent, LogContent...)

	f, err := os.Create("Output.csv")
	if err != nil {
		fmt.Print("Error Writeing Output")
	}

	defer f.Close()

	f.Write(FinalCOntent)
	f.Sync()

}

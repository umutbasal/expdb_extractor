package extract

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Entry struct {
	Id          string `json:"id"`
	File        string `json:"file"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Author      string `json:"author"`
	Type        string `json:"type"`
	Platform    string `json:"platform"`
	Port        string `json:"port"`
}

type db map[string]Entry

func csvRead(csvFile *os.File) (db db) {
	db = map[string]Entry{}
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		emp := Entry{
			Id:          line[0],
			File:        line[1],
			Description: line[2],
			Date:        line[3],
			Author:      line[4],
			Type:        line[5],
			Platform:    line[6],
			Port:        line[7],
		}
		db[line[1]] = emp
	}
	return db
}

func getRowByFileName(db db, fileName string) (e Entry, err error) {
	return db[fileName], nil
}

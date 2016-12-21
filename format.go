package main

import (
	_ "encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/pquerna/ffjson/ffjson"
)

type Report struct {
	ProjectName string `json: "ProjectName"`
	ProjectNo   string `json: "ProjectNo"`
	Name        string `json: "Name"`
	Description string `json: "Description"`
	Issue       string `json: "Issue"`
	Planning    string `json: "Planning"`
	LogTime     string `json: "LogTime"`
}

// text
func format2txt(rows *string) (lines string) {
	lines = strings.Replace(*rows, "/", "\t", -1)
	fmt.Println(lines)
	return
}

// json
func format2json(rows *string, sendtime int) (toj string) {
	lines := strings.Split(*rows, "/")
	if len(lines) != 6 {
		fmt.Println("Error format: ", *rows)
	}
	var logtime string
	logtime = time.Unix(int64(sendtime), 0).Format("2006-01-02 15:04:05")
	var rpt *Report = &Report{
		ProjectName: lines[0],
		ProjectNo:   lines[1],
		Name:        lines[2],
		Description: lines[3],
		Issue:       lines[4],
		Planning:    lines[5],
		LogTime:     logtime,
	}
	j, err := ffjson.Marshal(&rpt)
	if err != nil {
		panic(err)
	}
	toj = string(j)
	fmt.Println(toj)
	return
}

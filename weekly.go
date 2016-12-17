package main

import (
	_ "encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	_ "sort"
	"strings"
)

// type Report struct {
// 	ProjectName []string
// 	ProjectNo   []string
// 	Name        []string
// 	Description []string
// 	Question    []string
// 	Plan        []string
// }

var infile *string = flag.String("i", "", "输入文件")
var outfile *string = flag.String("o", "", "输出文件")

func main() {
	flag.Parse()

	content, err := readAll(*infile)
	if err != nil {
		panic(err)
	}

	rows := getRows(&content)
	//fmt.Println(rows)
	rows = rm_duplicate(&rows)

	lines := make([]string, 0)
	for _, row := range rows {
		lines = append(lines, format2txt(&row))
	}
	saveFile(*outfile, &lines)
}

// 读取文件内容
func readAll(filepath string) (content string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		return
	}
	defer file.Close()

	ra, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	content = string(ra)
	fmt.Printf("<<读取文件 %s 成功\n", filepath)
	return
}

// 保存文件
func saveFile(filepath string, contents *[]string) {
	//删除已存在文件
	if _, err := os.Stat(filepath); err == nil {
		err = os.Remove(filepath)
		if err != nil {
			fmt.Println("Failed to remove file: ", filepath)
			panic(err)
		}
	}

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Failed to open file:", filepath)
		panic(err)
	}
	defer file.Close()

	// 逐行写入
	for _, content := range *contents {
		file.WriteString(content)
		file.WriteString("\n")
	}
	fmt.Printf(">>保存文件 %s 成功\n", filepath)

}

// 分割文本行
func getRows(content *string) (rows []string) {
	for _, row := range strings.Split(*content, "\n") {
		//row = strings.Replace(row, "/", "", 1) //去除开头"/"符号
		row = strings.TrimSpace(strings.Replace(row, "\n", "", -1)) //去除换行符&去除空格
		if len(row) != 0 {
			rows = append(rows, row)
		}
	}
	return
}

// 去重
func rm_duplicate(list *[]string) (x []string) {
	for _, i := range *list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return
}

// 格式化文本
func format2txt(rows *string) (lines string) {
	lines = strings.Replace(*rows, "/", "\t", -1)
	return
}

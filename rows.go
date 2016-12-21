package main

import (
	_ "flag"
	"strings"

	_ "github.com/axgle/mahonia"
)

//var infile *string = flag.String("i", "", "输入文件")
//var outfile *string = flag.String("o", "", "输出文件")
//var redis_conf *string = flag.String("r", "", "Redis服务器IP及端口，默认：127.0.0.1:6379")

//func main() {
//	flag.Parse()

//	content, err := readAll(*infile)
//	if err != nil {
//		panic(err)
//	}

//	rows := getRows(&content)
//	rows = rm_duplicate(&rows)

//	lines := make([]string, 0)
//	for _, row := range rows {
//		lines = append(lines, format2txt(&row))
//		toredis(format2json(&row))
//	}
//	saveFile(*outfile, &lines)
//}

// 分割文本行
func getRows(content *string) (rows []string) {
	for _, row := range strings.Split(*content, "week") {
		//row = strings.Replace(row, "/", "", 1) //去除开头"/"符号
		row = strings.TrimSpace(strings.Replace(row, "\r\n", "", -1)) //去除换行符&去除空格
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

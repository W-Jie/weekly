package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"
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

var infile *string = flag.String("i", "", "输入文件（必须）")
var outfile *string = flag.String("o", "output.txt", "输出文件（可选）")

func main() {
	flag.Parse()
	//fmt.Println(os.Args[0])
	if len(os.Args) < 2 {
		// log.Fatal("infile can not be blank!")
		flag.PrintDefaults()
	}
	content := ReadAll(*infile)
	rows := GetRows(&content)
	rows = RmDuplicate(&rows)

	now := time.Now().Unix()
	rlist := ReportList(make([]Report, 0))

	for _, row := range rows {
		rpt := Format(&row, now)
		//fmt.Println(tojson(&rpt))
		rlist = append(rlist, rpt)
	}

	sort.Sort(rlist) // 返回[]Report
	//fmt.Println(rlist)
	SaveFile(*outfile, &rlist)
}

// 读取文件内容
func ReadAll(filepath string) (content string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ra, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	content = string(ra)
	log.Printf("读取文件 <<%s 成功\n", filepath)
	return
}

// 保存文件
func SaveFile(filepath string, contents *ReportList) bool {
	//删除已存在文件
	// if _, err := os.Stat(filepath); err == nil {
	// 	err = os.Remove(filepath)
	// 	if err != nil {
	// 		fmt.Println("Failed to remove file: ", filepath)
	// 		return false
	// 	}
	// }

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 逐行写入
	for _, content := range *contents {
		out := content.ProjectName + "\t" + content.ProjectNo + "\t" + content.Name + "\t" + content.Description + "\t" + content.Issue + "\t" + content.Planning + "\n"
		file.WriteString(out)
	}
	log.Printf("保存文件 >>%s 成功\n", filepath)
	return true
}

// 分割行
func GetRows(content *string) (rows []string) {
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
func RmDuplicate(list *[]string) (x []string) {
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

// 每行结构化
func Format(rows *string, sendtime int64) Report {
	lines := strings.Split(*rows, "/")
	if len(lines) != 6 {
		log.Fatal("格式异常: ", *rows)
	} else if len(lines[2]) > 12 {
		log.Fatal("姓名格式异常: ", *rows)
	}

	logtime := time.Unix(int64(sendtime), 0).Format("2006-01-02 15:04:05")
	rpt := Report{
		ProjectName: lines[0],
		ProjectNo:   lines[1],
		Name:        lines[2],
		Description: lines[3],
		Issue:       lines[4],
		Planning:    lines[5],
		LogTime:     logtime,
	}
	return rpt
}

// 结构体切片排序
type ReportList []Report // ReportList为sort接口的类型，需要实现Less()、Swap()和Len()方法;

func (list ReportList) Len() int {
	return len(list)
}

func (list ReportList) Less(i, j int) bool {
	if list[i].ProjectNo < list[j].ProjectNo {
		return true
	} else if list[i].ProjectNo > list[j].ProjectNo {
		return false
	} else if list[i].ProjectName < list[j].ProjectName {
		return true
	} else if list[i].ProjectName > list[j].ProjectName {
		return false
	} else {
		return list[i].Name < list[j].Name
	}
}

func (list ReportList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]

}

// json串行化
func Tojson(rpt *Report) (toj string) {
	j, err := json.Marshal(&rpt)
	if err != nil {
		log.Fatal(err)
	}
	toj = string(j)
	return
}

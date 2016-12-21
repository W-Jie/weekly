package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

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
	//enc:=mahonia.NewEncoder("utf8") //字符编码转换
	//content =　enc.ConvertString(content)
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

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to open file:", filepath)
		panic(err)
	}
	defer file.Close()

	// 逐行写入
	for _, content := range *contents {
		file.WriteString(content)
		file.WriteString("\r\n")
	}
	fmt.Printf(">>保存文件 %s 成功\n", filepath)

}

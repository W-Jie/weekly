package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/JamesWone/SmartQQ"
)

//使用自己封装的Http-Client包
var client_turing smartqq.Client = smartqq.Client{
	IsKeepCookie: true,
	Timeout:      5,
}

//调用图灵机器人Api
func getResponseByTuringRobot(request string) string {
	resp_turing, err := client_turing.Post("http://www.niurenqushi.com/app/simsimi/ajax.aspx", "txt="+request)
	if err != nil {
		return ""
	}
	return resp_turing.Body
}

func main() {
	//初始化一个QClient
	client := smartqq.QClient{}
	//当二维码图片变动后触发
	client.OnQRChange(func(qc *smartqq.QClient, image_bin []byte) {
		//将二维码保存至当前目录，打开手机QQ扫描二维码后即可登录成功
		fmt.Println("正在保存二维码图片.")
		file_image, err := os.OpenFile("二维码图片.jpg", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file_image.Close()
		if _, err := file_image.Write(image_bin); err != nil {
			fmt.Println(err)
			return
		}
	})
	//当登录成功后触发
	client.OnLogined(func(qc *smartqq.QClient) {
		fmt.Println("登录成功了！")
	})
	//当收到消息后触发
	client.OnMessage(func(qc *smartqq.QClient, qm smartqq.QMessage) {
		fmt.Println("收到新消息了：")
		fmt.Println(qm)
		fmt.Println("\n ==========\n")
		content := qm.Content
		if strings.Contains(qm.Content, "week") {
			switch qm.Poll_type {
			//QQ好友消息
			case "message":
				//发送给QQ好友
				//qc.SendToQQ(qm.From_uin, getResponseByTuringRobot(content)+"\n(by:ai)")
				//qc.SendToQQ(qm.From_uin, content)
			//QQ群消息
			case "group_message":
				//发送给QQ群
				//qc.SendToGroup(qm.From_uin, getResponseByTuringRobot(content)+"\n(by:ai)")
			//讨论组消息
			case "discu_message":
				//发送给讨论组
				rows := getRows(&content)
				rows = rm_duplicate(&rows)
				fmt.Println(rows)
				for _, row := range rows {
					//lines = append(lines, format2txt(&row))
					toredis("10.89.13.60:6379", format2json(&row, qm.Time))
				}
				qc.SendToDiscuss(qm.From_uin, "周报保存成功")
			}
		}
	})

	fmt.Println("开始登录.")
	//开始登录，并自动收发消息
	client.Run()
}

func send() {

}

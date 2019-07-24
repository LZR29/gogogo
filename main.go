package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	data, err := ioutil.ReadFile(`H:\log.txt`)
	if err != nil {
		fmt.Println("读取文件出错")
		return
	}
	now := time.Now().Unix() //获取时间戳
	fmt.Println(now)
	fileName := `H:\` + strconv.Itoa(int(now)) + `.txt`
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("创建文件失败")
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		fmt.Println("写入失败")
	}
	_, err = os.Stat(`H:\workspace\CBS-DeviceCenterService\src\git.shulaquan.com\cloud\device-service\log`)
	if err == nil {
		fmt.Println("目录存在！！")
		return
	}
	if os.IsNotExist(err) {
		fmt.Println("目录不存在")
		var path string
		if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
			path = "\\"
		} else {
			path = "/"
		}
		dir := `H:\workspace\CBS-DeviceCenterService\src\git.shulaquan.com\cloud\device-service\log`
		err := os.Mkdir(dir+path+"md", os.ModePerm) //在当前目录下生成md目录
		if err != nil {
			fmt.Println("创建失败！")
		}
	}
	t1 := time.Now().Year()   //年
	t2 := time.Now().Month()  //月
	t3 := time.Now().Day()    //日
	t4 := time.Now().Hour()   //小时
	t5 := time.Now().Minute() //分钟
	t6 := time.Now().Second() //秒
	curTime := strconv.Itoa(t1) + "年" + strconv.Itoa(int(t2)) + "月" + strconv.Itoa(t3) + "日" +
		strconv.Itoa(t4) + "时" + strconv.Itoa(t5) + "分" + strconv.Itoa(t6) + "秒"
	fmt.Println(curTime)
}

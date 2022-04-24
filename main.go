package main

import (
	"bufio"
	"flag"
	"os"
	"shu/shu"
)

var versionStr = "0.0.1"
var info = "互助吹水小组专用 " + versionStr

func main() {
	println(info)
	version := flag.Bool("v", false, "获取版本号")
	course := flag.String("course", "", "要获取答案的课程id")
	session := flag.String("session", "", "登录账号的session id")
	target := flag.String("o", "", "要存到的目标文件，如果不填则打印到控制台")
	
	flag.Parse()
	
	if *version {
		println(versionStr)
		return
	}
	if *course == "" {
		println("请传入course参数，可添加-h参数查看帮助")
		return
	}
	if *session == "" {
		println("请传入session参数，可添加-h参数查看帮助")
		return
	}
	
	answer, err := shu.GetAnswer(shu.NewApi(*course, *session))
	if err != nil {
		println(err)
		return
	}
	
	if *target == "" {
		println(info + "\n" + answer)
		return
	}
	
	file, err := os.OpenFile(*target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer shu.CloseCloser(file)
	write := bufio.NewWriter(file)
	_, err = write.WriteString(info + "\n" + answer)
	if err != nil {
		println(err)
		return
	}
	err = write.Flush()
	if err != nil {
		println(err)
		return
	}
}

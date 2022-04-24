# 上海大学继续教育作业答案获取工具

使用步骤：
1. 登录学习平台，获取到用户session（获取方式见最后）
2. 获取课程的id（获取方式见最后）
3. 用记事本打开一键获取答案.bat
4. 将session填入到-session ""的两个双引号中间
5. 将课程id填入到-course ""的两个双引号中间
6. 保存关闭
7. 双击一键获取答案.bat
8. 查看控制台或者生成的文件

```
Usage of ./release.exe:
-course string 要获取答案的课程id
-o string 要存到的目标文件，如果不填则打印到控制台
-session string 登录账号的session id
-v 获取版本号
```

【course】要获取答案的课程id
学习平台连接参数里就有，如：https://sdjj.ct-edu.com.cn/learning/student/studentIndex.action#!/index/course/home?courseId=ff8080817e6b5da2017e7a5b3002487a）


【session】用户的session
进入到学习平台后，打开开发者工具(F12)，应用程序 -> Cookie -> SESSION
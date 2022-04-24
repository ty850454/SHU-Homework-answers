package shu

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Api struct {
	courseId string
	session  string
	header   http.Header
}

func NewApi(courseId string, sessionId string) Api {
	header := make(http.Header)
	
	header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Mobile Safari/537.36 Edg/93.0.961.38")
	header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	header.Set("Accept", "application/json, text/plain, */*")
	header.Set("Cookie", "CLARITY_INDEX=0; homeUrl=http%3A%2F%2Fwww.ad.shu.edu.cn%2FIndexPage.html; SESSION="+sessionId)
	
	return Api{
		courseId: courseId,
		session:  sessionId,
		header:   header,
	}
	
}

func (api *Api) QueryHomeworkList() ([]HomeworkData, error) {
	data := "courseId=" + api.courseId
	
	var homeworkListRes HomeworkListRes
	api.doPost("https://sdjj.ct-edu.com.cn/learning/student/studentDataAPI.action?functionCode=queryHomeworkList", data, &homeworkListRes)
	
	if homeworkListRes.ReturnCode != "S0000" {
		return nil, errors.New(homeworkListRes.ReturnMessage)
	}
	
	return homeworkListRes.HomeworkDataList, nil
}

func (api *Api) ShowHomework(homeworkId string) (QuestionList, error) {
	data := "courseId=" + api.courseId + "&homeworkId=" + homeworkId
	
	var homeworkRes HomeworkRes
	api.doPost("https://sdjj.ct-edu.com.cn/learning/student/studentDataAPI.action?functionCode=showHomework", data, &homeworkRes)
	
	if homeworkRes.ReturnCode != "S0000" {
		return homeworkRes.QuestionObj, errors.New(homeworkRes.ReturnMessage)
	}
	
	return homeworkRes.QuestionObj, nil
}

func (api *Api) DoHomework(homeworkId string) (QuestionList, error) {
	data := "courseId=" + api.courseId + "&homeworkId=" + homeworkId
	
	var homeworkRes HomeworkRes
	api.doPost("https://sdjj.ct-edu.com.cn/learning/student/studentDataAPI.action?functionCode=doHomework", data, &homeworkRes)
	
	if homeworkRes.ReturnCode != "S0000" {
		return homeworkRes.QuestionObj, errors.New(homeworkRes.ReturnMessage)
	}
	
	return homeworkRes.QuestionObj, nil
}

func (api *Api) doPost(url string, data string, v any) {
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(data))
	req.Header = api.header
	res, _ := http.DefaultClient.Do(req)
	defer CloseCloser(res.Body)
	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(body, v)
}

func CloseCloser(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		println(err)
	}
}

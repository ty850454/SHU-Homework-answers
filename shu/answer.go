package shu

import (
	"strconv"
	"strings"
)

func GetAnswer(api Api) (string, error) {
	
	builder := strings.Builder{}
	
	homeworkList, err := api.QueryHomeworkList()
	if err != nil {
		return "", err
	}
	for _, homework := range homeworkList {
		builder.WriteString(homework.Title)
		builder.WriteByte('\n')
		
		var question QuestionList
		var err error
		if homework.HomeworkStatus == 3 {
			question, err = api.ShowHomework(homework.Id)
		} else {
			question, err = api.DoHomework(homework.Id)
		}
		if err != nil {
			return "", err
		}
		
		if len(question.DanxuanList) != 0 {
			// 单选题
			builder.WriteString("【单选题】\n")
			builder.WriteString(writeQuestions(question.DanxuanList))
		}
		
		if len(question.DuoxuanList) != 0 {
			// 多选题
			builder.WriteString("【多选题】\n")
			builder.WriteString(writeQuestions(question.DuoxuanList))
		}
		
		if len(question.PanduanList) != 0 {
			// 判断题
			builder.WriteString("【判断题】\n")
			builder.WriteString(writeQuestions(question.PanduanList))
		}
	}
	
	return builder.String(), nil
}

func writeQuestions(questions []Question) string {
	builder := strings.Builder{}
	
	for i, question := range questions {
		builder.WriteString(strconv.Itoa(i + 1))
		builder.WriteString("、")
		builder.WriteString(question.Sanswer)
		builder.WriteByte(' ')
		builder.WriteString(question.Title)
		builder.WriteByte('\n')
		for _, option := range question.OptionList {
			builder.WriteByte('\t')
			builder.WriteString(option.Index)
			builder.WriteByte(' ')
			builder.WriteString(option.Title)
			builder.WriteByte(' ')
			if option.IsAnswer {
				builder.WriteString("√")
			}
			builder.WriteByte('\t')
		}
		builder.WriteByte('\n')
	}
	builder.WriteByte('\n')
	builder.WriteByte('\n')
	return builder.String()
}

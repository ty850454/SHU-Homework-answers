package shu

type Course struct {
	courseId  string
	sessionId string
}

type HomeworkListRes struct {
	ReturnCode       string         `json:"returnCode"`
	ReturnMessage    string         `json:"returnMessage"`
	HomeworkDataList []HomeworkData `json:"homeworkDataList"`
}

type HomeworkData struct {
	ItemId             string `json:"itemId"`
	Sequence           int    `json:"sequence"`
	Id                 string `json:"id"`
	Title              string `json:"title"`
	StartDate          string `json:"startDate"`
	EndDate            string `json:"endDate"`
	HomeworkCategory   string `json:"homeworkCategory"`
	HomeworkType       string `json:"homeworkType"`
	AnswerShowType     string `json:"answerShowType"`
	QuestionCreateType string `json:"questionCreateType"`
	SectionTitle       string `json:"sectionTitle"`
	ChapterTitle       string `json:"chapterTitle"`
	AllowShowHomework  bool   `json:"allowShowHomework"`
	LockedIsExpire     int    `json:"lockedIsExpire"`
	IsExpire           int    `json:"isExpire"`
	RedoNum            int    `json:"redoNum"`
	AllowRedoNum       int    `json:"allowRedoNum"`
	HomeworkStatus     int    `json:"homeworkStatus"`
	Score              int    `json:"score"`
	HasWenda           bool   `json:"hasWenda"`
	AllowRedo          bool   `json:"allowRedo"`
	Message            string `json:"message"`
	ShowTkScore        bool   `json:"showTkScore"`
}

type QuestionOption struct {
	Content   string `json:"content"`
	Index     string `json:"index"`
	IsAnswer  bool   `json:"isAnswer"`
	IndexName string `json:"indexName"`
	Title     string `json:"title"`
}

type Question struct {
	Id          string           `json:"id"`
	Title       string           `json:"title"`
	Score       int              `json:"score"`
	Sscore      int              `json:"sscore"`
	Uscore      int              `json:"uscore"`
	Uanswer     string           `json:"uanswer"`
	Sanswer     string           `json:"sanswer"`
	Type        string           `json:"type"`
	AnswerAnaly string           `json:"answerAnaly"`
	OptionList  []QuestionOption `json:"optionList"`
}

type QuestionList struct {
	SwfUrl          string     `json:"swfUrl"`
	KeGuanTiScore   string     `json:"keGuanTiScore"`
	ZhuGuanTiScore  string     `json:"zhuGuanTiScore"`
	Score           string     `json:"score"`
	TotalNote       string     `json:"totalNote"`
	DanxuanList     []Question `json:"danxuanList"`
	DuoxuanList     []Question `json:"duoxuanList"`
	PanduanList     []Question `json:"panduanList"`
	IsTranscodeFile bool       `json:"isTranscodeFile"`
}

type HomeworkRes struct {
	ReturnCode    string       `json:"returnCode"`
	ReturnMessage string       `json:"returnMessage"`
	QuestionObj   QuestionList `json:"questionObj"`
	UseUeditor    bool         `json:"useUeditor"`
}

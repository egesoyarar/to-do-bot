package entity

import "time"

type User struct {
	UserId  string `json:"userid"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Job     string `json:"job"`
	ChatId  int64  `json:"chatid"`
}

type Task struct {
	TaskId      string        `json:"taskid"`
	TaskName    string        `json:"taskname"`
	Description string        `json:"description"`
	ShareDate   time.Time     `json:"sharedate"`
	Duration    time.Duration `json:"duration"`
	DueDate     time.Time     `json:"duedate"`
}

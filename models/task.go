package models

type Task struct {
	ID uint8 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
}
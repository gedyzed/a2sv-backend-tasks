package models

type Task struct {

	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Date string `json:"date"`
	Status string `json:"status"`
	
}


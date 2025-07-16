package data

import "task_manager/models"

var Tasks = []models.Task{
	{
		Id:          "1",
		Title:       "Implement login",
		Description: "Create login route and logic",
		DueDate:     "2025-07-20",
		Status:      "In Progress",
	},
	{
		Id:          "2",
		Title:       "Set up database",
		Description: "Configure PostgreSQL for dev environment",
		DueDate:     "2025-07-18",
		Status:      "Completed",
	},
}


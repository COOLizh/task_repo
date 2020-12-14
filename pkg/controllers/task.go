package controllers

import (
	"github.com/COOLizh/task_repo/pkg/db"
	"github.com/COOLizh/task_repo/pkg/models"
	"github.com/rs/zerolog/log"
)

// GetAll return all tasks found in database or error
func GetAll() (tasks []models.Task, err error) {
	tasks, err = db.GetAllTasks()
	if err != nil {
		return
	}
	if len(tasks) == 0 {
		log.Print("Empthy data")
	}
	return
}

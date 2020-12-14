// Package controllers for API
package controllers

import (
	"github.com/COOLizh/task_repo/pkg/client"
	"github.com/COOLizh/task_repo/pkg/db"
	"github.com/COOLizh/task_repo/pkg/models"
)

// SendSolution gets all needful info from DB, calls client
// and returns client response or error
func SendSolution(solutionRequest *models.SolutionRequest, taskID int) (solutionResult *models.SolutionResult, err error) {
	task, err := db.GetTaskByID(taskID)
	if err != nil {
		return
	}

	testCases, err := db.GetTestCasesByTaskID(task.ID)
	if err != nil {
		return
	}

	solutionResult, err = client.SendSolution(
		solutionRequest.Code,
		solutionRequest.Language,
		task.TimeLimit,
		task.Memory,
		testCases,
	)
	return
}

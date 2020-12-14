// Package handlers for API
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/COOLizh/task_repo/pkg/auth"
	"github.com/COOLizh/task_repo/pkg/controllers"
	"github.com/COOLizh/task_repo/pkg/models"
	"github.com/gin-gonic/gin"
)

// HealthCheck view
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

// SendSolutionHandler sends solution to executioner and returns the result
func SendSolutionHandler(c *gin.Context) {
	taskIDVal := c.Param("task_id")
	taskID, err := strconv.Atoi(taskIDVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong ID by task")
		return
	}

	solutionRequest := new(models.SolutionRequest)

	decoder := json.NewDecoder(c.Request.Body)
	err = decoder.Decode(solutionRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Wrong request body")
		return
	}

	solutionResult, err := controllers.SendSolution(solutionRequest, taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Something went wrong while receiving data")
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, solutionResult)
}

// Registration handler for Gin router
func Registration(c *gin.Context) {
	user := new(models.User)
	err := user.PopulateFromRequest(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	if !user.IsValid() {
		c.JSON(http.StatusUnprocessableEntity, "Invalid user data")
		return
	}

	err = auth.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "User did not register")
		return
	}
	response := user.PrepareResponse()

	c.JSON(http.StatusOK, response)
}

// Login handler for Gin router
func Login(c *gin.Context) {
	user := new(models.User)
	err := user.PopulateFromRequest(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	if !user.IsValid() {
		c.JSON(http.StatusUnprocessableEntity, "Invalid user data")
		return
	}
	response, err := auth.LoginUser(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	c.SetCookie("token", response.Authorization, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, response)
}

// AuthCheck needed to test authentication
func AuthCheck(c *gin.Context) {
	_, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "You must login in the system")
	} else {
		c.JSON(http.StatusOK, "secure endpoint")
	}
}

// GetAllTasks returns all tasks list
func GetAllTasks(c *gin.Context) {
	tasks, err := controllers.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error at Marshall data")
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskAPI interface {
	AddTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetTaskList(c *gin.Context)
	GetTaskListByCategory(c *gin.Context)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskRepo service.TaskService) *taskAPI {
	return &taskAPI{taskRepo}
}

func (t *taskAPI) AddTask(c *gin.Context) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.taskService.Store(&newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add task success"})
}

func (t *taskAPI) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")

	// Check if taskID is a valid integer
	id, err := strconv.Atoi(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	// Parse the JSON payload
	var updatedTask model.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	// Assign the taskID to the updatedTask struct
	updatedTask.ID = id

	// Update the task using taskService
	err = t.taskService.Update(&updatedTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "update task success"})
}


func (t *taskAPI) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")

	// Check if taskID is a valid integer
	id, err := strconv.Atoi(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	// Delete the task using taskService
	err = t.taskService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "delete task success"})
}


func (t *taskAPI) GetTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := t.taskService.GetByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskAPI) GetTaskList(c *gin.Context) {
	taskList, err := t.taskService.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, taskList)
}


func (t *taskAPI) GetTaskListByCategory(c *gin.Context) {
	taskList, err := t.taskService.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, taskList)
}


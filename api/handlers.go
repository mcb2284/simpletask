package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simpletask/database"
	"github.com/simpletask/types"
)

func CreateUser(c *gin.Context) {
	var newUser types.User

	err := c.BindJSON(&newUser)
	if err != nil {
		log.Fatal("error binding json")
	}

	resp, err := database.CreateUser(newUser)
	if err != nil {
		log.Fatal("couldn't create user")
	}

	c.JSON(http.StatusOK, resp)
}

func GetUser(c *gin.Context) {
	user_id := c.Param("userid")

	user, err := database.GetUser(user_id)
	if err != nil {
		log.Fatal("couldn't create user")
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {

	user_id := c.Param("userid")
	var updatedUser types.User

	err := c.BindJSON(&updatedUser)
	if err != nil {
		log.Fatal("error binding json")
	}

	updatedUser.ID = user_id
	user, err := database.UpdateUser(updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	user_id := c.Param("userid")

	err := database.DeleteUser(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, "Deleted user")
}

func GetTasks(c *gin.Context) {
	var req types.Request

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	fmt.Printf("query params: %v \n", req)

	tasks := database.GetTasks(req)

	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	task_id := c.Param("taskid")

	task, err := database.GetTask(task_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, task)

}

func CreateTask(c *gin.Context) {

	var newTask types.Task

	err := c.BindJSON(&newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	createdTask, err := database.CreateTask(newTask)

	c.JSON(http.StatusOK, createdTask)

}

func UpdateTask(c *gin.Context) {

	task_id := c.Param("taskid")
	var updatedTask types.Task

	err := c.BindJSON(&updatedTask)
	if err != nil {
		log.Fatal("error binding json")
	}

	updatedTask.ID = task_id
	task, err := database.UpdateTask(updatedTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	task_id := c.Param("taskid")

	err := database.DeleteTask(task_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, "Deleted task")
}

func Todo(c *gin.Context) {
	c.JSON(200, gin.H{"message": "TODO: implement this endpoint"})
}

package api

import "github.com/gin-gonic/gin"

func InitServer() {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("", CreateUser)
		user.GET("/:userid", GetUser)
		user.PATCH("/:userid", UpdateUser)
		user.DELETE("/:userid", DeleteUser)
	}

	task := router.Group("/tasks")
	{
		task.GET("", GetTasks)
		task.POST("", CreateTask)
		task.GET("/:taskid", GetTask)
		task.PATCH("/:taskid", UpdateTask)
		task.DELETE("/:taskid", DeleteTask)
	}

	router.Run(":8080")
}

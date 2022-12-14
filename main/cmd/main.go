package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/internal/handlers"
	"main/internal/storage"
	"net/http"
)

func main() {
	//tasks := storage.NewTaskMapStorage()
	tasks, err := storage.NewTasksMongoStorage()
	if err != nil {
		log.Fatalln(err)
	}
	defer tasks.Disconnect()
	router := gin.Default()

	router.GET("/check", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "It's okay"})
	})

	router.GET("/task/:task_name", handlers.HandleTaskAccess(tasks))

	router.POST("/task", handlers.HandleTaskCreation(tasks))
	router.POST("/task/:task_name", handlers.HandleTaskUpdate(tasks))
	router.DELETE("/task/:task_name", handlers.HandleTaskDelete(tasks))

	router.GET("/work/:task_name/:work_name", handlers.HandleWorkAccess(tasks))
	router.POST("/work/:task_name/:work_name", handlers.HandleWorkNeedsSetup(tasks))
	router.POST("/work/:task_name", handlers.HandleWorkCreation(tasks))
	router.DELETE("/work/:task_name/:work_name", handlers.HandleWorkDelete(tasks))

	err = router.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}

}

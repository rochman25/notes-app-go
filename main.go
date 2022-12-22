package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
	"notes-app-go/config"
	"notes-app-go/database"
	"notes-app-go/database/migrations"
	"notes-app-go/src/utils"
)

func main() {

	var m string
	flag.StringVar(
		&m, "m", "none", `for migration`,
	)
	flag.Parse()

	if m == "migrate" {
		migrations.Migrate()
	} else if m == "rollback" {
		migrations.Rollback()
	} else if m == "status" {
		database.InitDB()
		migrations.Status()
	}

	router := gin.New()
	router.HandleMethodNotAllowed = true

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Greeting from Go World.",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		response := utils.MethodNotAllowed
		c.JSON(http.StatusMethodNotAllowed, response)
	})

	router.NoRoute(func(c *gin.Context) {
		response := utils.NotFoundError
		c.JSON(http.StatusNotFound, response)
	})

	err := router.Run(":" + config.GetEnv("APP_PORT"))
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}

}

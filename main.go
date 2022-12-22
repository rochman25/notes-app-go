package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"notes-app-go/config"
	"notes-app-go/database"
	"notes-app-go/database/migrations"
	"notes-app-go/src/models"
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

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.HandleMethodNotAllowed = true

	router.GET("/users", func(c *gin.Context) {
		var users []models.Users
		database.InitDB()
		if err := database.DBCon.Find(&users).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, utils.SuccessResponse("", "", users))
		}
		db, _ := database.DBCon.DB()
		defer db.Close()
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

package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/gameiro/gila-test/interface/http/service"
	"gitlab.com/gameiro/gila-test/logging"
)

var logger logging.Logger

func postMessage(c *gin.Context) {
	var s service.NotificationService
	notifications := s.Send(c)
	if notifications != nil {
		s.Log(&logger, notifications)
		c.JSON(http.StatusOK, gin.H{"success": true})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
	}
}

func getLog(c *gin.Context) {
	logItems := logger.GetItems()
	var data []string
	for i := 0; i < len(logItems); i++ {
		data = append(data, logItems[len(logItems)-i-1])
	}
	if len(data) == 0 {
		data = append(data, "")
	}
	c.IndentedJSON(http.StatusOK, data)
}

func getCategories(c *gin.Context) {
	var s service.CategoryService
	c.IndentedJSON(http.StatusOK, s.FetchAll(c))
}

func main() {
	router := gin.Default()

	ipProtection := "127.0.0.1"
	if len(os.Args) > 1 {
		ipProtection = os.Args[1]
	}
	// It is the simplest way to prevent abuse. In a real scenario I would change this for an API Authentication
	router.SetTrustedProxies([]string{ipProtection})

	router.GET("/categories", getCategories)
	router.POST("/message", postMessage)
	router.GET("/log", getLog)

	router.Run(ipProtection + ":8080")
}

// pages.go

package main

import (
	"os"
	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	basePath := os.Getenv("MOC_ADMIN_BASEPATH")

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"basePath": basePath}, "index.html")
}

func showLogsPage(c *gin.Context) {
	basePath := os.Getenv("MOC_ADMIN_BASEPATH")
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":   "Logs Selection",
		"basePath": basePath}, 
		"logs.html")
}
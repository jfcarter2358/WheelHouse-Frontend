// api.go

package main

import (
	"net/http"
	// "os/exec"
	// "strings"
	"github.com/gin-gonic/gin"
	// "log"
	"fmt"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

func getCompassURLByNameAndVersion(c *gin.Context) {
	// Get the deployment that we want logs for
	name := c.Param("name")
	version:= c.Param("version")
	fmt.Println(name)
	fmt.Println(version)
	// Get the Compass object
	compass, err := getCompasssByNameAndVersion(name, version)
    if err == nil {
			// Return the URL for the COmpass
			c.JSON(http.StatusOK, gin.H{"url": compass.URL})
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func getCompassObjectsByName(c *gin.Context) {
	// Get the name that we want the compass for
	name := c.Param("name")
	// Get the Compass object
	compasses, err := getCompassesByName(name)
    if err == nil {
			// Return the URL for the COmpass
			c.JSON(http.StatusOK, gin.H{"compasses": compasses})
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func getCompassObjectsByFuzzyName(c *gin.Context) {
	// Get the name that we want the compass for
	name := c.Param("name")
	
	var targets []string
	for _, compass := range getAllCompasses() {
		targets = append(targets, compass.Name)
	}
	matches := fuzzy.Find(name, targets)
	matches = removeDuplicateStringValues(matches)

	c.JSON(http.StatusOK, gin.H{"matchess": matches})
}
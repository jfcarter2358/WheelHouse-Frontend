// models.article.go

package main

import (
	"errors"
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/google/uuid"
)

type Compasses struct {
    CompassList []Compass `json:"compasses"`
}

type Compass struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Version string `json:"version"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var compassList = []Compass{}

func readCompassDataFile(){
	// Open our jsonFile
	jsonFile, err := os.Open("data/compasses.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("Successfully Opened data/compasses.json")

	var compasses Compasses
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &compasses)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	for _, compass := range compasses.CompassList {
		fmt.Println("adding compass")
		fmt.Println(compass)
		compassList = append(compassList, compass)
	}
}

// Return a list of all the compasses
func getAllCompasses() []Compass {
	return compassList
}

// Fetch a compass based on the ID supplied
func getCompassByID(id string) (*Compass, error) {
	for _, c := range compassList {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, errors.New("Compass not found")
}

// Fetch a compass based on the name supplied
func getCompassesByName(name string) ([]Compass, error) {
	var returnList []Compass
	for _, c := range compassList {
		if c.Name == name {
			fmt.Println("Found match")
			returnList = append(returnList, c)
		}
	}
	if len(returnList) > 0{
		return returnList, nil
	}
	return nil, errors.New("Compass not found")
}

// Fetch a compass based on the name supplied
func getCompasssByNameAndVersion(name string, version string) (*Compass, error) {
	for _, c := range compassList {
		if c.Name == name && c.Version == version {
			return &c, nil
		}
	}
	return nil, errors.New("Compass not found")
}

// Create a new article with the title and content provided
func createNewCompass(name, version, url string) (*Compass, error) {
	// Set the ID of a new article to one more than the number of articles
	id := uuid.New().String()


	c := Compass{ID: id, Name: name, Version: version, URL: url}

	// Add the article to the list of articles
	compassList = append(compassList, c)

	return &c, nil
}
// Declare Package
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jay-SP/gcp/data_engineering/edgarfacts/internal/facts"
	"github.com/jay-SP/gcp/data_engineering/edgarfacts/internal/storage"
)

// Main function
func main() {

	//Parse Command Line Arguemnts
	var cik string
	var organization string
	var name string
	var email string

	//this we can use variables as cmd args

	flag.StringVar(&cik, "cik", "", "CIK Number")
	flag.StringVar(&organization, "organization", "", " Name of the organization")
	flag.StringVar(&name, "name", "", "Name")
	flag.StringVar(&email, "email", "", "Email")

	//parse varaibles from userinput to varibles declared
	flag.Parse()

	//Validate command line Arguments
	if len(cik) != 10 {
		panic("CIK must be of lenth 10")
	}
	if organization == "" {
		panic("Enter organization")
	}
	if name == "" {
		panic("Please provide your email address")
	}
	if email == "" {
		panic("Please provide your email address")
	}

	//Load Environment Variables
	bucketName := os.Getenv("BUCKET")
	folderPath := os.Getenv("STAGE")
	if bucketName == "" || folderPath == "" {
		panic("Eroor reading ENV")
	}

	//configure logger, stdflags for dates etc
	logger := log.New(os.Stdout, "", log.LstdFlags)

	//Load Data
	logger.Printf("Loading Facts for %s\n", cik)
	facts, err := facts.LoadFacts(cik, organization, name, email)
	if err != nil {
		panic(err)
	}

	//Upload to Google Storage
	fileName := fmt.Sprintf("%s.json", cik)
	filePath := filepath.Join(folderPath, fileName)

	logger.Printf("Uploading Facts to %s on bucket %s\n", fileName, bucketName)

	err = storage.UploadBytes(facts, bucketName, filePath)
	if err != nil {
		panic(err)
	}
}

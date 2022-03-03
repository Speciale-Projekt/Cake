package main

import (
	"Cake/filehandler"
	"flag"
	"log"
	"os"
	"path/filepath"
)

var (
	// Version is the version of the application
	Version = "0.0.0"
	// Build is the build number of the application
	Build = "0"
	// Commit is the commit hash of the application
	Commit = "0"
	// Branch is the branch of the application
	Branch = "0"

	readFolderPath string
	// TODO: Figure out how to interact with AFL
)

func init() {
	// Intialize the logger
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.StringVar(&readFolderPath, "read", "", "The folder to read from")
	flag.Parse()

	//if readFolderPath == "" {
	//	panic("No read folder specified")
	//}

}
func main() {
	// Get this file location
	thisLocation, _ := os.Getwd()
	readFolder := filehandler.New(filepath.Join(thisLocation, "test"))
	go readFolder.Monitor()

}

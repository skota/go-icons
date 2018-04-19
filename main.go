package main

import (
	"flag"
	"fmt"
	utils "icon-go/lib"
	"os"
)

func main() {
	//create flag vars
	var (
		filename  = flag.String("image", "foo", "Image name")
		operation = flag.String("operation", "icon", "operation name - create splashes or icons")
		osname    = flag.String("os", "ios", "Os name - ios or android")
	)
	flag.Parse()
	params := utils.Config{Image: *filename, Operation: *operation, Os: *osname}

	checkparams(params)
	utils.DoTasks(params)

	fmt.Printf("image name:  %s", *filename)
	fmt.Printf("operation:  %s", *operation)
	fmt.Printf("Os name:  %s", *osname)
}

func checkparams(params utils.Config) {
	//check file exists and is valid and is of right size
	_, err := utils.FileExists(params.Image)
	if err != nil {
		fmt.Println("File does not exist")
		os.Exit(-1)
	}

	//ensure file has correct extension
	fieldType := utils.CheckFileType(params.Image)
	if fieldType != 1 {
		fmt.Println("Invalid file. Only jpg and png file types are accepted")
		os.Exit(-1)
	}

	//and size isless than 1 mb
	fileSize := utils.CheckfileSize(params.Image)
	if fileSize > 1024 {
		fmt.Println("File size cannot exceed 1mb.")
		os.Exit(-1)
	}
}

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

}

func checkparams(params utils.Config) {
	//check file exists and is valid and is of right size
	_, err := utils.FileExists(params.Image)
	if err != nil {
		fmt.Println("File does not exist")
		os.Exit(-1)
	}
}

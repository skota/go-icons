package main

import (
	"flag"
	utils "lib/utilities"
)

func main() {
	//create flag vars
	var (
		filename  = flag.String("image", "foo", "Image name")
		operation = flag.String("operation", "icon", "operation name - create splashes or icons")
		osname    = flag.String("os", "ios", "Os name - ios or android")
	)
	flag.parse()
	params = utils.Config{Image: *filename, Operation: *operation, Os: *osname}

	checkparams(params)
}

func checkparams(params utils.Config) {
	//check file exists and is valid

	//launcj do tasks
}

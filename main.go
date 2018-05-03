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
	flag.Usage = usage

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
	//how do i use the first param returned ? ignore it?
	_, fieldType := utils.CheckFileType(params.Image)
	if fieldType != 1 {
		fmt.Println("Invalid file. Only jpg and png file types are accepted")
		os.Exit(-1)
	}

	//check file type..

	//and size isless than 1 mb
	fileSize := utils.CheckfileSize(params.Image)

	if fileSize > 12496341 {
		fmt.Println("File size cannot exceed 4mb.")
		os.Exit(-1)
	}

}

//
var usageinfo string = `icon-go is a Go implemented CLI tool for generating splashes and icons for android and ios.

Usage:

	go-icons [flags] [METHOD] URL [ITEM [ITEM]]

flags:
  -a, -auth=USER[:PASS]       Pass a username:password pair as the argument
  -b, -bench=false            Sends bench requests to URL
  -b.N=1000                   Number of requests to run
  -b.C=100                    Number of requests to run concurrently
  -body=""                    Send RAW data as body
  -f, -form=false             Submitting the data as a form
  -j, -json=true              Send the data in a JSON object
  -p, -pretty=true            Print Json Pretty Format
  -i, -insecure=false         Allow connections to SSL sites without certs
  -proxy=PROXY_URL            Proxy with host and port
  -print="A"                  String specifying what the output should contain, default will print all information
         "H" request headers
         "B" request body
         "h" response headers
         "b" response body
  -v, -version=true           Show Version Number

METHOD:
  bat defaults to either GET (if there is no request data) or POST (with request data).

URL:
  The only information needed to perform a request is a URL. The default scheme is http://,
  which can be omitted from the argument; example.org works just fine.

ITEM:
  Can be any of:
    Query string   key=value
    Header         key:value
    Post data      key=value
    File upload    key@/path/file

Example:

	bat beego.me

`

func usage() {
	fmt.Println(usageinfo)
	os.Exit(2)
}

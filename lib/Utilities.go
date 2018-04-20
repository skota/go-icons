package lib

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

//Config --
type Config struct {
	Image     string
	Operation string
	Os        string
}

type imagesize struct {
	width  int
	height int
}

//Dotasks ...
func DoTasks(params Config) {
	// if params.Operation == 'splashes' {
	// 	createSplashes(params)
	// } else {
	// 	createIcons(params)
	// }
	createIcons()

	// ResizeIcon()
}

//FileExists -------check if file exists --------------
func FileExists(fname string) (bool, error) {
	if _, err := os.Stat(fname); err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
		return false, err
	}
	return true, nil
}

//CheckFileType -------check if file has correct extension --------------
func CheckFileType(fname string) int {
	// is this really the correct way to check for file type
	// what if file is simply renamed to have a jpg extension
	fileparts := strings.SplitAfter(fname, ".")
	fileExtentions := map[string]bool{"PNG": true,
		"png":  true,
		"JPG":  true,
		"JPEG": true,
		"jpeg": true,
		"jpg":  true}

	if fileExtentions[fileparts[1]] {
		return 1 //file extension is valid
	}
	return -1
}

//CheckfileSize -------check if file is under 1 mb --------------
func CheckfileSize(fname string) (imgsize int64) {
	imgsize = 0
	fi, _ := os.Stat(fname)

	//what kind of error handling is needed here
	imgsize = fi.Size()
	return imgsize
}

//ResizeIcon ---
// paramaters - filename, osName (default IOS)
// read diff sizes to resize from toml file
// if output location is not present..create  see folder struct below
//
func ResizeIcon() {
	//by this time we know file aleady exists
	src, err := imaging.Open("/Users/sriramkota/Downloads/pics/IMG_3446.JPG")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	// make a copy of original
	imgCopy := imaging.Clone(src)

	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	destPath := path.Join(currDir, "copy.jpeg")

	// resize copy and save with new name
	err = imaging.Save(imgCopy, destPath)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	fmt.Println("Completed image resize -----------")
	// repeat
}

//func createIcons(image string, osname string) int {
func createIcons() int {
	//create dir structure
	currDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Unable to get current directory ")
		os.Exit(0)
	}

	imageRootDir := filepath.Join(currDir, "images")
	if err := os.MkdirAll(imageRootDir, 0711); err != nil {
		fmt.Println("Unable to create image root directory current directory ")
		os.Exit(0)
	}

	//create nested dirs to store images

	/*
		create map with image_name and dimensions

		loop through map
			clone main image
			resize cloned image
			save with  name
		end loop



	*/

	// for _, imgsiz = range imgSizes {
	// 	//clone image
	// 	//pass image and size params to resiz

	// }
	//open image
	//defer close
	// for _, imgsiz = range imgSizes {
	// 	//clone image
	// 	//pass image and size params to resiz
	// }
	return 1
}

// func createSplashes(image string, osname string) int {
// 	//returns 1 if success, -1 if error
// }

//createIconDir ...
// func createIconDir() {

// }

//createSplashDir ...
// func createSplashDir() {

// }

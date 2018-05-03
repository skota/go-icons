package lib

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/disintegration/imaging"
)

//Config --
type Config struct {
	Image     string
	Operation string
	Os        string
}

//----- image size
type imagesize struct {
	width  int
	height int
}

//--- struct mapping image name and size
type iconNames struct {
	name   string
	width  int
	height int
	os     string
}

//Dotasks ...
func DoTasks(params Config) {
	// if params.Operation == 'splashes' {
	// 	createSplashes(params)
	// } else {
	// 	createIcons(params)
	// }
	createIcons("ios")

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

// //CheckFileType -------check if file has correct extension --------------
// func CheckFileType(fname string) int {
// 	// is this really the correct way to check for file type
// 	// what if file is simply renamed to have a jpg extension
// 	fileparts := strings.SplitAfter(fname, ".")
// 	fileExtentions := map[string]bool{"PNG": true,
// 		"png":  true,
// 		"JPG":  true,
// 		"JPEG": true,
// 		"jpeg": true,
// 		"jpg":  true}

// 	if fileExtentions[fileparts[1]] {
// 		return 1 //file extension is valid
// 	}
// 	return -1
// }

//CheckFileType -------read file magic bytes to verify format --------------
func CheckFileType(fname string) (string, int) {
	buf := make([]byte, 3)
	f, err := os.Open(fname)

	_, err = f.Read(buf)
	if err != nil && err != io.EOF {
		return "unknown type", -1
	}
	//lets see the first three bytes
	fmt.Printf("these are the 3 bytes - %s ", string(buf))

	if len(buf) > 2 && buf[0] == 0xFF && buf[1] == 0xD8 && buf[2] == 0xFF {
		return "Jpg", 1
	} else {
		return "Not Jpg", -1
	}

}

//CheckfileSize -------check if file is under 1 mb --------------
func CheckfileSize(fname string) (imgsize int64) {
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
	//TODO  --  fix this to take image name as param
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

	//TODO  --  Destination file name..
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
func createIcons(osname string) int {
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

	//next create subfolder ios or android

	//then change to currFolder/images/osname   - ex ./images/ios

	imgList := getIconList("ios")

	for _, img := range imgList {
		fmt.Printf("Image name: is %s, width is %d, height is %d\n", img.name, img.width, img.height)
		//resizedIcon
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

	return 1
}

func getIconList(os string) []iconNames {
	if os == "ios" {
		return []iconNames{
			{"icon-40.png", 40, 40, "ios"},
			{"icon-40@2x.png", 80, 80, "ios"},
			{"icon-40@2x.png", 120, 120, "ios"},
			{"icon-50.png", 50, 50, "ios"},
			{"icon-50@2x.png", 100, 100, "ios"},
			{"icon-60.png", 60, 60, "ios"},
			{"icon-60@2x.png", 120, 120, "ios"},
			{"icon-60@3x.png", 180, 180, "ios"},
			{"icon-72.png", 72, 72, "ios"},
			{"icon-72@2x.png", 144, 144, "ios"},
			{"icon-76.png", 76, 76, "ios"},
			{"icon-72@2x.png", 152, 152, "ios"},
			{"icon-83.5@2x.png", 167, 167, "ios"},
			{"icon-1024.png", 1024, 1024, "ios"},
			{"icon-small.png", 29, 29, "ios"},
			{"icon-small@2x.png", 58, 58, "ios"},
			{"icon-small@3x.png", 87, 87, "ios"},
			{"icon.png", 57, 57, "ios"},
			{"icon.png", 114, 114, "ios"}}

	} else {
		//return images for android
		return []iconNames{{"drawable-hdpi-icon.png", 72, 72, "android"},
			{"drawable-ldpi-icon.png", 36, 36, "android"},
			{"drawable-mdpi-icon.png", 48, 48, "android"},
			{"drawable-xhdpi-icon.png", 96, 96, "android"},
			{"drawable-xxhdpi-icon.png", 144, 144, "android"},
			{"drawable-xxxhdpi-icon.png", 192, 192, "android"}}
	}

}

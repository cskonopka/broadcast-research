package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/cskonopka/gomu"
)

type FileObject struct {
	originFile []string
	stripedMP4 []string
	jpgRoot    []string
	jpgDir     []string
	slimDir    []string
	filename   []string
}

func main() {

	dir := []string{
		// "/Users/io/Desktop/2000-01-january",
		// "/Volumes/vs01_042015-102015/2015-04-april",
		"/Volumes/vs01_042015-102015/2015-05-may",
		// "/Volumes/vs01_042015-102015/2015-06-june",
		// "/Volumes/vs01_042015-102015/2015-07-july",
		// "/Volumes/vs01_042015-102015/2015-08-august",
		// "/Volumes/vs01_042015-102015/2015-09-september",
		// "/Volumes/vs01_042015-102015/2015-10-october",
	}

	for n := 0; n < len(dir); n++ {
		var stripMP4Files, captureMP4Files []string

		hi := gomu.ReadDirRmDups(dir[n])
		dirFiles := searchFiles("edits", hi)

		for p := 0; p < len(dirFiles); p++ {
			editName := strings.SplitAfter(dirFiles[p], "/edits")
			captureMP4Files = append(captureMP4Files, dirFiles[p])
			stripMP4Files = append(stripMP4Files, editName[0][:len(editName[0])-6])
			// fmt.Println(stripMP4Files)
		}
		// // Isolate and Create jpg directories
		slimDirectories := gomu.RemoveDuplicates(stripMP4Files)

		// Create Root Directories
		for lo := 0; lo < len(slimDirectories); lo++ {
			// fmt.Println(slimDirectories[lo] + "/histo")
			cmd3 := exec.Command("mkdir", slimDirectories[lo]+"/histo")
			cmd3.Run()
		}

		// Create a jpg bundle for each mp4 file within individual jpg dirs
		for so := 0; so < len(captureMP4Files); so++ {

			editName := strings.SplitAfter(captureMP4Files[so], "/edits")
			// fmt.Println("orig file : ", editName)
			// completeFile := editName[0] + editName[1]
			// fmt.Println(completeFile)
			name := editName[1][:len(editName[1])-4]
			// fmt.Println("name : ", name)
			histoRoot := stripMP4Files[so][:] + "/histo"
			// fmt.Println("png root : ", histoRoot)

			inputFile := stripMP4Files[so][:] + "/png" + name + ".png"
			fmt.Println(inputFile)
			outputHistogram := histoRoot + name + "-histogram.txt"
			fmt.Println(outputHistogram + "\n")
			CreateHistogram(inputFile, outputHistogram)
			fmt.Println("done!")
			time.Sleep(2 * time.Second)

		}

	}
}

// CreateHistogram : Generate a RGB pixel histogram based on an analysis PNG file
func CreateHistogram(inputFile string, outputHistogram string) {
	// convert /Users/csk/Documents/_REPO/VCCC/mirrorMan.png -format %c histogam:info:histofile.txt
	_, err := exec.Command("convert", inputFile, "-format", "%c", "histogam:info:", outputHistogram).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}
}

// searchFiles -> Search for a directory within a month directory and search the variable after duplicates have been removed.
// ----- types --> The type of directory to be searched (ex. /gifs, /edits, /png, /raw)
// ----- rmDups --> The incoming array after duplicates have been removed.
func searchFiles(typer string, rmDups []string) []string {
	var savePng []string

	// Receive a type
	switch typer {
	case "gifs":
		// iterate over the incoming array
		i := 0
		for i < len(rmDups)-1 {
			i++
			// If the string contains the incoming type of directory
			h := strings.Contains(rmDups[i], typer+"/")
			switch h {
			case true:
				// If the string does not contain .DS_Store
				if !strings.Contains(rmDups[i], ".DS_Store") {
					// append to savePng string array
					savePng = append(savePng, rmDups[i])
				}
			}
		}
	case "edits":
		// iterate over the incoming array
		i := 0
		for i < len(rmDups)-1 {
			i++
			// If the string contains the incoming type of directory
			h := strings.Contains(rmDups[i], typer+"/")
			switch h {
			case true:
				// If the string does not contain .DS_Store
				if !strings.Contains(rmDups[i], ".DS_Store") {
					// append to savePng string array
					savePng = append(savePng, rmDups[i])
				}
			}
		}
	case "png":
		// iterate over the incoming array
		i := 0
		for i < len(rmDups)-1 {
			i++
			// If the string contains the incoming type of directory
			h := strings.Contains(rmDups[i], typer+"/")
			switch h {
			case true:
				// If the string does not contain .DS_Store
				if !strings.Contains(rmDups[i], ".DS_Store") {
					// append to savePng string array
					savePng = append(savePng, rmDups[i])
				}
			}
		}
	case "raw":
		// iterate over the incoming array
		i := 0
		for i < len(rmDups)-1 {
			i++
			// If the string contains the incoming type of directory
			h := strings.Contains(rmDups[i], typer+"/")
			switch h {
			case true:
				// If the string does not contain .DS_Store
				if !strings.Contains(rmDups[i], ".DS_Store") {
					// append to savePng string array
					savePng = append(savePng, rmDups[i])
				}
			}
		}
	}
	return savePng
}

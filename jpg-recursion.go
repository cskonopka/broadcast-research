package main

import (
	"os/exec"
	"strings"

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
		"/Users/io/Desktop/2000-01-january",
	}

	for n := 0; n < len(dir); n++ {

		hi := gomu.ReadDirRmDups(dir[n])
		// fmt.Println(hi)
		dirFiles := searchFiles("edits", hi)

		// fmt.Println(dirFiles)

		var stripMP4Files, captureMP4Files []string
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
			// fmt.Println(slimDirectories[lo] + "/jpg")
			cmd3 := exec.Command("mkdir", slimDirectories[lo]+"/jpg")
			cmd3.Run()
		}

		// Create individual jpg dir for each mp4 file
		for lo := 0; lo < len(captureMP4Files); lo++ {
			// originFile := captureMP4Files[lo]
			// fmt.Println("*************DIRECTORY*************")

			editName := strings.SplitAfter(captureMP4Files[lo], "/edits")
			// fmt.Println("orig file : ", editName)

			jpgRoot := stripMP4Files[lo][:] + "/jpg"
			// fmt.Println("jpg root : ", jpgRoot)

			newNameSplit := strings.SplitAfter(editName[1], "/")
			newName := newNameSplit[1][:len(newNameSplit[1])-4]
			// fmt.Println("name : ", newName)

			jpgDir := jpgRoot + "/" + newName
			// fmt.Println("jpg dir : ", jpgDir)

			// frameOutput := jpgDir + "/" + newName + "-frame-%04d.jpg"
			// fmt.Println("frameOutput : ", frameOutput)

			cmd3 := exec.Command("mkdir", jpgDir)
			cmd3.Run()

		}

		//Create a jpg bundle for each mp4 file within individual jpg dirs
		for so := 0; so < len(captureMP4Files); so++ {
			// fmt.Println("*************BUNDLE*************")

			editName := strings.SplitAfter(captureMP4Files[so], "/edits")
			// fmt.Println("orig file : ", editName)

			jpgRoot := stripMP4Files[so][:] + "/jpg"
			// fmt.Println("jpg root : ", jpgRoot)

			newNameSplit := strings.SplitAfter(editName[1], "/")
			newName := newNameSplit[1][:len(newNameSplit[1])-4]
			// fmt.Println("name : ", newName)

			jpgDir := jpgRoot + "/" + newName
			// fmt.Println("jpg dir : ", jpgDir)

			frameOutput := jpgDir + "/" + newName + "-frame-%04d.jpg"
			// fmt.Println("frameOutput : ", frameOutput)

			realFile := editName[0] + editName[1]
			// fmt.Println("real file : ", realFile)

			cmd5 := exec.Command("ffmpeg", "-n", "-i", realFile, frameOutput)
			cmd5.Run()
		}
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

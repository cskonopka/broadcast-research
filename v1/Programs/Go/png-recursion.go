package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
		// "/Volumes/vs07-072018-052019/2018-07-july",
		// "/Volumes/vs07-072018-052019/2018-08-august",
		// "/Volumes/vs07-072018-052019/2018-09-september",
		// "/Volumes/vs07-072018-052019/2018-10-october",
		// "/Volumes/vs07-072018-052019/2018-11-november",
		// "/Volumes/vs07-072018-052019/2018-12-december",
		"/Volumes/vs07-072018-052019/2019-01-january",
		"/Volumes/vs07-072018-052019/2019-02-february",
		"/Volumes/vs07-072018-052019/2019-03-march",
		"/Volumes/vs07-072018-052019/2019-04-april",
		"/Volumes/vs07-072018-052019/2019-05-may",
		"/Volumes/vs07-072018-052019/2019-06-june",
		"/Volumes/vs07-072018-052019/2019-07-july",
	}

	for months := 0; months < len(dir); months++ {
		// Global variables
		// captureMP4Files -> capture all of the .mp4 files in a directory
		// stripMP4Files -> capture all files after removing the .mp4 suffix
		var captureMP4Files, stripMP4Files []string

		// readandremove -> Read directory and more duplicate entries
		readandremove := readDirRmDups(dir[months])

		// dirFiles -> search for the "edits" directory after removing duplicate entries
		dirFiles := searchFiles("edits", readandremove)

		// iterate over the directory
		for file := 0; file < len(dirFiles); file++ {
			// editName -> split the string after "/edits" in the string
			// ex: /Users/csk/Documents/_REPO/1987-06-may/06-01-1987/edits/thedayjohnLostHisEye.mp4
			editName := strings.SplitAfter(dirFiles[file], "/edits")
			// captureMP4Files -> collect split strings with .mp4 extension
			captureMP4Files = append(captureMP4Files, dirFiles[file])
			// stripMP4Files -> remove the .mp4 extension, add to stripMP4Files array
			stripMP4Files = append(stripMP4Files, editName[0][:len(editName[0])-6])
		}

		// slimDirectories -> remove duplicate directories
		// Each directory (day of the month) in the parent month directory can now be accessed individually
		slimDirectories := removeDuplicates(stripMP4Files)

		// Create new png for each day of the month
		for directory := 0; directory < len(slimDirectories); directory++ {
			cmd3 := exec.Command("mkdir", slimDirectories[directory]+"/png")
			cmd3.Run()
		}

		// Create a jpg bundle for each mp4 file within individual jpg dirs
		for so := 0; so < len(captureMP4Files); so++ {
			// splitEditString -> split the incoming string to make the directory and file name accessible
			splitEditString := strings.SplitAfter(captureMP4Files[so], "/edits")

			// pngRoot -> root gif directory for source file
			pngRoot := stripMP4Files[so][:] + "/png"

			// pngFilename -> define the .gif filename with it's directory (pwd)
			pngFilename := pngRoot + splitEditString[1][:len(splitEditString[1])-4] + ".png"

			createPNG(captureMP4Files[so], pngFilename)
		}
	}
}

// createPNG :  .png file from source content
func createPNG(mp4 string, png string) {
	// fmt.Println("-- CREATE PNG -- ", png)
	cmd := exec.Command("ffmpeg", "-y", "-ss", "0", "-t", "11", "-i", mp4, "-filter_complex", "[0:v] palettegen", png)
	cmd.Run()
}

// RemoveDuplicates : Remove duplicates from a []string
func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

func readDirRmDups(dir string) []string {
	var collectFiles []string

	// for looper := 0; looper < len(dir); looper++ {
	fileList := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	for _, file := range fileList {
		collectFiles = append(collectFiles, file)
	}

	rmDups := removeDuplicates(collectFiles)

	return rmDups
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

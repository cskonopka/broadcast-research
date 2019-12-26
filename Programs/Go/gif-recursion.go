package main

import (
	"os/exec"
	"strings"

	"github.com/cskonopka/gomu"
)

func main() {
	// source month directory
	dir := []string{
		"/Volumes/vs07-072018-052019/2018-07-july",
		"/Volumes/vs07-072018-052019/2018-07-july",
		"/Volumes/vs07-072018-052019/2018-08-august",
		"/Volumes/vs07-072018-052019/2018-09-september",
		"/Volumes/vs07-072018-052019/2018-10-october",
		"/Volumes/vs07-072018-052019/2018-11-november",
		"/Volumes/vs07-072018-052019/2018-12-december",
		// "/Volumes/vs07-072018-052019/2019-01-january",
		// "/Volumes/vs07-072018-052019/2019-02-february",
		// "/Volumes/vs07-072018-052019/2019-03-march",
		// "/Volumes/vs07-072018-052019/2019-04-april",
		// "/Volumes/vs07-072018-052019/2019-05-may",
		// "/Volumes/vs07-072018-052019/2019-06-june",
		// "/Volumes/vs07-072018-052019/2019-07-july",
	}

	// iterate over the months
	for months := 0; months < len(dir); months++ {
		// Global variables
		// captureMP4Files -> capture all of the .mp4 files in a directory
		// stripMP4Files -> capture all files after removing the .mp4 suffix
		var captureMP4Files, stripMP4Files []string

		// readandremove -> Read directory and more duplicate entries
		readandremove := gomu.ReadDirRmDups(dir[months])

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
		slimDirectories := gomu.RemoveDuplicates(stripMP4Files)

		// create a new "/gifs" directory in each day of the month directory
		for directory := 0; directory < len(slimDirectories); directory++ {
			cmd3 := exec.Command("mkdir", slimDirectories[directory]+"/gifs")
			cmd3.Run()
		}

		// Create a jpg bundle for each mp4 file within individual jpg dirs
		for so := 0; so < len(captureMP4Files); so++ {
			// splitEditString -> split the incoming string to make the directory and file name accessible
			splitEditString := strings.SplitAfter(captureMP4Files[so], "/edits")

			// tempGifRoot -> define the temporary location of the new .gif in the .png directory of it's source. (necessary since the .png for the new .gif needs to be in the same directory)
			tempGifRoot := stripMP4Files[so][:] + "/png" + splitEditString[1][:len(splitEditString[1])-4]

			// gifRoot -> root gif directory for source file
			gifRoot := stripMP4Files[so][:] + "/gifs"

			// gifFilename -> define the .gif filename with it's directory (pwd)
			gifFilename := gifRoot + splitEditString[1][:len(splitEditString[1])-4] + ".gif"

			// create a new .gif file using the source file and palettegen filter using ffmpeg
			createLowResGIF(captureMP4Files[so], tempGifRoot+".gif")

			// move the .gif to the associated /gif directory of the source file
			moveGIF(tempGifRoot+".gif", gifFilename)
		}
	}
}

// createLowResGIF -> Generate a new .gif from a video source using ffmpeg
func createLowResGIF(videoSource string, exportedGIF string) {
	cmd := exec.Command("ffmpeg", "-ss", "0", "-t", "11", "-i", videoSource, "-filter_complex", "[0:v] fps=24,scale=w=480:h=-1,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1", exportedGIF)
	cmd.Run()
}

// moveGif -> Move .gif file from /png directory of video source to designated video source /gif directory
func moveGIF(source string, destination string) {
	cmd := exec.Command("mv", source, destination)
	cmd.Run()
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

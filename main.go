// Go program to illustrate the
// concept of multiple interfaces
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type BroadcastResearch interface {
	DefineTargets()
}

type operations struct {
	title      string
	directory  string
	outputtype string
	optype     string
	jpg        struct {
		inputfile       string
		outputdirectory string
	}
	png struct {
		inputfile  string
		outputfile string
	}
	gif struct {
		inputfile  string
		outputfile string
	}
	histo struct {
		inputfile  string
		outputfile string
	}
}

func (o operations) DefineTargets() {
	// fmt.Println(o.directory)
	var collectFiles []string
	var jj []string
	var result []string

	fileList := []string{}
	filepath.Walk(o.directory, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	for _, file := range fileList {
		collectFiles = append(collectFiles, file)
	}

	dirFiles := searchFiles("edits", collectFiles)

	for stuff := 0; stuff < len(dirFiles); stuff++ {
		editName := strings.SplitAfter(dirFiles[stuff], "/edits/")
		// fmt.Println(editName)
		jj = append(jj, editName[0][:len(editName[0])-7])
		if stuff == len(dirFiles)-1 {
			encountered := map[string]bool{}
			for v := range jj {
				if encountered[jj[v]] == true {
				} else {
					encountered[jj[v]] = true
					result = append(result, jj[v])
					fmt.Println("created")
					// Create directories here
					fmt.Println(jj[v])

					cmd3 := exec.Command("mkdir", jj[v]+o.outputtype)
					cmd3.Run()
				}
			}
		}
		defer generateContent(o, dirFiles[stuff], editName)
	}
	collectFiles, jj, result = nil, nil, nil
}

func generateContent(o operations, dirFiles string, editName []string) {
	switch o.optype {
	case "png":
		fmt.Println("create pngs")
		mp4File := dirFiles
		pngFile := editName[0][:len(editName[0])-7] + "/png/" + editName[1][:len(editName[1])-4] + ".png"
		cmd := exec.Command("ffmpeg", "-y", "-ss", "0", "-t", "11", "-i", mp4File, "-filter_complex", "[0:v] palettegen", pngFile)
		cmd.Run()
	case "gif":
		// create gifs
		fmt.Println("create gifs")
		mp4File := dirFiles
		gifFile := editName[0][:len(editName[0])-7] + "/png/" + editName[1][:len(editName[1])-4] + ".gif"
		cmd := exec.Command("ffmpeg", "-ss", "0", "-t", "11", "-i", mp4File, "-filter_complex", "[0:v] fps=24,scale=w=480:h=-1,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1", gifFile)
		cmd.Run()

		fmt.Println("gif moved to directory")
		destinationDirectory := editName[0][:len(editName[0])-7] + "/gif"
		cmd2 := exec.Command("mv", gifFile, destinationDirectory)
		cmd2.Run()
	case "jpg":
		fmt.Println("create jpgs")
		jpgFolder := editName[0][:len(editName[0])-7] + "/jpg/" + editName[1][:len(editName[1])-4]
		cmd3 := exec.Command("mkdir", jpgFolder)
		cmd3.Run()

		mp4File := dirFiles
		jpgFrames := editName[0][:len(editName[0])-7] + "/jpg/" + editName[1][:len(editName[1])-4] + "/" + editName[1][:len(editName[1])-4] + "-frame-%04d.jpg"
		defer generateJpgs(mp4File, jpgFrames)
		fmt.Println(mp4File, jpgFrames)
	case "histogram1":
		fmt.Println("create histo1")
		histo1Folder := editName[0][:len(editName[0])-7] + "/histogram"
		cmd := exec.Command("mkdir", histo1Folder)
		cmd.Run()

		pngFile := editName[0][:len(editName[0])-7] + "/png/" + editName[1][:len(editName[1])-4] + ".png"
		histo1File := editName[0][:len(editName[0])-7] + "/histogram/" + editName[1][:len(editName[1])-4] + ".txt"
		defer generateHistogram1(pngFile, histo1File)
	}
}

func generateJpgs(mp4File string, jpgFrames string) {
	cmd5 := exec.Command("ffmpeg", "-n", "-i", mp4File, jpgFrames)
	cmd5.Run()
}

func generateHistogram1(pngFile string, histo1File string) {
	cmd := exec.Command("convert", pngFile, "-format", "%c", "histogam:info:", histo1File)
	cmd.Run()
}

func main() {
	var br operations
	globalDir := "/Users/csk/Documents/_REPO/1987-06-may/"

	// PNG
	br.directory = globalDir
	br.outputtype = "/png"
	br.optype = "png"
	var i1 BroadcastResearch = br
	i1.DefineTargets()

	// GIF
	br.directory = globalDir
	br.outputtype = "/gif"
	br.optype = "gif"
	var i2 BroadcastResearch = br
	i2.DefineTargets()

	// JPG
	br.directory = globalDir
	br.outputtype = "/jpg"
	br.optype = "jpg"
	var i3 BroadcastResearch = br
	i3.DefineTargets()

	// HISTOGRAM-1
	br.directory = globalDir
	br.outputtype = "/histogram"
	br.optype = "histogram1"
	var i4 BroadcastResearch = br
	i4.DefineTargets()
}

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

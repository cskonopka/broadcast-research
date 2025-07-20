package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/vansante/go-ffprobe"
)

var (
	fileInfo os.FileInfo
	err      error
)

type Ffprobe struct {
	Format struct {
		Filename       string `json:"filename"`
		NbStreams      int    `json:"nb_streams"`
		NbPrograms     int    `json:"nb_programs"`
		FormatName     string `json:"format_name"`
		FormatLongName string `json:"format_long_name"`
		StartTime      string `json:"start_time"`
		Duration       string `json:"duration"`
		Size           string `json:"size"`
		BitRate        string `json:"bit_rate"`
		ProbeScore     int    `json:"probe_score"`
		Tags           struct {
			MajorBrand       string    `json:"major_brand"`
			MinorVersion     string    `json:"minor_version"`
			CompatibleBrands string    `json:"compatible_brands"`
			CreationTime     time.Time `json:"creation_time"`
		} `json:"tags"`
	} `json:"format"`
}

type BroadcastResearch interface {
	DefineTargets()
}

type operations struct {
	title      string
	directory  string
	outputtype string
	optype     string
	monthlycsv bool
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

// Content Engine
func (o operations) DefineTargets() {
	var collectFiles, jj, result []string

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
					// Create directories

					if o.optype != "ffprobe" {
						fmt.Println("directory created")
						cmd3 := exec.Command("mkdir", jj[v]+o.outputtype)
						cmd3.Run()
					}
				}
			}
		}
		// Defer the creation of the content until all folders have been created
		defer generateContent(o, dirFiles[stuff], editName)
	}
	// Clear buffers
	collectFiles, jj, result = nil, nil, nil

	if o.monthlycsv != false {
		fmt.Println("generating csv analysis file~~~~~~~")
		generateCSV(dirFiles)
	}
}

func generateCSV(input []string) {
	var matrix [][]string

	for i := 0; i < len(input); i++ {
		// fmt.Println(input[i] + "s\n")

		editName := strings.SplitAfter(input[i], "/edits/")
		// fmt.Println(editName)

		data, err := ffprobe.GetProbeData(input[i], 5000*time.Millisecond)
		if err != nil {
			log.Panicf("Error getting data:  %v", err)
		}

		// MarhsalIndent the incoming data, accessible via buf variable (w/ error handling)
		buf, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Panicf("Error unmarshalling: %v", err)
		}

		// Connect struct to variable
		var probed Ffprobe
		// Unmarshal buffer into variable defined by the Ffprobe type
		if err := json.Unmarshal(buf, &probed); err != nil {
			panic(err)
		}

		// Base filename
		ffprobeFilename := probed.Format.Filename
		// Clean up the filename
		cleanFilename := filepath.Base(ffprobeFilename)
		// Unix date for the current file
		unixdate := string(probed.Format.Tags.CreationTime.Format(time.RFC850))
		// Split the Unix date by a comma
		s := strings.Split(unixdate, ",")
		// Title of file
		title := cleanFilename[:len(cleanFilename)-4]
		// Type of file
		typer := cleanFilename[len(cleanFilename)-4:]
		// Path of file
		path := editName[0]

		jo := strings.SplitAfter(input[i], "/edits/")
		again := jo[0][:len(jo[0])-7]

		splitagain := strings.SplitAfter(again, "/")
		again2 := splitagain[len(splitagain)-1]
		jj := strings.SplitAfter(again2, "-")

		// Folder month
		folderMonth := jj[0][:len(jj[0])-1]
		// Folder day
		folderDay := jj[1][:len(jj[1])-1]
		// Folder year
		folderYear := jj[2][:len(jj[2])]
		// Edit Month
		editMonth := s[1][4:7]
		// Edit date
		editDate := s[1][1:11]
		// Edit day (i.e. Monday)
		editDay := s[0]
		// Edit year
		editYear := "20" + s[1][8:11]
		// Timestamp of current file
		timestamp := s[1][11:19]
		// Location of the current file
		loc := s[1][20:23]

		matrix = append(matrix, []string{
			title,
			folderMonth,
			folderDay,
			folderYear,
			editMonth,
			editDay,
			editYear,
			editDate[:2],
			typer,
			path,
			timestamp,
			loc,
			probed.Format.Duration,
			probed.Format.Size,
			probed.Format.BitRate,
			probed.Format.FormatName,
			probed.Format.FormatLongName})
	}

	targetDirectory := strings.SplitAfter(input[0], "/edits/")
	fmt.Println(targetDirectory[0][:len(targetDirectory[0])-17] + "\n")

	again := strings.SplitAfter(targetDirectory[0][:len(targetDirectory[0])-17], "-")
	fmt.Println(again)

	year := again[0][len(again[0])-5:]
	month := again[1]
	day := again[2][:len(again[2])-1]

	fmt.Println(year + month + day)
	combine := year + month + day

	root := targetDirectory[0][:len(targetDirectory[0])-18]
	csvFile := root + "/" + combine + "-FFProbeAnalysis-V4.csv"
	fmt.Println(csvFile)
	file, err := os.Create(csvFile)

	if err != nil {
		log.Fatal("FAILED TO CREATE CSV", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Iterate over the FFmpeg matrix
	for _, value := range matrix {
		// Write data to new .csv file
		err := writer.Write(value)
		if err != nil {
			log.Fatal("FAILED TO WRITE CSV", err)
		}
	}
	fmt.Println("done!")
}

// generateContent
// Define an operation using "case" and pass files from the directory
func generateContent(o operations, dirFiles string, editName []string) {
	switch o.optype {
	case "png": // create pngs
		mp4File := dirFiles
		pngFile := editName[0][:len(editName[0])-7] + "/png/" + editName[1][:len(editName[1])-4] + ".png"
		cmd := exec.Command("ffmpeg", "-y", "-ss", "0", "-t", "11", "-i", mp4File, "-filter_complex", "[0:v] palettegen", pngFile)
		cmd.Run()
	case "gif": // create gifs
		mp4File := dirFiles
		gifFile := editName[0][:len(editName[0])-7] + "/png/" + editName[1][:len(editName[1])-4] + ".gif"
		cmd := exec.Command("ffmpeg", "-ss", "0", "-t", "11", "-i", mp4File, "-filter_complex", "[0:v] fps=24,scale=w=480:h=-1,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1", gifFile)
		cmd.Run()

		// Move .gif to gif directory
		destinationDirectory := editName[0][:len(editName[0])-7] + "/gif"
		cmd2 := exec.Command("mv", gifFile, destinationDirectory)
		cmd2.Run()
	case "jpg": // create jpgs
		fmt.Println("create jpgs")
		jpgFolder := editName[0][:len(editName[0])-7] + "/jpg/" + editName[1][:len(editName[1])-4]
		cmd3 := exec.Command("mkdir", jpgFolder)
		cmd3.Run()

		mp4File := dirFiles
		jpgFrames := editName[0][:len(editName[0])-7] + "/jpg/" + editName[1][:len(editName[1])-4] + "/" + editName[1][:len(editName[1])-4] + "-frame-%04d.jpg"
		defer generateJpgs(mp4File, jpgFrames)
	case "histogram1": // create histogram
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
	globalDir := "/Users/csk/Documents/_REPO/1987-06-may/"
	runSuite(globalDir)
}

func runSuite(globalDir string) {
	var br operations
	// // PNG test
	// br.directory = globalDir
	// br.outputtype = "/png"
	// br.optype = "png"
	// var i1 BroadcastResearch = br
	// i1.DefineTargets()

	// // GIF test
	// br.directory = globalDir
	// br.outputtype = "/gif"
	// br.optype = "gif"
	// var i2 BroadcastResearch = br
	// i2.DefineTargets()

	// // JPG test
	// br.directory = globalDir
	// br.outputtype = "/jpg"
	// br.optype = "jpg"
	// var i3 BroadcastResearch = br
	// i3.DefineTargets()

	// // HISTOGRAM-1 test
	// br.directory = globalDir
	// br.outputtype = "/histogram"
	// br.optype = "histogram1"
	// var i4 BroadcastResearch = br
	// i4.DefineTargets()

	// PNG test
	br.directory = globalDir
	br.outputtype = "/ffprobe"
	br.optype = "ffprobe"
	br.monthlycsv = true
	var i5 BroadcastResearch = br
	i5.DefineTargets()
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

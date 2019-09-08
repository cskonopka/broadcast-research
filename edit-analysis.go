package main

import (
	"os"
	"strings"
	"time"

	"github.com/cskonopka/gomu"
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

func main() {
	dir := []string{
		"/Users/io/Desktop/2000-01-january",
	}

	for k := 0; k < len(dir); k++ {
		scanEdits(dir[k])
	}

}

func scanEdits(dir string) {
	var folderdates []string
	var files []string

	a, b := gomu.CrawlAndCollect(dir, "/edits")

	for m := 0; m < len(a); m++ {
		collectedFiles := b[m]
		// fmt.Println("collected files per day : ", collectedFiles)
		// fmt.Println(len(collectedFiles))
		result := strings.Split(collectedFiles, "edits/")
		// fmt.Println(result[1])
		files = append(files, result[1])
		folderdates = append(folderdates, a[m][len(a[m])-16:len(a[m])-6])
	}
	// fmt.Println(files, folderdates)

	probedata := gomu.ProbeFiles(dir, files, folderdates)
	// fmt.Println(probedata)
	csvFile := dir + "/" + dir[27:len(dir)] + "-EDITS.csv"
	gomu.ExportCSV(probedata, csvFile)
	// fmt.Println("done!")
}

// func scanRaw(dir string) {
// 	var folderdates []string
// 	var files []string

// 	a, b := gomu.CrawlAndCollect(dir, "/raw")

// 	for m := 0; m < len(a); m++ {
// 		collectedFiles := b[m]
// 		result := strings.Split(collectedFiles, "raw/")
// 		// fmt.Println(result[1])
// 		files = append(files, result[1])
// 		folderdates = append(folderdates, a[m][len(a[m])-14:len(a[m])-6])
// 	}
// 	fmt.Println(files, folderdates)

// 	probedata := gomu.ProbeFiles(dir, files, folderdates)
// 	fmt.Println(probedata)
// 	csvFile := dir + "/" + dir[28:len(dir)] + "-RAW.csv"
// 	gomu.ExportCSV(probedata, csvFile)
// 	fmt.Println("done!")
// }

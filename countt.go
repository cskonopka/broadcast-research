package main

import (
	"fmt"
)

type Countt struct{}

func (operation Countt) ScanDirs(bc *BC, message string) error {
	_, err := fmt.Printf("Posting to Twitter \n ---> \n %s\n", bc.InputDirectory, message)

	// var collectFiles []string

	// fileList := []string{}
	// filepath.Walk(bc.InputDirectory, func(path string, f os.FileInfo, err error) error {
	// 	fileList = append(fileList, path)
	// 	return nil
	// })

	// for _, file := range fileList {
	// 	collectFiles = append(collectFiles, file)
	// }

	// dirFiles := searchFiles("edits", collectFiles)
	// var jj []string
	// var result []string
	// for stuff := 0; stuff < len(dirFiles); stuff++ {
	// 	editName := strings.SplitAfter(dirFiles[stuff], "/edits/")
	// 	jj = append(jj, editName[0][:len(editName[0])-7])
	// 	if stuff == len(dirFiles)-1 {
	// 		encountered := map[string]bool{}
	// 		for v := range jj {
	// 			if encountered[jj[v]] == true {
	// 			} else {
	// 				encountered[jj[v]] = true
	// 				result = append(result, jj[v])
	// 			}
	// 		}
	// 		// fmt.Println(result)
	// 	}
	// }
	// fmt.Println(result)

	return err
}

// func searchFiles(typer string, rmDups []string) []string {
// 	var savePng []string

// 	// Receive a type
// 	switch typer {
// 	case "gifs":
// 		// iterate over the incoming array
// 		i := 0
// 		for i < len(rmDups)-1 {
// 			i++
// 			// If the string contains the incoming type of directory
// 			h := strings.Contains(rmDups[i], typer+"/")
// 			switch h {
// 			case true:
// 				// If the string does not contain .DS_Store
// 				if !strings.Contains(rmDups[i], ".DS_Store") {
// 					// append to savePng string array
// 					savePng = append(savePng, rmDups[i])
// 				}
// 			}
// 		}
// 	case "edits":
// 		// iterate over the incoming array
// 		i := 0
// 		for i < len(rmDups)-1 {
// 			i++
// 			// If the string contains the incoming type of directory
// 			h := strings.Contains(rmDups[i], typer+"/")
// 			switch h {
// 			case true:
// 				// If the string does not contain .DS_Store
// 				if !strings.Contains(rmDups[i], ".DS_Store") {
// 					// append to savePng string array
// 					savePng = append(savePng, rmDups[i])
// 				}
// 			}
// 		}
// 	case "png":
// 		// iterate over the incoming array
// 		i := 0
// 		for i < len(rmDups)-1 {
// 			i++
// 			// If the string contains the incoming type of directory
// 			h := strings.Contains(rmDups[i], typer+"/")
// 			switch h {
// 			case true:
// 				// If the string does not contain .DS_Store
// 				if !strings.Contains(rmDups[i], ".DS_Store") {
// 					// append to savePng string array
// 					savePng = append(savePng, rmDups[i])
// 				}
// 			}
// 		}
// 	case "raw":
// 		// iterate over the incoming array
// 		i := 0
// 		for i < len(rmDups)-1 {
// 			i++
// 			// If the string contains the incoming type of directory
// 			h := strings.Contains(rmDups[i], typer+"/")
// 			switch h {
// 			case true:
// 				// If the string does not contain .DS_Store
// 				if !strings.Contains(rmDups[i], ".DS_Store") {
// 					// append to savePng string array
// 					savePng = append(savePng, rmDups[i])
// 				}
// 			}
// 		}
// 	}
// 	return savePng
// }

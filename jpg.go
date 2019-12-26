package main

import (
	"fmt"
	"os/exec"
)

type Jpg struct{}

func (operation Jpg) Create(bc *BC, message string) error {
	_, err := fmt.Printf("Posting to Twitter \n ---> \n %s\n %s\n %s\n", bc.Source, bc.OutputDirectory, message)
	cmd5 := exec.Command("ffmpeg", "-n", "-i", bc.Source, bc.OutputDirectory)
	cmd5.Run()
	return err
}

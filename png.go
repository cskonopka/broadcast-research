package main

import (
	"fmt"
	"os/exec"
)

type Png struct{}

func (operation Png) Create(bc *BC, message string) error {
	_, err := fmt.Printf("Posting to Twitter \n ---> \n %s\n %s\n %s\n", bc.Source, bc.Destination, message)
	cmd := exec.Command("ffmpeg", "-y", "-ss", "0", "-t", "11", "-i", bc.Source, "-filter_complex", "[0:v] palettegen", bc.Destination)
	cmd.Run()
	return err
}

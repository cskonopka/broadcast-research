package main

import (
	"fmt"
	"os/exec"
)

type LowResGif struct{}

func (operation LowResGif) Create(bc *BC, message string) error {
	_, err := fmt.Printf("Posting to Twitter \n ---> \n %s\n %s\n %s\n", bc.Source, bc.Destination, message)
	cmd := exec.Command("ffmpeg", "-ss", "0", "-t", "11", "-i", bc.Source, "-filter_complex", "[0:v] fps=24,scale=w=480:h=-1,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1", bc.Destination)
	cmd.Run()
	return err
}

# broadcast-research
## Programs

The *Programs* folder contains Go and Bash programs for creating various datapoints using FFprobe, FFmpeg and ImageMagick.

*bash-createpngcreategif.sh* : Recursively create *.png* analysis and generate a new *.gif* file. When the program starts, provide a directory of video files.
*gif-recursion.go* : Recursively create *.gif* files using the *.png* files as a reference.
*jpg-recursion.go* : Recursively extract *.jpg* frames into individuals edit folders.
*png-histogram-recursion.go* : Recursively create *histograms* for each video file based on *.png* files.
*png-recursion.go* : Recursively create *.png* analysis files using FFmpeg. 
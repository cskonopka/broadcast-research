#!/bin/bash
echo "source file please"
read sourcefile

# create an analsis .png file [palettegen]
ffmpeg -y -ss 0 -t 11 -i $sourcefile -filter_complex "[0:v] palettegen" "${sourcefile%.*}.png"
# generate a new .gif file
ffmpeg -ss 0 -t 11 -i $sourcefile -filter_complex "[0:v] fps=24,scale=w=480:h=-1,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1" "${sourcefile%.*}.gif"

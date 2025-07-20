#!/bin/bash
echo "Provide a directory please :"
read dir
# for f in $dir/*.mp4
for f in $dir/*.mov
do
    echo "Processing $f"
    echo "${f%.*}.png"
    ffmpeg -y -ss 0 -t 11 -i $f -filter_complex "[0:v] palettegen" "${f%.*}.png"
    ffmpeg -ss 0 -t 11 -i $f -filter_complex "[0:v] fps=24,scale=w=480:h=-1,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1" "${f%.*}.gif"
done
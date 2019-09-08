# broadcast-research
The *broadcast-research* project is a collection of data-mining programs written in Go meant to collect various forms of data about my personal analog video synthesis catalog spanning from 2015 - present.

## Scenario
When I started to explore analog video synthesis, I dedcided I wanted to archive all of my work in managable way for my future self and for potential ML projects. As the videos piled up, 

## Programs

### *png-recursion.go*
Create a .png file from an edit (.mp4 source) using FFmpeg and the palettegen filter.

<p align="center">
  <img width="80%" height="80%" src="https://i.ibb.co/p0zLPmC/demo-png-recursion.gif"/>
</p>

### *gif-recursion.go*
Use the .png file and the .mp4 source to create a new .gif file using FFmpeg.
<p align="center">
  <img width="80%" height="80%" src="https://i.ibb.co/qx8HYZY/demo-gif-recursion.gif"/>
</p>

### *jpg-recursion.go*
Extract individual frames as .jpg files from the .mp4 source using FFmpeg.
<p align="center">
  <img width="80%" height="80%" src="https://i.ibb.co/Mn0PkN0/demo-jpg-recursion.gif"/>
</p>

Example output of frame extraction to dedicated jpg directory.
<p align="center">
  <img width="80%" height="80%" src="https://i.ibb.co/zhxGL42/demo-jpg-output.png
"/>
</p>

### *histogram-recursion.go*
Using a .png file as the source, generate a .txt histogram file using ImageMagick
<p align="center">
  <img width="80%" height="80%" src="https://i.ibb.co/HqyXwTJ/demo-histogram-recursion.gif"/>
</p>

Example histogram .txt file
<p align="center">
  <img width="80%" height="80%" src="https://i.ibb.co/Hnbwr9T/demo-histogram-txt.png"/>
</p>

## Types of Data

### FFmpeg

### ImageMagick

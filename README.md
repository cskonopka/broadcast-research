# broadcast-research
The *broadcast-research* project is a collection of data-mining programs written in Go meant to collect various forms of data about my personal analog video synthesis catalog spanning from 2015 - present.

## Background
When I started to explore analog video synthesis, I consciously made the decision that I wanted to record my creative process each time I switched on. The next step was to create a folder structure that streamlined the daily archival process and made it easy to navigate the collection in the future. Below is an example of the folder structure I've been using since 2015.

```
year-#OfMonth-monthName
- year-#OfMonth-Date
```
Example
<p align="center">
  <img width="65%" height="65%" src="https://i.ibb.co/XYZ7Gy1/demo-folderstructure.png"/>
</p>

## Intention
Using the analog video synthesis collection, 
Design data-mining programs that analyze video files using FFprobe and image files using ImageMagick, outputting .csv files as open-source content. 

## Ideas
- Apply machine learning to source content and perform predictive modeling about my creative process.
- Develop generative data art projects based on source content using recursion.
- Learn about color
- Under


Once , effectively developing a machine learning feedback loop. The output 

## Program Types
1. Generation
2. Analysis

### Generation Programs

#### *png-recursion.go*
Create a .png file from an edit (.mp4 source) using FFmpeg and the palettegen filter.

<p align="center">
  <img width="80%" height="80%" src="https://i.ibb.co/p0zLPmC/demo-png-recursion.gif"/>
</p>

#### *gif-recursion.go*
Use the .png file and the .mp4 source to create a new .gif file using FFmpeg.

<p align="center">
  <img width="80%" height="80%" src="https://i.ibb.co/qx8HYZY/demo-gif-recursion.gif"/>
</p>

#### *jpg-recursion.go*
Extract individual frames as .jpg files from the .mp4 source using FFmpeg.
<p align="center">
  <img width="80%" height="80%" src="https://i.ibb.co/Mn0PkN0/demo-jpg-recursion.gif"/>
</p>

Example output of frame extraction to dedicated jpg directory.
<p align="center">
  <img width="80%" height="80%" src="https://i.ibb.co/zhxGL42/demo-jpg-output.png
"/>
</p>

### Analysis Programs

#### *histogram-recursion.go*
Using a .png file as the source, generate a .txt histogram file using ImageMagick
<p align="center">
  <img width="80%" height="80%" src="https://i.ibb.co/HqyXwTJ/demo-histogram-recursion.gif"/>
</p>

Example histogram .txt file
<p align="center">
  <img width="80%" height="80%" src="https://i.ibb.co/Hnbwr9T/demo-histogram-txt.png"/>
</p>







## Types of Data

### FFprobe Data
```
Data collected using
Filename	
FolderDate	
Folder Month	
Folder Day	
Folder 
Year	
Edit Date	
Edit Month	
Edit Day	
Edit Year	
Edit Day Number	
Timestamp	
Timezone	
Duration	
Size	
Bitrate	Format	
Formant Long
```

### ImageMagick
```
Image
Format
Mime type
Class
Geometry
Resolution
Print size
Units
Colorspace
Type
Base Type
Endianess
Depth
Red
Green
Blue
Alpha
Pixels
red min
red max
red mean
red standard deviation
red kurtosis
red skewness
red entropy
green min
green max
green mean
green standard deviation
green kurtosis
green skewness
green entropy
blue min
blue max
blue mean
blue standard deviation
blue kurtosis
blue skewness
blue entropy
alpha min
alpha max
alpha mean
alpha standard deviation
alpha kurtosis
alpha skewness
alpha entropy
imagstats min
imagstats max
imagstats mean
imagstats standard deviation
imagstats kurtosis
imagstats skewness
imagstats entropy
Colors
Rendering intent
Gamma
chromaticity red primary
chromaticity green primary
chromaticity blue primary
chromaticity white point
Matte color
Background color
Border color
Transparent color
Interlace
Intensity
Compose
Page geometry
Dispose
Iterations
Compression
Orientation
Prop date create
Prop date modify
png:IHDR.bit-depth-orig
png:IHDR.bit_depth
png:IHDR.color-type-orig
png:IHDR.color_type
png:IHDR.interlace_method
png:IHDR.width,height
png:pHYs
png:sRGB
Prop signature
Artifacts verbose
Tainted
Filesize
Number pixels
Pixels per second
User time
Elapsed time
Version
```

## Programs


# broadcast-research
*broadcast-research* is data art project that explores the potential of analog video synthesis as a generative feedback tool for creating artistic works. Written in Go, the project seeks to use data-mining, machine learning and NLP to gain new insights about the catalog and learn about interlaced subconscious psychological trends. 

<p align="center">
  <img width="65%" height="65%" src="https://i.ibb.co/WkrxHZj/cskonopka-04.png"/>  
</p>

A detailed overview of the project can be found in the [wiki](https://github.com/cskonopka/broadcast-research/wiki).

## Requirements
* Go
* FFprobe
* FFmpeg
* ImageMagick

*** 
## Project Background
When I started using analog video synthesizers in 2015, I actively decided I wanted to record video content each time I sat down for a studio session or live performance. Each piece is a first-hand documentation entry that provides insight about subconscious trends and conscious trends. Documenting daily experiences evolve into a library of emotions that extend the potential of video art as a psychological research topic. Below is an example of the folder structure I've been using since 2015.


### *Framework v1*
* **edits**: Edited source content for public consumption
* **raw**: Raw source content

``` 
year-#OfMonth-monthName
-- year-#OfMonth-Date
---- edits
---- raw
```

*in reality*
``` 
2015-04-april
-- 2015-04-01
---- edits
---- raw
```

Using tools I've created for this project, I've expanded the scope of the initial framework to include png, jpg, histograms and analysis files.

### *Framework v2*
* **year-#OfMonth-monthName-FFProbeAnalysis-V4.csv**: Monthly content analysis document containing the output of FFprobe for each video in the *edits* directory.
* **year-#OfMonth-monthName-png-imagemagick.csv**: Monthly content analysis document containing the output of ImageMagick *identify* for each video in the *edits* directory.
* **gifs**:  Gifs generated using FFmpeg from content in *edits* and using *.png* files as a reference.
* **histo**: Histograms generated using ImageMagick from the *.png* reference files.
* **jpg**: Extracted *.jpg* frames from content in *edits*
* **png**: Pngs generated using FFmpeg and the filter *palettegen* used to generated gifs.
* **stills**: Stills manually extracted from *raw* files
``` 
year-#OfMonth-monthName
-- year-#OfMonth-monthName-FFProbeAnalysis-V4.csv
-- year-#OfMonth-monthName-png-imagemagick.csv
-- year-#OfMonth-Date
---- edits
---- gifs
---- histo
---- jpg
---- png
---- raw
---- stills
```

*in reality*

``` 
2015-04-april
-- 2015-04-01
---- edits
---- gifs
---- histo
---- jpg
---- png
---- raw
---- stills
-- 2015-04-april-FFProbeAnalysis-V4.csv
-- 2015-04-april-png-imagemagick.csv
```

*** 
## Intention
Design data-mining programs intended to generate content and analysis files using source content from the edits folder from 2015 - present. Source content (.mp4) information is acquired using FFprobe and ImagickMagick's identify program and saved as monthly .csv files. Then, use cat to concatenate all of the .csv files into a singular master file.

*Note: CSV files have no headers. The data headers for .csv files are located in the README.md file of the specific data set*

## Data 

### FFprobe
Filename | FolderDate | Folder Month | Folder Day | Folder | Year | Edit Date | Edit Month | Edit Day | Edit Year | Edit Day Number | Timestamp | Timezone | Duration | Size | Bitrate Format | Formant Long
--- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | ---

### ImageMagick
Image | Format | Mime type | Class | Geometry | Resolution | Print size | Units | Colorspace | Type | Base Type | Endianess | Depth | Red | Green | Blue | Alpha | Pixels | red min | red max | red mean | red standard deviation | red kurtosis | red skewness | red entropy | green min | green max | green mean | green standard deviation | green kurtosis | green skewness | green entropy | blue min | blue max | blue mean | blue standard deviation | blue kurtosis | blue skewness | blue entropy | alpha min | alpha max | alpha mean | alpha standard deviation | alpha kurtosis | alpha skewness | alpha entropy | imagstats min | imagstats max | imagstats mean | imagstats standard deviation | imagstats kurtosis | imagstats skewness | imagstats entropy | Colors | Rendering intent | Gamma | chromaticity red primary | chromaticity green primary | chromaticity blue primary | chromaticity white point | Matte color | Background color | Border color | Transparent color | Interlace | Intensity | Compose | Page geometry | Dispose | Iterations | Compression | Orientation | Prop date create | Prop date modify | png:IHDR.bit-depth-orig | png:IHDR.bit_depth | png:IHDR.color-type-orig | png:IHDR.color_type | png:IHDR.interlace_method | png:IHDR.width,height | png:pHYs | png:sRGB | Prop signature | Artifacts verbose | Tainted | Filesize | Number pixels | Pixels per second | User time | Elapsed time | Version
--- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- 

### Titles
Title-raw | Title-spaced
--- | --- |

*** 
## Program Types
1. Generation -> *Programs used to generate new content sources (i.e. .png, .jpg, .gif)*
2. Analysis -> *Programs that use FFprobe or ImageMagick to create analysis files (.csv) based various source content.

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
  <img width="80%" height="80%" src="https://i.ibb.co/zhxGL42/demo-jpg-output.png"/>
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


*** 
# About

## Who am I?
My name is Christopher Konopka and I'm a video artist with a focus on modular analog video synthesizers. 

## Artist Statement
Christopher Konopka specializes in painting one-of-a-kind visual textures with analog video synthesizers. Evolving organically, each design is an encapsulated moment of experience used to terraform new planes of understanding. Humans store the experience of time in their own unique format by generating a memory and compiling all the internal and external senses. During the archival process of creating new memories, there is a significant amount of discarded nuance. These free-floating threads of minutia hold the potential choice of mindfully engaging unaltered emotional abstractions through creating personal coincidental attachment.

## What analog video synthesizers do you use?
* LZX Industries
* Brownshoesonly
* BPMC
* Dave Jones Design


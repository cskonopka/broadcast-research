#!/bin/bash
#!/bin/bash
echo input source?
read input

echo output location?
read output
echo "${output%%.*}"
newoutput=${output%%.*}

ffmpeg -n -i $input "${newoutput-frame-%04d.jpg}"
#!/bin/bash
echo input png?
read input

echo output histogram location?
read output

convert $input -format %c histogram:info:- > $output
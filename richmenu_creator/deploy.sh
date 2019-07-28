#!/usr/bin/env bash

source tokenloader.sh

go build -o rich_creator *.go

./rich_creator -width=2500 \
    -height=1686 \
    -selected=true \
    -name=Home \
    -chatBarText=Home \
    -areaFile=$(pwd)/areas/Home.json \
    -imageFile=$(pwd)/images/Home.png
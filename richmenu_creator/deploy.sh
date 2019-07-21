#!/usr/bin/env bash

source tokenloader.sh

./rich_creator -width 2500 \
    -height=1686 \
    -selected=true \
    -name=Home \
    -chatBarText=Home \
    -area=$(pwd)/areas/Home.json \
    -image=$(pwd)/images/Home.png
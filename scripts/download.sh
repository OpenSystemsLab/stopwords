#!/bin/bash

set -e

WORK_DIR=./build

rm -rf $WORK_DIR

mkdir -p $WORK_DIR

# Download the data
wget -O $WORK_DIR/data.zip https://github.com/natural-language-processing/stopwords/archive/refs/heads/main.zip

unzip $WORK_DIR/data.zip -d $WORK_DIR
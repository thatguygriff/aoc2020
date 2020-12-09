#!/bin/bash

mkdir $1

echo "package $1" >> $1/day_$1.go
echo "package $1" >> $1/day_$1_test.go
touch $1/sample.txt
touch $1/input.txt
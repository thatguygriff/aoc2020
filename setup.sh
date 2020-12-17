#!/bin/bash

mkdir $1

echo "package seventeen

import (
	\"bufio\"
	\"os\"
)

func load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	}

	return nil
}
" >> $1/day_$1.go
echo "package $1" >> $1/day_$1_test.go
touch $1/sample.txt
touch $1/input.txt
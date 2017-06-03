#!/usr/bin/env bash
if [ "$1" == "" ]; then
	echo "package name required!"
elif [ "$2" == "" ]; then
	go test github.com/Ganelon13/go-algorithms/"$1" -run=- -bench=. -benchmem > "$1"/benchmarks.txt
else
    go test github.com/Ganelon13/go-algorithms/"$1" -run=- -bench=. -benchmem
fi
#!/bin/bash
#author: Tim Sabsch <tim@sabsch.com>

> $1/output.out
find $1 -type f -name "*.in" | sort | while read inp; do
    printf "Output of $inp:\n" >> $1/output.out
    go run $1/solution.go $inp >> $1/output.out
    printf "\n\n" >> $1/output.out
done


#!/bin/bash
# author: Tim Sabsch <tim@sabsch.com>

find $1 -type f -name "*.txt" | sort | while read txt; do
    go run $1/solution.go $txt >> $1/output.out;
    sed -i -e '$a\' $1/output.out
done


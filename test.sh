#!/usr/bin/env bash

set -e

REPORT_COVERAGE="coverage.txt"
REPORT_PROFILE="profile.out"

echo "" > $REPORT_COVERAGE

for dir in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=$REPORT_PROFILE -covermode=atomic "$dir"
    if [ -f $REPORT_PROFILE ]; then
        cat $REPORT_PROFILE >> $REPORT_COVERAGE
        rm $REPORT_PROFILE
    fi
done
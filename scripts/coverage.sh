#!/usr/bin/env bash

COV_THRESHOLD: 0

echo "Running coverage tests..."
totalCoverage=`go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+'`
echo "Current test coverage : $totalCoverage %"
if (( $(echo "$totalCoverage $COV_THRESHOLD" | awk '{print ($1 > $2)}') )); then
    echo "OK"
else
    echo "Failed"
    exit 1
fi
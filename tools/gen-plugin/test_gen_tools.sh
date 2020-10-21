#!/usr/bin/env bash
set -e

dir=`pwd`

cd $dir
rm -rf ./staking
go build -o gen
./gen staking
if [ $? -ne 0 ]; then
  echo "Failed: all"
  exit 1
else
  rm gen
  rm -rf ./staking
fi
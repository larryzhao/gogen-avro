#!/bin/bash

# Rewrite references from github.com/larryzhao/gogen-avro to gopkg.in/alanctgardner/gogen-avro.<version>

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <version>"
  exit 1
fi
 
GITHUB_REPO="github.com/larryzhao/gogen-avro"
VERSION="$1"
GOPKG_REPO="gopkg.in/actgardner/gogen-avro.$VERSION"

sed -i "s|$GITHUB_REPO|$GOPKG_REPO|" container/*.go generator/*.go types/*.go gogen-avro/main.go example/*/*.go test.sh 

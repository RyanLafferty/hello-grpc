#!/bin/bash

if [[ $(command -v brew) == "" ]]; then
    echo "Please Install Hombrew"
fi

brew install go

echo 'Please ensure that GOPATH and GOBIN are set and added to the path'
echo 'export GOPATH=/Users/$USER/go'
echo 'export GOBIN=$GOPATH/bin'
echo 'export PATH=$PATH:$GOBIN'

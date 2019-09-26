#!/bin/bash

if [[ $(command -v brew) == "" ]]; then
    echo "Please Install Hombrew"
fi

brew install protobuf

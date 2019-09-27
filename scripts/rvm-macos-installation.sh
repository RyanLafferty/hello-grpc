#!/bin/bash

if [[ $(command -v curl) == "" ]]; then
    echo "Please install curl"
fi

curl -sSL https://get.rvm.io | bash -s stable --ruby

echo 'Please ensure that RVM is added to the path'
echo 'export PATH=$PATH:$HOME/.rvm/bin'

#/bin/bash

if [[ $(command -v gem) == "" ]]; then
    echo "Please Install Ruby"
fi

gem install grpc grpc-tools

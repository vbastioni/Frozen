#!/bin/bash
NAME="frozen_server"
case "$1" in
    "build")
    docker build -t $NAME .
    ;;
    "vm")
    docker run -it --rm -p 3306:3306 --name $NAME $NAME
    ;;
    "go")
    ./dockercommands.sh build && ./dockercommands.sh vm
    ;;
    "eval")
    eval $(docker-machine env Char)
    ;;
    "")
    echo "usage: $0 [build]"
    ;;
    *)
    echo "$1: commands not found" >/dev/stderr
    ;;
esac

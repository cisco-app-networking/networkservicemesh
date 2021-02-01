#!/usr/bin/env bash

#docker build . | tee /dev/stdout | tail -1
#docker build . | tee | tail -1

docker run --rm  -it --net=host -v /var/run/docker.sock:/var/run/docker.sock $(docker build -q .) bash


docker run --rm -it $(docker build -q .) bash

#TODO: remember to remove image for ci
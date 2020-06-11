#!/bin/bash

CURR_DIR=${PWD}

docker \
    build \
        --no-cache \
        -t eldius/game-manager-test:latest \
        . && \
    docker run -it --name game-manager-test --rm eldius/game-manager-test:latest $@

cd $CURR_DIR

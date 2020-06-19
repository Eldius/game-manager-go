#!/bin/bash

## -- header -- ##
eval "$(pyenv init -)"
eval "$(pyenv virtualenv-init -)"

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

CURR_DIR=${PWD}

## -- header -- ##

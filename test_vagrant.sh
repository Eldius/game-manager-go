#!/bin/bash

CURR_DIR=${PWD}

rm ~/dev/workspaces/vagrant/test-box/game-manager
go build -o ~/dev/workspaces/vagrant/test-box/game-manager

cd ~/dev/workspaces/vagrant/test-box/
vagrant destroy -f && vagrant up && vagrant ssh

cd $CURR_DIR

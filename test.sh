#!/bin/bash

CURR_DIR=${PWD}
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
VAGRANT_FOLDER=$SCRIPT_DIR/tests
function handle_error {
    echo $1
    cd $CURR_DIR
    exit 1
}

function destroy_box {
    if [ "${REBUILD}" -ne "0" ];then
        cd $VAGRANT_FOLDER
        echo ""
        echo "****************************"
        echo "* cleaning up previous box *"
        echo "****************************"
        echo ""
        vagrant destroy -f || \
        handle_error "Failed to destroy Box"
        cd $CURR_DIR
    fi
}

function start_box {
    echo ""
    echo "****************************"
    echo "* starting up box          *"
    echo "****************************"
    echo ""

    cd $VAGRANT_FOLDER
    vagrant up || \
        handle_error "Failed to start Box"
    SSH_PORT=$( vagrant ssh-config | grep Port | grep -Eo '[0-9]*' )
    cd $CURR_DIR
}

function manager_setup {

    if [ "${CLEAN_SETUP}" -ne "0" ];then
        echo ""
        echo "****************************"
        echo "* run game-manager setup   *"
        echo "****************************"
        echo ""

        cd $SCRIPT_DIR
        rm -rf ~/.game-manager
        go run main.go \
            setup || \
        handle_error "Failed to clean setup"
    fi
    cd $CURR_DIR
}

function provision_minecraft_server {
    echo ""
    echo "****************************"
    echo "* configuring minecraft    *"
    echo "****************************"
    echo ""

    cd $SCRIPT_DIR
    go run main.go \
        minecraft \
            setup \
                -u vagrant \
                -s 127.0.0.1 \
                -p $SSH_PORT \
                -k ./tests/.vagrant/machines/default/virtualbox/private_key \
                "minecraft_java_rcon_pass=ABC123" \
                "minecraft_java_enable_rcon=true" \
                "app_service_user=minecraft" \
                "broadcast-rcon-to-ops=true" \
                "app_service_user=minecraft" || \
                    handle_error "Failed to provision minecraft"

    cd $CURR_DIR
}

function test_minecraft_server {
    echo ""
    echo "****************************"
    echo "* testing minecraft server *"
    echo "****************************"
    echo ""

    cd $VAGRANT_FOLDER

    vagrant ssh-config > .vagrant/ssh-config && \
        py.test --hosts=default --ssh-config=.vagrant/ssh-config minecraft.py || \
            handle_error "Minecraft server  tests failed..."
    cd $CURR_DIR
}

REBUILD="0"
CLEAN_SETUP="0"

for var in "$@"
do
  case $var in
  "--build")
    REBUILD="1"
    ;;
  "--setup")
    CLEAN_SETUP="1"
    ;;
  esac
done

echo "**************************************************"
echo "* => STARTING TEST <=                            *"
if [ "${CLEAN_SETUP}" -ne "0" ];then
    echo "* CLEAN SETUP: ON                                *"
else
    echo "* CLEAN SETUP: OFF                               *"
fi
if [ "${REBUILD}" -ne "0" ];then
    echo "* REBUILD BOX: ON                                *"
else
    echo "* REBUILD BOX: OFF                               *"
fi
echo "* STARTING TEST                                  *"
echo "**************************************************"

destroy_box
start_box

sleep 10

manager_setup
provision_minecraft_server
sleep 10
test_minecraft_server

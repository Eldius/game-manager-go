# Game Manager #

A simple tool to manage game servers

## Build status ##

![Go](https://github.com/Eldius/game-manager-go/workflows/Go/badge.svg)

## Test snippets ##

```shell
go run main.go minecraft setup -u vagrant -s 127.0.0.1 -p 2222 -k ~/dev/workspaces/vagrant/test-box/.vagrant/machines/default/virtualbox/private_key "minecraft_java_rcon_pass=123" "minecraft_java_enable_rcon=true"
```

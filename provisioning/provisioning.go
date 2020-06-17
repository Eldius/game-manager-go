package provisioning

import (
	"github.com/Eldius/game-manager-go/command"
	"github.com/Eldius/game-manager-go/config"
	"github.com/Eldius/game-manager-go/scripts"
)

/*
Provision configure thea Minecraft game server
*/
func Provision(cfg config.ManagerConfig) {
	engine := scripts.NewScriptEngine(cfg)
	playbook := engine.GetProvisioningScript("minecraft_playbook")
	shell := engine.GetProvisioningScript("minecraft_playbook_run")

	playbook.SaveToFile()
	shell.SaveToFile()

	command.ExecuteScript(shell)
}

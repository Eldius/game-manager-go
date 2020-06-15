package scripts

import "github.com/Eldius/game-manager-go/config"

/*
NewScriptEngine sets up a new scripts.ScriptEngine
*/
func NewScriptEngine(cfg config.ManagerConfig) ScriptEngine {
	return ScriptEngine{
		Cfg: cfg,
	}
}

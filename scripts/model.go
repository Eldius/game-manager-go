package scripts

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Eldius/game-manager-go/config"
	"github.com/Eldius/game-manager-go/logger"
	"github.com/gobuffalo/packr"
)

const (
	scriptHeader = "shell/header.sh"
)

var (
	box    = packr.NewBox("./templates")
	engine = template.New("scripts_parser")
)

const (
	// AnsibleScript is ansible playbook script type
	AnsibleScript ScriptType = "Ansible"
	// ShellScript is shell script type
	ShellScript ScriptType = "Shell"
)

var setupScripts = map[string]map[string]string{
	"install_python_env": {
		"template": "shell/setup/setup_python_environment.sh",
		"type":     string(ShellScript),
	},
	"minecraft_ansible_requirements": {
		"template": "ansible/minecraft/roles/requirements.yml",
		"type":     string(AnsibleScript),
	},
}

var provisioningScripts = map[string]map[string]string{
	"minecraft_playbook": {
		"template": "ansible/minecraft/deploy-minecraft.yml",
		"type":     string(AnsibleScript),
	},
	"minecraft_playbook_run": {
		"template": "shell/minecraft/execute_ansible_playbook.sh",
		"type":     string(ShellScript),
	},
}

/*
ScriptTemplateVars is a representation of
vars to parse script templates
*/
type ScriptTemplateVars struct {
	WorkspacePath    string
	PythonVersion    string
	VenvName         string
	ProvisioningInfo *ServerProvisioningInfo
}

/*
ScriptType defines a type of script
*/
type ScriptType string

/*
ScriptDef represents a model to
render a script
*/
type ScriptDef struct {
	Name       string
	Template   string
	Path       string
	Type       ScriptType
	cfg        config.ManagerConfig
	ParamsInfo *ServerProvisioningInfo
}

/*
ServerProvisioningInfo is a model to
represents provisioning parameters
*/
type ServerProvisioningInfo struct {
	Game       string
	IP         string
	SSHPort    int
	SSHKey     string
	RemoteUser string
	Args       map[string]string
}

/*
NewServerProvisioning creates a server provisioning data
*/
func NewServerProvisioning(game string, ip string, sshPort int, remoteUser string, sshKey string, args []string) *ServerProvisioningInfo {
	return &ServerProvisioningInfo{
		Game:       game,
		IP:         ip,
		SSHPort:    sshPort,
		SSHKey:     sshKey,
		RemoteUser: remoteUser,
		Args:       getArgsMap(args),
	}
}

func getArgsMap(args []string) map[string]string {
	var argsMap = make(map[string]string)
	for _, a := range args {
		tmp := strings.Split(a, "=")
		argsMap[tmp[0]] = tmp[1]
	}

	return argsMap
}

/*
WithParams add the provisioning info
*/
func (s *ScriptDef) WithParams(i *ServerProvisioningInfo) *ScriptDef {
	s.ParamsInfo = i
	return s
}

/*
ScriptConfig returns the execution configuration
*/
func (s *ScriptDef) ScriptConfig() config.ManagerConfig {
	return s.cfg
}

/*
Render renders the script
*/
func (s *ScriptDef) Render() string {
	script := s.loadScriptTemplate(s.Template)

	var tmpl *template.Template
	var err error

	if strings.HasSuffix(s.Path, "sh") {
		header := s.loadScriptTemplate(scriptHeader)
		tmpl, err = engine.Parse(strings.Join([]string{header, script}, "\n#---\n"))
	} else {
		tmpl, err = engine.Parse(script)
	}
	if err != nil {
		log.Panic(err.Error())
	}

	buf := new(bytes.Buffer)
	tmpl.Execute(buf, s.loadTemplateVars())
	return buf.String()

}

func (s *ScriptDef) loadTemplateVars() ScriptTemplateVars {
	return ScriptTemplateVars{
		WorkspacePath:    s.cfg.Workspace,
		PythonVersion:    "3.8.0",
		VenvName:         "game-manager",
		ProvisioningInfo: s.ParamsInfo,
	}
}

func (s *ScriptDef) loadScriptTemplate(path string) string {
	script, err := box.FindString(path)
	if err != nil {
		log.Fatal(err)
	}
	return script
}

/*
SaveToFile saves rendered script to file
*/
func (s *ScriptDef) SaveToFile() {
	fileMode := s.getFileMode()
	scriptDir := filepath.Dir(s.Path)
	os.MkdirAll(scriptDir, os.ModePerm)
	if s.cfg.Verbose {
		log.Printf("\n---\nsaving to file\n- [%s] %s: %d\n---\n", s.Type, s.Template, fileMode)
	}
	ioutil.WriteFile(s.Path, []byte(s.Render()), fileMode)
}

func (s *ScriptDef) getFileMode() os.FileMode {
	if s.Type == ShellScript {
		return os.ModePerm
	}
	return 0666
}

/*
GetScriptExecutionEnvVars generates the env vars to execute
scripts
*/
func (s *ScriptDef) GetScriptExecutionEnvVars() []string {
	sysPath, _ := os.LookupEnv("PATH")
	newPath := fmt.Sprintf("PATH=%s:%s", s.cfg.GetPyenvBinFolder(), sysPath)
	workspace := s.cfg.Workspace
	newUserHome := fmt.Sprintf("HOME=%s", workspace)
	pyenvRoot := fmt.Sprintf("PYENV_ROOT=%s/pyenv", workspace)

	return append(os.Environ(), newPath, newUserHome, pyenvRoot)

}

/*
Execute executes script
*/
func (s *ScriptDef) Execute() {

	s.SaveToFile()

	execArgs := append([]string{s.Path})

	l := logger.NewLogWriter(logger.DefaultLogger())
	cmd := &exec.Cmd{
		Path:   s.Path,
		Args:   execArgs,
		Env:    s.GetScriptExecutionEnvVars(),
		Stdout: l,
		Stderr: l,
	}

	s.executeCmd(cmd)

}

func (s *ScriptDef) executeCmd(cmd *exec.Cmd) {

	if s.cfg.Verbose {
		log.Println("cmd:", cmd.String())
	}

	log.Println()
	log.Println("**********")
	if s.cfg.Verbose {
		log.Println("env vars:\n", cmd.Env)
	}
	if err := cmd.Run(); err != nil {
		log.Println("---")
		log.Println("Failed to execute script", s.Template)
		if s.cfg.Verbose {
			log.Println(s.Render())
		}
		log.Println(err.Error())
		os.Exit(1)
	}
	log.Println("**********")
}

/*
ScriptEngine is the script executor
*/
type ScriptEngine struct {
	Cfg config.ManagerConfig
}

/*
GetSetupScripts returns all the script models
*/
func (s *ScriptEngine) GetSetupScripts() []*ScriptDef {
	var scriptList []*ScriptDef
	for k := range setupScripts {
		scriptList = append(scriptList, s.GetScriptDef(k, s.Cfg, setupScripts))
	}

	return scriptList
}

/*
GetSetupScript returns a single setup script model
*/
func (s *ScriptEngine) GetSetupScript(scriptName string) *ScriptDef {
	return s.GetScriptDef(scriptName, s.Cfg, setupScripts)
}

/*
GetProvisioningScript returns a single provisioning script model
*/
func (s *ScriptEngine) GetProvisioningScript(scriptName string) *ScriptDef {
	return s.GetScriptDef(scriptName, s.Cfg, provisioningScripts)
}

/*
GetScriptDef returns
*/
func (s *ScriptEngine) GetScriptDef(scriptName string, cfg config.ManagerConfig, scriptsDef map[string]map[string]string) *ScriptDef {
	template := scriptsDef[scriptName]["template"]
	return &ScriptDef{
		Name:     scriptName,
		Template: template,
		Path:     filepath.Join(s.Cfg.GetScriptsFolder(), template),
		Type:     ScriptType(scriptsDef[scriptName]["type"]),
		cfg:      cfg,
	}
}

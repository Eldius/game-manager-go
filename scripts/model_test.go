package scripts

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/Eldius/game-manager-go/config"
)

func TestScriptDefGenerationWithHeader(t *testing.T) {
	cfg := config.ManagerConfig{
		Workspace: "/tmp/testing_script_engine",
	}

	engine := ScriptEngine{
		Cfg: cfg,
	}

	setup := engine.GetSetupScript("install_python_env")

	renderedScript := setup.Render()

	r1 := regexp.MustCompile(`(-- header --)`)
	t.Log(r1.FindStringSubmatch(renderedScript))

	if len(r1.FindStringSubmatch(renderedScript)) != 2 {
		t.Errorf("---\nExpected header to have two '-- header --' line\n---\n%s\n---\n", renderedScript)
	}

	r2 := regexp.MustCompile(`python -m pip install ansible`)
	t.Log(r2.FindStringSubmatch(renderedScript))

	if len(r2.FindStringSubmatch(renderedScript)) != 1 {
		t.Errorf("---\nExpected header to have a 'python -m pip install ansible' line\n---\n%s\n---", renderedScript)
	}
}

func TestScriptDefGenerationWithoutHeader(t *testing.T) {
	cfg := config.ManagerConfig{
		Workspace: "/tmp/testing_script_engine_2",
	}

	engine := ScriptEngine{
		Cfg: cfg,
	}

	setup := engine.GetSetupScript("minecraft_ansible_requirements")

	renderedScript := setup.Render()

	r1 := regexp.MustCompile(`(-- header --)`)
	t.Log(r1.FindStringSubmatch(renderedScript))

	if len(r1.FindStringSubmatch(renderedScript)) > 0 {
		t.Errorf("---\nExpected header to have two '-- header --' line\n---\n%s\n---\n", renderedScript)
	}

	r2 := regexp.MustCompile(`src: https://github.com/Eldius/minecraft-java-edition-ansible-role.git`)
	t.Log(r2.FindStringSubmatch(renderedScript))

	if len(r2.FindStringSubmatch(renderedScript)) != 1 {
		t.Errorf("---\nExpected header to have a 'python -m pip install ansible' line\n---\n%s\n---", renderedScript)
	}
}

func TestScriptExecutionEnvVars(t *testing.T) {
	home := "/tmp/my_test_home"

	workspace := filepath.Join(home, "workspace")
	cfg := config.ManagerConfig{
		Workspace: workspace,
	}

	engine := NewScriptEngine(cfg)

	s := engine.GetSetupScript("install_python_env")

	fmt.Println("cfg:", s.cfg)
	env := s.GetScriptExecutionEnvVars()

	var pathVars []string
	var homeVars []string

	for _, p := range env {
		if strings.HasPrefix(p, "PATH=") {
			pathVars = append(pathVars, p)
		}
		if strings.HasPrefix(p, "HOME=") {
			homeVars = append(homeVars, p)
		}
	}

	log.Println("path:", pathVars)
	if len(pathVars) != 2 {
		t.Errorf("Must have 2 elements starting with 'PATH', but have %d", len(pathVars))
	}

	if !strings.HasPrefix(pathVars[1], fmt.Sprintf("PATH=%s", home)) {
		t.Errorf("The second PATH var needs to start with '%s', but it's not %s", home, pathVars[1])
	}

	log.Println("home:", homeVars)
	if len(homeVars) != 2 {
		t.Errorf("Must have 2 elements starting with 'HOME', but have %d", len(homeVars))
	}

	if strings.Compare(homeVars[1], workspace) == 0 {
		t.Errorf("The second home var needs to be '%s', but is %v", workspace, homeVars[1])
	}

}

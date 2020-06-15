package scripts

import (
	"regexp"
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

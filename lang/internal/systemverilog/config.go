package systemverilog

import (
	"encoding/json"
	"fmt"
	"github.com/gasrodriguez/crowned/internal/util"
	"io/ioutil"
	"path/filepath"
)

const (
	ConfigFilename = ".crowned-lang-sv.json"
)

type Config struct {
	General struct {
		Includes []string
		Defines  []string
	}
	Slang   toolConfig
	Svlint  toolConfig
	Verible struct {
		Format toolConfig
		Lint   veribleLintConfig
	}
}

func (o *Config) ExpandEnv() {
	util.ExpandEnvList(&o.General.Includes)
	util.ExpandEnvList(&o.Slang.Arguments)
	util.ExpandEnvList(&o.Svlint.Arguments)
	util.ExpandEnvList(&o.Verible.Lint.Arguments)
}

type toolConfig struct {
	Enabled   bool
	Arguments []string
}

type veribleLintConfig struct {
	Enabled   bool
	Arguments []string
	Rules     []string
}

func (o *Handler) loadConfig() Config {
	// default config
	config := Config{}
	config.Slang.Enabled = true
	config.Svlint.Enabled = true
	config.Verible.Lint.Enabled = true
	config.Verible.Format.Enabled = true

	configPath := filepath.Join(o.workspacePath, ConfigFilename)
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		o.ShowError(fmt.Sprintf("Cannot read config from file '%s'\nerror: '%s'", configPath, err.Error()))
		return config
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		o.ShowError(fmt.Sprintf("Cannot load config from file '%s'\nerror: '%s'", configPath, err.Error()))
		return config
	}
	(&config).ExpandEnv()
	return config
}

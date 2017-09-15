package parser

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/benjamincaldwell/devctl/utilities"

	"github.com/ghodss/yaml"
)

type Version struct {
	Version string
}

type ProjectConfig struct {
	Node              Version
	Go                Version
	Python            Version
	Scripts           map[string]utilities.RunCommand
	DockerCompose     interface{} `json:"docker-compose"`
	DockerComposeFile string      `json:"docker-compose-file"`
	Services          []interface{}
	Dependencies      struct {
		Install []string
		Brew    struct {
			Install []string
		}
		Aptget struct {
			Install []string
		} `json:"apt-get"`
	}
}

func (c *ProjectConfig) ParseFileDefault() error {
	return c.ParseFile("devctl.yaml", "./devctl.yml")
}

func (c *ProjectConfig) ParseFile(paths ...string) (err error) {
	for _, path := range paths {
		data, err := ioutil.ReadFile(path)
		if err == nil {
			err = yaml.Unmarshal(data, c)
			return err
		}
	}
	return err
}

func (c *ProjectConfig) ProjectName() (string, error) {
	return os.Getwd()
}

func (c *ProjectConfig) ParseJson(data string) {
	json.Unmarshal([]byte(data), c)
}

func (c *ProjectConfig) ParseYaml(data string) {
	yaml.Unmarshal([]byte(data), c)
}
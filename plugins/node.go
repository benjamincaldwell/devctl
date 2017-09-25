package plugins

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/postCommand"
	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/benjamincaldwell/go-printer"
)

type Node struct {
	path    string
	version string
}

func (n *Node) Setup() {
	isNvmInstalled := nvmInstalled()
	if isNvmInstalled {
		printer.Success("nvm already installed")
		return
	}
	resp, err := http.Get("https://raw.githubusercontent.com/creationix/nvm/v0.31.3/install.sh")
	utilities.ErrorCheck(err, "nvm download")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	printer.Info("Installing nvm")
	err = shell.Command("bash").SetInput(string(body)).Run()
	utilities.ErrorCheck(err, "nvm install")
}

func UpTasks(*parser.ProjectConfigStruct) (tasks [][]Task, err error) {
	// client := dockerClient.New()
	stage1 := []Task{
	// &startContainer{client, "node:4.0"},
	}
	tasks = append(tasks, stage1)
	return tasks, err
}

func (n *Node) PreInstall(c *parser.ProjectConfigStruct) {
	printer.Info("setting node version to " + c.Node.Version)
	n.version = c.Node.Version

	// check if nvm is install
	isNvmInstalled := nvmInstalled()
	if !isNvmInstalled {
		printer.Fail("nvm not installed. Run devctl setup to install")
		return
	}

	// check if requested version is installed
	nodeVersions, _ := shell.Command("sh", "-c", "source ~/.nvm/nvm.sh && nvm version "+n.version).Output()
	if strings.Contains(string(nodeVersions), "N/A") {
		shell.Command("sh", "-c", "source ~/.nvm/nvm.sh && nvm install "+n.version).PrintOutput()
	}

	// set correct version in ENV
	postCommand.RunCommand("nvm use " + n.version)
}

func (n *Node) Install(c *parser.ProjectConfigStruct) {
	// npm install
	printer.Info("npm install")
	printer.InfoLineTop()
	shell.Command("sh", "-c", "source ~/.nvm/nvm.sh && nvm use "+n.version+" > /dev/null && npm install").PrintOutput()
	printer.InfoLineBottom()
}

func (n Node) PostInstall(c *parser.ProjectConfigStruct) {

}

func (n Node) PreScript(c *parser.ProjectConfigStruct) {
}

func (n Node) Scripts(c *parser.ProjectConfigStruct) map[string]utilities.RunCommand {
	// return scripts struct array
	scripts := make(map[string]utilities.RunCommand)

	packageJSON := parser.Parser{}
	if err := packageJSON.Parse("package.json"); err != nil {
		return scripts
	}

	packageJSONscripts := packageJSON.GetMap("scripts")

	for name := range packageJSONscripts {
		scripts[name] = utilities.RunCommand{
			Command: "npm run " + name,
		}
	}
	return scripts
}

func (n Node) PostScript(c *parser.ProjectConfigStruct) {
}

func (n Node) Down(c *parser.ProjectConfigStruct) {
}

func (n Node) IsProjectType(c *parser.ProjectConfigStruct) bool {
	if c.Node.Version != "" {
		return true
	}
	return false
}

func nvmInstalled() bool {
	err := shell.Command("sh", "-c", "source ~/.nvm/nvm.sh && command -v nvm").Run()
	if err != nil {
		return false
	}
	return true
}

package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"runtime-helper"
	"logger"
)

var (
	rt = runtime.RuntimeHelper{}
	lg = logger.GetLogger()
)

var (
	app = kingpin.New("Go Builder", "CI stuff")
	batchMode = kingpin.Flag("batch", "Runs commands without asking for any input").Bool()

	buildCommand = kingpin.Command("build", "Runs commands without asking for any input")

	cleanCommand = kingpin.Command("clean", "Cleans all packages and binaries")

	deployToCommand = kingpin.Command("deploy-to", "Runs commands without asking for any input")
	deployToNexus = deployToCommand.Command("nexus", "Deploy to Nexus")
	deployToNexusRepId = deployToCommand.Flag("repoId", "Repository ID").Default("nexus").String()
	deployToNexusUrl = deployToCommand.Flag("url", "Repository URL").Required().String()
	deployToNexusFile = deployToCommand.Flag("file", "Package to Deploy").Required().String()
	deployToServer = deployToCommand.Command("server", "Deploy to Server")

	distCommand = kingpin.Command("dist", "Creates distribution package")

	installCommand = kingpin.Command("install", "Install helper apps")
	installGBCommand = installCommand.Command("gb", "Installs GB")

	path = kingpin.Flag("path", "Go project path").Default(".").String()

	testCommand = kingpin.Command("test", "Run Tests")
)

func init() {

}

func main() {
	var cmds string
	cmds = kingpin.Parse()
	checkProjectFile()
	checkHome()
	switch cmds {
	case "build":
		build()
	case "clean":
		clean()
	case "deploy-to nexus":
		deploy2Nexus()
	case "deploy-to server":
		fmt.Println("server")
	case "dist":
		dist()
	case "test":
		runTests()
	}
}

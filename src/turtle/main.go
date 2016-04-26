package main

import (
	"fmt"
	"os"
	"gopkg.in/alecthomas/kingpin.v2"
	"os/signal"
	"syscall"
	"github.com/daviddengcn/go-colortext"
	"github.com/correia-io/goutils/src/utils"
	"github.com/correia-io/goutils/src/logd"
	"model"
	"plugin"
	"strings"
	"path/filepath"
)

var (
	TURTLE_FILE string
	TURTLE_HOME string
	TURTLE_VERSION = "0.0.1-SNAPSHOT"
	TURTLE_PROJECT_PATH string
	rt = utils.RuntimeHelper{}
	logger = logd.GetLogger()
	wiz = utils.Wizard{}
	compressor = utils.Compress{}
	project model.Project
	fileUtils = utils.FileUtils{}
	cmds string
	goBuilder plugin.GoBuilder
)

var (
	app = kingpin.New("turtle", "turtle - build, test, deploy, release, install build tools.")

	projectPath = kingpin.Flag("path", "Go project path").Short('p').Default(".").String()

	batchMode = kingpin.Flag("batch", "Runs commands without asking for any input").Bool()

	buildCommand = kingpin.Command("build", "Runs commands without asking for any input")

	cleanCommand = kingpin.Command("clean", "Cleans all packages and binaries")

	deployToCommand = kingpin.Command("deploy2", "Runs commands without asking for any input")
	deployToNexus = deployToCommand.Command("nexus", "Deploy to Nexus")
	deployToNexusRepoId = deployToCommand.Flag("repoId", "Repository ID").Default("default").Short('r').String()
	deployToNexusGeneratePom = deployToCommand.Flag("generate-pom", "Generate POM").Short('g').Default("true").String()
	deployToNexusFile = deployToCommand.Flag("file", "Package to Deploy").Short('f').String()
	deploy2NexusBuild = deployToNexus.Flag("build", "List of build ID's to deploy, separeted by commas.").Short('b').String()

	//deployToServer = deployToCommand.Command("server", "Deploy to Server")

	distCommand = kingpin.Command("dist", "Creates distribution package")

	installCommand = kingpin.Command("install", "Install helper apps")
	installGBCommand = installCommand.Command("gb", "Installs GB")
	testCommand = kingpin.Command("test", "Run Tests")

	releaseCommand = kingpin.Command("release", "Performs SCM release.")

	versionCommand = kingpin.Command("version", "Version")
)

func init() {
	kingpin.CommandLine.HelpFlag.Short('h')
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		ct.ResetColor()
		fmt.Println("Shutting down gracefully...")
		defer fmt.Println("Shutdown complete")
		os.Exit(0)
	}()

	cmds = kingpin.Parse()

	TURTLE_PROJECT_PATH, err := filepath.Abs(filepath.Dir(*projectPath + "/"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Chdir(TURTLE_PROJECT_PATH)

	TURTLE_HOME = os.Getenv("HOME") + "/" + app.Name
	TURTLE_FILE = TURTLE_PROJECT_PATH + "/turtle.json"
	tt := Turtle{}
	ct.Foreground(ct.Cyan, false)
	fmt.Println("Found Turtle file: " + TURTLE_FILE)
	ct.ResetColor()
	project = tt.GetProject()
	app.Version(project.Version)

	tt.CheckHome()
	goBuilder = plugin.GoBuilder{Project:project}
	switch cmds {
	case "build":
		tt.Build()
	case "clean":
		tt.Clean()
	case "deploy2 nexus":
		if *deploy2NexusBuild == "" && *deployToNexusFile == "" {
			ct.Foreground(ct.Red, true)
			fmt.Println("Error: You must provide build id list [--builds|-b] or a file to deploy [--file|-f]")
			ct.Foreground(ct.White, false)
			os.Exit(1)
		}
		builds := []string(strings.Split(*deploy2NexusBuild, ","))
		tt.Deploy2Nexus(builds)
	case "deploy2 server":
		fmt.Println("Coming soon...")
	case "dist":
		tt.Dist()
	case "install gb":
		tt.InstallGB()
	case "test":
		tt.RunTests()
	case "release":
		tt.Release()
	case "version":
		fmt.Println(app.Name, TURTLE_VERSION)

	//if (*EE) {
	//s.EE()
	//}
	}
	ct.ResetColor()
}

package main

import (
	"fmt"
	"os"
	"gopkg.in/alecthomas/kingpin.v2"
	"os/signal"
	"syscall"
	"github.com/daviddengcn/go-colortext"
	"model"
	"plugin"
	"strings"
	"path/filepath"
	"utils"
	"logd"
)

var (
	TURTLE_FILE string
	TURTLE_HOME string
	TURTLE_PROJECT_PATH string
	TURTLE_CONFIG_FILE string
	rt = &utils.RuntimeHelper{}
	logger = *logd.GetLogger()
	wiz = &utils.Wizard{}
	compressor = &utils.Compress{}
	project model.Project
	fileUtils = &utils.FileUtils{}
	cmds string
	goBuilder *plugin.GoBuilder
	distFolder string
)

var (
	app = kingpin.New("turtle", "turtle - build, test, deploy, release, install build tools.")

	projectPath = kingpin.Flag("path", "Go project path").Short('p').Default(".").String()

	batchMode = kingpin.Flag("batch", "Runs commands without asking for any input").Bool()

	buildCommand = kingpin.Command("build", "Runs GB Build")
	
	testCommand = kingpin.Command("test", "Executes GO tests")
	testCommandCoverage = kingpin.Flag("coverage", "Runs commands without asking for any input").Bool()

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

	releaseCommand = kingpin.Command("release", "Performs SCM release.")

	updateCommand = kingpin.Command("update", "Updates lots of different stuff")
	updateManageVersion = updateCommand.Command("version", "Updates Turtles Managed Version Constant")
	updateManagedNewVersion = updateManageVersion.Flag("new-version", "New Version to te applied").Short('n').Required().String()

	versionCommand = kingpin.Command("version", "Version")
)

func init() {
	kingpin.CommandLine.HelpFlag.Short('h')

}

func easyDeath() {
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
}

func main() {
	easyDeath()

	cmds = kingpin.Parse()

	TURTLE_PROJECT_PATH, err := filepath.Abs(filepath.Dir(*projectPath + "/"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Chdir(TURTLE_PROJECT_PATH)
	if os.Getenv("TURTLE_HOME") != "" {
		TURTLE_HOME = os.Getenv("TURTLE_HOME")
	} else {
		TURTLE_HOME = os.Getenv("HOME") + "/." + app.Name
	}

	TURTLE_FILE = TURTLE_PROJECT_PATH + "/turtle.json"
	tt := Turtle{}
	ct.Foreground(ct.Cyan, false)
	//fmt.Println("Found Turtle file: " + TURTLE_FILE)
	ct.ResetColor()
	project = tt.GetProject()
	distFolder = project.ArtifactId + "-" + project.Version
	TURTLE_CONFIG_FILE = TURTLE_HOME + "/config.json"
	tt.CheckHome()
	tt.LoadConfig()
	app.Version(project.Version)

	goBuilder = &plugin.GoBuilder{}

	switch cmds {
	case "build":
		tt.Build()

	case "clean":
		tt.Clean()
	case "deploy2 nexus":
		var bs string
		if *deploy2NexusBuild == "" {
			bs = "default"
		} else {
			bs = *deploy2NexusBuild
		}

		if bs == "" && *deployToNexusFile == "" {
			ct.Foreground(ct.Red, true)
			fmt.Println("Error: You must provide build id list [--builds|-b] or a file to deploy [--file|-f]")
			ct.Foreground(ct.White, false)
			os.Exit(1)
		}
		builds := []string(strings.Split(bs, ","))
		tt.Deploy2Nexus(builds)
	case "deploy2 server":
		fmt.Println("Coming soon...")
	case "dist":
		tt.Dist()
	case "install gb":
		tt.InstallGB()
	case "test":
		tt.RunTests(*testCommandCoverage)
	case "release":
		tt.Release()

	case "version":
		fmt.Println(app.Name, VERSION)


	}
	ct.ResetColor()
}

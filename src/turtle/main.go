package main

import (
	"fmt"
	"model"
	"os"
	"github.com/correia-io/goutils/src/utils"
	"github.com/correia-io/goutils/src/logd"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	TURTLE_FILE string
	TURTLE_HOME string
	TURTLE_VERSION = "0.0.1-SNAPSHOT"
	rt = utils.RuntimeHelper{}
	logger = logd.GetLogger()
	wiz = utils.Wizard{}
	compressor = utils.Compress{}
	project model.Project
	fileUtils = utils.FileUtils{}
	cmds string
)

var (
	app = kingpin.New("turtle", "turtle - build, test, deploy, release, install build tools.")

	TURTLE_PROJECT_PATH = kingpin.Flag("path", "Go project path").Default(".").String()
	batchMode = kingpin.Flag("batch", "Runs commands without asking for any input").Bool()
	buildCommand = kingpin.Command("build", "Runs commands without asking for any input")
	cleanCommand = kingpin.Command("clean", "Cleans all packages and binaries")

	deployToCommand = kingpin.Command("deploy", "Runs commands without asking for any input")
	deployToNexus = deployToCommand.Command("nexus", "Deploy to Nexus")
	deployToNexusRepId = deployToCommand.Flag("repoId", "Repository ID").Required().Short('r').String()
	deployToNexusGeneratePom = deployToCommand.Flag("generate-pom", "Generate POM").Short('g').Default("true").String()
	deployToNexusFile = deployToCommand.Flag("file", "Package to Deploy").Short('f').Required().String()

	deployToServer = deployToCommand.Command("server", "Deploy to Server")

	distCommand = kingpin.Command("dist", "Creates distribution package")

	installCommand = kingpin.Command("install", "Install helper apps")
	installGBCommand = installCommand.Command("gb", "Installs GB")
	testCommand = kingpin.Command("test", "Run Tests")

	releaseCommand = kingpin.Command("release", "Performs SCM release.")

	versionCommand = kingpin.Command("version", "Version")
	//EE = versionCommand.Flag("ee", "??").Bool()
	wtfCMD = kingpin.Command("wtf","WTF??")
)

func init() {
	kingpin.CommandLine.HelpFlag.Short('h')
	TURTLE_HOME = os.Getenv("HOME") + app.Name
	TURTLE_FILE = *TURTLE_PROJECT_PATH + "/turtle.json"
}

func main() {
	os.Chdir(*TURTLE_PROJECT_PATH)
	cmds = kingpin.Parse()
	app.Version(project.Version)
	s := Turtle{}

	s.CheckHome()

	switch cmds {
	case "build":
		s.Build()
	case "clean":
		s.Clean()
	case "deploy nexus":
		s.Deploy2Nexus()
	case "deploy server":
		fmt.Println("Coming soon...")
	case "dist":
		s.Dist()
	case "install gb":
		s.InstallGB()
	case "test":
		s.RunTests()
	case "release":
		s.Release()
	case "version":
		fmt.Println(app.Name, TURTLE_VERSION)
	case"wtf":
		resp:=wiz.Question("WTF??")
		fmt.Println(resp,"<---")

	//if (*EE) {
		//s.EE()
	//}
	}
}

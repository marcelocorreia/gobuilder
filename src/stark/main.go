package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/correia-io/goutils/src/logd"
	"github.com/correia-io/goutils/src/utils"
	"model"
	"os"
)

var (
	STARK_FILE string
	STARK_HOME string
	STARK_VERSION = "0.0.1-SNAPSHOT"
	rt = utils.RuntimeHelper{}
	logger = logd.GetLogger()
	wiz = utils.Wizard{}
	compressor = utils.Compress{}
	project model.Project
	fileUtils = utils.FileUtils{}
	cmds string
)

var (
	app = kingpin.New("stark", "stark - build, test, deploy, release, install build tools.")

	STARK_PROJECT_PATH = kingpin.Flag("path", "Go project path").Default(".").String()
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
)

func init() {
	cmds = kingpin.Parse()
	//if *versionCommand {

	//}
	STARK_HOME = os.Getenv("HOME") + app.Name
	STARK_FILE = *STARK_PROJECT_PATH + "/stark.json"
	os.Chdir(*STARK_PROJECT_PATH)
	//kingpin.CommandLine.HelpFlag.Short('h')
	//file, err := ioutil.ReadFile(*STARK_PROJECT_PATH)
	//if err != nil {
	//	logger.Error(err)
	//} else {
	//	var p model.Project
	//	err := json.Unmarshal(file, &p)
	//	if err != nil {
	//		logger.Fatal(err)
	//	}
	//	project = p
	//}

	app.Version(project.Version)
}

func main() {

	s := Stark{}

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
		fmt.Println(app.Name, STARK_VERSION)
	}
}

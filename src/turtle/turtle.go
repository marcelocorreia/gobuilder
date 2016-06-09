package main

import (
	"os"
	"path/filepath"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"model"
	"github.com/pborman/uuid"
	"strings"
	"github.com/correia-io/goutils/src/utils"
	"github.com/daviddengcn/go-colortext"
	"strconv"
)

type Tortuga interface {
	Build()
	CheckHome()
	CheckProjectFile()
	Clean()
	Deploy2Nexus()
	Dist()
	GetProject()
	InstallGB()
	LoadConfig()
	Release()
	RunTests()
	SaveConfig()
	UpdateManagedVersion()
}

type Turtle struct {
	config *model.TurtleConfig
}

func (t Turtle) SaveConfig() (error) {
	fileUtils.CopyFile(TURTLE_CONFIG_FILE, TURTLE_CONFIG_FILE + ".backup")
	cf := t.LoadConfig()
	resp, _ := json.MarshalIndent(&cf, "", "  ")
	wr := []byte(resp)
	logger.Debug("Writing config file", TURTLE_CONFIG_FILE)
	err := ioutil.WriteFile(TURTLE_CONFIG_FILE, wr, 0750)

	if (err != nil) {
		return err
	}

	return nil
}

func (t Turtle) LoadConfig() (model.TurtleConfig) {
	ct.Foreground(ct.Cyan, false)

	var cfg model.TurtleConfig
	cFile, err := ioutil.ReadFile(TURTLE_CONFIG_FILE)
	if err == nil {
		var t model.TurtleConfig
		err := json.Unmarshal(cFile, &t)
		if err != nil {
			logger.Error(err)
		}
		cfg = t
	}
	ct.ResetColor()

	t.config = &cfg

	return cfg
}

func (t Turtle) Build() {
	dir, _ := os.Getwd()
	fmt.Println(dir)

	p := t.GetProject()
	ct.Foreground(ct.Cyan, false)
	fmt.Println("Building ▶", p.Name + "." + p.Version)
	ct.Foreground(ct.Green, false)
	args := []string{"build", "-F", "-f", "-ldflags=-X " + p.VersionString + "=" + p.Version}

	errRun := rt.RunCommandLogStream("gb", args)

	if errRun != nil {
		logger.Error("Error building", p.ArtifactId, p.Version, errRun)

	}

	//if _, err := os.Stat(distFolder); os.IsNotExist(err) {
	//	os.Mkdir(distFolder, 00750)
	//}
	ct.ResetColor()
}

func (t Turtle) Clean() {
	fmt.Println("Cleaning the house")
	os.RemoveAll(TURTLE_PROJECT_PATH + distFolder)
	os.RemoveAll(TURTLE_PROJECT_PATH + "pkg")
	os.RemoveAll(TURTLE_PROJECT_PATH + "bin")
	files, err := ioutil.ReadDir(".")
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), "tar.gz") {
			fmt.Printf("Removing file:%s\n", file.Name())
			os.Remove(file.Name())
		}
	}
}

func (t Turtle) CheckHome() {
	if _, err := os.Stat(TURTLE_HOME); os.IsNotExist(err) {
		os.Mkdir(TURTLE_HOME, 00750)
		t.SaveConfig()
	}
}

func (t Turtle) Dist() {
	if project.ProjectType == "static" {
		dist := project.ArtifactId + "-dist"
		ct.Foreground(ct.Green, false)
		fmt.Println("Packaging Static Project", project.ArtifactId)
		tmpDir := "/tmp/" + uuid.New()

		fmt.Println(os.Getwd())
		source, _ := os.Getwd()

		fileUtils.CopyDir(source, tmpDir + "/" + project.ArtifactId)

		os.RemoveAll(dist)
		if e, _ := fileUtils.Exists(dist); !e {
			os.Mkdir(distFolder, 00750)
		}

		distName := fmt.Sprintf(source + "/%s-%s.%s", project.ArtifactId, project.Version, project.Packaging)
		os.Chdir(tmpDir)
		fmt.Println(tmpDir)
		fmt.Println(distName)
		compressor.Tar(project.ArtifactId, distName)
		os.RemoveAll(tmpDir)
		ct.ResetColor()
	} else if (project.ProjectType == "go") {
		goBuilder.Dist(&project)
		os.Unsetenv("GOOS")
		os.Unsetenv("GOARCH")
		os.RemoveAll("bin/")
		os.RemoveAll(distFolder)
		t.Build()
	}
}

func (t Turtle) InstallGB() {
	ct.Foreground(ct.Cyan, false)
	fmt.Println("Installing GB....")
	workdir := "/tmp/" + uuid.New()
	os.Chdir(workdir)
	os.Setenv("GOPATH", workdir)
	rt.RunCommandLogStream("go", []string{"get", "github.com/constabulary/gb/..."})
	fmt.Println("Copying GB binaries to /usr/local/bin/")
	rt.RunCommandLogStream("sudo", []string{"cp", workdir + "/bin/gb", "/usr/local/bin/gb"})
	rt.RunCommandLogStream("sudo", []string{"cp", workdir + "/bin/gb-vendor", "/usr/local/bin/gb-vendor"})
	os.RemoveAll(workdir)
	ct.Foreground(ct.Green, false)
	fmt.Println("Done.")
	ct.ResetColor()

}

func (t Turtle) RunTests() {
	dir, err := filepath.Abs(filepath.Dir(TURTLE_PROJECT_PATH))
	os.Chdir(dir)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug("Runing tests @", dir)
	rt.RunCommandLogStream("gb", []string{"test"})
}

func (t Turtle) Deploy2Nexus(builds []string) {
	if !rt.CheckBinaryInPath("mvn") {
		ct.Foreground(ct.Red, false)
		fmt.Println("Maven not found in PATH, please check your configuration.")
		os.Exit(1)
	}
	project := t.GetProject()
	var jobRepo model.Repository
	fmt.Println("Starting Deployment to Nexus Jobs:", builds)

	for _, build := range project.Builds {
		if utils.StringInSlice(build.ID, builds) {
			ct.Foreground(ct.Cyan, true)

			fmt.Println("Running build:", build.Type, build.ID)

			ct.Foreground(ct.Green, false)
			for _, rp := range t.LoadConfig().Repositories {
				if rp.Id == *deployToNexusRepoId {
					jobRepo = rp
				}
			}

			// Overrides repositories from Turtle Config
			for _, rp := range project.Repositories {
				if rp.Id == *deployToNexusRepoId {
					jobRepo = rp
				}
			}

			if jobRepo.Type == "releases" && strings.Contains(project.Version, "SNAPSHOT") {
				ct.Foreground(ct.Red, true)
				fmt.Println("Error: You are trying to deploy a unreleased package into repository of type\"releases\"")
				ct.Foreground(ct.Red, false)
				fmt.Println("Please check your project and try again once sorted.")
				ct.ResetColor()
				os.Exit(1)
			} else if jobRepo.URL == "" {
				ct.Foreground(ct.Red, true)
				fmt.Printf("Error: Repository ID -> %s not found in Turtle file\n", *deployToNexusRepoId)
				ct.Foreground(ct.Red, false)
				fmt.Println("Please check your project and try again once sorted.")
				ct.ResetColor()
				os.Exit(1)
			}
			var artifactId string

			if project.ProjectType == "go" {
				artifactId = fmt.Sprintf("%s-%s-%s", project.ArtifactId, build.OS, build.Arch)
			} else {
				artifactId = project.Version
			}

			file := fmt.Sprintf("%s-%s.%s", artifactId, project.Version, project.Packaging)
			fmt.Println("Deploying file:", file)
			args := []string{
				"deploy:deploy-file",
				"-DgroupId=" + project.GroupId,
				"-DartifactId=" + artifactId,
				"-Dversion=" + project.Version,
				"-Dpackaging=" + project.Packaging,
				"-Durl=" + jobRepo.URL,
				"-Dfile=" + file,
				"-DgeneratePom=" + *deployToNexusGeneratePom,
				"-DrepositoryId=" + jobRepo.Id,
			}
			fmt.Println(args)
			err := rt.RunCommandLogStream("mvn", args)
			if err != nil {
				logger.Fatal(err)
			}
		}
	}

}

type repoError struct {
	s string
}

func (e *repoError) Error() string {
	return e.s
}

func (t Turtle) getRepo(id string) (model.Repository, error) {
	for _, r := range project.Repositories {
		if r.Id == strings.TrimSpace(id) {
			return r, nil
		}
	}
	return model.Repository{}, &repoError{"Repository not found"}
}

func (t Turtle) GetProject() (model.Project) {
	projectFile, err := ioutil.ReadFile(TURTLE_FILE)
	var project model.Project
	if err != nil {
		ct.Foreground(ct.Red, true)
		//logger.Error("Turtle File Error.", err, TURTLE_FILE)
		ct.ResetColor()
	} else {
		var c model.Project
		err := json.Unmarshal(projectFile, &c)
		if err != nil {
			logger.Error(err)
		}
		project = c
	}
	ct.ResetColor()
	return project
}

func (t Turtle) Release() {
	prj := t.GetProject()
	v := prj.Version
	nextVersion := strings.Replace(v, "-SNAPSHOT", "", -1)
	prj.Version = nextVersion
	jsonStr, err := json.MarshalIndent(&prj, "", "  ")
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println("Updating:", TURTLE_FILE)
	err = ioutil.WriteFile(TURTLE_FILE, jsonStr, 00750)

	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println("Releasing project", app.Name, "-", prj.Version)
	t.Build()
	rt.RunCommandLogStream("git", []string{"tag", prj.Version})

	brkdwn := strings.Split(prj.Version, ".")

	var nextDevVersion string

	for i := 0; i < (len(brkdwn) - 1); i++ {
		nextDevVersion += brkdwn[i] + "."
	}
	res, _ := strconv.Atoi(brkdwn[(len(brkdwn) - 1)])

	nextDevVersion += strconv.Itoa(res + 1)

	prj.Version = nextDevVersion + "-SNAPSHOT"

	jsonStr, err = json.MarshalIndent(&prj, "", "  ")
	if err != nil {
		logger.Error(err)
	}

	err = ioutil.WriteFile(TURTLE_FILE, jsonStr, 00750)

	if err != nil {
		logger.Error(err)
	}
}

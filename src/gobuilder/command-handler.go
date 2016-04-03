package main

import (
	"os"
	"path/filepath"
	"utils"
	"strings"
	"fmt"
	"model"
	"encoding/json"
	"io/ioutil"
	"github.com/daviddengcn/go-colortext"
)

func build() {
	os.Chdir(*path)
	lg.Debug("Building go project @", *path)
	rt.RunCommandLogStream("gb", []string{"build"})

}

func clean() {
	dir, err := filepath.Abs(filepath.Dir(*path))
	os.Chdir(dir)
	if err != nil {
		lg.Fatal(err)
	}
	resp := utils.Question("You are about to clean all binaries and packages from " + dir + "\nProceed: [y/N] ")
	if strings.ToLower(resp) == "y" {
		os.Remove(dir + "/dist")
		os.Remove(dir + "/pkg")
		os.Remove(dir + "/bin")
		os.Mkdir(dir + "/dist", 00750)
		os.Mkdir(dir + "/dist/macos", 00750)
		os.Mkdir(dir + "/dist/linux", 00750)
		os.Mkdir(dir + "/dist/windows", 00750)
	}else {
		fmt.Println("Aborted")
	}
}

func checkProjectFile() {
	dir, err := filepath.Abs(filepath.Dir(*path))
	if err != nil {
		lg.Fatal(err)
	}

	if _, err := os.Stat(dir + "/gobuilder.json"); os.IsNotExist(err) {
		ct.Foreground(ct.Red, false)
		resp := utils.Question("Project doesn't have gobuilder.json. Would you like to create one? [y/N] ")
		project := model.Project{}
		project.Version = "0.0.1-SNAPSHOT"
		if strings.ToLower(resp) == "y" {
			slice := strings.Split(dir, "/")
			projectName := slice[len(slice) - 1]
			ct.Foreground(ct.Cyan, false)
			pName := utils.QuestionF("Project Name: [%s] ", projectName)
			if pName == "" {
				if pName == "" {
					project.Name = projectName
				}else {
					project.Name = pName
				}
			}

			pGroup := utils.QuestionF("GroupId: [%s] ", "com.company.my")
			if pGroup == "" {
				project.GroupId = "com.company.my"
			}else {
				project.GroupId = pGroup
			}

			pArti := utils.QuestionF("ArtifactId: [%s] ", projectName)
			if pArti == "" {
				project.ArtifactId = projectName
			}else {
				project.ArtifactId = pArti
			}


			file, _ := json.MarshalIndent(&project, "", "  ")
			wr := []byte(file)

			err := ioutil.WriteFile(dir + "/gobuilder.json", wr, 0750)
			if err != nil {
				lg.Fatal(err)
			}
			ct.Foreground(ct.Yellow,false)
			fmt.Println("Writing gobuilder config file...")
			fmt.Println(string(wr))
			ct.Foreground(ct.White, false)
		}else {
			fmt.Println("Aborted")
			os.Exit(1)
		}

	}

}
func dist() {
	dir, err := filepath.Abs(filepath.Dir(*path))
	os.Chdir(dir)
	if err != nil {
		lg.Fatal(err)
	}
	clean()
	os.Setenv("GOOS", "darwin")
	os.Setenv("GOARCH", "amd64")
	build()
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "amd64")
	build()
	os.Setenv("GOOS", "windows")
	os.Setenv("GOARCH", "amd64")
	build()
	utils.Tar(dir + "/dist", "dist.tar.gz")
}
func runTests() {
	dir, err := filepath.Abs(filepath.Dir(*path))
	os.Chdir(dir)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Debug("Runing tests @", dir)
	rt.RunCommandLogStream("gb", []string{"test"})
}

func deploy2Nexus() {
	if !rt.CheckBinaryInPath("mvn") {
		lg.Fatal("Maven not found in PATH, please check your configuration.")
	}else {
		project := getProject()
		args := []string{
			"-DgroupId=" + project.GroupId,
			"-DartifactId=" + project.ArtifactId,
			"-Dversion=" + project.Version,
			"-Dpackaging=tar.gz",
			"-DgeneratePom=true",
			"-Durl=" + *deployToNexusUrl,
			"-Dfile=" + project.ArtifactId + "-" + project.Version,
			"-DrepositoryId=" + *deployToNexusRepId,
		}
		err := rt.RunCommandLogStream("mvn", args)
		if err != nil {
			lg.Fatal(err)
		}
	}
}

func getProject() (model.Project) {
	dir, err := filepath.Abs(filepath.Dir(*path))
	os.Chdir(dir)
	if err != nil {
		lg.Fatal(err)
	}
	projectFile, err := ioutil.ReadFile(dir + "/gobuilder.json")
	var project model.Project
	if err != nil {
		lg.Error("Workspace busted")
	}else {
		var c model.Project
		err := json.Unmarshal(projectFile, &c)
		if err != nil {
			lg.Error(err)
		}
		project = c
	}

	return project
}
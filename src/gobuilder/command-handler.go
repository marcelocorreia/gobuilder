package main

import (
	"os"
	"path/filepath"
	"strings"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/daviddengcn/go-colortext"
	"model"
	"github.com/pborman/uuid"
)

func build() {
	os.Chdir(*path)
	logger.Debug("Building go project @", *path)
	rt.RunCommandLogStream("gb", []string{"build"})

}

func clean() {
	dir, err := filepath.Abs(filepath.Dir(*path))
	os.Chdir(dir)
	if err != nil {
		logger.Fatal(err)
	}
	resp := wiz.Question("You are about to clean all binaries and packages from " + dir + "\nProceed: [y/N] ")
	if strings.ToLower(resp) == "y" {
		os.Remove(dir + "/dist")
		os.Remove(dir + "/pkg")
		os.Remove(dir + "/bin")
	} else {
		fmt.Println("Aborted")
	}
}

func checkHome() {
	homeDir := os.Getenv("HOME") + "/.gobuilder"
	if _, err := os.Stat(homeDir); os.IsNotExist(err) {
		os.Mkdir(homeDir, 00750)
	}
}

func checkProjectFile() {
	dir, err := filepath.Abs(filepath.Dir(*path))
	if err != nil {
		logger.Fatal(err)
	}

	if _, err := os.Stat(dir + "/gobuilder.json"); os.IsNotExist(err) {
		ct.Foreground(ct.Red, false)
		resp := wiz.Question("Project doesn't have gobuilder.json. Would you like to create one? [y/N] ")
		project := model.Project{}
		project.Version = "0.0.1-SNAPSHOT"
		if strings.ToLower(resp) == "y" {
			slice := strings.Split(dir, "/")
			projectName := slice[len(slice) - 1]
			ct.Foreground(ct.Cyan, false)
			pName := wiz.QuestionF("Project Name: [%s] ", projectName)
			if pName == "" {
				if pName == "" {
					project.Name = projectName
				} else {
					project.Name = pName
				}
			}

			pGroup := wiz.QuestionF("GroupId: [%s] ", "com.company.my")
			if pGroup == "" {
				project.GroupId = "com.company.my"
			} else {
				project.GroupId = pGroup
			}

			pArti := wiz.QuestionF("ArtifactId: [%s] ", projectName)
			if pArti == "" {
				project.ArtifactId = projectName
			} else {
				project.ArtifactId = pArti
			}

			packaging := wiz.QuestionF("Packaging: [%s] ", "tar.gz")
			if pArti == "" {
				project.Packaging = "tar.gz"
			} else {
				project.Packaging = packaging
			}

			file, _ := json.MarshalIndent(&project, "", "  ")
			wr := []byte(file)

			err := ioutil.WriteFile(dir + "/gobuilder.json", wr, 0750)
			if err != nil {
				logger.Fatal(err)
			}
			ct.Foreground(ct.Yellow, false)
			fmt.Println("Writing gobuilder config file...")
			fmt.Println(string(wr))
			ct.Foreground(ct.White, false)
		} else {
			fmt.Println("Aborted")
			ct.Foreground(ct.White, false)
			os.Exit(1)
		}

	}

}

func dist() {
	if (project.ProjectType == "go") {
		dir, err := filepath.Abs(filepath.Dir(*path))
		os.Chdir(dir)
		if err != nil {
			logger.Fatal(err)
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
		compressor.Tar(dir + "/dist", "dist.tar.gz")
	} else if project.ProjectType == "da-template" {
		fmt.Println("Packaging DA Template")
		tmpDir := "/tmp/" + uuid.New()

		fmt.Println(os.Getwd())
		source, _ := os.Getwd()

		fileUtils.CopyDir(source, tmpDir + "/" + project.ArtifactId)

		os.RemoveAll("dist")
		if e, _ := fileUtils.Exists("dist"); !e {
			os.Mkdir("dist", 00750)
		}

		distName := fmt.Sprintf(source + "/dist/%s-%s.%s", project.ArtifactId, project.Version, project.Packaging)
		os.Chdir(tmpDir)
		fmt.Println(tmpDir)
		fmt.Println(distName)
		compressor.Tar(project.ArtifactId, distName)
		os.RemoveAll(tmpDir)
	}
}

func runTests() {
	dir, err := filepath.Abs(filepath.Dir(*path))
	os.Chdir(dir)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug("Runing tests @", dir)
	rt.RunCommandLogStream("gb", []string{"test"})
}

func deploy2Nexus() {
	if !rt.CheckBinaryInPath("mvn") {
		logger.Fatal("Maven not found in PATH, please check your configuration.")
	} else {
		project := getProject()
		args := []string{
			"-DgroupId=" + project.GroupId,
			"-DartifactId=" + project.ArtifactId,
			"-Dversion=" + project.Version,
			"-Dpackaging=tar.gz",
			"-DgeneratePom=true",
			"-Durl=" + *deployToNexusUrl,
			"-Dfile=" + *deployToNexusFile,
			"-DrepositoryId=" + *deployToNexusRepId,
		}
		err := rt.RunCommandLogStream("mvn", args)
		if err != nil {
			logger.Fatal(err)
		}
	}
}

func getProject() (model.Project) {
	dir, err := filepath.Abs(filepath.Dir(*path))
	os.Chdir(dir)
	if err != nil {
		logger.Fatal(err)
	}
	projectFile, err := ioutil.ReadFile(dir + "/gobuilder.json")
	var project model.Project
	if err != nil {
		logger.Error("Workspace busted")
	} else {
		var c model.Project
		err := json.Unmarshal(projectFile, &c)
		if err != nil {
			logger.Error(err)
		}
		project = c
	}

	return project
}
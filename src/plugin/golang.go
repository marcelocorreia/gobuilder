package plugin

import (
	"model"
	"fmt"
	"os"
	"github.com/daviddengcn/go-colortext"
	"io/ioutil"
	"strings"
	"github.com/correia-io/goutils/src/utils"
)

type GoBuilder struct {}

var (
	rt = utils.RuntimeHelper{}
	fileUtils = utils.FileUtils{}
	compressor = utils.Compress{}
)

func (s GoBuilder) Build(p *model.Project) error {
	dir, _ := os.Getwd()
	fmt.Println(dir)

	ct.Foreground(ct.Cyan, false)
	fmt.Println("Building ▶", p.Name + "." + p.Version)
	ct.Foreground(ct.Green, false)

	args := []string{"build", "-F", "-f", "-ldflags=-X " + p.VersionString + "=" + p.Version}

	errRun := rt.RunCommandLogStream("gb", args)
	defer ct.ResetColor()

	if errRun != nil {
		return errRun

	}
	return nil
}

func (s GoBuilder) Dist(project *model.Project) error {
	distFolder := project.ArtifactId + "-" + project.Version
	for _, build := range project.Builds {
		if build.Type == "go" {
			packBase := fmt.Sprintf("%s-%s-%s-%s", project.ArtifactId, build.OS, build.Arch, project.Version)
			pack := fmt.Sprintf("%s.%s", packBase, project.Packaging)
			os.Setenv("GOOS", build.OS)
			os.Setenv("GOARCH", build.Arch)
			ct.Foreground(ct.Cyan, true)
			fmt.Println("Building: ", packBase)
			ct.Foreground(ct.Green, false)
			s.Build(project)
			ct.Foreground(ct.Cyan, true)
			fmt.Println("Generating package:", pack)
			ct.Foreground(ct.Green, false)
			files, err := ioutil.ReadDir("./bin")
			if err != nil {
				fmt.Println(err, "Error searching for binaries. ")
				ct.Foreground(ct.Red, false)
				fmt.Println("Perhaps the project might not generate an executable")
				fmt.Println("If you think is a good idea to create package for libraries only, please contact the author or send a pull request at https://github.com/marcelocorreia/turtle")
				ct.ResetColor()
				return err
			}

			os.Mkdir(distFolder,00750)
			for _, file := range files {
				fmt.Println("Renaming:", file.Name())
				rpl := fmt.Sprintf("-%s-%s", build.OS, build.Arch)
				os.Rename("bin/" + file.Name(), distFolder + "/" + strings.Replace(file.Name(), rpl, "", -1))
			}
			fileUtils.CopyFile("README.md", distFolder + "/README.md")
			fileUtils.CopyFile("turtle.json", distFolder + "/turtle.json")

			compressor.Tar(distFolder + "/", pack)
			os.RemoveAll("bin/")
			os.RemoveAll(distFolder)
			s.Build(project)
		} else {
			fmt.Println("Skipping non \"go\" build type:", build.Type)
		}
	}
	return nil
}

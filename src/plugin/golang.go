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

type GoBuilder struct {
	Project model.Project
}

var (
	rt = utils.RuntimeHelper{}
	fileUtils = utils.FileUtils{}
	compressor = utils.Compress{}
)

func (s GoBuilder) Build() error {
	err := rt.RunCommandLogStream("gb", []string{"build"})
	if err != nil {
		fmt.Println(err)
		return err
	}
	if _, err := os.Stat("dist"); os.IsNotExist(err) {
		os.Mkdir("dist", 00750)
	}
	return nil
}
func (s GoBuilder) Dist(project model.Project) error {
	for _, build := range s.Project.Builds {
		if build.Type == "go" {
			packBase := fmt.Sprintf("%s-%s-%s-%s", s.Project.ArtifactId, build.OS, build.Arch, s.Project.Version)
			pack := fmt.Sprintf("%s.%s", packBase, s.Project.Packaging)
			os.Setenv("GOOS", build.OS)
			os.Setenv("GOARCH", build.Arch)
			ct.Foreground(ct.Cyan, true)
			fmt.Println("Building: ", packBase)
			ct.Foreground(ct.Green, false)
			s.Build()
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

			for _, file := range files {
				fmt.Println("Renaming:", file.Name())
				rpl := fmt.Sprintf("-%s-%s", build.OS, build.Arch)
				os.Rename("bin/" + file.Name(), "dist/" + strings.Replace(file.Name(), rpl, "", -1))
			}
			fileUtils.CopyFile("README.md", "dist/README.md")
			fileUtils.CopyFile("turtle.json", "dist/turtle.json")

			compressor.Tar("dist/", pack)
		} else {
			fmt.Println("Skipping non \"go\" build type:", build.Type)
		}
	}
	return nil
}

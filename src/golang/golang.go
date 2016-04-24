package golang

import (
	"model"
	"fmt"
	"os"
	"github.com/daviddengcn/go-colortext"
	"io/ioutil"
	"strings"
	"github.com/pborman/uuid"
	"github.com/correia-io/goutils/src/utils"
)

type GolangPlugin interface {

}

type GoBuilder struct  {
	Project model.Project
}

func (s GoBuilder) Dist() {
	if (s.Project.ProjectType == "go") {
		for _, build := range s.Project.Builds {
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
				logger.Error(err, "Error searching for binaries. ")
				ct.Foreground(ct.Red,false)
				fmt.Println("Perhaps the project might not generate an executable")
				fmt.Println("If you think is a good idea to create package for libraries only, please contact the author or send a pull request at https://github.com/marcelocorreia/turtle")
				ct.Foreground(ct.White,false)
				os.Exit(1)
			}

			for _, file := range files {
				fmt.Println("Renaming:", file.Name())
				rpl := fmt.Sprintf("-%s-%s", build.OS, build.Arch)
				os.Rename("bin/" + file.Name(), "dist/" + strings.Replace(file.Name(), rpl, "", -1))
			}
			fileUtils.CopyFile("README.md","dist/README.md")
			fileUtils.CopyFile("turtle.json","dist/turtle.json")

			compressor.Tar("dist/", pack)
		}
	} else if s.Project.ProjectType == "static" {
		fmt.Println("Packaging Static Project")
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

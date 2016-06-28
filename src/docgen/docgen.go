package docgen

import (
	"os"
	"path/filepath"
	"strings"
	"io/ioutil"
	"fmt"
	"sort"
)

type DocHelper interface {

}

type DocGen struct{}

func (dg DocGen) ProcessMarkdown(path string) (error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	list:=[]string{}
	lines := strings.Split(string(file), "\n")
	for _,line := range lines {
		if strings.HasPrefix(line, "###") {
			list = append(list,line)
		}
	}

	sort.Strings(list)
	for _,item := range list {
		fmt.Println(item)
	}
	return nil
}

func (dg DocGen) ScanMarkdownRecursive(path string) ([]string) {
	resp := []string{}
	fileList := []string{}

	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	for _, file := range fileList {
		if strings.HasSuffix(strings.ToLower(file), "md") {
			resp = append(resp, file)
		}
	}
	return resp
}
package utils

import (
	"os"
	"fmt"
	"io"
	"io/ioutil"
)

type FileUtil interface {
	CopyDir(source string, dest string) (err error)
	Exists(path string) (bool, error)
	ListDir(dir string) ([]os.FileInfo)
	ListDirWithExceptions(dir string, exceptions []string) ([]os.FileInfo)
	CopyFile(source string, dest string) (err error)
}

type FileUtils struct{}

func (f FileUtils) CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

func (f FileUtils) ListDir(dir string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []os.FileInfo{}, err
	}
	return files, nil
}

func (f FileUtils) ListDirWithExceptions(dir string, exceptions []string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {

		return []os.FileInfo{}, err
	}

	for _, file := range files {
		if file.IsDir() {
			if !StringInSlice(file.Name(), exceptions) {
				fmt.Println(file.Name())
			}
		}
	}

	return files, nil
}

func (f FileUtils) CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = f.CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = f.CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func (f FileUtils) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}


package utils

import (
	"os"

	"logger"
	"runtime-helper"
)

var (
	lg = logger.GetLogger()
)

func Tar(source, target string) error {
	rt := runtime.RuntimeHelper{}
	lg.Debugf("Archiving: %s as %s", source, target)
	args := []string{"-cvzf", target, source}
	err := rt.RunCommandLogStream("tar", args)
	if err != nil {
		return err
	}
	return nil
}

func Untar(tarball, targetDir string) error {
	rt := runtime.RuntimeHelper{}
	os.Chdir(targetDir)
	lg.Debugf("Extracting: %s as %s", tarball)
	args := []string{"-xvzf", targetDir}
	err := rt.RunCommandLogStream("tar", args)
	if err != nil {
		return err
	}
	return nil
}
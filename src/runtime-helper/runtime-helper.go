package runtime

import (
	"os/exec"
	"bytes"
	"os"

)

type Runtime interface {
	RunCommand(name string, arg ...string) string
	CheckBinaryInPath(binary string) bool
}


type RuntimeHelper struct {

}

func (r RuntimeHelper) RunCommand(command string, arg []string) (string, error) {
	cmd := exec.Command(command, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), err

}
func (r RuntimeHelper) CheckBinaryInPath(binary string) bool {
	_, err := exec.LookPath(binary)
	if err != nil {
		return false
	}
	return true
}
func (r RuntimeHelper) RunCommandLogStream(command string, arg []string) (error) {
	cmd := exec.Command(command, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}
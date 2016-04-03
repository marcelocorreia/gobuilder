package runtime

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestRunCommand(t *testing.T) {
	r := RuntimeHelper{}
	args := []string{"-ls"}
	resp, e := r.RunCommand("ls", args)
	fmt.Println(resp, e)
	assert.NotEmpty(t, resp)
}

func TestCheckBinaryInPath(t *testing.T) {
	r := RuntimeHelper{}
	assert.True(t, r.CheckBinaryInPath("ls"))
	assert.False(t, r.CheckBinaryInPath("dfghdhnedtumdfychb56urth45bertaw34bt "))
}
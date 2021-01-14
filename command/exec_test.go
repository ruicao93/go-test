package command

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestCommandError(t *testing.T) {
	cmd := exec.Command("powershell.exe", "error-ls")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	stdout, err := cmd.Output()
	t.Logf("stdout: %s\n", string(stdout))
	//t.Logf("err: %v\n", err)
	//t.Logf("err: %s\n", err)
	t.Logf("stderr: %s", string(stderr.Bytes()))
	t.Logf("err: %s\n", string(err.(*exec.ExitError).Stderr))
}

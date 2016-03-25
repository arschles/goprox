package githttp

import (
	"fmt"
	"os/exec"
	"strings"
)

type cmdErr struct {
	cmd *exec.Cmd
	out []byte
	err error
}

func (c *cmdErr) Error() string {
	return fmt.Sprintf("error running command %s in %s: %s (%s)", strings.Join(c.cmd.Args, " "), c.cmd.Dir, string(c.out), c.err)
}

func newCmdErr(cmd *exec.Cmd, out []byte, err error) *cmdErr {
	return &cmdErr{cmd: cmd, out: out, err: err}
}

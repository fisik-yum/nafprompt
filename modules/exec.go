package modules

import (
	"os/exec"
	"strings"
)

type execModule struct {
}

func (e execModule) data(args []string) string {
	insplit := strings.Split(args[1], " ")
	d := exec.Command(insplit[0], insplit[1:]...)
	if d.ProcessState.ExitCode() != -1 {
		return ""
	}
	output, _ := d.Output()
	return strings.ReplaceAll(args[0], "%s", strings.TrimRight(string(output), "\r\n"))
}

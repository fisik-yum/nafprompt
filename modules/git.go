/*
nafprompt - a simple bash prompt
Copyright (C) 2022  fisik_yum
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package modules

import (
	"os/exec"
	"strings"
)

type gitModule struct {
	branch string //git rev-parse --abbrev-ref HEAD
	commit string //git rev-parse --short HEAD
}

func (g gitModule) data(args []string) string {
	g.getBranch()
	g.getCommit()
	if g.branch == "" || g.commit == "" {
		return ""
	}
	if args[0] == "" {
		args = make([]string, 0)
	}
	for len(args) < 1 {
		args = append(args, "%b %c")
	}
	fstring := strings.ReplaceAll(strings.ReplaceAll(args[0], "%b", g.branch), "%c", g.commit)
	return fstring

}

func (g *gitModule) getBranch() bool {
	b := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	if b.ProcessState.ExitCode() != -1 {
		g.branch = ""
		return false
	}
	output, _ := b.Output()
	g.branch = strings.TrimRight(string(output), "\r\n") //string(output)
	return true

}

func (g *gitModule) getCommit() bool {
	c := exec.Command("git", "rev-parse", "--short", "HEAD")
	if c.ProcessState.ExitCode() != -1 {
		g.commit = ""
		return false
	}
	output, _ := c.Output()
	g.commit = strings.TrimRight(string(output), "\r\n") //string(output)
	return true

}

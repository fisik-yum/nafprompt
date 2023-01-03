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

type goModule struct{}

func (g goModule) data(args []string) string {
	c, _ := exec.Command("go", "list", "-m").Output()
	if strings.HasPrefix(string(c), "command-line-arguments") {
		return ""
	}
	if args[0] == "" {
		args = make([]string, 0)
	}
	for len(args) < 1 {
		args = append(args, "%s")
	}
	return strings.ReplaceAll(args[0], "%s", strings.TrimRight(string(c), "\r\n"))
	//return strings.TrimRight(string(c), "\r\n")

}

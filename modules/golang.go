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
	"fmt"
	"os/exec"
	"strings"
)

func go_mod(opts []string) string {
	c, _ := exec.Command("go", "list", "-m").Output()
	if strings.HasPrefix(string(c), "command-line-arguments") {
		return ""
	}
	return fmt.Sprintf(opts[0], strings.TrimRight(string(c), "\r\n"))

}

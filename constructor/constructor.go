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
package constructor

import (
	"nafprompt/modules"
	"strings"
)

func Lookup(keyword string) string {
	return replaceTree(keyword)
}

func pmodname(s string) string {
	r := []rune(s)
	for i := range r {
		if r[i] == ':' && i != 0 {
			return string(r[0 : i+1])
		}
	}
	return ""
}

func replaceTree(kw string) string { //keyword replacement logic
	if strings.HasPrefix(kw, ":") { //now we only have one janky branch for module logic
		modname := pmodname(kw)
		return modules.Request(modname[1:len(modname)-1], strings.Split(kw[len(modname):], ";"))
	}
	switch kw {
	case "user":
		return `\u`
	case "host":
		return `\h`
	case "date":
		return `\d`
	case "device":
		return `\l`
	case "shellname":
		return `\s`
	case "time24":
		return `\t`
	case "time12":
		return `\T`
	case "time12m":
		return `\@`
	case "cwd":
		return `\w`
	case "basename}":
		return `\W`
	case "cmdnum":
		return `\#`

	default:
		return kw
	}
}

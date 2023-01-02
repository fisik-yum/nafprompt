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
	"nafprompt/effects"
	"nafprompt/modules"
	"regexp"
	"strings"
)

func Parse_string(prompt string) string {
	pwords := pkeywords(prompt)
	for x := range pwords {
		a := decide_replacement(pwords[x], replaceTree)
		prompt = strings.ReplaceAll(prompt, pwords[x], a)
	}
	return prompt
}

func decide_replacement(keyword string, repfunc func(k string) string) string {
	return repfunc(keyword)
}

func pkeywords(s string) []string {
	re := regexp.MustCompile(`\{[^}]*\}`) //keyword enclosed in curly braces
	return re.FindAllString(s, -1)
}

func replaceTree(kw string) string { //keyword replacement logic
	if strings.HasPrefix(kw, "{text:") {
		return EscSeq(strings.ToLower(kw[6 : len(kw)-1]))
	} else if strings.HasPrefix(kw, "{module:") {
		return modules.RequestModuleData(strings.Split(strings.ToLower(kw[8:len(kw)-1]), ";"))
	}
	switch kw {
	case "{user}":
		return `\u`
	case "{host}":
		return `\h`
	case "{date}":
		return `\d`
	case "{device}":
		return `\l`
	case "{shellname}":
		return `\s`
	case "{time24}":
		return `\t`
	case "{time12}":
		return `\T`
	case "{time12m}":
		return `\@`
	case "{cwd}":
		return `\w`
	case "{basename}":
		return `\W`
	case "{cmdnum}":
		return `\#`

	default:
		/*a := strings.Fields(kw[1 : len(kw)-1])
		c, _ := exec.Command(a[0], a[1:]...).Output()
		return strings.TrimRight(string(c), "\r\n")*/
		return kw
	}
}

func EscSeq(option string) string { //color syntax is {text:fgcolor;bgcolor;formatting}
	l := strings.Split(option, ";")

	for len(l) < 3 {
		l = append(l, "")
	}

	return effects.CreateCode(effects.LookUp(l[0], l[1], l[2]))

}

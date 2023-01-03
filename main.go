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

package main

import (
	"nafprompt/constructor"
	"os"
	"strings"
)

var prompt string

func init() {
	prompt = os.Getenv("PROMPT")
	if strings.ReplaceAll(prompt, " ", "") == "" {
		prompt = "{:text:red;;bold}{user}@{host}{:text:green;;bold} {:go:(go: %s) }{:text:cyan;;bold}{:git:(git: %b) }{:text:blue;;bold}/#/ {:text:;;}"
	}
}

func main() {
	os.Stdout.Write([]byte(constructor.Parse_string(prompt)))
}

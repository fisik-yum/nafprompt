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
	"nafprompt/parser3"
	"os"
)

var prompt string

func init() {
	prompt = os.Getenv("PROMPT")
}

func main() {
	a, _ := parser3.Parse([]rune(prompt), 0)
	os.Stdout.Write([]byte(a))
}

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

var modlist = make(map[string]module)

type module interface {
	data(args []string) string
}

// modules are compiled into the binary, and must be registered using the function register(name, module)
// the name must be unique as it is used as an identifier to request information
func init() {
	register("go", goModule{})
	register("text", textModule{})
	register("git", gitModule{})
}

func Request(id string, opts []string) string {
	return modlist[id].data(opts)
}

func register(id string, mod module) {
	modlist[id] = mod
}

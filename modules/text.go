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

import "nafprompt/escapes"

type textModule struct{}

func (t textModule) data(opts []string) string {
	for len(opts) < 3 {
		opts = append(opts, "")
	}

	return escapes.CreateCode(escapes.LookUp(opts[0], opts[1], opts[2]))

}

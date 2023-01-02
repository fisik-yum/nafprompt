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
package effects

import "fmt"

// formatting
const (
	Default = iota
	Bold
	Dim
	_
	Underlined
	Blink
	_
	Invert
	Hidden
)

// foreground (text) colors
const (
	FGBlack = iota + 30
	FGRed
	FGGreen
	FGYellow
	FGBlue
	FGMagenta
	FGCyan
	FGLGrey
	_
	FGDefault
)

const (
	FGGrey = iota + 90
	FGLRed
	FGLGreen
	FGLYellow
	FGLBlue
	FGLMagenta
	FGLCyan
	FGWhite
)

//background colors

const (
	BGBlack = iota + 40
	BGRed
	BGGreen
	BGYellow
	BGBlue
	BGMagenta
	BGCyan
	BGLGrey
	_
	BGDefault
)

const (
	BGGrey = iota + 100
	BGLRed
	BGLGreen
	BGLYellow
	BGLBlue
	BGLMagenta
	BGLCyan
	BGWhite
)

func lookFG(name string) int {
	switch name {
	default:
		return FGDefault
	case "black":
		return FGBlack
	case "red":
		return FGRed
	case "green":
		return FGGreen
	case "yellow":
		return FGYellow
	case "blue":
		return FGBlue
	case "magenta":
		return FGMagenta
	case "cyan":
		return FGCyan
	case "grey":
		return FGGrey

	case "lred":
		return FGLRed
	case "lgreen":
		return FGLGreen
	case "lyellow":
		return FGLYellow
	case "lblue":
		return FGLBlue
	case "lmagenta":
		return FGLMagenta
	case "lcyan":
		return FGLCyan
	case "lgrey":
		return FGLGrey
	case "white":
		return FGWhite
	}
}

func lookBG(name string) int {
	switch name {
	default:
		return BGDefault
	case "black":
		return BGBlack
	case "red":
		return BGRed
	case "green":
		return BGGreen
	case "yellow":
		return BGYellow
	case "blue":
		return BGBlue
	case "magenta":
		return BGMagenta
	case "cyan":
		return BGCyan
	case "grey":
		return BGGrey

	case "lred":
		return BGLRed
	case "lgreen":
		return BGLGreen
	case "lyellow":
		return BGLYellow
	case "lblue":
		return BGLBlue
	case "lmagenta":
		return BGLMagenta
	case "lcyan":
		return BGLCyan
	case "lgrey":
		return BGLGrey
	case "white":
		return BGWhite
	}
}

func lookStyle(name string) int {
	switch name {
	default:
		return Default
	case "bold":
		return Bold
	case "dim":
		return Dim
	case "underline":
		return Underlined
	case "blink":
		return Blink
	case "invert":
		return Invert
	case "hide":
		return Hidden
	}

}

func LookUp(fg string, bg string, style string) (int, int, int) {
	return lookFG(fg), lookBG(bg), lookStyle(style)

}
func CreateCode(fg int, bg int, style int) string {
	return fmt.Sprintf(`\[\e[%d;%d;%dm\]`, style, fg, bg)
}

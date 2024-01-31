package parser3

import "nafprompt/constructor"

func Parse(p []rune, ind int) (string, int) {
	var istr string
	for ; ind < len(p); ind++ {
		if p[ind] == '{' && ind != 0 {
			ind++
			var a string
			a, ind = Parse(p, ind)
			istr += a
		} else if p[ind] == '}' && ind != 0 {
			return constructor.Lookup(istr), ind
		} else {
			istr += string(p[ind])
		}
	}
	return istr, ind
}

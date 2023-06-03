package parser2

import (
	"log"
	"nafprompt/constructor"
)

func Parse(p []rune) []rune {
	fp := make([]rune, 0)
	ta := make([]rune, 0)
    var ind int
	for i :=0; i<len(p);i++ {
		if p[i] == '{' && i !=0{
			ta, ind = parseinside(p[i:]) //fix i+1
			i=i+ind
			fp = append(fp, ta...)
		} else {
			fp = append(fp, p[i])
		}
	}
	return fp
}

// parses inside template
// returns processed template and where to pick up from.
func parseinside(p []rune) ([]rune, int) { //parsed inside length
	fp := make([]rune, 0)
	t := make([]rune, 0)
    var ind int
	for i :=0; i<len(p);i++ {
		//t = ta
		if p[i] == '{' && i !=0{
			t, ind = parseinside(p[i:]) //fix i+1
			i=i+ind
            fp = append(fp, t...)
		} else if p[i] == '}' {
			//look up assignment
			fp = append(fp, p[i])
            t = append(t, []rune(constructor.Lookup(string(fp)))...)
			return t, i
		} else {
			fp = append(fp, p[i])
		}
	}
	log.Fatal("ERROR: No closing brace!")
	return p, 0
}


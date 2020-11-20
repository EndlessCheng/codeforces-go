package main

import "strings"

// null => nil
// while => for
// remove () after for and if
// .xxx => .Xxx
// type a = val  =>  a := val
// todo type a = val, b = val2  =>  a, b := val, val2
func transJava(code string) string {
	code = strings.ReplaceAll(code, "null", "nil")
	code = strings.ReplaceAll(code, "while", "for")

	code = strings.ReplaceAll(code, "for (", "for ")
	code = strings.ReplaceAll(code, "if (", "if ")
	code = strings.ReplaceAll(code, ") {", " {")

	code = strings.ReplaceAll(code, ".next", ".Next")
	code = strings.ReplaceAll(code, ".left", ".Left")
	code = strings.ReplaceAll(code, ".right", ".Right")
	code = strings.ReplaceAll(code, ".val", ".Val")

	lines := strings.Split(code, "\n")
	for i, s := range lines {
		// todo multi assign
		sp := strings.SplitN(s, " = ", 2)
		if len(sp) == 1 {
			continue
		}
		h := sp[0]
		for j := len(h) - 2; j > 0; j-- {
			if h[j] == ' ' && h[j-1] != ' ' && h[j+1] != ' ' {
				k := j - 1
				for ; k >= 0 && h[k] != ' '; k-- {
				}
				sp[0] = h[:k+1] + h[j+1:]
				lines[i] = strings.Join(sp, " := ")
				break
			}
		}
	}
	code = strings.Join(lines, "\n")
	return code
}

package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf858C(in io.Reader, out io.Writer) {
	s := ""
	Fscan(in, &s)
	st, cnt, diff := 0, 0, false
	for i, b := range s {
		if 0x208222>>(b&31)&1 > 0 {
			cnt, diff = 0, false
			continue
		}
		cnt++
		if cnt > 1 && s[i] != s[i-1] {
			diff = true
		}
		if diff && cnt > 2 {
			Fprint(out, s[st:i], " ")
			st = i
			cnt, diff = 1, false
		}
	}
	Fprint(out, s[st:])
}

//func main() { cf858C(os.Stdin, os.Stdout) }

package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func CF489C(in io.Reader, out io.Writer) {
	var n, s int
	Fscan(in, &n, &s)
	if n == 1 {
		if s > 9 {
			Fprint(out, -1, -1)
		} else {
			Fprint(out, s, s)
		}
		return
	}
	if s == 0 || s > n*9 {
		Fprint(out, -1, -1)
		return
	}
	maxS := strings.Repeat("9", s/9)
	if s%9 > 0 {
		maxS += string('0' + byte(s%9))
	}
	maxS += strings.Repeat("0", n-len(maxS))
	var minS string
	if (s+8)/9 == n {
		minS = string('0'+byte(s-(n-1)*9)) + strings.Repeat("9", n-1)
	} else {
		s--
		minS = "1" + strings.Repeat("0", n-2-s/9) + string('0'+byte(s%9)) + strings.Repeat("9", s/9)
	}
	Fprintln(out, minS, maxS)
}

//func main() { CF489C(os.Stdin, os.Stdout) }

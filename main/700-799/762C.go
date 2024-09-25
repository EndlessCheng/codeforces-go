package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf762C(in io.Reader, out io.Writer) {
	var s, t string
	Fscan(in, &s, &t)
	n, m := len(s), len(t)
	suf := make([]int, n+1)
	suf[n] = m
	for i, j := n-1, m-1; i >= 0; i-- {
		if s[i] == t[j] {
			j--
		}
		if j < 0 {
			Fprint(out, t)
			return
		}
		suf[i] = j + 1
	}

	l, r := 0, suf[0]
	j := 0
	for i := range s {
		if s[i] == t[j] {
			j++
			if suf[i+1]-j < r-l {
				l, r = j, suf[i+1]
			}
		}
	}
	if r-l == m {
		Fprint(out, "-")
	} else {
		Fprint(out, t[:l], t[r:])
	}
}

//func main() { cf762C(bufio.NewReader(os.Stdin), os.Stdout) }

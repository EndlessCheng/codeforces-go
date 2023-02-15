package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func f95(s string) (ans int) {
	c, i1, neg := 0, -1, false
	for i, b := range s {
		if b == '(' {
			c++
		} else {
			c--
		}
		if c < 0 {
			if c < -2 {
				return
			}
			if i1 < 0 && c == -1 {
				i1 = i
			}
			neg = true
		}
	}
	if c == -2 {
		return strings.Count(s[:i1+1], ")")
	}
	if c != 2 || neg {
		return
	}
	for i := len(s) - 1; c > 1; i-- {
		if s[i] == '(' {
			ans++
			c--
		} else {
			c++
		}
	}
	return
}

func CF1095E(in io.Reader, out io.Writer) {
	var s string
	Fscan(bufio.NewReader(in), &s, &s)
	Fprint(out, f95(s))
}

//func main() { CF1095E(os.Stdin, os.Stdout) }

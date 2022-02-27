package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF214B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	c := [10]int{}
	var n, v, s int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		c[v]++
		s += v
	}
	if c[0] == 0 {
		Fprint(out, -1)
		return
	}
	if s%3 > 0 {
		v1, v2 := []int{1, 4, 7}, []int{2, 5, 8}
		if s%3 == 2 {
			v1, v2 = v2, v1
		}
		for _, v := range v1 {
			if c[v] > 0 {
				c[v]--
				s -= v
				goto print
			}
		}
		for i, v := range v2 {
			for _, w := range v2[:i] {
				if c[v] > 0 && c[w] > 0 {
					c[v]--
					c[w]--
					s -= v + w
					goto print
				}
			}
			if c[v] > 1 {
				c[v] -= 2
				s -= v * 2
				goto print
			}
		}
		s = 0
	}
print:
	if s == 0 {
		Fprint(out, 0)
		return
	}
	for i := 9; i >= 0; i-- {
		Fprint(out, strings.Repeat(string('0'+byte(i)), c[i]))
	}
}

//func main() { CF214B(os.Stdin, os.Stdout) }

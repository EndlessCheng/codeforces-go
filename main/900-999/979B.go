package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF979B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var k, ans, ansI, cntAns int
	var s string
	Fscan(in, &k)
	for i := 0; i < 3; i++ {
		Fscan(in, &s)
		n := len(s)
		c := ['z' + 1]int{}
		for _, b := range s {
			c[b]++
		}
		mx := 0
		if c[s[0]] < n {
			for _, c := range c {
				if c > mx {
					mx = c
				}
			}
			if mx+k < n {
				mx += k
			} else {
				mx = n
			}
		} else if k == 1 && n > 1 {
			mx = n - 1
		} else {
			mx = n
		}
		if mx > ans {
			ans, ansI, cntAns = mx, i, 1
		} else if mx == ans {
			cntAns++
		}
	}
	if cntAns > 1 {
		Fprint(out, "Draw")
	} else {
		Fprint(out, []string{"Kuro", "Shiro", "Katie"}[ansI])
	}
}

//func main() { CF979B(os.Stdin, os.Stdout) }

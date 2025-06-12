package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf526E(in io.Reader, out io.Writer) {
	var n, q, mx int
	Fscan(in, &n, &q)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
	}
	tot := s[n]
	f := make([]struct{ f, st int }, n+1)
	for range q {
		Fscan(in, &mx)
		if tot <= mx {
			Fprintln(out, 1)
			continue
		}
		left := 0
		for i := 1; ; i++ {
			for s[i]-s[left] > mx {
				left++
			}
			f[i].f = f[left].f + 1
			if f[i].f == 2 {
				f[i].st = left
			} else {
				f[i].st = f[left].st
			}
			if tot-(s[i]-s[f[i].st]) <= mx {
				Fprintln(out, max(f[i].f, 2))
				break
			}
		}
	}
}

//func main() { cf526E(bufio.NewReader(os.Stdin), os.Stdout) }

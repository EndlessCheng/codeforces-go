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
	f := make([]int, n+1)
	st := make([]int, n+1)
	for range q {
		Fscan(in, &mx)
		if tot <= mx {
			Fprintln(out, 1)
			continue
		}
		l := 0
		for i := 1; ; i++ {
			for s[i]-s[l] > mx {
				l++
			}
			f[i] = f[l] + 1
			if f[i] == 2 {
				st[i] = l
			} else {
				st[i] = st[l]
			}
			if tot-(s[i]-s[st[i]]) <= mx {
				Fprintln(out, max(f[i], 2))
				break
			}
		}
	}
}

//func main() { cf526E(bufio.NewReader(os.Stdin), os.Stdout) }

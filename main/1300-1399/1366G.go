package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1366G(in io.Reader, out io.Writer) {
	var s, t string
	Fscan(in, &s, &t)
	n, m := len(s), len(t)
	s = " " + s
	t = " " + t
	nxt := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if s[i] == '.' {
			d := -1
			for j := i - 1; j > 0; j-- {
				if s[j] == '.' {
					d--
				} else {
					d++
				}
				if d == 0 {
					nxt[i] = j
					break
				}
			}
		}
	}

	f := make([]int, n+1)
	for i := range f {
		f[i] = 1e9
	}

	pre := make([]int, n+1)
	for j := 0; j <= m; j++ {
		pre[0] = f[0]
		if j != 0 {
			f[0] = 1e9
		} else {
			f[0] = 0
		}
		for i := 1; i <= n; i++ {
			pre[i] = f[i]
			if s[i] == t[j] {
				f[i] = pre[i-1]
			} else {
				f[i] = 1e9
			}
			if nxt[i] != 0 {
				f[i] = min(f[i], f[nxt[i]-1])
			}
			f[i] = min(f[i], f[i-1]+1)
		}
	}
	Fprint(out, f[n])
}

//func main() { cf1366G(bufio.NewReader(os.Stdin), os.Stdout) }

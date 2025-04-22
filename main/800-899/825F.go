package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf825F(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	n := len(s)
	sz := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sz[i] = sz[i/10] + 1
	}

	f := make([]int, n+1)
	pi := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		t := s[i:]
		m := len(t)
		f[i] = min(m+1, f[i+1]+2)
		cnt := 0
		for j := 1; j < m; j++ {
			b := t[j]
			for cnt > 0 && t[cnt] != b {
				cnt = pi[cnt-1]
			}
			if t[cnt] == b {
				cnt++
			}
			pi[j] = cnt
			if k := j + 1 - cnt; cnt > 0 && (j+1)%k == 0 {
				f[i] = min(f[i], f[i+j+1]+sz[(j+1)/k]+k)
			} else {
				f[i] = min(f[i], f[i+j+1]+j+2)
			}
		}
	}
	Fprint(out, f[0])
}

//func main() { cf825F(bufio.NewReader(os.Stdin), os.Stdout) }

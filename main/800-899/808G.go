package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf808G(in io.Reader, out io.Writer) {
	var s, t []byte
	Fscan(in, &s, &t)
	n, m := len(s), len(t)

	pi := make([]int, m)
	match := 0
	for i := 1; i < m; i++ {
		v := t[i]
		for match > 0 && t[match] != v {
			match = pi[match-1]
		}
		if t[match] == v {
			match++
		}
		pi[i] = match
	}

	preMax := make([]int, n+1)
	f := make([]int, n+1)
o:
	for i := m; i <= n; i++ {
		preMax[i] = preMax[i-1]
		for j, b := range s[i-m : i] {
			if b != '?' && b != t[j] {
				continue o
			}
		}
		f[i] = preMax[i-m] + 1
		for j := pi[m-1]; j > 0; j = pi[j-1] {
			f[i] = max(f[i], f[i-m+j]+1)
		}
		preMax[i] = max(preMax[i], f[i])
	}
	Fprint(out, preMax[n])
}

//func main() { cf808G(bufio.NewReader(os.Stdin), os.Stdout) }

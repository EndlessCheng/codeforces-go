package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf607C(in io.Reader, out io.Writer) {
	var n, c int
	var s, t []byte
	Fscan(in, &n, &s, &t)
	slices.Reverse(t)
	for i, b := range t {
		if b == 'N' || b == 'S' {
			t[i] ^= 'N' ^ 'S'
		} else {
			t[i] ^= 'E' ^ 'W'
		}
	}
	t = append(t, '#')
	t = append(t, s...)

	m := len(t)
	pi := make([]int, m)
	for i := 1; i < m; i++ {
		b := t[i]
		for c > 0 && t[c] != b {
			c = pi[c-1]
		}
		if t[c] == b {
			c++
		}
		pi[i] = c
	}
	if pi[m-1] > 0 {
		Fprint(out, "NO")
	} else {
		Fprint(out, "YES")
	}
}

//func main() { cf607C(bufio.NewReader(os.Stdin), os.Stdout) }

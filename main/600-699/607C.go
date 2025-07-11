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

	pi := make([]int, len(t))
	for i := 1; i < len(pi); i++ {
		b := t[i]
		for c > 0 && t[c] != b {
			c = pi[c-1]
		}
		if t[c] == b {
			c++
		}
		pi[i] = c
	}
	if pi[len(pi)-1] > 0 {
		Fprintln(out, "NO")
	} else {
		Fprintln(out, "YES")
	}
}

//func main() { cf607C(bufio.NewReader(os.Stdin), os.Stdout) }

package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1644E(in io.Reader, out io.Writer) {
	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		base := s[0]
		ans := n * n
		h := 0
		for _, b := range s {
			if b != base {
				h++
			} else {
				ans -= h
			}
		}

		m := len(s)
		h = 0
		for i := m - 1; i >= 0; i-- {
			if s[i] != base {
				h++
			} else {
				ans -= h
			}
		}

		i := 0
		for i < m && s[i] == base {
			i++
		}
		if i < m {
			Fprintln(out, ans-(n-1-h)*i)
		} else {
			Fprintln(out, n)
		}
	}
}

//func main() { cf1644E(bufio.NewReader(os.Stdin), os.Stdout) }

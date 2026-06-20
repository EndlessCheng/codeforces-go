package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1658F(in io.Reader, out io.Writer) {
	var T, n, m int
	var t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &t)
		t = " " + t + t
		s := make([]int, 2*n)
		for i := 1; i < 2*n; i++ {
			s[i] = s[i-1]
			if t[i] == '1' {
				s[i]++
			}
		}

		tar := m * s[n]
		if tar%n != 0 {
			Fprintln(out, -1)
			continue
		}
		tar /= n

		for i := m; i < 2*n; i++ {
			if s[i]-s[i-m] != tar {
				continue
			}
			if i <= n {
				Fprintln(out, 1)
				Fprintln(out, i-m+1, i)
			} else {
				Fprintln(out, 2)
				Fprintln(out, 1, i-n)
				Fprintln(out, i-m+1, n)
			}
			break
		}
	}
}

//func main() { cf1658F(bufio.NewReader(os.Stdin), os.Stdout) }

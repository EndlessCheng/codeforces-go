package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2189E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ss := 0
		fa := false
		fb := false
		for i := 0; i < n; i++ {
			if s[i] == '1' {
				ss++
			} else {
				ss--
			}
			if ss > 0 {
				fa = true
			}
			if ss >= 0 {
				fb = true
			}
			if i > 0 && s[i] == '1' && s[i-1] == '1' {
				fb = true
			}
		}

		ss = 0
		for i := n - 1; i >= 0; i-- {
			if s[i] == '1' {
				ss++
			} else {
				ss--
			}
			if ss > 0 {
				fa = true
			}
			if ss >= 0 {
				fb = true
			}
		}

		if ss == -n {
			Fprintln(out, -1)
		} else if n == 1 && s[0] == '1' {
			Fprintln(out, 0)
		} else if ss >= 0 {
			Fprintln(out, n)
		} else if fa || ss == -1 {
			Fprintln(out, n+1)
		} else if fb {
			Fprintln(out, n+2)
		} else {
			Fprintln(out, n+3)
		}
	}
}

//func main() { cf2189E(bufio.NewReader(os.Stdin), os.Stdout) }

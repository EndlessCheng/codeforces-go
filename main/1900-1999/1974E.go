package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1974E(in io.Reader, out io.Writer) {
	var T, n, x, c, h int
	f := [5e4 + 1]int{}
	for i := 1; i < len(f); i++ {
		f[i] = 1e18
	}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		s := 0
		for i := 0; i < n; i++ {
			Fscan(in, &c, &h)
			s += h
			for j := s; j >= h; j-- {
				if f[j-h] <= i*x-c {
					f[j] = min(f[j], f[j-h]+c)
				}
			}
		}
		for i := s; i >= 0; i-- {
			if f[i] < 1e18 {
				Fprintln(out, i)
				break
			}
		}
		for i := s; i > 0; i-- {
			f[i] = 1e18
		}
		f[0] = 0
	}
}

//func main() { cf1974E(bufio.NewReader(os.Stdin), os.Stdout) }

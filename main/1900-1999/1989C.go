package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1989C(in io.Reader, out io.Writer) {
	var T, n, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		var s0, s1, inc, dec int
		for _, v := range a {
			Fscan(in, &w)
			if v == w {
				if v == 1 {
					inc++
				} else {
					dec -= v
				}
			} else if v > w {
				s0 += v
			} else {
				s1 += w
			}
		}
		if s0 > s1 {
			s0, s1 = s1, s0
		}
		if s0+inc <= s1-dec {
			Fprintln(out, s0+inc)
		} else {
			Fprintln(out, (s0+s1+inc-dec)>>1)
		}
	}
}

//func main() { cf1989C(bufio.NewReader(os.Stdin), os.Stdout) }

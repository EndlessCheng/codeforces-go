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
			continue
		}
		newS1 := max(s0, s1-dec)
		dec -= s1 - newS1
		if inc > dec {
			Fprintln(out, newS1+(inc-dec-newS1+s0)/2)
		} else {
			Fprintln(out, s0-(dec-inc+1)/2)
		}
	}
}

//func main() { cf1989C(bufio.NewReader(os.Stdin), os.Stdout) }

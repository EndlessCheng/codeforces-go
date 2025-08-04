package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2014D(in io.Reader, out io.Writer) {
	var T, n, d, k, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &d, &k)
		diff := make([]int, n+1)
		for range k {
			Fscan(in, &l, &r)
			diff[l]++
			if r+d <= n {
				diff[r+d]--
			}
		}

		mx, mxI := -1, 0
		mn, mnI := k+1, 0
		s := 0
		for i := 1; i <= n; i++ {
			s += diff[i]
			l := i + 1 - d
			if l <= 0 {
				continue
			}
			if s > mx {
				mx, mxI = s, l
			}
			if s < mn {
				mn, mnI = s, l
			}
		}
		Fprintln(out, mxI, mnI)
	}
}

//func main() { cf2014D(bufio.NewReader(os.Stdin), os.Stdout) }

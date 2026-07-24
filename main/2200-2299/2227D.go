package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2227D(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n*2)
		pos0 := []int{}
		for i := range a {
			Fscan(in, &a[i])
			if a[i] == 0 {
				pos0 = append(pos0, i)
			}
		}

		f := func(i, j int) (mex int) {
			has := make([]bool, n+1)
			for i >= 0 && j < n*2 && a[i] == a[j] {
				has[a[i]] = true
				i--
				j++
			}

			for has[mex] {
				mex++
			}
			return
		}

		p, q := pos0[0], pos0[1]
		Fprintln(out, max(f(p, p), f(q, q), f((p+q)/2, (p+q+1)/2)))
	}
}

//func main() { cf2227D(bufio.NewReader(os.Stdin), os.Stdout) }

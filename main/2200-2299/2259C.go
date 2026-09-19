package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2259C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		var firstN1, last1, p, q int
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			if firstN1 == 0 && a[i] < 0 {
				firstN1 = i
			}
			if a[i] != 0 {
				j := last1
				if j == 0 {
					j = firstN1
				}
				if j > 0 && i-j >= q-p {
					p, q = j, i
				}
			}
			if a[i] > 0 {
				last1 = i
			}
		}

		for i, v := range a {
			if v >= 0 {
				continue
			}
			if i == p || i == q {
				a[i] = 1
			} else {
				a[i] = 0
			}
		}
		for _, v := range a[1:] {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2259C(bufio.NewReader(os.Stdin), os.Stdout) }

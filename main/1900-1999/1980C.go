package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1980C(in io.Reader, out io.Writer) {
	var T, n, m, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		left := 0
		need := map[int]int{}
		has := map[int]bool{}
		for _, v := range a {
			Fscan(in, &w)
			if v != w {
				need[w]++
				left++
			} else {
				has[w] = true
			}
		}

		Fscan(in, &m)
		d := make([]int, m)
		for i := range d {
			Fscan(in, &d[i])
		}

		if need[d[m-1]] == 0 && !has[d[m-1]] {
			Fprintln(out, "NO")
			continue
		}

		for i := m - 1; i >= 0; i-- {
			v := d[i]
			if need[v] > 0 {
				need[v]--
				left--
			}
		}
		if left == 0 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1980C(bufio.NewReader(os.Stdin), os.Stdout) }

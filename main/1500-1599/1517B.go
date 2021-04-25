package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1517B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, m)
			for j := range a[i] {
				Fscan(in, &a[i][j])
			}
			sort.Ints(a[i])
		}
		for j := m; j > 1; j-- {
			p := 0
			for i, r := range a {
				if r[0] < a[p][0] {
					p = i
				}
			}
			a[p] = append(append(append([]int(nil), a[p][1:j]...), a[p][0]), a[p][j:]...)
		}
		for _, r := range a {
			for _, v := range r {
				Fprint(out, v, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { CF1517B(os.Stdin, os.Stdout) }

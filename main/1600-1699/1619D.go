package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1619D(_r io.Reader, _w io.Writer) {
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
		}
		Fprintln(out, sort.Search(1e9, func(low int) bool {
		o:
			for j := range a[0] {
				for _, r := range a {
					if r[j] > low {
						continue o
					}
				}
				return true
			}
			for _, r := range a {
				ok := false
				for _, v := range r {
					if v > low {
						if ok {
							return false
						}
						ok = true
					}
				}
			}
			return true
		}))
	}
}

//func main() { CF1619D(os.Stdin, os.Stdout) }

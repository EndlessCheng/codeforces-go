package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1385C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		Fprintln(out, sort.Search(n-2, func(i int) bool {
			for j, v := n-1, 0; i < j; {
				if a[i] < a[j] {
					if a[i] < v {
						return false
					}
					v = a[i]
					i++
				} else {
					if a[j] < v {
						return false
					}
					v = a[j]
					j--
				}
			}
			return true
		}))
	}
}

//func main() { CF1385C(os.Stdin, os.Stdout) }

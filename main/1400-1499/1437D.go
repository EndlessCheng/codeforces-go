package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1437D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := []int{}
		for i := 1; i < n; {
			st := i
			for i++; i < n && a[i] > a[i-1]; i++ {
			}
			b = append(b, i-st)
		}
		q := []int{b[0]}
		ans := 1
		for i, n := 1, len(b); i < n; {
			ans++
			qq := q
			q = nil
			for _, sz := range qq {
				for ; sz > 0 && i < n; sz-- {
					q = append(q, b[i])
					i++
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1437D(os.Stdin, os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF732D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n, &m)
	d := make([]int, n)
	for i := range d {
		Fscan(in, &d[i])
	}
	a := make([]int, m)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := sort.Search(n, func(mx int) bool {
		pos := make([]int, m)
		for i, v := range d[:mx+1] {
			if v > 0 {
				pos[v-1] = i + 1
			}
		}
		for _, p := range pos {
			if p == 0 {
				return false
			}
		}
		cnt := 0
		for i, v := range d[:mx+1] {
			if v > 0 && i+1 == pos[v-1] {
				if a[v-1] > cnt {
					return false
				}
				cnt -= a[v-1]
			} else {
				cnt++
			}
		}
		return true
	}) + 1
	if ans > n {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF732D(os.Stdin, os.Stdout) }

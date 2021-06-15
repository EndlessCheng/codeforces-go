package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1156E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n+1)
	pos := make([]int, n+1)
	l := make([]int, n+2)
	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		pos[a[i]] = i
		l[i], r[i] = i-1, i+1
	}
	for v := 1; v <= n; v++ {
		p := pos[v]
		r[l[p]] = r[p]
		l[r[p]] = l[p]
	}
	for i := 1; i <= n; i++ {
		if v := a[i]; i-l[i] < r[i]-i {
			for _, w := range a[l[i]+1 : i] {
				if p := pos[v-w]; i < p && p < r[i] {
					ans++
				}
			}
		} else {
			for _, w := range a[i+1 : r[i]] {
				if p := pos[v-w]; l[i] < p && p < i {
					ans++
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1156E(os.Stdin, os.Stdout) }

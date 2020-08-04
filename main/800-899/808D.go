package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF808D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int64, n)
	s := make([]int64, n+1)
	for i := range a {
		Fscan(in, &a[i])
		s[i+1] = s[i] + a[i]
	}
	if s[n]&1 > 0 {
		Fprint(out, "NO")
		return
	}
	mp := map[int64]bool{0: true}
	for i, v := range a {
		mp[v] = true
		if mp[s[i+1]-s[n]/2] {
			Fprint(out, "YES")
			return
		}
	}
	mp = map[int64]bool{0: true}
	for i := len(a) - 1; i >= 0; i-- {
		mp[a[i]] = true
		if mp[s[n]/2-s[i]] {
			Fprint(out, "YES")
			return
		}
	}
	Fprint(out, "NO")
}

//func main() { CF808D(os.Stdin, os.Stdout) }

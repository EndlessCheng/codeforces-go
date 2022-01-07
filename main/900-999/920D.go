package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF920D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const inf int = 1e9

	var n, k, tar, tot, s1 int
	Fscan(in, &n, &k, &tar)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		tot += a[i]
	}
	if tot < tar {
		Fprint(out, "NO")
		return
	}
	if tar%k == 0 {
		Fprintln(out, "YES")
		for i := 2; i <= n; i++ {
			Fprintln(out, inf, i, 1)
		}
		if tar > 0 {
			Fprintln(out, tar/k, 1, 2)
		}
		return
	}

	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, k)
	}
	tarPos1 := 0
	used := make([]bool, n)
	var f func(int, int) bool
	f = func(i, v int) bool {
		if v == tar%k {
			Fprintln(out, "YES")
			tarPos1 = i - 1
			used[i-1] = true
			s1 += a[i-1]
			return true
		}
		if i == n || vis[i][v] {
			return false
		}
		vis[i][v] = true
		if f(i+1, v) {
			return true
		}
		if f(i+1, (v+a[i])%k) {
			if i != tarPos1 {
				used[i] = true
				s1 += a[i]
				Fprintln(out, inf, i+1, tarPos1+1)
			}
			return true
		}
		return false
	}
	if !f(0, 0) {
		Fprint(out, "NO")
		return
	}

	tarPos2 := -1
	for i, u := range used {
		if !u {
			tarPos2 = i
			break
		}
	}
	if tarPos2 < 0 {
		if s1 > tar {
			tarPos2 = 0
			if tarPos1 == 0 {
				tarPos2 = 1
			}
			Fprintln(out, (s1-tar)/k, tarPos1+1, tarPos2+1)
		}
		return
	}
	for i, u := range used {
		if !u && i != tarPos2 {
			Fprintln(out, inf, i+1, tarPos2+1)
		}
	}
	if s1 < tar {
		Fprintln(out, (tar-s1)/k, tarPos2+1, tarPos1+1)
	} else if s1 > tar {
		Fprintln(out, (s1-tar)/k, tarPos1+1, tarPos2+1)
	}
}

//func main() { CF920D(os.Stdin, os.Stdout) }

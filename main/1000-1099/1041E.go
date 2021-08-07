package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1041E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	cnt := make([]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		if w != n {
			Fprint(out, "NO")
			return
		}
		cnt[v]++
	}
	var a, b []int
	for i := n - 1; i > 0; i-- {
		if cnt[i] > i {
			Fprint(out, "NO")
			return
		}
		if cnt[i] > 0 {
			b = append(b, i)
		} else {
			a = append(a, i)
		}
	}
	ans := make([]int, n)
	ans[0] = n
	j := 0
	for _, i := range b {
		c := cnt[i]
		for k := j + 1; k < j+c; k++ {
			if a[0] > i {
				Fprint(out, "NO")
				return
			}
			ans[k] = a[0]
			a = a[1:]
		}
		j += c
		ans[j] = i
	}
	Fprintln(out, "YES")
	for i := 1; i < n; i++ {
		Fprintln(out, ans[i-1], ans[i])
	}
}

//func main() { CF1041E(os.Stdin, os.Stdout) }

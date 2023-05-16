package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1827C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		s = "^" + s + "$"
		n += 2
		h := make([]int, n)
		f := make([]int, n)
		t := []int{}
		S := int64(0)
		for i, m := 1, 0; i+1 < n; i++ {
			if h[m]+m > i {
				h[i] = min(h[m]+m-i, h[m*2-i])
			}
			for s[i+h[i]+1] == s[i-h[i]] {
				h[i]++
				m = i
			}
			t = append(t, i-1)
			for len(t) > 0 {
				m := t[len(t)-1]
				if h[m]+m >= i {
					f[i] = f[m*2-i] + 1
					break
				}
				t = t[:len(t)-1]
			}
			S += int64(f[i])
		}
		Fprintln(out, S)
	}
}

//func main() { CF1827C(os.Stdin, os.Stdout) }

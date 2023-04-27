package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1798E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := make([]byte, n-1)
		f := make([]int, n+1)
		f1 := make([]int, n+1) // 修改一次
		for i, mx := n-1, 0; i > 0; i-- {
			f1[i] = mx + 1
			j := i + a[i] + 1
			if j > n {
				f[i] = -1e9
			} else {
				f1[i] = max(f1[i], f1[j]+1)
				f[i] = f[j] + 1
				mx = max(mx, f[i])
			}
			m := a[i-1]
			if f[i] != m {
				if f[i] > 0 || f1[i] >= m {
					ans[i-1] = 1
				} else {
					ans[i-1] = 2
				}
			}
		}
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1798E(os.Stdin, os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1181C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	f := make([]int, n)
	for j := 0; j < m; j++ {
		for i, row := range a {
			if j > 0 && row[j] == row[j-1] {
				f[i]++
			} else {
				f[i] = 1
			}
		}
	next:
		for i := 0; i < n; {
			i0 := i
			mn := f[i] // 左侧最短同色长度
			for i++; i < n && a[i][j] == a[i0][j]; i++ {
				mn = min(mn, f[i])
			}

			size := i - i0
			if i0 < size || i+size > n {
				continue
			}

			// 前一段同色 size
			for k := i0 - 1; k >= i0-size; k-- {
				if a[k][j] != a[i0-1][j] {
					continue next
				}
				mn = min(mn, f[k])
			}

			// 后一段同色 size
			for k := i; k < i+size; k++ {
				if a[k][j] != a[i][j] {
					continue next
				}
				mn = min(mn, f[k])
			}
			ans += mn
		}
	}
	Fprint(out, ans)
}

//func main() { CF1181C(os.Stdin, os.Stdout) }

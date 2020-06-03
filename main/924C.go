package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF924C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if v > res {
				res = v
			}
		}
		return res
	}

	var n, mx int
	Fscan(in, &n)
	m := make([]int, n)
	for i := range m {
		Fscan(in, &m[i])
	}
	maxs := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		mx = max(mx-1, m[i]+1)
		maxs[i] = mx
	}
	cnt := make([]int, n)
	cnt[0] = 1
	for i := 1; i < n; i++ {
		cnt[i] = max(cnt[i-1], m[i]+1, maxs[i])
	}
	ans := int64(0)
	for i, v := range cnt {
		ans += int64(v - 1 - m[i])
	}
	Fprint(out, ans)
}

//func main() { CF924C(os.Stdin, os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1154E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, v int
	Fscan(in, &n, &k)
	pos := make([]int, n+1)
	left := make([]int, n+2)
	right := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		pos[v] = i
		left[i], right[i] = i-1, i+1
	}
	del := func(i int) {
		l, r := left[i], right[i]
		right[l] = r
		left[r] = l
	}

	ans := make([]byte, n+1)
	for v, cur := n, byte('1'); v > 0; v-- {
		i := pos[v]
		if ans[i] > 0 {
			continue
		}
		for j, k := right[i], k; j <= n && k > 0; k-- {
			ans[j] = cur
			del(j)
			j = right[j]
		}
		for j, k := i, k+1; j > 0 && k > 0; k-- {
			ans[j] = cur
			del(j)
			j = left[j]
		}
		cur ^= 3
	}
	Fprintf(out, "%s", ans[1:])
}

//func main() { CF1154E(os.Stdin, os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF602B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, ans, l int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	cnt := map[int]int{}
	for r, v := range a {
		cnt[v]++
		for len(cnt) > 2 {
			cnt[a[l]]--
			if cnt[a[l]] == 0 {
				delete(cnt, a[l])
			}
			l++
		}
		ans = max(ans, r-l+1)
	}
	Fprint(out, ans)
}

//func main() { CF602B(os.Stdin, os.Stdout) }

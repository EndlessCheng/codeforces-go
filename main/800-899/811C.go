package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF811C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	cnt := [5001]int{}
	for i := range a {
		Fscan(in, &a[i])
		cnt[a[i]]++
	}
	f := make([]int, n+1)
	vis := [5001]int{}
	for i := 1; i <= n; i++ {
		f[i] = f[i-1]
		for j, c, xor := i, 0, 0; j > 0; j-- {
			v := a[j-1]
			if vis[v] < i {
				vis[v] = i
				c += cnt[v]
				xor ^= v
			}
			if c--; c == 0 {
				f[i] = max(f[i], f[j-1]+xor)
			}
		}
	}
	Fprint(out, f[n])
}

//func main() { CF811C(os.Stdin, os.Stdout) }

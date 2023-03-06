package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF148E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, m, k, t int
	Fscan(in, &n, &m)
	f := make([]int, m+1)
	for ; n > 0; n-- {
		Fscan(in, &k)
		s := make([]int, k+1) // 前缀和
		for i := 1; i <= k; i++ {
			Fscan(in, &s[i])
			s[i] += s[i-1]
		}
		mx := make([]int, k+1)
		for i := 1; i <= k; i++ { // 计算选 i 个物品的最大价值和
			for j := 0; j <= i; j++ { // 枚举前缀长度
				mx[i] = max(mx[i], s[j]+s[k]-s[k-i+j])
			}
		}
		t = min(t+k, m) // 优化循环上界
		for c := t; c >= 0; c-- {
			for j := 1; j <= k && j <= c; j++ {
				f[c] = max(f[c], f[c-j]+mx[j])
			}
		}
	}
	Fprint(out, f[m])
}

//func main() { CF148E(os.Stdin, os.Stdout) }

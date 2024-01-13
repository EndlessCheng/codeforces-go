package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1670C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		to := make([]int, n+1)
		for _, v := range a {
			Fscan(in, &to[v])
		}
		has := make([]bool, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			has[v] = true
		}

		ans := 1
		vis := make([]bool, n+1)
		for i := 1; i <= n; i++ {
			if vis[i] {
				continue
			}
			vis[i] = true
			allZero := !has[i]
			for x := to[i]; x != i; x = to[x] {
				vis[x] = true
				if has[x] {
					allZero = false
				}
			}
			if allZero && i != to[i] {
				ans = ans * 2 % 1_000_000_007
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1670C(os.Stdin, os.Stdout) }

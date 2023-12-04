package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf620C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	Fscan(in, &n)
	ans := [][2]int{}
	for i := 1; i <= n; i++ {
		st := i
		vis := map[int]bool{}
		for ; i <= n; i++ {
			Fscan(in, &v)
			if vis[v] {
				break
			}
			vis[v] = true
		}
		if i > n {
			if len(ans) == 0 {
				Fprint(out, -1)
				return
			}
			ans[len(ans)-1][1] = n
		} else {
			ans = append(ans, [2]int{st, i})
		}
	}
	Fprintln(out, len(ans))
	for _, p := range ans {
		Fprintln(out, p[0], p[1])
	}
}

//func main() { cf620C(os.Stdin, os.Stdout) }

package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1592E(in io.Reader, out io.Writer) {
	var n, ans, ts, xor int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	var vis, pre [1 << 20]int
	for b := range 20 {
		ts++
		for i, v := range a {
			if v>>b&1 == 0 {
				ts++
				continue
			}
			if vis[xor] < ts {
				vis[xor] = ts
				pre[xor] = i
			}
			xor ^= v >> b
			if vis[xor] == ts {
				ans = max(ans, i-pre[xor]+1)
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1592E(bufio.NewReader(os.Stdin), os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1579G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const inf int = 1e9
	const mx = 2001
	pre := make([]int, mx)
	f := make([]int, mx)

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		for i := range pre {
			pre[i] = inf
		}
		pre[0] = 0
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			for i := range f {
				f[i] = inf
			}
			for j, dis := range pre {
				if pre[j] == inf {
					continue
				}
				if j+v < mx {
					f[j+v] = min(f[j+v], max(dis, j+v))
				}
				if j >= v {
					f[j-v] = min(f[j-v], dis)
				} else {
					f[0] = min(f[0], dis-j+v)
				}
			}
			pre, f = f, pre
		}
		ans := inf
		for _, x := range pre {
			ans = min(ans, x)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1579G(os.Stdin, os.Stdout) }

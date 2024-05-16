package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1884D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := make([]int, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			cnt[v]++
		}
		hasD := make([]bool, n+1)
		res := make([]int, n+1)
		for i := n; i > 0; i-- {
			c := 0
			for j := i; j <= n; j += i {
				if cnt[i] > 0 {
					hasD[j] = true
				}
				c += cnt[j]
				res[i] -= res[j]
			}
			res[i] += c * (c - 1) / 2
		}
		ans := 0
		for i, b := range hasD {
			if !b {
				ans += res[i]
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1884D(bufio.NewReader(os.Stdin), os.Stdout) }

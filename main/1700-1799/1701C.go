package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1701C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		cnt := make([]int, n)
		for i := 0; i < m; i++ {
			Fscan(in, &v)
			cnt[v-1]++
		}
		ans := sort.Search(m*2, func(t int) bool {
			ex := 0
			for _, c := range cnt {
				if c > t {
					ex += c - t
				} else {
					ex -= (t - c) / 2
				}
			}
			return ex <= 0
		})
		Fprintln(out, ans)
	}
}

//func main() { cf1701C(os.Stdin, os.Stdout) }

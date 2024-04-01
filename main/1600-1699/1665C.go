package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1665C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := make([]int, n)
		for w := 1; w < n; w++ {
			Fscan(in, &v)
			cnt[v-1]++
		}

		sort.Sort(sort.Reverse(sort.IntSlice(cnt)))
		low := 1
		for _, c := range cnt {
			if c > 0 {
				low++
			}
		}
		ans := low + sort.Search(n-low, func(t int) bool {
			t += low
			need := n - t
			for i, c := range cnt {
				if c <= 1 {
					break
				}
				need -= min(t-1-i, c-1)
			}
			return need <= 0
		})
		Fprintln(out, ans)
	}
}

//func main() { CF1665C(os.Stdin, os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func cf1777C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 100001
	ds := [mx][]int{}
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i {
			ds[j] = append(ds[j], i)
		}
	}

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		a = slices.Compact(a)

		ans := int(1e9)
		todo := m
		l := 0
		cnt := make([]int, m+1) // 无需 map，题目保证 sum(m) <= 1e5
		for _, v := range a {
			for _, d := range ds[v] {
				if d > m {
					break
				}
				if cnt[d] == 0 {
					todo--
				}
				cnt[d]++
			}
			for todo == 0 {
				ans = min(ans, v-a[l])
				for _, d := range ds[a[l]] {
					if d > m {
						break
					}
					cnt[d]--
					if cnt[d] == 0 {
						todo++
					}
				}
				l++
			}
		}
		if ans == 1e9 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1777C(os.Stdin, os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1770C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int64, n)
		has := map[int64]bool{}
		for i := range a {
			Fscan(in, &a[i])
			if has[a[i]] {
				has[0] = true
			}
			has[a[i]] = true
		}
		if has[0] {
			Fprintln(out, "NO")
			continue
		}
	o2:
		for i := 2; i <= n/2; i++ {
			cnt := make([]int, i)
			for _, v := range a {
				cnt[v%int64(i)]++
			}
			for _, c := range cnt {
				if c < 2 {
					continue o2
				}
			}
			Fprintln(out, "NO")
			continue o
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1770C(os.Stdin, os.Stdout) }

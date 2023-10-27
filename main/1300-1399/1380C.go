package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1380C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, low int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &low)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Sort(sort.Reverse(sort.IntSlice(a)))
		ans := 0
		for i := 0; i < n; i++ {
			for st := i; i < n; i++ {
				if a[i]*(i-st+1) >= low {
					ans++
					break
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1380C(os.Stdin, os.Stdout) }

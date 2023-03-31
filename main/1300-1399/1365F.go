package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1365F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ x, y int }
		cnt := map[pair]int{}
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			if i >= (n+1)/2 {
				cnt[pair{a[i], a[n-1-i]}]++
				cnt[pair{a[n-1-i], a[i]}]++
			}
		}
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
			if i >= (n+1)/2 {
				cnt[pair{b[i], b[n-1-i]}]--
				cnt[pair{b[n-1-i], b[i]}]--
			}
		}
		if n%2 > 0 && a[n/2] != b[n/2] {
			Fprintln(out, "No")
			continue
		}
		for _, c := range cnt {
			if c != 0 {
				Fprintln(out, "No")
				continue o
			}
		}
		Fprintln(out, "Yes")
	}
}

//func main() { CF1365F(os.Stdin, os.Stdout) }

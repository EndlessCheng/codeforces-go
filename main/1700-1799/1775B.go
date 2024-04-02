package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1775B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
NEXT:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([][]int, n)
		cnt := map[int]int{}
		for i := range a {
			Fscan(in, &k)
			a[i] = make([]int, k)
			for j := range a[i] {
				Fscan(in, &a[i][j])
				cnt[a[i][j]]++
			}
		}
	nxt:
		for _, r := range a {
			for _, v := range r {
				if cnt[v] == 1 {
					continue nxt
				}
			}
			Fprintln(out, "Yes")
			continue NEXT
		}
		Fprintln(out, "No")
	}
}

//func main() { cf1775B(os.Stdin, os.Stdout) }

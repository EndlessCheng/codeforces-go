package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1579B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := append(a[:0:0], a...)
		sort.Ints(b)
		ans := [][2]int{}
		for i, v := range b {
			if a[i] != v {
				j := i + 1
				for a[j] != v {
					j++
				}
				ans = append(ans, [2]int{i + 1, j + 1})
				copy(a[i+1:], a[i:j])
			}
		}
		Fprintln(out, len(ans))
		for _, p := range ans {
			Fprintln(out, p[0], p[1], p[1]-p[0])
		}
	}
}

//func main() { CF1579B(os.Stdin, os.Stdout) }

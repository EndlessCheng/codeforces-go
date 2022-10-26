package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, s, ans int
	Fscan(in, &n, &m)
	a := make([]int, n-1)
	for i := range a {
		Fscan(in, &a[i])
	}
	cnt := map[int]int{}
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
		cnt[b[i]]++
	}
	for i, v := range a {
		if i&1 > 0 {
			s -= v
			for _, x := range b {
				cnt[s+x]++
			}
		} else {
			s += v
			for _, x := range b {
				cnt[s-x]++
			}
		}
	}
	for _, c := range cnt {
		if c > ans {
			ans = c
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

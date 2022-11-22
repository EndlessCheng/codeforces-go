package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	cnt := make([]int, n+1)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		cnt[a[i]]++
	}
	b := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		Fscan(in, &b[i])
		cnt[b[i]]++
	}
	for _, c := range cnt {
		if c > n {
			Fprint(out, "No")
			return
		}
	}
	j := 0
	for i, v := range a {
		w := b[i]
		if v != w {
			continue
		}
		for b[j] == v || w == a[j] {
			j++
		}
		b[j], b[i] = b[i], b[j]
		j++
	}
	Fprintln(out, "Yes")
	for _, v := range b {
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }

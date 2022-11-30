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
	cnt := map[int]int{}
	dup := false
	var n, inv int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		cnt[a[i]]++
		if cnt[a[i]] > 1 {
			dup = true
		}
		for _, v := range a[:i] {
			if v > a[i] {
				inv ^= 1
			}
		}
	}
	for i := range a {
		Fscan(in, &a[i])
		cnt[a[i]]--
		for _, v := range a[:i] {
			if v > a[i] {
				inv ^= 1
			}
		}
	}
	if !dup && inv > 0 {
		Fprint(out, "No")
		return
	}
	for _, c := range cnt {
		if c != 0 {
			Fprint(out, "No")
			return
		}
	}
	Fprint(out, "Yes")
}

func main() { run(os.Stdin, os.Stdout) }

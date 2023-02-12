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

	var n, x, v int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ban := [1e5 + 1]bool{}
	for Fscan(in, &x); x > 0; x-- {
		Fscan(in, &v)
		ban[v] = true
	}
	Fscan(in, &x)

	f := make([]bool, x+1)
	f[0] = true
	for i := 1; i <= x; i++ {
		if ban[i] {
			continue
		}
		for _, v := range a {
			if v > i {
				break
			}
			if f[i-v] {
				f[i] = true
				break
			}
		}
	}
	if f[x] {
		Fprintln(out, "Yes")
	} else {
		Fprintln(out, "No")
	}
}

func main() { run(os.Stdin, os.Stdout) }

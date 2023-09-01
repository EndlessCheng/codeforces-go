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
	var n, ans int
	var s string
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &s)

	var pre, suf [3]int
	for i, c := range s {
		if c == 'X' {
			suf[a[i]]++
		}
	}

	mex := []int{0, 1, 0, 2, 0, 1, 0, 3}
	for i, v := range a {
		if s[i] == 'M' {
			pre[v]++
		} else if s[i] == 'E' {
			for j, cj := range pre {
				for k, ck := range suf {
					ans += mex[1<<v|1<<j|1<<k] * cj * ck
				}
			}
		} else {
			suf[v]--
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

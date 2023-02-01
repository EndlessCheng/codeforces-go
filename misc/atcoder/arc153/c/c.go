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

	var n, s int
	Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		b[i] = i
		s += i * a[i]
	}
	output := func() {
		Fprintln(out, "Yes")
		for _, v := range b {
			Fprint(out, v, " ")
		}
	}
	if s == 0 {
		output()
		return
	}
	if s > 0 == (a[0] > 0) {
		b[0] = -s * a[0]
		output()
		return
	}
	if s > 0 == (a[n-1] < 0) {
		b[n-1] = -(s - (n-1)*a[n-1]) * a[n-1]
		output()
		return
	}

	pre := 0
	for i, v := range a {
		pre += v
		if pre != 0 && pre > 0 == (a[n-1] > 0) {
			s = 0
			for j := i + 1; j < n; j++ {
				s += j * a[j]
			}
			for j := 0; j < i; j++ {
				b[j] = j - 1e12
				s += b[j] * a[j]
			}
			b[i] = -s * a[i]
			output()
			return
		}
	}

	suf := 0
	for i := n - 1; i >= 0; i-- {
		suf += a[i]
		if suf != 0 && suf > 0 == (a[0] > 0) {
			s = 0
			for j := 1; j < i; j++ {
				s += j * a[j]
			}
			for j := i + 1; j < n; j++ {
				b[j] = 1e12 - (n - 1 - j)
				s += b[j] * a[j]
			}
			b[i] = -s * a[i]
			output()
			return
		}
	}
	Fprint(out, "No")
}

func main() { run(os.Stdin, os.Stdout) }

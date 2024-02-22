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

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		z := make([]int, n)
		diff := make([]int, n)
		l, r := 0, 0
		for i := 1; i < n; i++ {
			if i <= r {
				z[i] = min(z[i-l], r-i+1)
			}
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				l, r = i, i+z[i]
				z[i]++
			}
			diff[z[i]]--
		}
		sumD := n
		for _, d := range diff {
			sumD += d
			Fprint(out, sumD, " ")
		}
		Fprintln(out)
	}
}

func main() { run(os.Stdin, os.Stdout) }

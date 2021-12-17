package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1593G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, q, l, r int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &q)
		sum := make([]int, len(s)+1)
		for i, b := range s {
			sum[i+1] = sum[i]
			if b < 42 {
				sum[i+1] += i&1<<1 - 1
			}
		}
		for ; q > 0; q-- {
			Fscan(in, &l, &r)
			Fprintln(out, abs(sum[r]-sum[l-1]))
		}
	}
}

//func main() { CF1593G(os.Stdin, os.Stdout) }

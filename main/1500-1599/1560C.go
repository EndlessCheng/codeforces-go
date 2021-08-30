package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1560C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &k)
		c := 1
		for ; k > c; c += 2 {
			k -= c
		}
		if r := (c + 1) / 2; k < r {
			Fprintln(out, k, r)
		} else {
			Fprintln(out, r, c-k+1)
		}
	}
}

//func main() { CF1560C(os.Stdin, os.Stdout) }

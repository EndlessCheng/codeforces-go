package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1383B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := uint(0)
		a := make([]uint, n)
		for i := range a {
			Fscan(in, &a[i])
			s ^= a[i]
		}
		if s == 0 {
			Fprintln(out, "DRAW")
			continue
		}
		b := bits.Len(s) - 1
		o := 0
		for _, v := range a {
			if v>>b&1 > 0 {
				o++
			}
		}
		if o&3 == 3 && (n-o)&1 == 0 {
			Fprintln(out, "LOSE")
		} else {
			Fprintln(out, "WIN")
		}
	}
}

//func main() { CF1383B(os.Stdin, os.Stdout) }

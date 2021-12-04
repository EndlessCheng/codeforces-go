package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1110C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n&(n+1) > 0 {
			Fprintln(out, 1<<bits.Len(uint(n))-1)
			continue
		}
		for d := 2; d*d < n; d++ {
			if n%d == 0 {
				Fprintln(out, n/d)
				continue o
			}
		}
		Fprintln(out, 1)
	}
}

//func main() { CF1110C(os.Stdin, os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, s, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
	}
	for i := 1 << (bits.Len(uint(s)) - 1); i > 0; i >>= 1 {
		cnt := 0
		for l := range a {
			s := 0
			for _, v := range a[l:] {
				s += v
				if s&i > 0 && s&ans == ans {
					cnt++
				}
			}
		}
		if cnt >= k {
			ans |= i
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

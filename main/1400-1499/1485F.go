package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1485F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int = 1e9 + 7

	var T, n, f int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		sf := make([]int, n+1)
		last := map[int64]int{}
		s := int64(0)
		for i := n - 1; i >= 0; i-- {
			s += int64(b[i])
			if j := last[s]; j > 0 {
				f = (sf[i+1] - sf[j+1] + mod) % mod
			} else {
				f = (sf[i+1] + 1) % mod
			}
			sf[i] = (sf[i+1] + f) % mod
			last[s] = i
		}
		Fprintln(out, f)
	}
}

//func main() { CF1485F(os.Stdin, os.Stdout) }

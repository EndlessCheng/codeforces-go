package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// https://space.bilibili.com/206214
func CF1744D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := 0
		c2 := make([]int, n)
		for i := range c2 {
			Fscan(in, &v)
			s += bits.TrailingZeros(uint(v))
			c2[i] = bits.TrailingZeros(uint(i + 1))
		}
		sort.Ints(c2)
		i := n - 1
		for ; i >= 0 && s < n; i-- {
			s += c2[i]
		}
		if i < 0 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, n-1-i)
		}
	}
}

//func main() { CF1744D(os.Stdin, os.Stdout) }

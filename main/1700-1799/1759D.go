package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf1759D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		c2 := bits.TrailingZeros(uint(n))
		c5 := 0
		for x := n; x%5 == 0; x /= 5 {
			c5++
		}
		k := 1
		for c2 < c5 && k*2 <= m {
			c2++
			k *= 2
		}
		for c5 < c2 && k*5 <= m {
			c5++
			k *= 5
		}
		for k*10 <= m {
			k *= 10
		}
		Fprintln(out, n*(m-m%k))
	}
}

//func main() { cf1759D(os.Stdin, os.Stdout) }

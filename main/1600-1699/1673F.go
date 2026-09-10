package main

import (
	"bufio"
	. "fmt"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
func cf1673F() {
	out := bufio.NewWriter(os.Stdout)
	var n, m int
	Scan(&n, &m)
	for range n {
		for j := 1; j < n; j++ {
			t := bits.TrailingZeros(uint(j))
			Fprint(out, 1<<(t*2+1), " ")
		}
		Fprintln(out)
		out.Flush()
	}
	for i := 1; i < n; i++ {
		for range n {
			t := bits.TrailingZeros(uint(i))
			Fprint(out, 1<<(t*2), " ")
		}
		Fprintln(out)
		out.Flush()
	}
	var x, y int
	for range m {
		var w int
		Scan(&w)
		for j := 4; j >= 0; j-- {
			if w>>(j*2)&1 != 0 {
				x ^= (1 << (j + 1)) - 1
			}
			if w>>(j*2+1)&1 != 0 {
				y ^= (1 << (j + 1)) - 1
			}
		}
		Fprintln(out, x+1, y+1)
		out.Flush()
	}
}

//func main() { cf1673F() }

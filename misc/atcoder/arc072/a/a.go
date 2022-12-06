package main

import (
	"bufio"
	. "fmt"
	"os"
)

// https://space.bilibili.com/206214
func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f := func(to int) (res int) {
		s := 0
		for _, v := range a {
			s += v
			if s*to <= 0 {
				res += abs(s - to)
				s = to
			}
			to *= -1
		}
		return
	}
	Print(min(f(1), f(-1)))
}

func abs(x int) int { if x < 0 { return -x }; return x }
func min(a, b int) int { if a > b { return b }; return a }

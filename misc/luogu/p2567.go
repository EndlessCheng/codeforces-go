package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p2567(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var l, r, ans int
	Fscan(in, &l, &r)
	a := []int{}
	var init func(int)
	init = func(v int) {
		if v > r {
			return
		}
		a = append(a, v)
		init(v*10 + 6)
		init(v*10 + 8)
	}
	init(0)
	a = a[1:]
	slices.Sort(a)

	b := a[:0]
o:
	for _, v := range a {
		for _, w := range b {
			if v%w == 0 {
				continue o
			}
		}
		b = append(b, v)
	}

	var f func(int, int, int)
	f = func(i, lcm, sign int) {
		for j, v := range b[:i] {
			l2 := lcm / gcd(lcm, v)
			if l2 > r/v {
				continue
			}
			l2 *= v
			ans += sign * (r/l2 - (l-1)/l2)
			f(j, l2, -sign)
		}
	}
	f(len(b), 1, 1) // 大的数可以提前退出，如果从小的数开始，后面会做大量无用功
	Fprint(out, ans)
}

//func main() { p2567(os.Stdin, os.Stdout) }

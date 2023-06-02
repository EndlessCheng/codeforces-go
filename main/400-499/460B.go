package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF460B(in io.Reader, out io.Writer) {
	var a int
	var b, c int64
	Fscan(in, &a, &b, &c)
	ans := []int64{}
	for ds := int64(1); ds < 82; ds++ {
		x := b
		for i := 0; i < a; i++ {
			x *= ds
		}
		x += c
		if x >= 1e9 {
			break
		}
		if x > 0 {
			s := ds
			for x := x; x > 0; x /= 10 {
				s -= x % 10
			}
			if s == 0 {
				ans = append(ans, x)
			}
		}
	}
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF460B(os.Stdin, os.Stdout) }

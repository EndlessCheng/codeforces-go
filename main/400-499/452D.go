package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF452D(in io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var k, n1, n2, n3, t1, t2, t3, finish int
	Fscan(in, &k, &n1, &n2, &n3, &t1, &t2, &t3)
	f1 := make([]int, n1)
	f2 := make([]int, n2)
	f3 := make([]int, n3)
	for i := 0; i < k; i++ {
		finish = max(max(f1[i%n1]+t1+t2+t3, f2[i%n2]+t2+t3), f3[i%n3]+t3)
		f1[i%n1] = finish - t2 - t3
		f2[i%n2] = finish - t3
		f3[i%n3] = finish
	}
	Fprint(out, finish)
}

//func main() { CF452D(os.Stdin, os.Stdout) }

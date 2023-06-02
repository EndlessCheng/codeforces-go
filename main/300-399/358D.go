package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF358D(_r io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	// a[i][0]: 213 / 312
	// a[i][1]: 123 / 321
	// a[i][2]: 132 / 231 对于 0 和 n-1 来说不可能
	a := make([][3]int, n)
	for j := 0; j < 3; j++ {
		for i := range a {
			Fscan(in, &a[i][j])
		}
	}
	// f0：先 i 后 i+1
	// f1：先 i+1 后 i
	f0, f1 := a[0][0], a[0][1]
	for _, p := range a[1:] {
		f0, f1 = max(f1+p[0], f0+p[1]), max(f1+p[1], f0+p[2])
	}
	// 假设 n-1 的右边还有一个点，那么肯定是先取 n-1，所以答案是 f0
	Fprint(out, f0)
}

//func main() { CF358D(os.Stdin, os.Stdout) }

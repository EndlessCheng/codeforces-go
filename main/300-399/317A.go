package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF317A(in io.Reader, out io.Writer) {
	var x, y, m int64
	Fscan(in, &x, &y, &m)
	if x > y {
		x, y = y, x
	}
	if y >= m {
		Fprint(out, 0)
	} else if m <= 1 || y <= 0 {
		Fprint(out, -1)
	} else {
		ans := int64(0)
		if x <= 0 {
			ans = -x/y + 1
			x += ans * y
		}
		for y < m {
			x, y = y, x+y
			ans++
		}
		Fprint(out, ans)
	}
}

//func main() { CF317A(os.Stdin, os.Stdout) }

package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1709A(in io.Reader, out io.Writer) {
	var T, x int
	var a [4]int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &a[1], &a[2], &a[3])
		for i := 0; i < 2; i++ {
			x = a[x]
			if x == 0 {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1709A(os.Stdin, os.Stdout) }

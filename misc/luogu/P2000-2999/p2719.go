package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p2719(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	ans := 1.
	for i := 1; i <= n/2-1; i++ {
		ans = ans * float64(n-1-i) / float64(i*4)
	}
	Fprintf(out, "%.4f", 1-ans)
}

//func main() { p2719(os.Stdin, os.Stdout) }

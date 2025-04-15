package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p1365(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s, &s)
	ans, l := 0., 0.
	for _, b := range s {
		if b == 'o' {
			ans += l*2 + 1
			l++
		} else if b == 'x' {
			l = 0
		} else {
			ans += l + 0.5
			l = l/2 + 0.5
		}
	}
	Fprintf(out, "%.4f", ans)
}

//func main() { p1365(bufio.NewReader(os.Stdin), os.Stdout) }

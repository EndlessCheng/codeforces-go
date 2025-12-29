package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p10031(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n%2 == 0 {
			n ^= n / 2
		}
		Fprintln(out, n)
	}
}

//func main() { p10031(bufio.NewReader(os.Stdin), os.Stdout) }

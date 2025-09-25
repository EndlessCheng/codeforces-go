package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1574B(in io.Reader, out io.Writer) {
	var T, a, b, c, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c, &m)
		s := a + b + c
		if max(a, b, c)*2-s-1 <= m && m < s-2 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1574B(bufio.NewReader(os.Stdin), os.Stdout) }

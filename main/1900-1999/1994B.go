package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf1994B(in io.Reader, out io.Writer) {
	var T, n int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		i := strings.IndexByte(s, '1')
		j := strings.IndexByte(t, '1')
		if j < 0 || i >= 0 && i <= j {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1994B(bufio.NewReader(os.Stdin), os.Stdout) }

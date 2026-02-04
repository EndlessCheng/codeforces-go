package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf1333A(in io.Reader, out io.Writer) {
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		Fprintln(out, strings.Repeat("B", m-1)+"W")
		for range n - 1 {
			Fprintln(out, strings.Repeat("B", m))
		}
	}
}

//func main() { cf1333A(bufio.NewReader(os.Stdin), os.Stdout) }

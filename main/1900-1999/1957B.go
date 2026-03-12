package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"strings"
)

// https://github.com/EndlessCheng
func cf1957B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		if n == 1 {
			Fprintln(out, k)
			continue
		}
		w := bits.Len(uint(k+1)) - 1
		x := 1<<w - 1
		Fprintln(out, x, k-x, strings.Repeat("0 ", n-2))
	}
}

//func main() { cf1957B(os.Stdin, os.Stdout) }

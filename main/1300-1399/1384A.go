package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1384A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, lcp int
	s := bytes.Repeat([]byte{'b'}, 51)
	for Fscan(in, &T); T > 0; T-- {
		Fprintf(out, "%s\n", s)
		Fscan(in, &n)
		for range n {
			Fscan(in, &lcp)
			s[lcp] ^= 1
			Fprintf(out, "%s\n", s)
		}
	}
}

//func main() { cf1384A(bufio.NewReader(os.Stdin), os.Stdout) }

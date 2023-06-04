package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1839A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		ones := 0
		for i := n; i > 0; i-- {
			if (i-1)%k == 0 {
				ones++
			}
			if (n-i)/k+1 > ones {
				ones++
			}
		}
		Fprintln(out, ones)
	}
}

//func main() { CF1839A(os.Stdin, os.Stdout) }

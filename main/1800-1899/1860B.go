package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1860B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, m, k, a1, ak int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &m, &k, &a1, &ak)
		ak += max(a1-m%k, 0) / k
		Fprintln(out, max(m%k-a1, 0)+max(m/k-ak, 0))
	}
}

//func main() { cf1860B(bufio.NewReader(os.Stdin), os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1926E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		for i := 0; ; i++ {
			m := (n>>i + 1) / 2
			if k <= m {
				Fprintln(out, (k*2-1)<<i)
				break
			}
			k -= m
		}
	}
}

//func main() { cf1926E(bufio.NewReader(os.Stdin), os.Stdout) }

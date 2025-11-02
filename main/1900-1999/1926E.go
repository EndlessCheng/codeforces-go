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
		for mul := 1; ; mul *= 2 {
			m := (n + 1) / 2
			if k <= m {
				Fprintln(out, (k*2-1)*mul)
				break
			}
			k -= m
			n /= 2
		}
	}
}

//func main() { cf1926E(bufio.NewReader(os.Stdin), os.Stdout) }

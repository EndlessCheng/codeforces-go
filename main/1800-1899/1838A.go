package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1838A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		neg, mx := 0, 0
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			if a[i] < 0 {
				neg = a[i]
			} else if a[i] > mx {
				mx = a[i]
			}
		}
		if neg < 0 {
			Fprintln(out, neg)
		} else {
			Fprintln(out, mx)
		}
	}
}

//func main() { CF1838A(os.Stdin, os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	k2 := k * 2
	m := n - n%k2 - k2
	if m < 0 {
		Fprint(out, -1)
		return
	}
	i := 1
	for ; i <= m; i += k2 {
		for j := 0; j < k; j++ {
			Fprint(out, i+k+j, " ")
		}
		for j := 0; j < k; j++ {
			Fprint(out, i+j, " ")
		}
	}
	if n%k2 <= k {
		for j := 0; j < k+n%k2; j++ {
			Fprint(out, i+k+j, " ")
		}
		for j := 0; j < k; j++ {
			Fprint(out, i+j, " ")
		}
	} else {
		for j := 0; j < k; j++ {
			Fprint(out, i+k+j, " ")
		}
		for j := 0; j < n%k; j++ {
			Fprint(out, i+j, " ")
		}
		for j := n % k; j < k+n%k; j++ {
			Fprint(out, i+k2+j, " ")
		}
		for j := n % k; j < k; j++ {
			Fprint(out, i+j, " ")
		}
		for j := 0; j < n%k; j++ {
			Fprint(out, i+k2+j, " ")
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }

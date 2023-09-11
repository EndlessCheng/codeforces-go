package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		x ^= a[i]
	}
	for _, v := range a {
		Fprint(out, x^v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }

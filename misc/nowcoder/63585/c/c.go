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

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	j := 1
	for i := 0; i < n; i += 2 {
		a[i] = j
		j++
	}
	j = n
	for i := 1; i < n; i += 2 {
		a[i] = j
		j--
	}
	for _, v := range a {
		Fprint(out, v, " ")
	}
	Fprintln(out)
}

func main() { run(os.Stdin, os.Stdout) }

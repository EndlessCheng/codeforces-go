package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k, s int
	Fscan(in, &n, &k, &s)
	for i := 0; i < k; i++ {
		Fprint(out, s, " ")
	}
	t := 1
	if s < 1e9 {
		t = 1e9
	}
	for i := k; i < n; i++ {
		Fprint(out, t, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] == 1 {
			Fprintln(out, "YES")
			return
		}
	}
	Fprintln(out, "NO")
}

func main() { run(os.Stdin, os.Stdout) }

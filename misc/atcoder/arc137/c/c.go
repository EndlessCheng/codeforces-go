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
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	if a[n-1]+1 < a[n] || (a[n]-n)%2 == 0 {
		Fprintln(out, "Alice")
	} else {
		Fprintln(out, "Bob")
	}
}

func main() { run(os.Stdin, os.Stdout) }

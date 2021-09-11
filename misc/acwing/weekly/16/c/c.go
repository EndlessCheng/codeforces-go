package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := 1; i < n-1; i++ {
		if a[i] < a[i+1] && a[i] < a[0] || a[i] > a[i+1] && a[i] > a[0] {
			Fprintln(out, 3)
			Fprintln(out, 1, i+1, i+2)
			return
		}
	}
	Fprintln(out, 0)
}

func main() { run(os.Stdin, os.Stdout) }

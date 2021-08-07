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
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	win := make([]bool, k+1)
	for i := a[0]; i <= k; i++ {
		for _, v := range a {
			if i >= v && !win[i-v] {
				win[i] = true
				break
			}
		}
	}
	if win[k] {
		Fprint(out, "First")
	} else {
		Fprint(out, "Second")
	}
}

func main() { run(os.Stdin, os.Stdout) }

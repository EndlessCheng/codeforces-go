package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	sum := [3]int{}
	var n, v int
	for Fscan(in, &n); n > 0; n-- {
		for j := 0; j < 3; j++ {
			Fscan(in, &v)
			sum[j] += v
		}
	}
	for _, s := range sum {
		if s != 0 {
			Fprint(out, "NO")
			return
		}
	}
	Fprint(out, "YES")
}

func main() { run(os.Stdin, os.Stdout) }

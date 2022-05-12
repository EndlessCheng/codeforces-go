package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF710C(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var n int
	Fscan(in, &n)
	odd, even := 1, 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if abs(i-n/2)+abs(j-n/2) <= n/2 {
				Fprint(out, odd, " ")
				odd += 2
			} else {
				Fprint(out, even, " ")
				even += 2
			}
		}
		Fprintln(out)
	}
}

//func main() { CF710C(os.Stdin, os.Stdout) }

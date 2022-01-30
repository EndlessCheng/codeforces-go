package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF263A(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var v int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if Fscan(in, &v); v > 0 {
				Fprint(out, abs(i-2)+abs(j-2))
				return
			}
		}
	}
}

//func main() { CF263A(os.Stdin, os.Stdout) }

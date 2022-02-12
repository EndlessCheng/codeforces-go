package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1059C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	for k := 1; n > 0; k *= 2 {
		if n == 3 {
			Fprint(out, k, k, k*3)
			return
		}
		for i := 0; i < (n+1)/2; i++ {
			Fprint(out, k, " ")
		}
		n /= 2
	}
}

//func main() { CF1059C(os.Stdin, os.Stdout) }

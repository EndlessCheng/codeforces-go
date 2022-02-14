package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF959C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	if n < 6 {
		Fprintln(out, -1)
	} else {
		Fprintln(out, 1, 2)
		Fprintln(out, 1, 3)
		Fprintln(out, 1, 4)
		for i := 5; i <= n; i++ {
			Fprintln(out, 2, i)
		}
	}
	for i := 2; i <= n; i++ {
		Fprintln(out, 1, i)
	}
}

//func main() { CF959C(os.Stdin, os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == i {
				Fprint(out, "0 ")
			} else if i < n-1 && j < n-1 {
				Fprint(out, (i+j)%(n-1)+1, " ")
			} else {
				Fprint(out, (i+j)*2%(n-1)+1, " ")
			}
		}
		Fprintln(out)
	}
}

func main() { run(os.Stdin, os.Stdout) }

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
	n--
	block := (n-1)/6 + 1
	Fprintln(out, block*5)
	m := block * 3
	Fprint(out, m)
	for i := 1; i <= m; i++ {
		Fprint(out, " ", i)
	}
	Fprintln(out)
	for i := 0; i < m && n > 0; n-- {
		Fprintln(out, 2, i+1, m+1+i/3*2)
		i++
	}
	for i := 0; n > 0; n-- {
		Fprintln(out, 2, i+1, m+2+i/3*2)
		i++
	}
}

func main() { run(os.Stdin, os.Stdout) }

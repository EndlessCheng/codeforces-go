package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF715A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int64
	Fscan(in, &n)
	Fprintln(out, 2)
	for i := int64(2); i <= n; i++ {
		Fprintln(out, (i+1)*(i+1)*i-i+1)
	}
}

//func main() { CF715A(os.Stdin, os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1316C(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n, m, p, x, y := read(), read(), read(), 0, 0
	for i := 0; i < n; i++ {
		if read()%p > 0 {
			x = i
		}
	}
	for i := 0; i < m; i++ {
		if read()%p > 0 {
			y = i
		}
	}
	Fprint(_w, x+y)
}

//func main() { CF1316C(os.Stdin, os.Stdout) }

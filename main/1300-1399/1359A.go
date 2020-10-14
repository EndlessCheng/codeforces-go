package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1359A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, c, j, p int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &c, &j, &p)
		c /= p
		if j <= c {
			Fprintln(out, j)
		} else {
			Fprintln(out, c-(j-c-1)/(p-1)-1)
		}
	}
}

//func main() { CF1359A(os.Stdin, os.Stdout) }

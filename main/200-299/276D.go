package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol276D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var l, r int64
	Fscan(in, &l, &r)
	if l == r {
		Fprint(out, 0)
		return
	}
	for m := l ^ r; ; m &= m - 1 {
		if m&(m-1) == 0 {
			Fprint(out, m<<1-1)
			return
		}
	}
}

//func main() {
//	Sol276D(os.Stdin, os.Stdout)
//}

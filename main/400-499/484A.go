package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF484A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, l, r int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r)
		for l|(l+1) <= r {
			l |= l + 1 // 把最低的 0 改成 1
		}
		Fprintln(out, l)
	}
}

//func main() { CF484A(os.Stdin, os.Stdout) }

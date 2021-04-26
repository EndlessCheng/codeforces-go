package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF985D(in io.Reader, out io.Writer) {
	var n, H int64
	Fscan(in, &n, &H)
	ans := sort.Search(2e9, func(L int) bool {
		l := int64(L)
		if l <= H {
			return (l+1)*l/2 >= n
		}
		h := (l + H) / 2
		if (l+H)&1 > 0 {
			return ((H+h)*(h-H+1)+(h+1)*h)/2 >= n
		}
		return ((H+h-1)*(h-H)+(h+1)*h)/2 >= n
	})
	Fprint(out, ans)
}

//func main() { CF985D(os.Stdin, os.Stdout) }

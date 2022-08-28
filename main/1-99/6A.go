package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF6A(in io.Reader, out io.Writer) {
	a := make([]int, 4)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	if a[0]+a[1] > a[2] || a[1]+a[2] > a[3] {
		Fprint(out, "TRIANGLE")
	} else if a[0]+a[1] == a[2] || a[1]+a[2] == a[3] {
		Fprint(out, "SEGMENT")
	} else {
		Fprint(out, "IMPOSSIBLE")
	}
}

//func main() { CF6A(os.Stdin, os.Stdout) }

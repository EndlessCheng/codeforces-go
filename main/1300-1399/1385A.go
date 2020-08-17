package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1385A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	t, a := 0, make([]int, 3)
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &a[0], &a[1], &a[2])
		sort.Ints(a)
		if a[1] < a[2] {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
			Fprintln(out, a[0], a[0], a[1])
		}
	}
}

//func main() { CF1385A(os.Stdin, os.Stdout) }

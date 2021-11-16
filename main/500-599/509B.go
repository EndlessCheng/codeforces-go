package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF509B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, mx int
	mi := 100
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] > mx {
			mx = a[i]
		}
		if a[i] < mi {
			mi = a[i]
		}
	}
	if mx-mi > k {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for _, v := range a {
		Fprint(out, strings.Repeat("1 ", mi))
		for i := 1; i <= v-mi; i++ {
			Fprint(out, i, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF509B(os.Stdin, os.Stdout) }

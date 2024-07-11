package main

import (
	. "fmt"
	"io"
	"slices"
)

func cf985E(in io.Reader, out io.Writer) {
	var n, k, d int
	Fscan(in, &n, &k, &d)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)

	f := make([]bool, n+1)
	f[0] = true
	left, cnt := 0, 0
	for i := k; i <= n; i++ {
		if f[i-k] {
			cnt++
		}
		for a[i-1]-a[left] > d {
			if f[left] {
				cnt--
			}
			left++
		}
		f[i] = cnt > 0
	}
	if f[n] {
		Fprintln(out, "YES")
	} else {
		Fprintln(out, "NO")
	}
}

//func main() { cf985E(bufio.NewReader(os.Stdin), os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
	"strings"
)

// https://github.com/EndlessCheng
func cf1028E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, mxI int
	Fscan(in, &n)
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}

	mx := slices.Max(b)
	if slices.Min(b) == mx {
		if mx == 0 {
			Fprint(out, "YES\n", strings.Repeat("1 ", n))
		} else {
			Fprint(out, "NO")
		}
		return
	}

	for i, v := range b {
		if v == mx && b[(i-1+n)%n] < mx {
			mxI = i
			break
		}
	}

	a := make([]int, n)
	a[mxI] = mx
	a[(mxI-1+n)%n] = mx*2 + b[(mxI-1+n)%n]
	for i := 2; i < n; i++ {
		j := (mxI - i + n) % n
		a[j] = a[(j+1)%n] + b[j]
	}
	Fprintln(out, "YES")
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { cf1028E(bufio.NewReader(os.Stdin), os.Stdout) }

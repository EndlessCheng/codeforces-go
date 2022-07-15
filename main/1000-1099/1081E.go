package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1081E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 2e5
	ds := [mx + 1][]int{}
	for i := 1; i*i <= mx; i++ {
		for j := i * (i + 1); j <= mx; j += i {
			ds[j] = append(ds[j], i)
		}
	}

	var n int
	Fscan(in, &n)
	a := make([]int64, n)
	s := int64(0)
o:
	for i := 1; i < n; i += 2 {
		Fscan(in, &a[i])
		ds := ds[a[i]]
		for j := len(ds) - 1; j >= 0; j-- {
			d := int64(ds[j])
			d = a[i]/d - d
			if d&1 == 0 {
				d /= 2
				a[i-1] = d*d - s
				if a[i-1] > 0 {
					s = d*d + a[i]
					continue o
				}
			}
		}
		Fprint(out, "No")
		return
	}
	Fprintln(out, "Yes")
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { CF1081E(os.Stdin, os.Stdout) }

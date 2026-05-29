package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1469F(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)

	limit := a[n-1] + 100
	d := make([]int, limit+2)
	d[0] = 1
	d[1] = -1
	sd := 0
	j := n - 1
	for i := range limit {
		sd += d[i]
		d[i+1] += d[i]
		if sd+d[i+1] >= k {
			Fprint(out, i+1)
			return
		}
		for d[i] > 0 && j >= 0 {
			d[i]--
			l1 := a[j] / 2
			l2 := (a[j] - 1) / 2
			d[i+2]++
			d[i+2+l1]--
			d[i+2]++
			d[i+2+l2]--
			sd--
			j--
		}
	}
	Fprint(out, -1)
}

//func main() { cf1469F(bufio.NewReader(os.Stdin), os.Stdout) }

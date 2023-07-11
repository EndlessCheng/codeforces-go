package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func CF1516C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, total, idx int
	Fscan(in, &n)
	a := make([]int, n)
	minLowbit := math.MaxInt
	for i := range a {
		Fscan(in, &a[i])
		total += a[i]
		lb := a[i] & -a[i]
		if lb < minLowbit {
			minLowbit = lb
			idx = i
		}
	}

	total /= minLowbit
	if total%2 > 0 {
		Fprint(out, 0)
		return
	}

	f := make([]bool, total+1)
	f[0] = true
	for _, x := range a {
		x /= minLowbit
		for j := total; j >= x; j-- {
			f[j] = f[j] || f[j-x]
		}
	}

	if f[total/2] {
		Fprintln(out, 1)
		Fprint(out, idx+1)
	} else {
		Fprint(out, 0)
	}
}

//func main() { CF1516C(os.Stdin, os.Stdout) }

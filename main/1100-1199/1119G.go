package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1119G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, sum int
	Fscan(in, &n, &m)
	a := make([]int, m)
	b := make([]int, m)
	for i := range a {
		Fscan(in, &a[i])
		b[i] = (sum + a[i]) % n
		sum += a[i]
	}
	b[m-1] = n

	slices.Sort(b)
	for i := m - 1; i > 0; i-- {
		b[i] -= b[i-1]
	}

	Fprintln(out, (sum+n-1)/n)
	for _, v := range b {
		Fprint(out, v, " ")
	}
	Fprintln(out)

	j := 0
	for i := range m {
		for a[i] > 0 || i == m-1 && j > 0 {
			a[i] -= b[j]
			j = (j + 1) % m
			Fprint(out, i+1, " ")
		}
	}
}

//func main() { cf1119G(bufio.NewReader(os.Stdin), os.Stdout) }

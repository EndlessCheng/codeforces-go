package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1323D(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n := read()
	a := make([]int, n)
	for i := range a {
		a[i] = read()
	}

	b := make([]int, n)
	search := func(l, x int) int {
		r := n
		for l < r {
			m := (l + r) >> 1
			if b[m] >= x {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}
	ans := 0
	for i := uint(0); i < 25; i++ {
		for j, v := range a {
			b[j] = v &^ (-1 << (i + 1))
		}
		sort.Ints(b)
		cnt := 0
		for j, v := range b {
			cnt ^= n - search(j+1, 3<<i-v) + search(j+1, 2<<i-v) - search(j+1, 1<<i-v)
		}
		ans |= cnt & 1 << i
	}
	Fprint(_w, ans)
}

//func main() { CF1323D(os.Stdin, os.Stdout) }

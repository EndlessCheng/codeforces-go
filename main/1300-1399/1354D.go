package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1354D(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	r := func() (x int) {
		in.Scan()
		data := in.Bytes()
		if data[0] == '-' {
			for _, b := range data[1:] {
				x = x*10 + int(b&15)
			}
			return -x
		}
		for _, b := range data {
			x = x*10 + int(b&15)
		}
		return
	}

	n, q := r(), r()
	a := make([]int, n)
	for i := range a {
		a[i] = r()
	}
	b := make([]int, q)
	for i := range b {
		b[i] = r()
	}
	ans := sort.Search(1e6+1, func(x int) bool {
		c := 0
		for _, v := range a {
			if v <= x {
				c++
			}
		}
		for _, v := range b {
			if v > 0 {
				if v <= x {
					c++
				}
			} else if -v <= c {
				c--
			}
		}
		return c > 0
	})
	if ans == 1e6+1 {
		ans = 0
	}
	Fprint(_w, ans)
}

//func main() { CF1354D(os.Stdin, os.Stdout) }

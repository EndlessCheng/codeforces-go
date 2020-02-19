package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF687B(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n, k := read(), read()
	if k == 1 {
		Fprint(out, "Yes")
		return
	}

	factors := []int{}
	exponents := []int{}
	kk := k
	for i := 2; i*i <= kk; i++ {
		cnt := 0
		for ; kk%i == 0; kk /= i {
			cnt++
		}
		if cnt > 0 {
			factors = append(factors, i)
			exponents = append(exponents, cnt)
		}
	}
	if kk > 1 {
		factors = append(factors, kk)
		exponents = append(exponents, 1)
	}

	match := make([]bool, len(factors))
	for i := 0; i < n; i++ {
		c := read()
		for j, f := range factors {
			cnt := 0
			for ; c%f == 0; c /= f {
				cnt++
			}
			if cnt >= exponents[j] {
				match[j] = true
			}
		}
	}
	for _, ok := range match {
		if !ok {
			Fprint(out, "No")
			return
		}
	}
	Fprint(out, "Yes")
}

//func main() { CF687B(os.Stdin, os.Stdout) }

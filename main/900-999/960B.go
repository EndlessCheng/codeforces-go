package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF960B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	
	var n, k, v int
	Fscan(in, &n, &k, &v)
	k += v
	a := make([]int, n, n+1)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := range a {
		Fscan(in, &v)
		a[i] = abs(a[i] - v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	a = append(a, 0)

	s := 0
	for i := 0; a[i] > 0; {
		v := a[i]
		for ; a[i] == v; i++ {
			s += v
		}
		if s-k >= a[i]*i {
			s -= k
			h, ex := int64(s/i), s%i
			ans := int64(ex)*(h+1)*(h+1) + int64(i-ex)*h*h
			for _, v := range a[i:] {
				ans += int64(v) * int64(v)
			}
			Fprint(out, ans)
			return
		}
	}
	Fprint(out, (k-s)%2)
}

//func main() { CF960B(os.Stdin, os.Stdout) }

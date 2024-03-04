package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1514C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	
	var n int
	Fscan(in, &n)
	ans := []int{}
	mul := int64(1)
	for i := 1; i < n; i++ {
		if gcd(i, n) == 1 {
			ans = append(ans, i)
			mul = mul * int64(i) % int64(n)
		}
	}
	if mul != 1 {
		i := sort.SearchInts(ans, int(mul))
		ans = append(ans[:i], ans[i+1:]...)
	}
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF1514C(os.Stdin, os.Stdout) }

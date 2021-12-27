package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF777E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]struct{ r, R, h int }, n)
	for i := range a {
		Fscan(in, &a[i].r, &a[i].R, &a[i].h)
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.R > b.R || a.R == b.R && a.r > b.r })

	sum := int64(a[0].h)
	ans := sum
	s := []int{0}
	for i := 1; i < n; i++ {
		for len(s) > 0 && a[s[len(s)-1]].r >= a[i].R {
			sum -= int64(a[s[len(s)-1]].h)
			s = s[:len(s)-1]
		}
		s = append(s, i)
		sum += int64(a[i].h)
		if sum > ans {
			ans = sum
		}
	}
	Fprint(out, ans)
}

//func main() { CF777E(os.Stdin, os.Stdout) }

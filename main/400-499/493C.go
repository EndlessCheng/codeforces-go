package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf493C(in io.Reader, out io.Writer) {
	var n, m, d int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		a[i] <<= 1
	}
	Fscan(in, &m)
	a = slices.Grow(a, m)
	for range m {
		Fscan(in, &d)
		a = append(a, d<<1|1)
	}
	slices.Sort(a)

	s := [2]int{n * 3, m * 3}
	ans0, ans1 := s[0], s[1]
	for i, v := range a {
		s[v&1]--
		if (i == len(a)-1 || v>>1 != a[i+1]>>1) && s[0]-s[1] > ans0-ans1 {
			ans0, ans1 = s[0], s[1]
		}
	}
	Fprint(out, ans0, ":", ans1)
}

//func main() { cf493C(bufio.NewReader(os.Stdin), os.Stdout) }

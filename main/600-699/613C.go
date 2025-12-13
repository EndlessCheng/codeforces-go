package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf613C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, odd, g int
	Fscan(in, &n)
	cnt := make([]int, n)
	for i := range cnt {
		Fscan(in, &cnt[i])
		odd += cnt[i] % 2
		g = gcd(g, cnt[i])
	}
	if odd > 1 {
		Fprintln(out, 0)
	} else {
		Fprintln(out, g)
	}

	s := []byte{}
	for i, c := range cnt {
		c /= g
		s = append(s, bytes.Repeat([]byte{'a' + byte(i)}, c/2)...)
	}
	rev := slices.Clone(s)
	slices.Reverse(rev)
	for i, c := range cnt {
		if c/g%2 > 0 {
			s = append(s, 'a'+byte(i))
		}
	}
	s = append(s, rev...)

	for range g {
		Fprintf(out, "%s", s)
		slices.Reverse(s)
	}
}

//func main() { cf613C(os.Stdin, os.Stdout) }

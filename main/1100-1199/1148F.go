package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1148F(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	v := make([]int, n)
	mask := make([]int, n)
	width := make([]int, n)
	s := 0
	for i := range n {
		Fscan(in, &v[i], &mask[i])
		s += v[i]
		width[i] = bits.Len(uint(mask[i])) - 1
	}

	if s < 0 {
		for i := range v {
			v[i] = -v[i]
		}
	}

	ans := 0
	for i := range 62 {
		s = 0
		for j, w := range width {
			if w == i {
				s += v[j]
			}
		}
		if s <= 0 {
			continue
		}
		ans |= 1 << i
		for j, m := range mask {
			if m>>i&1 > 0 {
				v[j] = -v[j]
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1148F(bufio.NewReader(os.Stdin), os.Stdout) }

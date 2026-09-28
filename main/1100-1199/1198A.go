package main

import (
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://github.com/EndlessCheng
func cf1198A(in io.Reader, out io.Writer) {
	var n, mx, ans, uni, left int
	Fscan(in, &n, &mx)
	mx *= 8
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)

	for i, x := range a {
		if i == 0 || x != a[i-1] {
			uni++
		}
		for n*bits.Len(uint(uni-1)) > mx {
			if a[left] != a[left+1] {
				uni--
			}
			left++
		}
		ans = max(ans, i-left+1)
	}
	Fprint(out, n-ans)
}

//func main() { cf1198A(bufio.NewReader(os.Stdin), os.Stdout) }

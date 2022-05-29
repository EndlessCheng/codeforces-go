package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF960C(in io.Reader, out io.Writer) {
	var x, d int
	Fscan(in, &x, &d)
	ans := []int64{}
	for v := int64(1); x > 0; v += int64(d) {
		n := bits.Len(uint(x))
		if 1<<n-1 > x {
			n--
		}
		x -= 1<<n - 1
		for ; n > 0; n-- {
			ans = append(ans, v)
		}
	}
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF960C(os.Stdin, os.Stdout) }

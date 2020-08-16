package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// github.com/EndlessCheng/codeforces-go
func Sol1278B(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var t int
	Fscan(in, &t)
	for case_ := 0; case_ < t; case_++ {
		Fprintln(out, solveCase(in, out))
	}
}

// https://oeis.org/A140358
// Smallest positive integer k such that n = +-1+-2+-...+-k for some choice of +'s and -'s.
func solveCase(in io.Reader, out io.Writer) int64 {
	var n, b int64
	Fscan(in, &n, &b)
	n -= b
	if n < 0 {
		n = -n
	}
	if n == 0 {
		return 0
	}
	d := math.Sqrt(float64(8*n + 1))
	k := int64((d - 1) / 2)
	tk := (k + 1) * k / 2
	if tk == n {
		return k
	}
	if k&1 == 1 {
		if (n-tk)&1 == 1 {
			return k + 2
		}
		return k + 1
	}
	if (n-tk)&1 == 1 {
		return k + 1
	}
	return k + 3
}

//func main() {
//	Sol1278B(os.Stdin, os.Stdout)
//}

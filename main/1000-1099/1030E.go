package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1030E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	ans := int64(0)
	var n, sum int
	var v uint64
	Fscan(in, &n)
	a := make([]int, n)
	cntS := [2]int{1}
	for i := range a {
		Fscan(in, &v)
		a[i] = bits.OnesCount64(v)
		sum ^= a[i] & 1
		ans += int64(cntS[sum])
		cntS[sum]++
		for j, mx, s := i, 0, 0; j >= 0 && s < 120; j-- {
			if a[j] > mx {
				mx = a[j]
			}
			s += a[j]
			if s&1 == 0 && mx*2 > s {
				ans--
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1030E(os.Stdin, os.Stdout) }

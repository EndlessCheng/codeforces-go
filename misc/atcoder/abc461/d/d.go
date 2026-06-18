package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func f(a []int, k int) (res int) {
	var s1, s2, l1, l2 int
	for i, x := range a {
		s1 += x
		s2 += x
		for l1 <= i && s1 >= k {
			s1 -= a[l1]
			l1++
		}
		for s2 > k {
			s2 -= a[l2]
			l2++
		}
		res += l1 - l2
	}
	return
}

func run(in io.Reader, out io.Writer) {
	var n, m, k, ans int
	Fscan(in, &n, &m, &k)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
		s := make([]int, m)
		for i2 := i; i2 >= 0; i2-- {
			for j, b := range a[i2] {
				s[j] += int(b - '0')
			}
			ans += f(s, k)
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

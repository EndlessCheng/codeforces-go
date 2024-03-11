package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf749E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, inv int
	Fscan(in, &n)
	sumI := 0.
	t := make([][2]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		v = n + 1 - v // 后缀变前缀
		s := 0
		for j := v; j > 0; j &= j - 1 {
			inv += t[j][0]
			s += t[j][1]
		}
		sumI += float64(s * (n + 1 - i))
		for j := v; j <= n; j += j & -j {
			t[j][0]++
			t[j][1] += i
		}
	}
	m := float64(n)
	Fprintf(out, "%.9f", float64(inv)-sumI/m/(m+1)*2+(m+2)*(m-1)/24)
}

//func main() { cf749E(os.Stdin, os.Stdout) }

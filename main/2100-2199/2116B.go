package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2116B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	pow2 := [1e5]int{1}
	for i := 1; i < len(pow2); i++ {
		pow2[i] = pow2[i-1] * 2 % mod
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		p := make([]int, n)
		for i := range p {
			Fscan(in, &p[i])
		}
		q := make([]int, n)
		j, k := 0, 0
		for i, v := range p {
			Fscan(in, &q[i])
			if v > p[j] {
				j = i
			}
			if q[i] > q[k] {
				k = i
			}
			if max(p[j], q[i-j])<<32|min(p[j], q[i-j]) > max(p[i-k], q[k])<<32|min(p[i-k], q[k]) {
				Fprint(out, (pow2[p[j]]+pow2[q[i-j]])%mod, " ")
			} else {
				Fprint(out, (pow2[p[i-k]]+pow2[q[k]])%mod, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf2116B(bufio.NewReader(os.Stdin), os.Stdout) }

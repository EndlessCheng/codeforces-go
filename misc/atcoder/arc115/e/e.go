package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	f := make([]int, n+1)
	sumF := make([]int, n+2)
	sumF[1] = 1
	st := []int{0}
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		for a[st[len(st)-1]] >= a[i] {
			st = st[:len(st)-1]
		}
		j := st[len(st)-1]
		k := i%2*2 - 1
		f[i] = (k*a[i]*(sumF[i]-sumF[j]) + (1-(i-j)%2*2)*f[j]) % mod
		sumF[i+1] = (sumF[i] - k*f[i]) % mod
		st = append(st, i)
	}
	Fprint(out, (f[n]+mod)%mod)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

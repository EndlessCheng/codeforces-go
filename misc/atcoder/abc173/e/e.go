package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
const mod = 1_000_000_007

func run(in io.Reader, out io.Writer) {
	var n, k, neg, posL, negL, posR, negR int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Slice(a, func(i, j int) bool { return abs(a[i]) > abs(a[j]) })
	if a[k-1] == 0 {
		Fprint(out, 0)
		return
	}

	mul := 1
	for _, v := range a[:k] {
		if v < 0 {
			v = -v
			neg++
			negL = v
		} else {
			posL = v
		}
		mul = mul * v % mod
	}
	if neg%2 == 0 {
		Fprint(out, mul)
		return
	}

	ans := mul
	for _, v := range a[k:] {
		if v > 0 {
			posR = v
			break
		}
	}
	for _, v := range a[k:] {
		if v < 0 {
			negR = -v
			break
		}
	}
	if (posL == 0 || negR == 0) && posR > 0 || posL*posR > negL*negR {
		ans = mul * pow(negL, mod-2) % mod * posR % mod
	} else if posL > 0 && negR > 0 {
		ans = mul * pow(posL, mod-2) % mod * negR % mod
	} else if a[n-1] == 0 {
		ans = 0
	} else {
		ans = 1
		for _, v := range a[n-k:] {
			ans = ans * abs(v) % mod
		}
		ans = mod - ans
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
func abs(x int) int { if x < 0 { return -x }; return x }
func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

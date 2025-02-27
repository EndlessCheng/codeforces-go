package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1406B(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	const k = 5
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		slices.SortFunc(a, func(a, b int) int { return abs(b) - abs(a) })
		if a[k-1] == 0 {
			Fprintln(out, 0)
			continue
		}

		var neg, posL, negL, posR, negR int
		mul := 1
		for _, v := range a[:k] {
			if v < 0 {
				v = -v
				neg++
				negL = v
			} else {
				posL = v
			}
			mul = mul * v
		}
		if neg%2 == 0 {
			Fprintln(out, mul)
			continue
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
			ans = mul / negL * posR
		} else if posL > 0 && negR > 0 {
			ans = mul / posL * negR
		} else if a[n-1] == 0 {
			ans = 0
		} else {
			ans = 1
			for _, v := range a[n-k:] {
				ans = ans * abs(v)
			}
			ans = -ans
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1406B(bufio.NewReader(os.Stdin), os.Stdout) }

package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1034A(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}
	calcGCD := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	calcGCDN := func(nums []int) (gcd int) {
		gcd = nums[0]
		for _, v := range nums[1:] {
			gcd = calcGCD(gcd, v)
		}
		return
	}

	const mx int = 1.5e7
	lpf := [mx + 1]int{}
	lpf[1] = 1
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	n := read()
	a := make([]int, n)
	for i := range a {
		a[i] = read()
	}

	cnt := [mx + 1]int{}
	g := calcGCDN(a)
	for i := range a {
		for v := a[i] / g; v > 1; {
			p := lpf[v]
			cnt[p]++
			for v /= p; lpf[v] == p; v /= p {
			}
		}
	}
	ans := 0
	for _, c := range cnt {
		if c > ans {
			ans = c
		}
	}
	if ans == 0 {
		Fprint(out, -1)
	} else {
		Fprint(out, n-ans)
	}
}

//func main() {
//	CF1034A(os.Stdin, os.Stdout)
//}

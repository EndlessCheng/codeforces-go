package main

import (
	"bufio"
	. "fmt"
	"io"
)

func Sol1152C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var a, b int64
	Fscan(in, &a, &b)
	if a > b {
		a, b = b, a
	}

	delta := b - a
	minLCM := int64(1e18)
	var k int64
	calcK := func(x int64) {
		ak := (a + x - 1) / x * x
		if lcm := ak / x * (ak + delta); lcm <= minLCM {
			minLCM = lcm
			k = ak - a
		}
	}
	// 设 a<b，由于 gcd(a+k, b+k) = gcd(a+k, b-a)，
	// 若 a+k 是 b-a 的某个因子 x 的倍数，则根据 gcd 的定义，gcd(a+k, b-a) 是 x 的倍数
	// 这样遍历所有 b-a 的因子，可以找到最小的 k，使得 gcd(a+k, b-a) 尽可能大，从而 lcm 尽可能小
	for i := int64(1); i*i <= delta; i++ {
		if delta%i == 0 {
			calcK(i)
			calcK(delta / i)
		}
	}
	Fprint(out, k)
}

//func main() {
//	Sol1152C(os.Stdin, os.Stdout)
//}

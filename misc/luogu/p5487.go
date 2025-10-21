package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
const mod5487 = 998244353

func pow5487(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod5487
		}
		x = x * x % mod5487
	}
	return res
}

func berlekampMassey5487(a []int) (coef []int) {
	var preC []int
	preI, preD := -1, 0
	for i, v := range a {
		d := v
		for j, c := range coef {
			d = (d - c*a[i-1-j]) % mod5487
		}
		if d == 0 {
			continue
		}

		// 首次算错，初始化 coef
		if preI < 0 {
			coef = make([]int, i+1)
			preI, preD = i, d
			continue
		}

		bias := i - 1 - preI
		oldSz := len(coef)
		sz := bias + len(preC) + 1
		var oldCoef []int
		if sz > oldSz {
			oldCoef = slices.Clone(coef)
			coef = slices.Grow(coef, sz-oldSz)[:sz]
		}

		// 上一次算错告诉我们，preD = a[preI] - sum_j preC[j]*a[preI-1-j]
		// 现在 a[i] = sum_j coef[j]*a[i-1-j] + d
		// 联立得 a[i] = sum_j coef[j]*a[i-1-j] + d/preD * (a[preI] - sum_j preC[j]*a[preI-1-j])
		// 其中 a[preI] 的系数 d/preD 位于当前（i）的 bias=i-1-preI 处
		// 注意：preI 之前的数据符合旧公式，即 0 = a[<preI] - sum_j preC[j]*a[<preI-1-j]
		//      对于新公式，i 之前的每一项增加了 d/preD * 0 = 0，所以也符合新公式
		delta := d * pow5487(preD, mod5487-2) % mod5487
		coef[bias] = (coef[bias] + delta) % mod5487
		for j, c := range preC {
			coef[bias+1+j] = (coef[bias+1+j] - delta*c) % mod5487
		}

		if sz > oldSz {
			preC = oldCoef
			preI, preD = i, d
		}
	}

	// 去掉不必要的 0
	for len(coef) > 0 && coef[len(coef)-1] == 0 {
		coef = coef[:len(coef)-1]
	}

	// 把负数调整为非负数
	// 比如后面计算递推式第 n 项，这可以保证不会产生负数（但那样的话，可以最后输出时再调整，所以下面的循环其实没必要）
	for i, c := range coef {
		coef[i] = (c + mod5487) % mod5487
	}

	return
}

func kitamasa5487(a, coef []int, n int) (ans int) {
	defer func() { ans = (ans%mod5487 + mod5487) % mod5487 }()
	if n < len(a) {
		return a[n]
	}

	k := len(coef)
	if k == 0 {
		return
	}
	if k == 1 {
		return a[0] * pow5487(coef[0], n)
	}

	// 比如 f(4) = 3*f(2) + 2*f(1) + f(0)
	// 或者说 f(n) = 3*f(n-2) + 2*f(n-3) + f(n-4)
	// 那么 f(8) = 3*f(6) + 2*f(5) + f(4)
	// 其中 f(5) = 3*f(3) + 2*f(2) + f(1)
	//           = 3*(用 f(2) f(1) f(0) 表出) + 2*f(2) + f(1)
	// f(6) 同理
	// 这样可以用 f(2) f(1) f(0)，也就是 a[2] a[1] a[0] 表出 f(8)
	mul := func(a, b []int) []int {
		c := make([]int, k)
		for _, v := range a {
			for j, w := range b {
				c[j] = (c[j] + v*w) % mod5487
			}
			// 原地计算下一组系数，比如上面已知 f(4) 的各项系数，现在要计算 f(5) 的各项系数
			bk := b[k-1]
			for i := k - 1; i > 0; i-- {
				b[i] = (b[i-1] + bk*coef[i]) % mod5487
			}
			b[0] = bk * coef[0] % mod5487
		}
		return c
	}

	// 计算 resC，以表出 f(n) = recC[k-1] * a[k-1] + recC[k-2] * a[k-2] + ... + resC[0] + a[0]
	resC := make([]int, k)
	resC[0] = 1
	c := make([]int, k)
	c[1] = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			resC = mul(c, resC)
		}
		c = mul(c, slices.Clone(c))
	}

	for i, c := range resC {
		ans = (ans + c*a[i]) % mod5487
	}
	return
}

func p5487(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	coef := berlekampMassey5487(a)
	for _, v := range coef {
		Fprint(out, v, " ")
	}
	Fprintln(out)

	slices.Reverse(coef)
	Fprint(out, kitamasa5487(a, coef, m))
}

//func main() { p5487(bufio.NewReader(os.Stdin), os.Stdout) }

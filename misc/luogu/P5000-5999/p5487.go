package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
const mod487 = 998244353

func pow487(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod487
		}
		x = x * x % mod487
	}
	return res
}

func berlekampMassey487(a []int) (coef []int) {
	var preC []int
	preI, preD := -1, 0
	for i, v := range a {
		// d = a[i] - 递推式算出来的值
		d := v
		for j, c := range coef {
			d = (d - c*a[i-1-j]) % mod487
		}
		if d == 0 { // 递推式正确
			continue
		}

		// 首次算错，初始化 coef 为 i+1 个 0
		if preI < 0 {
			coef = make([]int, i+1)
			preI, preD = i, d
			continue
		}

		bias := i - preI
		oldSz := len(coef)
		sz := bias + len(preC)
		var tmp []int
		if sz > oldSz { // 递推式变长了
			tmp = slices.Clone(coef)
			coef = slices.Grow(coef, sz-oldSz)[:sz]
		}

		// 上一次算错告诉我们，preD = a[preI] - sum_j preC[j]*a[preI-1-j]
		// 现在 a[i] = sum_j coef[j]*a[i-1-j] + d
		// 联立得 a[i] = sum_j coef[j]*a[i-1-j] + d/preD * (a[preI] - sum_j preC[j]*a[preI-1-j])
		// 其中 a[preI] 的系数 d/preD 位于当前（i）的 bias-1 = i-preI-1 处
		// 注意：preI 之前的数据符合旧公式，即 a[(<preI)] = sum_j preC[j]*a[(<preI)-1-j]
		//      对于新公式，i 之前的每个公式增加了 d/preD * (a[(<preI)] - sum_j preC[j]*a[(<preI)-1-j]) = d/preD * 0 = 0，所以也符合新公式
		delta := d * pow487(preD, mod487-2) % mod487
		coef[bias-1] = (coef[bias-1] + delta) % mod487
		for j, c := range preC {
			coef[bias+j] = (coef[bias+j] - delta*c) % mod487
		}

		if sz > oldSz {
			preC = tmp
			preI, preD = i, d
		}
	}

	// 计算完后，可能 coef 的末尾有 0，这些 0 不能去掉
	// 比如数列 {1,2,4,2,4,2,4,...} 的系数为 [0,1,0]，表示 f_n = 0*f_{n-1} + f_{n-2} + 0*f_{n-3} = f_{n-2} (n>=3)
	// 如果把末尾的 0 去掉，变成 [0,1]，就表示 f_n = 0*f_{n-1} + f_{n-2} = f_{n-2} (n>=2)
	// 看上去一样，但按照这个式子算出来的数列是错误的 {1,2,1,2,1,2,...}

	// 把负数调整为非负数（可以省略）
	for i, c := range coef {
		coef[i] = (c + mod487) % mod487
	}

	return
}

func kitamasa487(a, coef []int, n int) (ans int) {
	defer func() { ans = (ans%mod487 + mod487) % mod487 }()
	if n < len(a) {
		return a[n]
	}

	k := len(coef)
	if k == 0 {
		return
	}
	if k == 1 {
		return a[0] * pow487(coef[0], n)
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
				c[j] = (c[j] + v*w) % mod487
			}
			// 原地计算下一组系数，比如上面已知 f(4) 的各项系数，现在要计算 f(5) 的各项系数
			bk := b[k-1]
			for i := k - 1; i > 0; i-- {
				b[i] = (b[i-1] + bk*coef[i]) % mod487
			}
			b[0] = bk * coef[0] % mod487
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
		ans = (ans + c*a[i]) % mod487
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

	coef := berlekampMassey487(a)
	for _, v := range coef {
		Fprint(out, v, " ")
	}
	Fprintln(out)

	slices.Reverse(coef)
	Fprint(out, kitamasa487(a, coef, m))
}

//func main() { p5487(bufio.NewReader(os.Stdin), os.Stdout) }

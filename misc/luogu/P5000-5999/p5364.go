package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
const mod364 = 1_000_000_007

func pow364(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod364
		}
		x = x * x % mod364
	}
	return res
}

func kitamasa(coef, a []int, n int) (ans int) {
	defer func() { ans = (ans%mod364 + mod364) % mod364 }()
	if n < len(a) {
		return a[n]
	}

	k := len(coef)
	if k == 0 {
		return
	}
	if k == 1 {
		return a[0] * pow364(coef[0], n)
	}

	// 已知 f(n) 的各项系数为 a，f(m) 的各项系数为 b
	// 计算并返回 f(n+m) 的各项系数 c
	compose := func(a, b []int) []int {
		c := make([]int, k)
		for _, v := range a {
			// 累加 a[i] * f(m+i) 的各项系数
			for j, w := range b {
				c[j] = (c[j] + v*w) % mod364
			}
			// 从 f(m+i) 到 f(m+i+1)
			bk1 := b[k-1]
			for j := k - 1; j > 0; j-- {
				b[j] = (b[j-1] + bk1*coef[j]) % mod364
			}
			b[0] = bk1 * coef[0] % mod364
		}
		return c
	}

	// 计算 resC，以表出 f(n) = resC[k-1] * a[k-1] + resC[k-2] * a[k-2] + ... + resC[0] * a[0]
	resC := make([]int, k)
	resC[0] = 1
	c := make([]int, k)
	c[1] = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			resC = compose(c, resC)
		}
		c = compose(c, slices.Clone(c))
	}

	for i, c := range resC {
		ans = (ans + c*a[i]) % mod364
	}
	return
}

func berlekampMassey(a []int) (coef []int) {
	var preC []int
	preI, preD := -1, 0
	for i, v := range a {
		// d = a[i] - 递推式算出来的值
		d := v
		for j, c := range coef {
			d = (d - c*a[i-1-j]) % mod364
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
		oldLen := len(coef)
		newLen := bias + len(preC)
		var tmp []int
		if newLen > oldLen { // 递推式变长了
			tmp = slices.Clone(coef)
			coef = slices.Grow(coef, newLen-oldLen)[:newLen] // coef.resize(newLen)
		}

		// 历史错误为 preD = a[preI] - sum_j preC[j]*a[preI-1-j]
		// 现在 a[i] = sum_j coef[j]*a[i-1-j] + d
		// 联立得 a[i] = sum_j coef[j]*a[i-1-j] + d/preD * (a[preI] - sum_j preC[j]*a[preI-1-j])
		// 其中 a[preI] 的系数 d/preD 位于当前（i）的 bias-1 = i-preI-1 处
		// 注意：preI 之前的数据符合旧公式，即 a[(<preI)] = sum_j preC[j]*a[(<preI)-1-j]
		//      对于新公式，i 之前的每个公式增加了 d/preD * (a[(<preI)] - sum_j preC[j]*a[(<preI)-1-j]) = d/preD * 0 = 0，所以也符合新公式
		delta := d * pow364(preD, mod364-2) % mod364
		coef[bias-1] = (coef[bias-1] + delta) % mod364
		for j, c := range preC {
			coef[bias+j] = (coef[bias+j] - delta*c) % mod364
		}

		if newLen > oldLen {
			preC = tmp
			preI, preD = i, d
		}
	}

	// 计算完后，可能 coef 的末尾有 0，这些 0 不能去掉
	// 比如数列 (1,2,4,2,4,2,4,...) 的系数为 [0,1,0]，表示 f(n) = 0*f(n-1) + 1*f(n-2) + 0*f(n-3) = f(n-2)   (n >= 3)
	// 如果把末尾的 0 去掉，变成 [0,1]，就表示 f(n) = 0*f(n-1) + f(n-2) = f(n-2)   (n >= 2)
	// 看上去一样，但按照这个式子算出来的数列是错误的 (1,2,1,2,1,2,...)

	// 手动找规律用
	for i, c := range coef {
		if c < -mod364/2 {
			c += mod364
		} else if c > mod364/2 {
			c -= mod364
		}
		coef[i] = c
	}

	return
}

func guessNth(a []int, n int) int {
	coef := berlekampMassey(a)
	slices.Reverse(coef)
	return kitamasa(coef, a, n)
}

func p5364(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := []int{}
	s := 0
	for i := 1; i <= k*2+2; i++ {
		a = append(a, (s+pow364(i, k))%mod364)
		s = (s + a[len(a)-1]) % mod364
	}
	Fprint(out, guessNth(a, n-1))
}

//func main() { p5364(bufio.NewReader(os.Stdin), os.Stdout) }

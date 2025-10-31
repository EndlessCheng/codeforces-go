## 引言

你在小学一定做过这样的题目：

- 找规律，在括号内填入合适的数：$1,\ 1,\ 2,\ 3,\ 5,\ 8,\ (\ \ \ )$。

「合适」的意思是找到最简单的规律。由于从 $2$ 开始，每一项都是前两项之和，所以答案是 $5+8=13$。

如果数列变得更复杂，规律就不好找了：

- $1,\ 1,\ 2,\ 5,\ 10,\ 22,\ 47,\ (\ \ \ )$。

能不能发明一台机器，输入数列的前若干项，就能显示这个数列的规律？

来试一试吧，我将带你发明这台机器！

设数列为 $a$。我们要做的，是用一个统一的式子刻画每个 $a_n$ 与前面 $a_{n-1},a_{n-2},\ldots$ 若干个数的关系，这个式子叫做**递推式**。

> **注**：本文限定递推式为**常系数齐次线性递推式**，例如 $a_n = a_{n-1}+a_{n-2}$、$a_n = 2a_{n-1}-3a_{n-3}$，等号右侧每个 $a_i$ 只能乘以一个常数系数。特别注意，递推式**不包含常数项**，例如 $a_n = 2a_{n-1} + a_{n-2} + 3$，其中 $3$ 为常数项。

## 例子

设 $a = (1,1,2,5,10,22,47)$。我们要计算 $a$ 的最短递推式，根据递推式算出 $47$ 的下一项。

#### 1)

前两项是 $a_0 = 1$，$a_1 = 1$，猜测递推式为 $a_n = a_{n-1}\ (n\ge 1)$。

#### 2)

用猜测的递推式计算下一项 $a'_2=a_1 = 1\ne 2$，比实际值少 $1$，即实际上 $a_2 = a_1 + 1$。

尝试修复，把 $1$ 用 $a_0$ 代替，得 $a_2=a_1+a_0$，猜测递推式为 $a_n = a_{n-1} + a_{n-2}\ (n\ge 2)$。

> **注**：也可以猜测递推式为 $a_n = 2a_{n-1}\ (n\ge 2)$。注意这里限定 $n\ge 2$，因为当 $n=1$ 时递推式不符合。限定 $n\ge 2$ 也意味着把 $a_0=1$ 和 $a_1=1$ 作为初始值。

#### 3)

继续计算下一项 $a'_3 = a_2 + a_1 = 3\ne 5$，比实际值少 $2$，即实际上 $a_3 = a_2 + a_1 + 2$。

如何修复这个错误？你能找到一个合适的递推式吗？

请注意，不能随意修复这个错误，比如猜测递推式为 $a_n = 2a_{n-1}+a_{n-2}\ (n\ge 2)$，虽然 $a_3$ 没问题了，但对于之前的项，用递推式算出的 $a'_2 = 2a_1+a_0 = 3\ne 2$。我们猜测的递推式不仅要能修复当前的错误，还要保证之前的项也符合递推式。不能修复了一个 bug，又引入了新的 bug。

Berlekamp-Massey 算法（BM 算法）的**关键想法**来了：

- 上次算错的时候，只有 $a_2 = a_1$ 是错的，前面的 $a_1 = a_0$ 是正确的。
- 如果把之前的错误「**平移**」到当前错误上，我们不仅修复了当前错误，也保证新的递推式适用于之前的项。

之前算错的时候，我们知道 $a_2 = a_1 + 1$，也就是 $1 = a_2 - a_1$。

现在的错误表明 $a_3 = a_2+a_1 + 2$。把 $1 = a_2 - a_1$ 两边同时乘以 $2$，得到 $2 = 2(a_2-a_1)$，代入 $a_3 = a_2+a_1 + 2$ 中的 $2$，得到

$$
\begin{aligned}
a_3 &= a_2 + a_1 + 2(a_2-a_1)      \\
&= 3a_2 - a_1                  \\
\end{aligned}
$$

同样地，猜测递推式为

$$
\begin{aligned}
a_n &= a_{n-1} + a_{n-2} +2(a_{n-1}-a_{n-2})       \\
&= 3a_{n-1} - a_{n-2}\ (n\ge 2)                \\
\end{aligned}
$$

当 $n=3$ 时，递推式是正确的。$n=2$ 呢？

也是正确的，从化简前的式子可以看出来。令 $a_n = a_{n-1} + a_{n-2} +2(a_{n-1}-a_{n-2})$ 中的 $n=2$，得

$$
\begin{aligned}
a_2 &= a_1 + a_0 + 2(a_1 - a_0)      \\
&= a_1 + a_0 + 2\cdot 0           \\
&= a_1 + a_0           \\
\end{aligned}
$$

所以新递推式兼容旧数据（$a_3$ 之前的项）。

#### 4)

继续计算下一项 $a'_4 = 3a_3 - a_2 = 13\ne 10$，比实际值多 $3$，即实际上 $a_4 = 3a_3 - a_2 - 3$。

同样地，把 $1 = a_2 - a_1$ 两边同时乘以 $-3$，得到 $-3 = -3(a_2-a_1)$，代入 $a_4 = 3a_3 - a_2 - 3$ 中的 $-3$，得到

$$
\begin{aligned}
a_4 &= 3a_3 - a_2 - 3              \\
&= 3a_3 - a_2 - 3(a_2-a_1)                 \\
&= 3a_3 - 4a_2 + 3a_1               \\
\end{aligned}
$$

猜测递推式为 $a_n = 3a_{n-1} - 4a_{n-2} + 3a_{n-3}\ (n\ge 3)$。

> **注**：也可以用上个错误 $2 = a_3 - (a_2+a_1)$ 修复问题。如果有多个历史错误，选哪个最好？请继续阅读。

#### 5)

继续计算下一项 $a'_5 = 3a_4 - 4a_3 + 3a_2 = 16\ne 22$，比实际值少 $6$，即实际上 $a_5 = 3a_4 - 4a_3 + 3a_2 + 6$。

我们有多个历史错误可以修复该问题：

- 选择用 $1 = a_2 - a_1$？
- 还是用 $2 = a_3 - (a_2+a_1)$？
- 还是用 $-3 = a_4 - (3a_3 - a_2)$？

都可以。但别忘了，我们还要让递推式尽量短。如果用 $1 = a_2 - a_1$ 或者 $2 = a_3 - (a_2+a_1)$ 修复，得到的 $a_5$ 会包含 $a_1$，或者说 $a_n$ 的递推式会包含 $a_{n-4}$。而用 $-3 = a_4 - (3a_3 - a_2)$ 修复，包含的最小项 $a_2$ 的下标 $2$ 是最大的，可以得到尽量短的递推式。

把 $-3 = a_4 - (3a_3 - a_2)$ 两边同时乘以 $-2$，得到 $6 = -2[a_4 - (3a_3 - a_2)]$，代入 $a_5 = 3a_4 - 4a_3 + 3a_2 + 6$ 中的 $6$，得到

$$
\begin{aligned}
a_5 &= 3a_4 - 4a_3 + 3a_2 + 6              \\
&= 3a_4 - 4a_3 + 3a_2 -2[a_4 - (3a_3 - a_2)]                \\
&= a_4 + 2a_3 + a_2               \\
\end{aligned}
$$

猜测递推式为 $a_n = a_{n-1} + 2a_{n-2} + a_{n-3}\ (n\ge 3)$。

继续计算下一项 $a'_6 = a_5 + 2a_4 + a_3 = 47 = a_6$，符合。

至此，我们找到了符合 $a$ 的递推式 $a_n = a_{n-1} + 2a_{n-2} + a_{n-3}\ (n\ge 3)$，初始项 $a_0=1,a_1=1,a_2=2$。

所以 $a_6=47$ 的合适的下一项为 $a_7 = a_6 + 2a_5 + a_4 = 101$。

#### 小结

1. 把历史错误「**平移**」到当前错误上，不仅能修复当前错误，还能保证新递推式兼容旧数据。
2. 选择满足「**最小项的下标最大**」的历史错误，可以得到尽量短的递推式。

读者可以根据上述内容，尝试编程实现这台机器。

---

在讲解一般公式和代码之前，我还想带你挖掘更多的性质，希望能让你更清晰地理解这个算法。

上面计算 $a_4,a_5,a_6$ 时，式子的长度没变，即

$$
\begin{aligned}
a_4 &= 3a_3 - 4a_2 + 3a_1   \\
a_5 &= a_4 + 2a_3 + a_2     \\
a_6 &= a_5 + 2a_4 + a_3     \\
\end{aligned}
$$

观察等号右边的项，下标范围依次为 $[1,3]\to [2,4]\to [3,5]$，这是个 [滑动窗口](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/)，等号右边最小项的下标在逐渐变大。

这意味着什么？一旦等号右边最小项的下标超过了历史错误最小项的下标，那么出错的时候，得到的递推式会**变长**。

对于上文的例子，假如实际上 $a_7=107$，也就是 $a_7 = a_6 + 2a_5 + a_4 + 6$，那么把历史错误 $-3 = a_4 - (3a_3 - a_2)$ 代入，得到

$$
\begin{aligned}
a_7 &= a_6 + 2a_5 + a_4 + 6              \\
&= a_6 + 2a_5 + a_4 -2[a_4 - (3a_3 - a_2)]                \\
&= a_6 + 2a_5 - a_4 + 6a_3 - 2a_2               \\
\end{aligned}
$$

新的递推式为 $a_n = a_{n-1} + 2a_{n-2} - a_{n-3} + 6a_{n-4} -2a_{n-5}\ (n\ge 5)$。

这里同时发生了两件事情：

1. 递推式变长了。
2. 产生了一个新的错误 $6 = a_7 - (a_6 + 2a_5 + a_4)$，并且最小项的下标比其他历史错误更大。

其实，这两件事情一定会同时发生。

我们修正 $a_7 = a_6 + 2a_5 + a_4 + 6$ 时，最小项为 $a_4$，其下标 $4$ 比历史错误的最小项 $a_2$ 的下标 $2$ 还要大，那么把递推式和历史错误拼起来（联立），递推式就一定会变长。并且，新产生的错误 $6 = a_7 - (a_6 + 2a_5 + a_4)$ 的最小项也是 $a_4$，$4 > 2$，新错误的最小项的下标是更大的。

所以**只要递推式变长，就可以把新产生的错误保存起来，用于修正后续错误**。如果递推式没有变长，那么旧的历史错误仍然是最优的，无需替换。

## 一般公式

设保存的历史错误为

$$
a_{\textit{preI}} = \left(\sum_{j=0}^{k'-1} c'_j a_{\textit{preI}-1-j}\right) + \textit{preD}
$$

也就是

$$
\textit{preD} = a_{\textit{preI}} - \sum_{j=0}^{k'-1} c'_j a_{\textit{preI}-1-j}
$$

如果当前计算 $a_i$ 出错，即

$$
a_{i} = \left(\sum_{j=0}^{k-1} c_j a_{i-1-j}\right) + d
$$

联立上面两个式子，得

$$
a_{i} = \left(\sum_{j=0}^{k-1} c_j a_{i-1-j}\right) + \dfrac{d}{\textit{preD}} \left(a_{\textit{preI}} - \sum_{j=0}^{k'-1} c'_j a_{\textit{preI}-1-j}\right)
$$

设 $\textit{bias} = i - \textit{preI}$。上式 $a_{\textit{preI}}$ 的系数 $\delta = \dfrac{d}{\textit{preD}}$ 位于系数列表 $c$ 的下标 $\textit{bias}-1$ 处。

我们在系数列表 $c$ 上修改：

- 系数列表的长度调整为 $\max(\textit{bias}+k',k)$。
- 下标为 $\textit{bias}-1$ 的系数增加了 $\delta$。
- 对于 $j=0,1,2,\ldots,k'-1$，下标为 $\textit{bias}+j$ 的系数减少了 $\delta\cdot c'_j$。

如果新系数列表的长度比原来的长，即 $\textit{bias}+k' > k$，那么记录系数 $c$（更新前的）、当前下标 $i$ 和偏差值 $d$，作为历史错误。

结合 [Kitamasa 算法](https://zhuanlan.zhihu.com/p/1964051212304364939)，我们可以预测数列第 $n$ 项（$n$ 从 $0$ 开始）的值。

## 算法实现

```go
package main
import("bufio";."fmt";"os";"slices")

const mod = 998244353

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

// 给定数列的前 m 项 a，返回符合 a 的最短常系数齐次线性递推式的系数 coef（模 mod 意义下）
// 设 coef 长为 k，当 n >= k 时，有递推式 f(n) = coef[0] * f(n-1) + coef[1] * f(n-2) + ... + coef[k-1] * f(n-k)  （注意 coef 的顺序）
// 初始值 f(n) = a[n]  (0 <= n < k)
// 时间复杂度 O(m^2)，其中 m 是 a 的长度
func berlekampMassey(a []int) (coef []int) {
	var preC []int
	preI, preD := -1, 0

	for i, v := range a {
		// d = a[i] - 递推式算出来的值
		d := v
		for j, c := range coef {
			d = (d - c*a[i-1-j]) % mod
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
		delta := d * pow(preD, mod-2) % mod // pow(preD, mod-2) 为 preD 的逆元
		coef[bias-1] = (coef[bias-1] + delta) % mod
		for j, c := range preC {
			coef[bias+j] = (coef[bias+j] - delta*c) % mod
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

	// 把负数调整为非负数
	for i, c := range coef {
		coef[i] = (c + mod) % mod
	}

	return
}

// 给定常系数齐次线性递推式 f(n) = coef[k-1] * f(n-1) + coef[k-2] * f(n-2) + ... + coef[0] * f(n-k)
// 以及初始值 f(i) = a[i] (0 <= i < k)
// 返回 f(n) % mod，其中参数 n 从 0 开始
// 注意 coef 的顺序
// 时间复杂度 O(k^2 log n)，其中 k 是 coef 的长度
func kitamasa(coef, a []int, n int) (ans int) {
	defer func() { ans = (ans + mod) % mod }() // 保证结果非负
	if n < len(a) {
		return a[n] % mod
	}

	k := len(coef)
	// 特判 k = 0, 1 的情况
	if k == 0 {
		return 0
	}
	if k == 1 {
		return a[0] * pow(coef[0], n) % mod
	}

	// 已知 f(n) 的各项系数为 a，f(m) 的各项系数为 b
	// 计算并返回 f(n+m) 的各项系数 c
	compose := func(a, b []int) []int {
		c := make([]int, k)
		for _, v := range a {
			for j, w := range b {
				c[j] = (c[j] + v*w) % mod
			}
			// 原地计算下一组系数，比如已知 f(4) 的各项系数，现在要计算 f(5) 的各项系数
			// 倒序遍历，避免提前覆盖旧值
			bk1 := b[k-1]
			for i := k - 1; i > 0; i-- {
				b[i] = (b[i-1] + bk1*coef[i]) % mod
			}
			b[0] = bk1 * coef[0] % mod
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
		// 由于会修改 compose 的第二个参数，这里把 c 复制一份再传入
		c = compose(c, slices.Clone(c))
	}

	for i, c := range resC {
		ans = (ans + c*a[i]) % mod
	}

	return
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	coef := berlekampMassey(a)
	for _, v := range coef {
		Fprint(out, v, " ")
	}
	Fprintln(out)

	slices.Reverse(coef) // 注意 kitamasa 入参的顺序
	Fprint(out, kitamasa(coef, a, m))
}
```

**时间复杂度**：$\mathcal{O}(n^2\log m)$，其中 $n$ 是 $a$ 的长度。Berlekamp-Massey 算法需要 $\mathcal{O}(n^2)$ 时间，Kitamasa 算法需要 $\mathcal{O}(n^2\log m)$ 时间。

谢谢阅读，欢迎点赞~

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1903D2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, v, k, total, maxVal int
	Fscan(in, &n, &q)

	const w = 20
	const u = 1 << w
	cnt := [u]int{}
	f := [u][w]int{}
	for range n {
		Fscan(in, &v)
		total += v
		maxVal = max(maxVal, v)
		cnt[v]++
		// 遍历 v 的每个 0
		for j := uint32(u - 1 ^ v); j > 0; j &= j - 1 {
			i := bits.TrailingZeros32(j)
			f[v][i] += v & (1<<i - 1) // 累加 v 的低 i 位
		}
	}

	// 算完 SOS DP（从超集转移到当前状态）后：
	// 二进制包含 s 的元素，有 cnt[s] 个
	// 对于二进制包含 s 且第 i 位是 0 的元素，累加这些元素的低 i 位之和，即 f[s][i]
	for i := range w {
		for s := 0; s < u; s++ {
			s |= 1 << i
			cnt[s^1<<i] += cnt[s]
			for j := range w {
				f[s^1<<i][j] += f[s][j]
			}
		}
	}

	for range q {
		Fscan(in, &k)
		avg := (total + k) / n
		if avg >= maxVal {
			Fprintln(out, avg)
			continue
		}
		ans := 0
		for i := w - 1; i >= 0; i-- {
			// 现在我们要计算，让答案第 i 位是 1，代价是多少。也就是元素要包含 ans|1<<i
			// 对于（在所有操作之前）已经包含 ans|1<<i 的元素，无需操作，代价是 0。设这样的元素有 cnt[ans|1<<i]] 个
			// 其余 n-cnt[ans|1<<i]] 个元素呢？可以分为两类：
			// 第一类是（在所有操作之前）包含 ans，但不包含 1<<i 的元素。这些元素要增大，比如从 001 到 100，需要 +3，而不是 +4
			// 第二类是（在所有操作之前）不包含 ans 的元素。由于我们已经在之前的循环中增加了这些数，这些元素的低 i 位现在都是 0。这些元素要增大，比如从 000 到 100，直接 +4
			// 我们可以先增加 n-cnt 个 2^i，再减去第一类元素多操作的次数，即第一类元素的低 i 位之和，记作 f[ans][i]
			// 综上，预处理 cnt 和 f，就可以 O(1) 求出让答案第 i 位是 1 的代价：(n-cnt[ans|1<<i])<<i - f[ans][i]
			cost := (n-cnt[ans|1<<i])<<i - f[ans][i]
			if cost <= k {
				k -= cost
				ans |= 1 << i
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1903D2(bufio.NewReader(os.Stdin), os.Stdout) }

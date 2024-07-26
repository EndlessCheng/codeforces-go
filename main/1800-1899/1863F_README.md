## 一

首先，有如下 $\mathcal{O}(n^3)$ 区间 DP 写法：

从 $f(1, n)$ 开始递归。

对于 $f(L,R)$，枚举在 $[L,R]$ 中的下标 $k$：

- 如果 $[L,k]$ 中的异或和比 $[k+1,R]$ 中的异或和大，则往下递归 $f(L,k)$，表示我们可以得到子数组 $[L,k]$。
- 如果 $[L,k]$ 中的异或和比 $[k+1,R]$ 中的异或和小，则往下递归 $f(k+1,R)$，表示我们可以得到子数组 $[k+1,R]$。
- 如果一样大，则往下递归 $f(L,k)$ 以及 $f(k+1,R)$，表示子数组 $[L,k]$ 和 $[k+1,R]$ 都可以得到。

如果递归到 $L=R$ 的状态，则说明能通过操作得到 $a_L$。

如何优化这个区间 DP？

## 二

从 $f(i,j)$ 的角度上看，需要知道：

- 是否存在 $R$，使得我们可以从 $f(i,R)$ 递归到 $f(i,j)$。
- 是否存在 $L$，使得我们可以从 $f(L,j)$ 递归到 $f(i,j)$。

当然，前提是我们能递归到 $f(i,R)$ 或者 $f(L,j)$。

如果可以 $\mathcal{O}(1)$ 算出能否递归到 $f(i,j)$，就可以用 $\mathcal{O}(n^2)$ 的时间顺利通过本题。

设 $s_1$ 为下标在 $[i,R]$ 的元素异或和，$s_2$ 为下标在 $[i,j]$ 的元素异或和。

如果可以从 $f(i,R)$ 递归到 $f(i,j)$，说明 $[j+1,R]$ 的异或和小于等于 $s_2$，即 $s_1 \operatorname{xor} s_2 \le s_2$。

什么情况下 $s_1 \operatorname{xor} s_2 \le s_2$ 成立呢？

## 三

首先考虑特殊情况，如果 $s_1 = 0$，那么 $s_1 \operatorname{xor} s_2 \le s_2$ 恒成立。

如果 $s_1 > 0$，**从高到低**考虑每个比特位：

- 如果 $s_1$ 这一位是 $0$，那么 $s_2$ 这一位是 $0$ 是 $1$ 都可以，并且没法判断 $s_1 \operatorname{xor} s_2$ 和 $s_2$ 谁大谁小。
- 如果 $s_1$ 这一位是 $1$，那么 $s_2$ 这一位也必须是 $1$，且此时 $s_1 \operatorname{xor} s_2 < s_2$ 成立。如果 $s_2$ 这一位是 $0$，那么 $s_1 \operatorname{xor} s_2 > s_2$。由于确定清楚了大小关系，所以考虑到此就可以盖棺定论了。由于我们是从高到低遍历的，所以 $s_1$ 的**最高位**的 $1$ 必须在 $s_2$ 中。

根据 [从集合论到位运算](https://leetcode.cn/circle/discuss/CaOJ45/)，用集合语言描述，我们要满足：集合 $s_1$ 是空集，或者集合 $s_1$ 的最大元素一定在集合 $s_2$ 中。如果不满足该条件，则 $s_1 \operatorname{xor} s_2 > s_2$，我们无法从 $f(i,R)$ 递归到 $f(i,j)$。

对于 $[i,j+1], [i,j+2], \cdots, [i,n-1]$ 中的**能递归到的区间**，计算每个区间的异或和，把异或和的二进制视作集合，把集合的最大元素，记录到集合 $\textit{rightBits}$ 中。如果有异或和是 $0$，则直接把 $\textit{rightBits}$ 标记为 $-1$。
如果 $\textit{rightBits}=-1$ 或者 $\textit{rightBits}$ 和 $s_2$ 的交集不为空，即 $\textit{rightBits} \operatorname{and} s_2 \ne 0$，则说明我们可以递归到 $f(i,j)$。

上面讨论的是从 $f(i,R)$ 递归到 $f(i,j)$ 的情况，对于从 $f(L,j)$ 递归到 $f(i,j)$ 的情况也同理。

下面 Go 代码下标从 $0$ 开始。

```go
package main
import ("bufio";."fmt";"math/bits";"os")

func main() {
	in := bufio.NewReader(os.Stdin)
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		sum := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &sum[i])
			sum[i] ^= sum[i-1]
		}
		leftBits := make([]int, n)
		for i := 0; i < n; i++ {
			rightBits := 0
			for j := n - 1; j >= i; j-- {
				s2 := sum[j+1] ^ sum[i]
				ok := i == 0 && j == n-1 || // 递归入口
					rightBits < 0 || rightBits&s2 != 0 ||  // 能从 f(i,R) 递归到 f(i,j)
					leftBits[j] < 0 || leftBits[j]&s2 != 0 // 能从 f(L,j) 递归到 f(i,j)
				if ok {
					if s2 == 0 {
						leftBits[j] = -1
						rightBits = -1
					} else {
						high := 1 << (bits.Len(uint(s2)) - 1)
						leftBits[j] |= high
						rightBits |= high
					}
				}
				if j == i {
					if ok {
						Print("1")
					} else {
						Print("0")
					}
				}
			}
		}
		Println()
	}
}
```

**时间复杂度**：$\mathcal{O}(n^2)$。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

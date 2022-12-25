推荐先看我写的这篇题解：[878. 第 N 个神奇数字](https://leetcode.cn/problems/nth-magical-number/solutions/1984641/er-fen-da-an-rong-chi-yuan-li-by-endless-9j34/)。

---

一看到「最大值的最小值」就想到二分答案。

下文把 $\textit{divisor}_1$ 和 $\textit{divisor}_2$ 简写成 $d_1$ 和 $d_2$，记 $\textit{LCM}$ 为 $d_1$ 和 $d_2$ 的最小公倍数。

由于：

- 能被 $d_2$ 整除但不能被 $d_1$ 整除的数，能在 $\textit{arr}_1$ 中且不能在 $\textit{arr}_2$ 中；
- 能被 $d_1$ 整除但不能被 $d_2$ 整除的数，能在 $\textit{arr}_2$ 中且不能在 $\textit{arr}_1$ 中；
- 既不能被 $d_1$ 整除也不能被 $d_2$ 整除的数，可以在 $\textit{arr}_1$ 和 $\textit{arr}_2$ 中。

因此二分答案 $x$，则：

- 有 $\left\lfloor\dfrac{x}{d_2}\right\rfloor - \left\lfloor\dfrac{x}{\textit{LCM}}\right\rfloor$ 个数是 $\textit{arr}_1$ 独享的；
- 有 $\left\lfloor\dfrac{x}{d_1}\right\rfloor - \left\lfloor\dfrac{x}{\textit{LCM}}\right\rfloor$ 个数是 $\textit{arr}_2$ 独享的；
- 有 $x - \left\lfloor\dfrac{x}{d_1}\right\rfloor - \left\lfloor\dfrac{x}{d_2}\right\rfloor + \left\lfloor\dfrac{x}{\textit{LCM}}\right\rfloor$ 个数（根据容斥原理）是 $\textit{arr}_1$ 和 $\textit{arr}_2$ 共享的。

去掉独享的，剩余的数字只能在共享中选择，因此

$$
x - \left\lfloor\dfrac{x}{d_1}\right\rfloor - \left\lfloor\dfrac{x}{d_2}\right\rfloor + \left\lfloor\dfrac{x}{\textit{LCM}}\right\rfloor \ge \max(\textit{uniqueCnt}_1 - \left\lfloor\dfrac{x}{d_2}\right\rfloor + \left\lfloor\dfrac{x}{\textit{LCM}}\right\rfloor, 0) + \max(\textit{uniqueCnt}_2 - \left\lfloor\dfrac{x}{d_1}\right\rfloor + \left\lfloor\dfrac{x}{\textit{LCM}}\right\rfloor, 0)
$$

为二分判定条件。

代码实现时，二分上界可以取 $(\textit{uniqueCnt}_1 + \textit{uniqueCnt}_2)\cdot 2-1$，因为最坏情况下 $d_1=d_2=2$，只能取奇数。

有关二分的写法，可以看我的 [【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/) 这期视频。

```py [sol1-Python3]
class Solution:
    def minimizeSet(self, d1: int, d2: int, uniqueCnt1: int, uniqueCnt2: int) -> int:
        lcm = math.lcm(d1, d2)
        def check(x: int) -> bool:
            left1 = max(uniqueCnt1 - x // d2 + x // lcm, 0)
            left2 = max(uniqueCnt2 - x // d1 + x // lcm, 0)
            common = x - x // d1 - x // d2 + x // lcm
            return common >= left1 + left2
        return bisect_left(range((uniqueCnt1 + uniqueCnt2) * 2 - 1), True, key=check)
```

```go [sol1-Go]
func minimizeSet(d1, d2, uniqueCnt1, uniqueCnt2 int) int {
	lcm := d1 / gcd(d1, d2) * d2
	return sort.Search((uniqueCnt1+uniqueCnt2)*2-1, func(x int) bool {
		left1 := max(uniqueCnt1-x/d2+x/lcm, 0)
		left2 := max(uniqueCnt2-x/d1+x/lcm, 0)
		common := x - x/d1 - x/d2 + x/lcm
		return common >= left1+left2
	})
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$O(\log(\textit{divisor}_1+\textit{divisor}_2) + \log(\textit{uniqueCnt}_1+\textit{uniqueCnt}_2))$。
- 空间复杂度：$O(1)$，仅用到若干变量。

#### 相似题目

- [878. 第 N 个神奇数字](https://leetcode.cn/problems/nth-magical-number/)
- [1201. 丑数 III](https://leetcode.cn/problems/ugly-number-iii/)

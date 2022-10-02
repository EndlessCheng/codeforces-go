[视频讲解](https://www.bilibili.com/video/BV1kd4y1q7fC) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

设 $n_1$ 为 $\textit{num}_1$ 的二进制表示的长度，$c_1$ 为 $\textit{num}_1$ 的置位数，$c_2$ 为 $\textit{num}_2$ 的置位数。

基本思路：

$x$ 的置位数和 $\textit{num}_2$ 相同，意味着 $x$ 的二进制表示中有 $c_2$ 个 $1$，我们需要合理地分配这 $c_2$ 个 $1$。

为了让异或和尽量小，这些 $1$ 应当从高位到低位匹配 $\textit{num}_1$ 中的 $1$；如果匹配完了还有多余的 $1$，那么就从低位到高位把 $0$ 改成 $1$。

分类讨论：

- 如果 $c_2\ge n_1$，$x$ 只能是 $2^{c_2}-1$，任何其他方案都会使异或和变大；
- 如果 $c_2=c_1$，那么 $x=\textit{num}_1$；
- 如果 $c_2<c_1$，那么将 $\textit{num}_1$ 的最低的 $c_1-c_2$ 个 $1$ 变成 $0$，其结果就是 $x$；
- 如果 $c_2>c_1$，那么将 $\textit{num}_1$ 的最低的 $c_2-c_1$ 个 $0$ 变成 $1$，其结果就是 $x$；

```py [sol1-Python3]
class Solution:
    def minimizeXor(self, num1: int, num2: int) -> int:
        c1 = num1.bit_count()
        c2 = num2.bit_count()
        while c2 < c1:
            num1 &= num1 - 1  # 最低的 1 变成 0
            c2 += 1
        while c2 > c1:
            num1 |= num1 + 1  # 最低的 0 变成 1
            c2 -= 1
        return num1
```

```go [sol1-Go]
func minimizeXor(num1, num2 int) int {
	c1 := bits.OnesCount(uint(num1))
	c2 := bits.OnesCount(uint(num2))
	for ; c2 < c1; c2++ {
		num1 &= num1 - 1 // 最低的 1 变成 0
	}
	for ; c2 > c1; c2-- {
		num1 |= num1 + 1 // 最低的 0 变成 1
	}
	return num1
}
```

#### 复杂度分析

- 时间复杂度：$O(|\log\textit{num}_1 - \log\textit{num}_2|)$。
- 空间复杂度：$O(1)$，仅用到若干变量。

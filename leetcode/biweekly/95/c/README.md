下午两点在B站讲这场双周赛的题目，[欢迎关注](https://space.bilibili.com/206214)~

---

经典技巧：由于每个比特位互不相干，所以拆分成每个比特位分别计算。

由于只有 $0$ 和 $1$，这样就好算了。

对异或有影响的是 $1$，所以只需要统计 $(a|b)\&c=1$ 的情况。

那么 $c$ 必须是 $1$，$a$ 和 $b$ 不能都是 $0$。

设有 $y$ 个 $1$，那么就有 $x=n-y$ 个 $0$。

那么这个比特一共可以产生

$$
\textit{one} = (n^2-x^2)y = (n^2-(n-y)^2)y = (2ny-y^2)y
$$

个 $1$。

由于异或只在乎 $\textit{one}$ 的奇偶性，所以 $2ny$ 可以去掉，那么就变成看 $y^3$ 的奇偶性，也就是 $y$ 的奇偶性。

如果 $y$ 是奇数，那么这个比特位的异或值就是 $1$。

这实际上就是看每个比特位的异或值是否为 $1$。

那么把 $\textit{nums}$ 的每个数异或起来，就是答案。

```py [sol1-Python3]
class Solution:
    def xorBeauty(self, nums: List[int]) -> int:
        return reduce(xor, nums)
```

```go [sol1-Go]
func xorBeauty(nums []int) (ans int) {
	for _, x := range nums {
		ans ^= x
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$，仅用到若干额外变量。

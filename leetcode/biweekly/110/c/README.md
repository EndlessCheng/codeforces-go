下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

---

### 提示 1

最终所有元素一定变成了一个在 $\textit{nums}$ 中的数。

枚举这个数。

### 提示 2

考虑把数字 $x$「扩散」到其它位置，那么每一秒 $x$ 都可以向左右扩散一位。

多个相同数字 $x$ 同时扩散，那么扩散完整个数组的耗时，就取决于相距最远的两个相邻的 $x$。

假设这两个 $x$ 的下标分别为 $i$ 和 $j$，且 $i<j$，那么耗时为：

$$
\left\lfloor\dfrac{j-i}{2}\right\rfloor
$$

取所有耗时的最小值，即为答案。

### 提示 3

统计所有相同数字的下标，记到一个哈希表 $\textit{pos}$ 中。

本题数组可以视作是环形的，假设最左边的 $x$ 的下标是 $i$，只需要在 $\textit{pos}[x]$ 列表末尾添加一个 $i+n$，就可以转换成非环形数组处理了。

```py [sol-Python3]
class Solution:
    def minimumSeconds(self, nums: List[int]) -> int:
        pos = defaultdict(list)
        for i, x in enumerate(nums):
            pos[x].append(i)

        ans = n = len(nums)
        for a in pos.values():
            a.append(a[0] + n)
            mx = max((j - i) // 2 for i, j in pairwise(a))
            ans = min(ans, mx)
        return ans
```

```go [sol-Go]
func minimumSeconds(nums []int) int {
	pos := map[int][]int{}
	for i, x := range nums {
		pos[x] = append(pos[x], i)
	}

	n := len(nums)
	ans := n
	for _, a := range pos {
		a = append(a, a[0]+n)
		mx := 0
		for i := 1; i < len(a); i++ {
			mx = max(mx, (a[i]-a[i-1])/2)
		}
		ans = min(ans, mx)
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

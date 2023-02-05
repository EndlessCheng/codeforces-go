下午两点【bilibili@灵茶山艾府】直播讲题，记得关注哦~

---

首先，把两个数组中都有的数去掉，那么每个剩余数字的出现次数必须为偶数。这可以用哈希表来统计。

设处理后的剩余数组分别 $a$ 和 $b$。

贪心地想，如果要交换 $a$ 中最小的数，那么找一个 $b$ 中最大的数是最合适的；对于 $b$ 中最小的数也同理。

那么把 $a$ 从小到大排序，$b$ 从大到小排序，两两匹配。

但是，还有一种方案。

把 $\textit{basket}_1$ 和 $\textit{basket}_2$ 中的最小值 $\textit{mn}$ 当作「工具人」，对于 $a[i]$ 和 $b[i]$ 的交换，可以分别和 $\textit{mn}$ 交换一次，就相当于 $a[i]$ 和 $b[i]$ 交换了。

因此每次交换的代价为

$$
\min(a[i], b[i], 2\cdot\textit{mn})
$$

累加代价，即为答案。

```py [sol1-Python3]
class Solution:
    def minCost(self, basket1: List[int], basket2: List[int]) -> int:
        cnt = Counter()
        for x, y in zip(basket1, basket2):
            cnt[x] += 1
            cnt[y] -= 1
        mn = min(cnt)
        a, b = [], []
        for x, c in cnt.items():
            if c % 2: return -1
            if c > 0: a.extend([x] * (c // 2))
            else:     b.extend([x] * (-c // 2))
        a.sort()
        b.sort(reverse=True)
        return sum(min(x, y, mn * 2) for x, y in zip(a, b))
```

```go [sol1-Go]
func minCost(basket1, basket2 []int) (ans int64) {
	cnt := map[int]int{}
	for i, x := range basket1 {
		cnt[x]++
		cnt[basket2[i]]--
	}

	mn := math.MaxInt
	var a, b []int
	for x, c := range cnt {
		if c%2 != 0  {
			return -1
		}
		mn = min(mn, x)
		if c > 0 {
			for i := 0; i < c/2; i++ {
				a = append(a, x)
			}
		} else {
			for i := 0; i < -c/2; i++ {
				b = append(b, x)
			}
		}
	}
	sort.Ints(a)
	sort.Sort(sort.Reverse(sort.IntSlice(b)))

	for i, x := range a {
		ans += int64(min(min(x, b[i]), mn*2))
	}
	return
}

func min(a, b int) int { if a > b { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{basket}_1$ 的长度。
- 空间复杂度：$O(n)$。

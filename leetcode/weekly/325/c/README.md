欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)，下午两点在B站讲这场周赛的题目。

---

由于随着甜蜜度的增大，能选择的糖果数量变小，有单调性，所以可以用二分答案来做。

对 $\textit{price}$ 从小到大排序，二分答案 $d$。最小的数 $x$ 可以选，下一个可以选的数是第一个 $\ge x+d$ 的数，依此类推。

如果可以选的数 $< k$，说明 $d$ 取大了，否则说明 $d$ 取小了，根据这一点来二分。

二分上界可以取 $\max(\textit{price})$。

有关二分的写法，可以看我的 [【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/) 这期视频。

```py [sol1-Python3]
class Solution:
    def maximumTastiness(self, price: List[int], k: int) -> int:
        price.sort()
        def check(d: int) -> bool:
            cnt, x0 = 1, price[0]
            for x in price:
                if x >= x0 + d:
                    cnt += 1
                    x0 = x
            return cnt >= k

        left, right = 0, price[-1]
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid): left = mid
            else: right = mid
        return left
```

```go [sol1-Go]
func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	return sort.Search(price[len(price)-1], func(d int) bool {
		d++
		cnt, x0 := 1, price[0]
		for _, x := range price[1:] {
			if x >= x0+d {
				cnt++
				x0 = x
			}
		}
		return cnt < k
	})
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n + n\log U)$，其中 $n$ 为 $\textit{price}$ 的长度，$U=\max(\textit{price})$。
- 空间复杂度：$O(1)$，忽略排序的空间，仅用到若干额外变量。

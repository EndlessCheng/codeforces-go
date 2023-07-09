下午两点[【b站@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，欢迎关注！

注意每次操作，将 $\textit{moveFrom}[i]$ 的**所有石块**都移走，所以无需维护 $\textit{nums}[i]$ 的个数。

```py [sol-Python3]
class Solution:
    def relocateMarbles(self, nums: List[int], moveFrom: List[int], moveTo: List[int]) -> List[int]:
        cnt = set(nums)
        for f, t in zip(moveFrom, moveTo):
            cnt.remove(f)
            cnt.add(t)
        return sorted(cnt)
```

```go [sol-Go]
func relocateMarbles(nums, moveFrom, moveTo []int) []int {
	set := map[int]struct{}{}
	for _, x := range nums {
		set[x] = struct{}{}
	}
	for i, x := range moveFrom {
		delete(set, x)
		set[moveTo[i]] = struct{}{}
	}
	ans := make([]int, 0, len(set))
	for x := range set {
		ans = append(ans, x)
	}
	sort.Ints(ans)
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m + n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{moveFrom}$ 的长度。排序是瓶颈。
- 空间复杂度：$\mathcal{O}(n)$。

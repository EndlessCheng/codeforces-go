雇佣过程可以用两个最小堆来模拟。

如果两个堆要取的元素重叠了，则可以合并这两个堆，然后取前 $k'$ 小的元素（$k'$ 表示此时剩余待雇佣的人数）。

代码实现时，如果忽略切片的空间开销，则可以做到 $O(1)$ 空间复杂度（Python 切片是有拷贝的，而 Go 没有）。

```py [sol1-Python3]
class Solution:
    def totalCost(self, costs: List[int], k: int, candidates: int) -> int:
        ans, n = 0, len(costs)
        if candidates * 2 < n:
            pre = costs[:candidates]
            heapify(pre)
            suf = costs[-candidates:]
            heapify(suf)
            i, j = candidates, n - 1 - candidates
            while k and i <= j:
                if pre[0] <= suf[0]:
                    ans += heapreplace(pre, costs[i])
                    i += 1
                else:
                    ans += heapreplace(suf, costs[j])
                    j -= 1
                k -= 1
            costs = pre + suf
        costs.sort()
        return ans + sum(costs[:k])  # 也可以用快速选择算法求前 k 小
```

```go [sol1-Go]
func totalCost(costs []int, k, candidates int) int64 {
	ans := 0
	if n := len(costs); candidates*2 < n {
		pre := hp{costs[:candidates]}
		heap.Init(&pre) // 原地建堆
		suf := hp{costs[n-candidates:]}
		heap.Init(&suf)
		for i, j := candidates, n-1-candidates; k > 0 && i <= j; k-- {
			if pre.IntSlice[0] <= suf.IntSlice[0] {
				ans += pre.IntSlice[0]
				pre.IntSlice[0] = costs[i]
				heap.Fix(&pre, 0)
				i++
			} else {
				ans += suf.IntSlice[0]
				suf.IntSlice[0] = costs[j]
				heap.Fix(&suf, 0)
				j--
			}
		}
		costs = append(pre.IntSlice, suf.IntSlice...)
	}
	sort.Ints(costs)
	for _, c := range costs[:k] { // 也可以用快速选择算法求前 k 小
		ans += c
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }
func (hp) Push(interface{})     {} // 没有用到，留空即可
func (hp) Pop() (_ interface{}) { return }
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{costs}$ 的长度。
- 空间复杂度：$O(n)$ 或 $O(1)$。

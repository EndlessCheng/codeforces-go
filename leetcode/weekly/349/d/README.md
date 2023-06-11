### 视频讲解

见个人主页。

### 思路

为方便处理，可以先把 $\textit{nums}_1$ 和询问中的 $x_i$ 排序。

这样就可以把重心放在 $\textit{nums}_2[j]$ 与 $y_i$ 的大小关系上。

我们可以按照 $x_i$ 从大到小、$\textit{nums}_1[j]$ 从大到小的顺序处理，同时**增量地**维护满足 $\textit{nums}_1[j]\ge x_i$ 的 $\textit{nums}_2[j]$。

如何维护？分类讨论：

- 如果 $\textit{nums}_2[j]$ 比之前遍历过的 $\textit{nums}_2[j']$ 还要小，那么由于 $\textit{nums}_1[j]$ 是从大到小处理的，所以 $\textit{nums}_1[j]+\textit{nums}_2[j]$ 也比之前遍历过的 $\textit{nums}_1[j']+\textit{nums}_2[j']$ 要小。那么在回答 $\le \textit{nums}_2[j]$ 的 $y_i$ 时，最大值不可能是 $\textit{nums}_1[j]+\textit{nums}_2[j]$，所以无需考虑这样的 $\textit{nums}_2[j]$。（这种单调性启发我们用**单调栈**来维护。）
- 如果是相等，那么同理，无需考虑。
- 如果是大于，那么就可以入栈。在入栈前还要去掉一些无效数据：
    - 如果 $\textit{nums}_1[j]+\textit{nums}_2[j]$ 不低于栈顶的 $\textit{nums}_1[j']+\textit{nums}_2[j']$，那么可以弹出栈顶。因为更大的 $\textit{nums}_2[j]$ 更能满足 $\ge y_i$ 的要求，栈顶的 $\textit{nums}_1[j']+\textit{nums}_2[j']$ 在后续的询问中，永远不会是最大值。
    - 代码实现时，可以直接比较 $\textit{nums}_1[j]+\textit{nums}_2[j]$ 与栈顶的值，这是因为如果这一条件成立，由于 $\textit{nums}_1[j]$ 是从大到小处理的，$\textit{nums}_1[j]+\textit{nums}_2[j]$ 能比栈顶的大，说明 $\textit{nums}_2[j]$ 必然不低于栈顶的 $\textit{nums}_2[j']$。

这样我们会得到一个从栈底到栈顶，$\textit{nums}_2[j]$ 递增，$\textit{nums}_1[j]+\textit{nums}_2[j]$ 递减的单调栈。

最后在单调栈中二分 $\ge y_i$ 的最小的 $\textit{nums}_2[j]$，对应的 $\textit{nums}_1[j]+\textit{nums}_2[j]$ 就是最大的。

```py [sol-Python3]
class Solution:
    def maximumSumQueries(self, nums1: List[int], nums2: List[int], queries: List[List[int]]) -> List[int]:
        ans = [-1] * len(queries)
        st = []
        a = sorted((a, b) for a, b in zip(nums1, nums2))
        i = len(a) - 1
        for qid, (x, y) in sorted(enumerate(queries), key=lambda p: -p[1][0]):
            while i >= 0 and a[i][0] >= x:
                ax, ay = a[i]
                while st and st[-1][1] <= ax + ay:  # ay >= st[-1][0]
                    st.pop()
                if not st or st[-1][0] < ay:
                    st.append((ay, ax + ay))
                i -= 1
            j = bisect_left(st, (y,))
            if j < len(st):
                ans[qid] = st[j][1]
        return ans
```

```go [sol-Go]
func maximumSumQueries(nums1, nums2 []int, queries [][]int) (ans []int) {
	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(queries, func(i, j int) bool { return queries[i][0] > queries[j][0] })

	ans = make([]int, len(queries))
	st := []pair{}
	i := len(a) - 1
	for _, q := range queries {
		for i >= 0 && a[i].x >= q[0] {
			for len(st) > 0 && st[len(st)-1].y <= a[i].x+a[i].y {
				st = st[:len(st)-1] // a[i].y >= st[len(st)-1].x
			}
			if len(st) == 0 || st[len(st)-1].x < a[i].y {
				st = append(st, pair{a[i].y, a[i].x + a[i].y})
			}
			i--
		}
		j := sort.Search(len(st), func(i int) bool { return st[i].x >= q[1] })
		if j < len(st) {
			ans[q[2]] = st[j].y
		} else {
			ans[q[2]] = -1
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计。

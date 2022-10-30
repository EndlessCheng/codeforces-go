遍历，哈希表统计每个人的：

- 播放量的总和 $\textit{viewSum}$；
- 最大播放量 $\textit{maxView}$；
- 最大播放量对应的视频 $\textit{id}$。

最后遍历哈希表，得到答案。

```py [sol1-Python3]
class Solution:
    def mostPopularCreator(self, creators: List[str], ids: List[str], views: List[int]) -> List[List[str]]:
        m, max_view_sum = {}, 0
        for name, id, view in zip(creators, ids, views):
            if name in m:
                t = m[name]
                t[0] += view
                if view > t[1] or view == t[1] and id < t[2]:
                    t[1], t[2] = view, id
            else: m[name] = [view, view, id]
            max_view_sum = max(max_view_sum, m[name][0])
        return [[name, id] for name, (view_sum, _, id) in m.items() if view_sum == max_view_sum]
```

```go [sol1-Go]
func mostPopularCreator(creators, ids []string, views []int) (ans [][]string) {
	type tuple struct {
		viewSum, maxView int
		id               string
	}
	m, maxViewSum := map[string]tuple{}, 0
	for i, name := range creators {
		id, view := ids[i], views[i]
		t := m[name]
		if t.id == "" {
			t = tuple{view, view, id}
		} else {
			t.viewSum += view
			if view > t.maxView || view == t.maxView && id < t.id {
				t.maxView = view
				t.id = id
			}
		}
		maxViewSum = max(maxViewSum, t.viewSum)
		m[name] = t
	}
	for name, t := range m {
		if t.viewSum == maxViewSum {
			ans = append(ans, []string{name, t.id})
		}
	}
	return
}

func max(a, b int) int { if a < b { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{creators}$ 的长度。字符串的长度视作 $O(1)$。
- 空间复杂度：$O(n)$。

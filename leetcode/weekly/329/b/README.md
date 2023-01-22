下午两点【bilibili@灵茶山艾府】直播讲题，记得关注哦~

---

语法题。

```py [sol1-Python3]
class Solution:
    def sortTheStudents(self, score: List[List[int]], k: int) -> List[List[int]]:
        score.sort(key=lambda s: -s[k])
        return score
```

```go [sol1-Go]
func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool { return score[i][k] > score[j][k] })
	return score
}
```

### 复杂度分析

- 时间复杂度：$O(m\log m)$，其中 $m$ 为 $\textit{score}$ 的长度。
- 空间复杂度：$O(1)$，忽略排序时的栈开销，仅用到若干额外变量。

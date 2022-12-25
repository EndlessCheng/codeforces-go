把 $\textit{feedback}$ 及其分数存到哈希表 $\textit{score}$ 中，对每个 $\textit{report}_i$，按照空格分割，然后用 $\textit{score}$ 计算分数之和。

最后按照题目规则排序，取前 $k$ 个 $\textit{studentId}$ 为答案。

```py [sol1-Python3]
class Solution:
    def topStudents(self, positive_feedback: List[str], negative_feedback: List[str], report: List[str], student_id: List[int], k: int) -> List[int]:
        score = defaultdict(int)
        for w in positive_feedback: score[w] = 3
        for w in negative_feedback: score[w] = -1
        a = sorted((-sum(score[w] for w in r.split()), i) for r, i in zip(report, student_id))
        return [i for _, i in a[:k]]
```

```go [sol1-Go]
func topStudents(positiveFeedback, negativeFeedback, report []string, studentId []int, k int) []int {
	score := map[string]int{}
	for _, w := range positiveFeedback {
		score[w] = 3
	}
	for _, w := range negativeFeedback {
		score[w] = -1
	}
	type pair struct{ score, id int }
	a := make([]pair, len(report))
	for i, r := range report {
		s := 0
		for _, w := range strings.Split(r, " ") {
			s += score[w]
		}
		a[i] = pair{s, studentId[i]}
	}
	sort.Slice(a, func(i, j int) bool {
		a, b := a[i], a[j]
		return a.score > b.score || a.score == b.score && a.id < b.id
	})
	ans := make([]int, k)
	for i, p := range a[:k] {
		ans[i] = p.id
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O((p+q+n)L+n\log n)$，其中 $p$ 为 $\textit{positiveFeedback}$ 的长度，$q$ 为 $\textit{negativeFeedback}$ 的长度，$n$ 为 $\textit{report}$ 的长度，$L$ 为字符串的平均长度。
- 空间复杂度：$O((p+q)L+n)$。

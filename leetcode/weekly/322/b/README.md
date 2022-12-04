如果最小的不和最大的匹配，那么最大的和一个比最小数更大的数匹配，就会导致技能点之和不相等。

因此最小的一定和最大的匹配。

那么排序后模拟即可。

```py [sol1-Python3]
class Solution:
    def dividePlayers(self, skill: List[int]) -> int:
        skill.sort()
        ans, s = 0, skill[0] + skill[-1]
        for i in range(len(skill) // 2):
            x, y = skill[i], skill[-1 - i]
            if x + y != s: return -1
            ans += x * y
        return ans
```

```go [sol1-Go]
func dividePlayers(skill []int) (ans int64) {
	sort.Ints(skill)
	n := len(skill)
	sum := skill[0] + skill[n-1]
	for i := 0; i < n/2; i++ {
		x, y := skill[i], skill[n-1-i]
		if x+y != sum {
			return -1
		}
		ans += int64(x * y)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{skill}$ 的长度。
- 空间复杂度：$O(1)$，忽略排序的栈空间，只用到若干额外变量。

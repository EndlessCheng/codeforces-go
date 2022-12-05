[视频讲解](https://www.bilibili.com/video/BV15d4y147YF) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

# 方法一：排序

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

# 方法二：哈希表

设 $\textit{total}$ 为 $\textit{skill}$ 所有数之和，$m$ 为 $\textit{skill}$ 长度的一半。

那么 $\textit{total}$ 必须为 $m$ 的倍数，否则返回 $-1$。

统计 $\textit{skill}$ 每个数的出现次数，记在哈希表 $\textit{cnt}$ 中。

设 $s=\dfrac{\textit{total}}{m}$，遍历哈希表，如果 $\textit{cnt}[x]\ne cnt[s-x]$，则无法匹配，返回 $-1$。

否则把 $\textit{cnt}[x]\cdot x\cdot(s-x)$ 记到答案中。

最后返回答案的一半（因为重复记录了对称的部分）。

```py [sol2-Python3]
class Solution:
    def dividePlayers(self, skill: List[int]) -> int:
        total, m = sum(skill), len(skill) // 2
        if total % m: return -1
        ans, s = 0, total // m
        cnt = Counter(skill)
        for x, c in cnt.items():
            if c != cnt[s - x]: return -1
            ans += c * x * (s - x)
        return ans // 2
```

```go [sol2-Go]
func dividePlayers(skill []int) (ans int64) {
	total := 0
	cnt := map[int]int{}
	for _, x := range skill {
		total += x
		cnt[x]++
	}
	m := len(skill) / 2
	if total%m > 0 {
		return -1
	}
	s := total / m
	for x, c := range cnt {
		if c != cnt[s-x] {
			return -1
		}
		ans += int64(c * x * (s - x))
	}
	return ans / 2
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{skill}$ 的长度。
- 空间复杂度：$O(n)$。

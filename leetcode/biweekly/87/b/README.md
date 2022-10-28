[视频讲解](https://www.bilibili.com/video/BV1MT411u7fW) 已出炉，欢迎点赞三连，在评论区分享你对这场双周赛的看法~

---

考虑能力值最小的运动员 $\textit{player}[i]$，他应当匹配训练能力值大于等于 $\textit{player}[i]$ 且最接近 $\textit{player}[i]$ 的训练师（如果选了一个训练能力值更大的，可能会导致能力值更大的运动员无法匹配）。

那么对 $\textit{players}$ 和 $\textit{trainers}$ 从小到大排序，然后双指针模拟即可。

```py [sol1-Python3]
class Solution:
    def matchPlayersAndTrainers(self, players: List[int], trainers: List[int]) -> int:
        players.sort()
        trainers.sort()
        j, m = 0, len(trainers)
        for i, p in enumerate(players):
            while j < m and trainers[j] < p:
                j += 1
            if j == m:  # 无法找到匹配的训练师
                return i
            j += 1  # 匹配一位训练师
        return len(players)  # 所有运动员都能匹配
```

```go [sol1-Go]
func matchPlayersAndTrainers(players, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	j, m := 0, len(trainers)
	for i, p := range players {
		for j < m && trainers[j] < p {
			j++
		}
		if j == m { // 无法找到匹配的训练师
			return i
		}
		j++ // 匹配一位训练师
	}
	return len(players) // 所有运动员都能匹配
}
```

也可以遍历 $\textit{trainers}$ 去找对应的 $\textit{players}$。

```py [sol1-Python3]
class Solution:
    def matchPlayersAndTrainers(self, players: List[int], trainers: List[int]) -> int:
        players.sort()
        trainers.sort()
        j, m = 0, len(players)
        for t in trainers:
            if j < m and players[j] <= t:
                j += 1
        return j
```

```go [sol1-Go]
func matchPlayersAndTrainers(players, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	j, m := 0, len(players)
	for _, t := range trainers {
		if j < m && players[j] <= t {
			j++
		}
	}
	return j
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n + m\log m)$，其中 $n$ 为 $\textit{players}$ 的长度，$m$ 为 $\textit{trainers}$ 的长度。
- 空间复杂度：$O(1)$，忽略排序时的栈开销，仅用到若干变量。

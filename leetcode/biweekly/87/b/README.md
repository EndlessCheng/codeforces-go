下午 2 点在 B 站直播讲周赛和双周赛的题目，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

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
func matchPlayersAndTrainers(players []int, trainers []int) int {
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

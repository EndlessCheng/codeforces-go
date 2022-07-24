用哈希表统计每行出现的次数，然后遍历列，累加哈希表中列出现的次数。

```py [sol1-Python3]
class Solution:
    def equalPairs(self, grid: List[List[int]]) -> int:
        cnt = Counter(tuple(row) for row in grid)
        return sum(cnt[col] for col in zip(*grid))
```

```go [sol1-Go]
func equalPairs(grid [][]int) (ans int) {
	cnt := map[[200]int]int{}
	for _, row := range grid {
		a := [200]int{}
		for j, v := range row {
			a[j] = v
		}
		cnt[a]++
	}
	for j := range grid[0] {
		a := [200]int{}
		for i, row := range grid {
			a[i] = row[j]
		}
		ans += cnt[a]
	}
	return
}
```

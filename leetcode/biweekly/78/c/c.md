将 $\textit{tiles}$ 按左端点 $l_i$ 排序后，我们可以枚举毯子放置位置，然后计算能覆盖多少块瓷砖。

实际上，毯子右端点放在瓷砖中间，不如直接放在瓷砖右端点（因为从中间向右移动，能覆盖的瓷砖数不会减少），所以可以枚举每个瓷砖的右端点来摆放毯子。

这样就可以双指针了，左指针 $\textit{left}$ 需要满足其指向的瓷砖的右端点被毯子覆盖。设毯子右端点在瓷砖 $i$ 上，则有

$$
\textit{tiles}[\textit{left}][0] \ge \textit{tiles}[i][1] - \textit{carpetLen}+1
$$

如果毯子左端点在瓷砖 $\textit{tiles}[\textit{left}]$ 内部，则需要减去这块瓷砖没被覆盖的部分，即减去

$$
\textit{tiles}[i][1] - \textit{carpetLen}+1-\textit{tiles}[\textit{left}][0]
$$

```Python [sol1-Python3]
class Solution:
    def maximumWhiteTiles(self, tiles: List[List[int]], carpetLen: int) -> int:
        tiles.sort(key=lambda x: x[0])
        ans = cover = left = 0
        for tl, tr in tiles:
            cover += tr - tl + 1
            while tiles[left][1] < tr - carpetLen + 1:
                cover -= tiles[left][1] - tiles[left][0] + 1
                left += 1
            ans = max(ans, cover - max(tr - carpetLen + 1 - tiles[left][0], 0))  # 0 表示毯子左端点不在瓷砖内的情况
        return ans
```

```go [sol1-Go]
func maximumWhiteTiles(tiles [][]int, carpetLen int) (ans int) {
	sort.Slice(tiles, func(i, j int) bool { return tiles[i][0] < tiles[j][0] })
	cover, left := 0, 0
	for _, t := range tiles {
		tl, tr := t[0], t[1]
		cover += tr - tl + 1
		for tiles[left][1]+carpetLen-1 < tr {
			cover -= tiles[left][1] - tiles[left][0] + 1
			left++
		}
		ans = max(ans, cover-max(tr-carpetLen+1-tiles[left][0], 0)) // 0 表示毯子左端点不在瓷砖内的情况
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```


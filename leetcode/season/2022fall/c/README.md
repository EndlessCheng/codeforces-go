个人赛五道题目的 [视频讲解](https://www.bilibili.com/video/BV1zN4y1K762) 已出炉，欢迎点赞三连，在评论区分享你对这场比赛的看法~

---

由于路径是唯一的，一个入口只会对应一个唯一的出口；一个出口+进入出口的方向，可以找到唯一的入口。

因此，从不同入口出发的弹珠，走过的路径是不会有**完全重叠**的部分的（方向相反不算重叠），且路径不存在环（可以画图理解）。

枚举所有入口，模拟即可。

```py [sol1-Python3]
DIRS = ((0, 1), (1, 0), (0, -1), (-1, 0))  # 右下左上（顺时针）

class Solution:
    def ballGame(self, num: int, plate: List[str]) -> List[List[int]]:
        m, n = len(plate), len(plate[0])

        def check(x, y, d):
            left = num
            while plate[x][y] != 'O':
                if left == 0: return False  # 无剩余步数
                if plate[x][y] == 'W':   d = (d + 3) % 4  # 逆时针
                elif plate[x][y] == 'E': d = (d + 1) % 4  # 顺时针
                x += DIRS[d][0]
                y += DIRS[d][1]
                if not (0 <= x < m and 0 <= y < n): return False  # 出界
                left -= 1
            return True

        ans = []
        for j in range(1, n - 1):
            if plate[0][j] == '.' and check(0, j, 1): ans.append([0, j])  # 上边
            if plate[-1][j] == '.' and check(m - 1, j, 3): ans.append([m - 1, j])  # 下边
        for i in range(1, m - 1):
            if plate[i][0] == '.' and check(i, 0, 0): ans.append([i, 0])  # 左边
            if plate[i][-1] == '.' and check(i, n - 1, 2): ans.append([i, n - 1])  # 右边
        return ans
```

```go [sol1-Go]
var dirs = []struct{ x, y int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上（顺时针）

func ballGame(num int, plate []string) (ans [][]int) {
	m, n := len(plate), len(plate[0])
	check := func(x, y, d int) bool {
		for left := num; plate[x][y] != 'O'; left-- {
			if left == 0 { // 无剩余步数
				return false
			}
			if plate[x][y] == 'W' { // 逆时针
				d = (d + 3) % 4
			} else if plate[x][y] == 'E' { // 顺时针
				d = (d + 1) % 4
			}
			x += dirs[d].x
			y += dirs[d].y
			if x < 0 || x >= m || y < 0 || y >= n { // 从另一边出去了
				return false
			}
		}
		return true
	}
	for j := 1; j < n-1; j++ {
		if plate[0][j] == '.' && check(0, j, 1) {
			ans = append(ans, []int{0, j})
		}
		if plate[m-1][j] == '.' && check(m-1, j, 3) {
			ans = append(ans, []int{m - 1, j})
		}
	}
	for i := 1; i < m-1; i++ {
		if plate[i][0] == '.' && check(i, 0, 0) {
			ans = append(ans, []int{i, 0})
		}
		if plate[i][n-1] == '.' && check(i, n-1, 2) {
			ans = append(ans, []int{i, n - 1})
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{plate}$ 的行数和列数。每个状态（横坐标，纵坐标，方向）至多被访问一次，总共有 $4mn$ 个状态。
- 空间复杂度：$O(1)$。返回值不计入。

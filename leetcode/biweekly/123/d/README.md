[视频讲解](https://www.bilibili.com/video/BV14C411r7nN/) 第四题。

将 $\textit{points}$ 按照横坐标**从小到大**排序，横坐标相同的，按照纵坐标**从大到小**排序。

如此一来，在枚举 $\textit{points}[i]$ 和 $\textit{points}[j]$ 时（$i<j$），就只需要关心纵坐标的大小。

固定 $\textit{points}[i]$，然后枚举 $\textit{points}[j]$：

- 如果 $\textit{points}[j][1]$ 比之前枚举的点的纵坐标都大，那么矩形内没有其它点，符合要求，答案加一。
- 如果 $\textit{points}[j][1]$ 小于等于之前枚举的某个点的纵坐标，那么矩形内有其它点，不符合要求。

所以在枚举 $\textit{points}[j]$ 的同时，需要维护纵坐标的最大值 $\textit{maxY}$。这也解释了为什么横坐标相同的，按照纵坐标**从大到小**排序。这保证了横坐标相同时，我们总是优先枚举更靠上的点，不会误把包含其它点的矩形也当作符合要求的矩形。

```py [sol-Python3]
class Solution:
    def numberOfPairs(self, points: List[List[int]]) -> int:
        points.sort(key=lambda p: (p[0], -p[1]))
        ans = 0
        for i, (_, y0) in enumerate(points):
            max_y = -inf
            for (_, y) in points[i + 1:]:
                if max_y < y <= y0:
                    max_y = y
                    ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int numberOfPairs(int[][] points) {
        Arrays.sort(points, (p, q) -> p[0] != q[0] ? p[0] - q[0] : q[1] - p[1]);
        int ans = 0;
        for (int i = 0; i < points.length; i++) {
            int y0 = points[i][1];
            int maxY = Integer.MIN_VALUE;
            for (int j = i + 1; j < points.length; j++) {
                int y = points[j][1];
                if (y <= y0 && y > maxY) {
                    maxY = y;
                    ans++;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfPairs(vector<vector<int>> &points) {
        ranges::sort(points, [](const auto &p, const auto &q) {
            return p[0] != q[0] ? p[0] < q[0] : p[1] > q[1];
        });
        int ans = 0, n = points.size();
        for (int i = 0; i < n; i++) {
            int y0 = points[i][1];
            int max_y = INT_MIN;
            for (int j = i + 1; j < n; j++) {
                int y = points[j][1];
                if (y <= y0 && y > max_y) {
                    max_y = y;
                    ans++;
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfPairs(points [][]int) (ans int) {
	slices.SortFunc(points, func(p, q []int) int {
		if p[0] != q[0] {
			return p[0] - q[0]
		}
		return q[1] - p[1]
	})
	for i, p := range points {
		y0 := p[1]
		maxY := math.MinInt
		for _, q := range points[i+1:] {
			y := q[1]
			if y <= y0 && y > maxY {
				maxY = y
				ans++
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{points}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销和 Python 切片开销。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)

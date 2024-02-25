[视频讲解](https://www.bilibili.com/video/BV1qx421179t/)。

枚举两个矩形。

如果矩形有交集，那么交集一定是矩形。求出这个交集矩形的左下角和右上角。

- 左下角横坐标：两个矩形左下角横坐标的最大值。
- 左下角纵坐标：两个矩形左下角纵坐标的最大值。
- 右上角横坐标：两个矩形右上角横坐标的最小值。
- 右上角纵坐标：两个矩形右上角纵坐标的最小值。

知道坐标就可以算出矩形的长和宽，取二者最小值作为正方形的边长。

如果矩形没有交集，那么长和宽是负数，在计算面积前判断。

```py [sol-Python3]
class Solution:
    def largestSquareArea(self, bottomLeft: List[List[int]], topRight: List[List[int]]) -> int:
        ans = 0
        for ((x1, y1), (x2, y2)), ((x3, y3), (x4, y4)) in combinations(zip(bottomLeft, topRight), 2):
            width = min(x2, x4) - max(x1, x3)  # 注：改成用 if-else 计算 min 和 max 会更快
            height = min(y2, y4) - max(y1, y3)
            size = min(width, height)
            if size > 0:
                ans = max(ans, size * size)
        return ans
```

```java [sol-Java]
class Solution {
    public long largestSquareArea(int[][] bottomLeft, int[][] topRight) {
        long ans = 0;
        for (int i = 0; i < bottomLeft.length; i++) {
            int[] b1 = bottomLeft[i];
            int[] t1 = topRight[i];
            for (int j = i + 1; j < bottomLeft.length; j++) {
                int[] b2 = bottomLeft[j];
                int[] t2 = topRight[j];
                int height = Math.min(t1[1], t2[1]) - Math.max(b1[1], b2[1]);
                int width = Math.min(t1[0], t2[0]) - Math.max(b1[0], b2[0]);
                int size = Math.min(width, height);
                if (size > 0) {
                    ans = Math.max(ans, (long) size * size);
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
    long long largestSquareArea(vector<vector<int>> &bottomLeft, vector<vector<int>> &topRight) {
        long long ans = 0;
        for (int i = 0; i < bottomLeft.size(); i++) {
            auto &b1 = bottomLeft[i];
            auto &t1 = topRight[i];
            for (int j = i + 1; j < bottomLeft.size(); j++) {
                auto &b2 = bottomLeft[j];
                auto &t2 = topRight[j];
                int height = min(t1[1], t2[1]) - max(b1[1], b2[1]);
                int width = min(t1[0], t2[0]) - max(b1[0], b2[0]);
                int size = min(width, height);
                if (size > 0) {
                    ans = max(ans, (long long) size * size);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func largestSquareArea(bottomLeft, topRight [][]int) (ans int64) {
	for i, b1 := range bottomLeft {
		t1 := topRight[i]
		for j := i + 1; j < len(bottomLeft); j++ {
			b2, t2 := bottomLeft[j], topRight[j]
			height := min(t1[1], t2[1]) - max(b1[1], b2[1])
			width := min(t1[0], t2[0]) - max(b1[0], b2[0])
			size := min(width, height)
			if size > 0 {
				ans = max(ans, int64(size)*int64(size))
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{bottomLeft}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)

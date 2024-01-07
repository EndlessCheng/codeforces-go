[本题视频讲解](https://www.bilibili.com/video/BV1ae411e7fn/)

设长宽分别为 $x$ 和 $y$。

根据勾股定理，对角线长度的平方为

$$
x^2+y^2
$$

本题是双关键字比较，第一关键字是对角线长度，我们直接用其平方值。如果遍历到更大的长度，则覆盖矩形面积。

第二关键字是矩形面积，即 $xy$。如果遇到了和最长长度一样长的矩形，那么更新面积的最大值。

```py [sol-Python3]
class Solution:
    def areaOfMaxDiagonal(self, dimensions: List[List[int]]) -> int:
        return max((x * x + y * y, x * y) for x, y in dimensions)[1]
```

```java [sol-Java]
class Solution {
    public int areaOfMaxDiagonal(int[][] dimensions) {
        int ans = 0, maxL = 0;
        for (int[] d : dimensions) {
            int x = d[0], y = d[1];
            int l = x * x + y * y;
            if (l > maxL || (l == maxL && x * y > ans)) {
                maxL = l;
                ans = x * y;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int areaOfMaxDiagonal(vector<vector<int>> &dimensions) {
        int ans = 0, max_l = 0;
        for (auto &d: dimensions) {
            int x = d[0], y = d[1];
            int l = x * x + y * y;
            if (l > max_l || (l == max_l && x * y > ans)) {
                max_l = l;
                ans = x * y;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func areaOfMaxDiagonal(dimensions [][]int) (ans int) {
	maxL := 0
	for _, d := range dimensions {
		x, y := d[0], d[1]
		l := x*x + y*y
		if l > maxL || l == maxL && x*y > ans {
			maxL = l
			ans = x * y
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{dimensions}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

周赛总结更新啦！请看 [2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)

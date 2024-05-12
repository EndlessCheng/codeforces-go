由于正方形边长越大，越不合法，有单调性，所以可以二分边长。

在二分中统计点数，如果正方形合法，则更新答案的最大值。

关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)

代码用到了一些位运算技巧，原理请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

其它写法见 [视频讲解](https://www.bilibili.com/video/BV1cz421m786/) 第二题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def maxPointsInsideSquare(self, points: List[List[int]], s: str) -> int:
        ans = 0
        def check(size: int) -> bool:
            vis = set()
            for (x, y), c in zip(points, s):
                if abs(x) <= size and abs(y) <= size:
                    if c in vis:
                        return True
                    vis.add(c)
            nonlocal ans
            ans = len(vis)
            return False
        bisect_left(range(1_000_000_001), True, key=check)
        return ans
```

```java [sol-Java]
class Solution {
    private int ans;

    public int maxPointsInsideSquare(int[][] points, String S) {
        char[] s = S.toCharArray();
        int left = -1, right = 1_000_000_001;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (check(mid, points, s)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return ans;
    }

    boolean check(int size, int[][] points, char[] s) {
        int vis = 0;
        for (int i = 0; i < points.length; i++) {
            int x = points[i][0];
            int y = points[i][1];
            int c = s[i] - 'a';
            if (Math.abs(x) <= size && Math.abs(y) <= size) {
                if ((vis >> c & 1) > 0) {
                    return false;
                }
                vis |= 1 << c;
            }
        }
        ans = Integer.bitCount(vis);
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxPointsInsideSquare(vector<vector<int>>& points, string s) {
        int ans = 0;
        auto check = [&](int size) -> bool {
            int vis = 0;
            for (int i = 0; i < points.size(); i++) {
                int x = points[i][0];
                int y = points[i][1];
                char c = s[i] - 'a';
                if (abs(x) <= size && abs(y) <= size) {
                    if (vis >> c & 1) {
                        return false;
                    }
                    vis |= 1 << c;
                }
            }
            ans = __builtin_popcount(vis);
            return true;
        };
        int left = -1, right = 1'000'000'001;
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            (check(mid) ? left : right) = mid;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxPointsInsideSquare(points [][]int, s string) (ans int) {
	sort.Search(1_000_000_001, func(size int) bool {
		vis := 0
		for i, p := range points {
			if abs(p[0]) <= size && abs(p[1]) <= size {
				c := s[i] - 'a'
				if vis>>c&1 > 0 {
					return true
				}
				vis |= 1 << c
			}
		}
		ans = bits.OnesCount(uint(vis))
		return false
	})
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $s$ 的长度，$U=\max(|x_i|,|y_i|)$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$ 或 $\mathcal{O}(1)$。其中 $|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

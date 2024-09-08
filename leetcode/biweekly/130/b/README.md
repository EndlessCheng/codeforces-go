## 方法一：二分间接值

由于正方形边长越大，越不合法，有单调性，所以可以**二分边长的一半**。

在二分中统计遇到的字符，如果没有遇到重复的字符，说明正方形合法，用字符个数更新答案的最大值。

关于二分算法的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)

代码用到了一些位运算技巧，原理请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

#### 答疑

**问**：为什么可以直接更新 $\textit{ans}$，为什么不需要写 `ans = max(ans, ...)`？

**答**：更新 $\textit{ans}$ 时必然会伴随着二分区间左边界 $\textit{left}$ 的更新，那么下一次更新 $\textit{ans}$ 的时候，正方形的边长一定更长，包含的点数不会变少，所以可以直接更新，无需写 `max`。

```py [sol-Python3]
class Solution:
    def maxPointsInsideSquare(self, points: List[List[int]], s: str) -> int:
        ans = 0
        def check(size: int) -> bool:
            vis = set()
            for (x, y), c in zip(points, s):
                if abs(x) <= size and abs(y) <= size:  # 点在正方形中
                    if c in vis:
                        return True
                    vis.add(c)
            nonlocal ans
            ans = len(vis)
            return False
        # 注意 range 并不会创建 list，它是 O(1) 的
        bisect_left(range(1_000_000_001), True, key=check)
        return ans
```

```java [sol-Java]
class Solution {
    private int ans;

    public int maxPointsInsideSquare(int[][] points, String S) {
        char[] s = S.toCharArray();
        int left = -1;
        int right = 1_000_000_001;
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
            // 判断点是否在正方形中
            if (Math.abs(points[i][0]) <= size && Math.abs(points[i][1]) <= size) {
                int c = s[i] - 'a';
                if ((vis >> c & 1) > 0) { // c 在集合中
                    return false;
                }
                vis |= 1 << c; // 把 c 加入集合
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
                // 判断点是否在正方形中
                if (abs(points[i][0]) <= size && abs(points[i][1]) <= size) {
                    char c = s[i] - 'a';
                    if (vis >> c & 1) { // c 在集合中
                        return false;
                    }
                    vis |= 1 << c; // 把 c 加入集合
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
            if abs(p[0]) <= size && abs(p[1]) <= size { // 点在正方形中
                c := s[i] - 'a'
                if vis>>c&1 > 0 { // c 在集合中
                    return true
                }
                vis |= 1 << c // 把 c 加入集合
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

## 方法二：维护次小距离的最小值

定义点 $(x,y)$ 到 $(0,0)$ 的**切比雪夫距离**为

$$
\max(|x|,|y|)
$$

定义 $\textit{minD}[c]$ 为标签为 $c$ 的所有点到 $(0,0)$ 的**最小**切比雪夫距离。

定义 $\textit{minD}_2[c]$ 为标签为 $c$ 的所有点到 $(0,0)$ 的**次小**切比雪夫距离。

![lc3143.png](https://pic.leetcode.cn/1722695389-Arudlg-lc3143.png)

换句话说，可以在正方形中的点需要满足

$$
\textit{minD}[c] < \textit{min}_2
$$

代码实现时，无需维护 $\textit{minD}_2[c]$，而是直接维护 $\textit{min}_2$。

```py [sol-Python3]
class Solution:
    def maxPointsInsideSquare(self, points: List[List[int]], s: str) -> int:
        min_d = defaultdict(lambda: inf)
        min2 = inf
        for (x, y), c in zip(points, s):
            d = max(abs(x), abs(y))
            if d < min_d[c]:
                # d 是目前最小的，那么 min_d[c] 是次小的
                min2 = min(min2, min_d[c])
                min_d[c] = d
            else:
                # d 可能是次小的
                min2 = min(min2, d)
        return sum(d < min2 for d in min_d.values())
```

```java [sol-Java]
class Solution {
    public int maxPointsInsideSquare(int[][] points, String s) {
        int[] minD = new int[26];
        Arrays.fill(minD, Integer.MAX_VALUE);
        int min2 = Integer.MAX_VALUE;
        for (int i = 0; i < points.length; i++) {
            int d = Math.max(Math.abs(points[i][0]), Math.abs(points[i][1]));
            int c = s.charAt(i) - 'a';
            if (d < minD[c]) {
                // d 是目前最小的，那么 minD[c] 是次小的
                min2 = Math.min(min2, minD[c]);
                minD[c] = d;
            } else {
                // d 可能是次小的
                min2 = Math.min(min2, d);
            }
        }
        int ans = 0;
        for (int d : minD) {
            if (d < min2) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxPointsInsideSquare(vector<vector<int>>& points, string s) {
        int min_d[26], min2 = INT_MAX;
        ranges::fill(min_d, INT_MAX);
        for (int i = 0; i < points.size(); i++) {
            int d = max(abs(points[i][0]), abs(points[i][1]));
            char c = s[i] - 'a';
            if (d < min_d[c]) {
                // d 是目前最小的，那么 min_d[c] 是次小的
                min2 = min(min2, min_d[c]);
                min_d[c] = d;
            } else {
                // d 可能是次小的
                min2 = min(min2, d);
            }
        }
        int ans = 0;
        for (int d : min_d) {
            ans += d < min2;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxPointsInsideSquare(points [][]int, s string) (ans int) {
    minD := [26]int{}
    for i := range minD {
        minD[i] = math.MaxInt
    }
    min2 := math.MaxInt
    for i, p := range points {
        d := max(abs(p[0]), abs(p[1]))
        c := s[i] - 'a'
        if d < minD[c] {
            // d 是目前最小的，那么 minD[c] 是次小的
            min2 = min(min2, minD[c])
            minD[c] = d
        } else {
            // d 可能是次小的
            min2 = min(min2, d)
        }
    }
    for _, d := range minD {
        if d < min2 {
            ans++
        }
    }
    return
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

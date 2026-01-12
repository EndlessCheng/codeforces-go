膨胀长度越短，越能满足要求（不重叠）；膨胀长度越长，越不能满足要求（重叠）。据此可以**二分答案**。关于二分的原理，请看视频 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

现在问题变成：

- 给定膨胀长度 $m$，膨胀后，能否让所有城墙均无重叠？

⚠**注意**：题目说每个城墙可以向两个方向膨胀，这两个方向的膨胀长度可以不相等，只要两个方向的膨胀长度之和等于 $m$ 就行。

我们可以**贪心**。每段城墙**先向左膨胀，再向右膨胀**，如果向右膨胀后，与右侧的城墙重叠，则说明 $m$ 过大。

#### 细节

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的。

- 开区间左端点初始值：$0$。一定满足要求。
- 开区间左端点初始值（优化）：相邻城墙的最小间距。一定满足要求。
- 开区间右端点初始值：设间距之和为 $s$，那么平均值加一，即 $\left\lfloor\dfrac{s}{n-2}\right\rfloor+1$，一定无法满足要求。其中 $n-2$ 是去掉最左最右的城墙（这两个城墙向外膨胀即可），剩余 $n-2$ 个城墙分配内部的 $s$。

```py [sol-Python3]
class Solution:
    def rampartDefensiveLine(self, rampart: List[List[int]]) -> int:
        def check(m: int) -> bool:
            pre_r = -inf
            for (l, r), (nxt_l, _) in pairwise(rampart):
                space = m - (l - pre_r)  # 向左膨胀后的剩余长度
                if space > 0:
                    r += space  # 向右膨胀
                    if r > nxt_l:  # 无法膨胀
                        return False
                pre_r = r
            return True

        s = sum(l2 - r1 for (_, r1), (l2, _) in pairwise(rampart))
        left, right = 0, s // (len(rampart) - 2) + 1
        while left + 1 < right:
            mid = (left + right) // 2
            if check(mid):
                left = mid
            else:
                right = mid
        return left
```

```py [sol-Python3 库函数]
class Solution:
    def rampartDefensiveLine(self, rampart: List[List[int]]) -> int:
        def check(m: int) -> bool:
            m += 1
            pre_r = -inf
            for (l, r), (nxt_l, _) in pairwise(rampart):
                space = m - (l - pre_r)  # 向左膨胀后的剩余长度
                if space > 0:
                    r += space  # 向右膨胀
                    if r > nxt_l:  # 无法膨胀
                        return True
                pre_r = r
            return False

        s = sum(l2 - r1 for (_, r1), (l2, _) in pairwise(rampart))
        right = s // (len(rampart) - 2)
        return bisect_left(range(right), True, key=check)
```

```java [sol-Java]
class Solution {
    public int rampartDefensiveLine(int[][] rampart) {
        int n = rampart.length;
        int s = 0;
        for (int i = 0; i < n - 1; i++) {
            s += rampart[i + 1][0] - rampart[i][1];
        }

        int left = 0;
        int right = s / (n - 2) + 1;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (check(mid, rampart)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left;
    }

    private boolean check(int m, int[][] rampart) {
        int pre_r = Integer.MIN_VALUE / 2;
        for (int i = 0; i < rampart.length - 1; i++) {
            int l = rampart[i][0];
            int r = rampart[i][1];
            int space = m - (l - pre_r); // 向左膨胀后的剩余长度
            if (space > 0) {
                r += space; // 向右膨胀
                if (r > rampart[i + 1][0]) { // 无法膨胀
                    return false;
                }
            }
            pre_r = r;
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int rampartDefensiveLine(vector<vector<int>>& rampart) {
        int n = rampart.size();

        auto check = [&](int m) -> bool {
            int pre_r = INT_MIN / 2;
            for (int i = 0; i + 1 < n; i++) {
                int l = rampart[i][0], r = rampart[i][1];
                int space = m - (l - pre_r); // 向左膨胀后的剩余长度
                if (space > 0) {
                    r += space; // 向右膨胀
                    if (r > rampart[i + 1][0]) { // 无法膨胀
                        return false;
                    }
                }
                pre_r = r;
            }
            return true;
        };

        int s = 0;
        for (int i = 0; i + 1 < n; i++) {
            s += rampart[i + 1][0] - rampart[i][1];
        }
        int left = 0, right = s / (n - 2) + 1;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left;
    }
};
```

```go [sol-Go]
func rampartDefensiveLine(rampart [][]int) int {
	n := len(rampart)
	s := rampart[n-1][0] - rampart[0][1]
	for _, p := range rampart[1 : n-1] {
		s -= p[1] - p[0]
	}
	return sort.Search(s/(n-2), func(m int) bool {
		m++
		preR := math.MinInt / 2
		for i, p := range rampart[:n-1] {
			r := p[1]
			space := m - (p[0] - preR) // 向左膨胀后的剩余长度
			if space > 0 {
				r += space // 向右膨胀
				if r > rampart[i+1][0] { // 无法膨胀
					return true
				}
			}
			preR = r
		}
		return false
	})
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U$ 为二分的初始长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

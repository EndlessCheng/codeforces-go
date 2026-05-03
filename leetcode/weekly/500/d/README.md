固定点是元素值等于其下标的元素。

$\textit{nums}[i]$ 要想成为固定点，它必须位于下标 $\textit{nums}[i]$。例如 $x = \textit{nums}[5] = 3$，删除 $x$ 左侧的 $5-3=2$ 个数后，$x$ 就从下标 $5$ 移动到下标 $3$ 了，成为固定点。

一般地，删除 $x = \textit{nums}[i]$ 左侧的 $i - x$ 个数，就可以让 $x$ 位于下标 $x$，成为固定点。如果 $i-x < 0$，也就是 $i < x$，那么 $x$ 无法成为固定点。

对于 $\textit{nums}$ 中的满足 $i\ge \textit{nums}[i]$ 的元素，我们需要从中选出一个最长子序列，满足：

1. 由于固定点的值等于其下标，所以**子序列必须是严格递增的**。
2. 对于两个固定点 $i$ 和 $j$（$i<j$），由于 $j$ 左侧删除的元素，包含 $i$ 左侧删除的所有元素，所以对于这些固定点来说，其左侧「删除的元素个数」是递增的，即 $i-\textit{nums}[i]$ **是递增的**（允许相邻相等）。

把 $\textit{nums}$ 中的满足 $i\ge \textit{nums}[i]$ 的元素视作一个二元组 $(\textit{nums}[i],i-\textit{nums}[i])$，问题相当于：

- 想象有一些信封，每个信封的长宽用二元组表示。
- 如果信封 $A$ 的长宽都比另一个信封 $B$ 大（允许宽相等），那么 $A$ 可以装下 $B$，如同俄罗斯套娃一样。
- 计算最多能有多少个信封能组成一组俄罗斯套娃信封（不允许旋转信封）。

这题是 [354. 俄罗斯套娃信封问题](https://leetcode.cn/problems/russian-doll-envelopes/)，[我的题解](https://leetcode.cn/problems/russian-doll-envelopes/solutions/3785353/qiao-miao-pai-xu-zhuan-hua-cheng-yi-wei-25mb0/)。

注意二元组的第二个数允许相等，所以排序后，我们求的是第二个数的**最长非降子序列**的长度。非降的意思是：递增，允许相邻元素相等。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    # 354. 俄罗斯套娃信封问题
    def maxEnvelopes(self, envelopes: list[tuple[int, int]]) -> int:
        envelopes.sort(key=lambda e: (e[0], -e[1]))

        g = []
        for _, h in envelopes:
            j = bisect_right(g, h)  # 允许 LIS 相邻元素相等
            if j < len(g):
                g[j] = h
            else:
                g.append(h)
        return len(g)

    def maxFixedPoints(self, nums: list[int]) -> int:
        a = [(x, i - x) for i, x in enumerate(nums) if i >= x]
        return self.maxEnvelopes(a)
```

```java [sol-Java]
class Solution {
    public int maxFixedPoints(int[] nums) {
        List<int[]> a = new ArrayList<>();
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (i >= x) {
                a.add(new int[]{x, i - x});
            }
        }
        return maxEnvelopes(a);
    }

    // 354. 俄罗斯套娃信封问题
    private int maxEnvelopes(List<int[]> envelopes) {
        envelopes.sort((a, b) -> {
            if (a[0] == b[0]) {
                return b[1] - a[1];
            }
            return a[0] - b[0];
        });

        List<Integer> g = new ArrayList<>();
        for (int[] e : envelopes) {
            int h = e[1];
            int j = upperBound(g, h); // 允许 LIS 相邻元素相等
            if (j < g.size()) {
                g.set(j, h);
            } else {
                g.add(h);
            }
        }
        return g.size();
    }

    private int upperBound(List<Integer> g, int target) {
        int left = -1, right = g.size(); // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[right] > target
            // nums[left] <= target
            int mid = left + (right - left) / 2;
            if (g.get(mid) > target) {
                right = mid; // 范围缩小到 (left, mid)
            } else {
                left = mid; // 范围缩小到 (mid, right)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 354. 俄罗斯套娃信封问题
    int maxEnvelopes(vector<pair<int, int>>& envelopes) {
        ranges::sort(envelopes, {}, [](auto& e) { return pair(e.first, -e.second); });

        vector<int> g;
        for (auto& [_, h] : envelopes) {
            auto it = ranges::upper_bound(g, h); // 允许 LIS 相邻元素相等
            if (it != g.end()) {
                *it = h;
            } else {
                g.push_back(h);
            }
        }
        return g.size();
    }

public:
    int maxFixedPoints(vector<int>& nums) {
        vector<pair<int, int>> a;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            if (i >= x) {
                a.emplace_back(x, i - x);
            }
        }
        return maxEnvelopes(a);
    }
};
```

```go [sol-Go]
// 354. 俄罗斯套娃信封问题
func maxEnvelopes(envelopes [][2]int) int {
	slices.SortFunc(envelopes, func(a, b [2]int) int {
		return cmp.Or(a[0]-b[0], b[1]-a[1])
	})

	g := []int{}
	for _, e := range envelopes {
		h := e[1]
		j := sort.SearchInts(g, h+1) // 允许 LIS 相邻元素相等
		if j < len(g) {
			g[j] = h
		} else {
			g = append(g, h)
		}
	}
	return len(g)
}

func maxFixedPoints(nums []int) int {
	a := [][2]int{}
	for i, x := range nums {
		if i >= x {
			a = append(a, [2]int{x, i - x})
		}
	}
	return maxEnvelopes(a)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面动态规划题单的「**§4.2 最长递增子序列（LIS）**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

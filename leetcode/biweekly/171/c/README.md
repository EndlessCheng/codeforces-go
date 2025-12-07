为方便描述，下文把 $\textit{technique}_1$ 和 $\textit{technique}_2$ 简称为 $a$ 和 $b$。

同时考虑两个数组有些麻烦，能不能合并成一个数组呢？那样思考难度会降低很多。

我们可以先选择所有 $a[i]$，然后再从中撤销至多 $n-k$ 个，改成选择 $b$ 中的元素。

撤销一个 $a[i]$，改选 $b[i]$，分数增加了 $d[i] = b[i]-a[i]$。

要让总得分尽量大，应当选择最大的 $n-k$ 个 $d[i]$。注意 $d[i]\le 0$ 的数就不用选了。

只保留大于 $0$ 的 $d[i]$，然后把 $d$ 从大到小排序，选择前 $n-k$ 个数（不超过 $d$ 的大小），作为额外增加的分数。

[本题视频讲解](https://www.bilibili.com/video/BV1wr2fBpENB/?t=22m17s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxPoints(self, a: List[int], b: List[int], k: int) -> int:
        d = sorted((y - x for x, y in zip(a, b) if y > x), reverse=True)
        return sum(a) + sum(d[:len(a) - k])
```

```java [sol-Java]
class Solution {
    public long maxPoints(int[] a, int[] b, int k) {
        int n = a.length;
        long ans = 0;
        // ArrayList 比较慢，更快的写法见【Java 原地】
        List<Integer> diff = new ArrayList<>();

        for (int i = 0; i < n; i++) {
            ans += a[i];
            int d = b[i] - a[i];
            if (d > 0) {
                diff.add(d);
            }
        }

        diff.sort(Collections.reverseOrder());
        int limit = Math.min(n - k, diff.size());
        for (int i = 0; i < limit; i++) {
            ans += diff.get(i);
        }
        return ans;
    }
}
```

```java [sol-Java 原地]
class Solution {
    public long maxPoints(int[] a, int[] b, int k) {
        int n = a.length;
        long ans = 0;
        int idx = 0;

        for (int i = 0; i < n; i++) {
            ans += a[i];
            int d = b[i] - a[i];
            if (d > 0) {
                a[idx++] = d; // 把 a 当作 diff
            }
        }

        Arrays.sort(a, 0, idx);
        for (int i = Math.max(idx - (n - k), 0); i < idx; i++) {
            ans += a[i];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxPoints(vector<int>& a, vector<int>& b, int k) {
        int n = a.size();
        long long ans = 0;
        vector<int> diff;

        for (int i = 0; i < n; i++) {
            ans += a[i];
            int d = b[i] - a[i];
            if (d > 0) {
                diff.push_back(d);
            }
        }

        ranges::sort(diff, greater()); // 更快的写法见【C++ 快速选择】
        int limit = min(n - k, (int) diff.size());
        ans += reduce(diff.begin(), diff.begin() + limit, 0LL);
        return ans;
    }
};
```

```cpp [sol-C++ 快速选择]
class Solution {
public:
    long long maxPoints(vector<int>& a, vector<int>& b, int k) {
        int n = a.size();
        long long ans = 0;
        vector<int> diff;

        for (int i = 0; i < n; i++) {
            ans += a[i];
            int d = b[i] - a[i];
            if (d > 0) {
                diff.push_back(d);
            }
        }

        int limit = min(n - k, (int) diff.size());
        ranges::nth_element(diff, diff.begin() + limit, greater());
        ans += reduce(diff.begin(), diff.begin() + limit, 0LL);
        return ans;
    }
};
```

```go [sol-Go]
func maxPoints(a, b []int, k int) (ans int64) {
	n := len(a)
	d := a[:0]
	for i, x := range a {
		ans += int64(x)
		v := b[i] - x
		if v > 0 {
			d = append(d, v)
		}
	}

	slices.SortFunc(d, func(a, b int) int { return b - a })
	for _, x := range d[:min(n-k, len(d))] {
		ans += int64(x)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 是 $\textit{a}$ 的长度。快速选择可以做到 $\mathcal{O}(n)$，见 C++ 第二份代码。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$，取决于实现。

## 相似题目

- [2611. 老鼠和奶酪](https://leetcode.cn/problems/mice-and-cheese/)
- [3367. 移除边之后的权重最大和](https://leetcode.cn/problems/maximize-sum-of-weights-after-edge-removals/)

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

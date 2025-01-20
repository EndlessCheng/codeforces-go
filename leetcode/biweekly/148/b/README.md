把 $\textit{arr}$ 简记为 $a$，$\textit{brr}$ 简记为 $b$。

如果不使用操作一，那么答案为所有 $|a[i]-b[i]|$ 之和。

如果使用操作一，那么直接把 $a$ 分成 $n$ 个长为 $1$ 的子数组，这样 $a$ 就可以随意排了。

最优配对方式是最小的 $a[i]$ 与最小的 $b[i]$ 一对，次小的 $a[i]$ 与次小的 $b[i]$ 一对。用交换论证法可以证明这样做是最优的，详细解释请看 [视频讲解](https://www.bilibili.com/video/BV1xBwBeEEie/?t=1m49s)。

**优化**：如果 $k$ 很大，比只用操作二还大（或者相等），那么使用操作一一定不会得到更优的答案。所以可以据此提前返回，这样可以省去排序的时间。

```py [sol-Python3]
class Solution:
    def minCost(self, a: List[int], b: List[int], k: int) -> int:
        ans2 = sum(abs(x - y) for x, y in zip(a, b))
        if ans2 <= k:
            return ans2

        a.sort()
        b.sort()
        ans1 = sum(abs(x - y) for x, y in zip(a, b)) + k
        return min(ans1, ans2)
```

```java [sol-Java]
class Solution {
    public long minCost(int[] a, int[] b, long k) {
        long ans2 = 0;
        for (int i = 0; i < a.length; i++) {
            ans2 += Math.abs(a[i] - b[i]);
        }
        if (ans2 <= k) {
            return ans2;
        }

        Arrays.sort(a);
        Arrays.sort(b);
        long ans1 = k;
        for (int i = 0; i < a.length; i++) {
            ans1 += Math.abs(a[i] - b[i]);
        }

        return Math.min(ans1, ans2);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(vector<int>& a, vector<int>& b, long long k) {
        long long ans2 = 0;
        for (int i = 0; i < a.size(); i++) {
            ans2 += abs(a[i] - b[i]);
        }
        if (ans2 <= k) {
            return ans2;
        }

        ranges::sort(a);
        ranges::sort(b);
        long long ans1 = k;
        for (int i = 0; i < a.size(); i++) {
            ans1 += abs(a[i] - b[i]);
        }

        return min(ans1, ans2);
    }
};
```

```go [sol-Go]
func minCost(a, b []int, k int64) int64 {
	ans2 := int64(0)
	for i, x := range a {
		ans2 += int64(abs(x - b[i]))
	}
	if ans2 <= k {
		return ans2
	}

	slices.Sort(a)
	slices.Sort(b)
	ans1 := k
	for i, x := range a {
		ans1 += int64(abs(x - b[i]))
	}

	return min(ans1, ans2)
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $a$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

更多相似题目，见下面贪心题单中的「**§1.1 从最小/最大开始贪心**」和「**§1.7 交换论证法**」。

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
10. 【本题相关】[贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

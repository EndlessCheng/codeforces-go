对于相同的 $x[i]$，在这些 $x[i]$ 对应的 $y[i]$ 中，至多选一个。

看示例 1：

- $x[0]=x[2]=1$，对应的 $y[0]=5$ 和 $y[2]=4$，这两个数至多选一个。
- $x[1]=x[4]=2$，对应的 $y[1]=3$ 和 $y[4]=2$，这两个数至多选一个。
- $x[3]=3$，对应的 $y[3]=6$，至多选一个。

也就是说，把相同 $x[i]$ 对应的 $y[i]$ 放到同一组，每组至多选一个数。

贪心地，每组选最大的。所有最大值取前 $3$ 大。

示例 1 的答案为 $\max(5,4) + \max(3,2) + 6 = 5+3+6=14$。

如果不足 $3$ 组，返回 $-1$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def maxSumDistinctTriplet(self, x: List[int], y: List[int]) -> int:
        mx = defaultdict(int)
        for a, b in zip(x, y):
            mx[a] = max(mx[a], b)
        if len(mx) < 3:
            return -1
        return sum(nlargest(3, mx.values()))
```

```java [sol-Java]
class Solution {
    public int maxSumDistinctTriplet(int[] x, int[] y) {
        Map<Integer, Integer> mx = new HashMap<>();
        for (int i = 0; i < x.length; i++) {
            mx.merge(x[i], y[i], Math::max);
        }
        if (mx.size() < 3) {
            return -1;
        }
        List<Integer> a = new ArrayList<>(mx.values());
        a.sort((p, q) -> q - p);
        return a.get(0) + a.get(1) + a.get(2);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSumDistinctTriplet(vector<int>& x, vector<int>& y) {
        unordered_map<int, int> mx;
        for (int i = 0; i < x.size(); i++) {
            mx[x[i]] = max(mx[x[i]], y[i]);
        }
        if (mx.size() < 3) {
            return -1;
        }
        vector<int> a;
        for (auto& [_, val] : mx) {
            a.push_back(val);
        }
        ranges::sort(a, greater<>());
        return a[0] + a[1] + a[2];
    }
};
```

```go [sol-Go]
func maxSumDistinctTriplet(x, y []int) int {
	mx := map[int]int{}
	for i, v := range x {
		mx[v] = max(mx[v], y[i])
	}
	if len(mx) < 3 {
		return -1
	}
	a := slices.SortedFunc(maps.Values(mx), func(a, b int) int { return b - a })
	return a[0] + a[1] + a[2]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 是 $x$ 的长度。用堆或者手动维护前三大可以做到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

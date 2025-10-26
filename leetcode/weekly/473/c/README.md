**前置题目**：[560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/)，[我的题解](https://leetcode.cn/problems/subarray-sum-equals-k/solutions/2781031/qian-zhui-he-ha-xi-biao-cong-liang-ci-bi-4mwr/)。

设 $\textit{capacity}$ 的前缀和数组为 $s$。关于 $s$ 的定义，见 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)。

题目要求的式子，等价于

$$
\textit{capacity}[l] = \textit{capacity}[r] = s[r] - s[l+1]
$$

这等价于如下两个式子**同时成立**（第二个式子做了移项）

$$
\begin{aligned}
& \textit{capacity}[l] = \textit{capacity}[r]    \\
& \textit{capacity}[l] + s[l+1] = s[r]  \\
\end{aligned}
$$

枚举 $r$，问题变成：

- 有多少个左端点 $l$ 满足子数组长度 $r-l+1\ge 3$，且二元组 $(\textit{capacity}[l],\textit{capacity}[l] + s[l+1])$ 等于二元组 $(\textit{capacity}[r],s[r])$？

枚举 $r$，用哈希表维护左边的 $(\textit{capacity}[l],\textit{capacity}[l] + s[l+1])$ 的个数。

为保证 $r-l + 1\ge 3$，可以在枚举 $r$ 的同时，先查询哈希表更新答案，再把二元组 $(\textit{capacity}[r-1],\textit{capacity}[r-1] + s[r])$ 加到哈希表，这样对于下一轮循环的 $r+1$ 来说，把添加的 $r-1$ 作为 $l$，与 $r+1$ 构成的子数组长度就是 $(r+1)-(r-1)+1 = 3$，满足要求。

此外，按照上述方式更新哈希表，$s$ 可以简化成一个变量。

[本题视频讲解](https://www.bilibili.com/video/BV1eqxNzXE8v/?t=10m52s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countStableSubarrays(self, capacity: List[int]) -> int:
        cnt = defaultdict(int)
        s = capacity[0]  # 前缀和
        ans = 0
        for last, x in pairwise(capacity):
            ans += cnt[(x, s)]
            cnt[(last, last + s)] += 1
            s += x
        return ans
```

```java [sol-Java]
class Solution {
    private record Pair(int x, long s) {
    }

    public long countStableSubarrays(int[] capacity) {
        Map<Pair, Integer> cnt = new HashMap<>();
        long sum = capacity[0]; // 前缀和
        long ans = 0;
        for (int r = 1; r < capacity.length; r++) {
            ans += cnt.getOrDefault(new Pair(capacity[r], sum), 0);
            cnt.merge(new Pair(capacity[r - 1], capacity[r - 1] + sum), 1, Integer::sum);
            sum += capacity[r];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countStableSubarrays(vector<int>& capacity) {
        map<pair<int, long long>, int> cnt; // 另见【C++ 自定义哈希】
        long long sum = capacity[0]; // 前缀和
        long long ans = 0;
        for (int r = 1; r < capacity.size(); r++) {
            ans += cnt[{capacity[r], sum}];
            cnt[{capacity[r - 1], capacity[r - 1] + sum}]++;
            sum += capacity[r];
        }
        return ans;
    }
};
```

```cpp [sol-C++ 自定义哈希]
struct TupleHash {
    template<typename T>
    static void hash_combine(size_t& seed, const T& v) {
        // 参考 boost::hash_combine
        seed ^= hash<T>()(v) + 0x9e3779b9 + (seed << 6) + (seed >> 2);
    }

    template<typename Tuple, size_t Index = 0>
    static void hash_tuple(size_t& seed, const Tuple& t) {
        if constexpr (Index < tuple_size_v<Tuple>) {
            hash_combine(seed, get<Index>(t));
            hash_tuple<Tuple, Index + 1>(seed, t);
        }
    }

    template<typename... Ts>
    size_t operator()(const tuple<Ts...>& t) const {
        size_t seed = 0;
        hash_tuple(seed, t);
        return seed;
    }
};

class Solution {
public:
    long long countStableSubarrays(vector<int>& capacity) {
        unordered_map<tuple<int, long long>, int, TupleHash> cnt;
        long long sum = capacity[0]; // 前缀和
        long long ans = 0;
        for (int r = 1; r < capacity.size(); r++) {
            ans += cnt[{capacity[r], sum}];
            cnt[{capacity[r - 1], capacity[r - 1] + sum}]++;
            sum += capacity[r];
        }
        return ans;
    }
};
```

```go [sol-Go]
func countStableSubarrays(capacity []int) (ans int64) {
	type pair struct{ x, s int }
	cnt := map[pair]int{}
	sum := capacity[0] // 前缀和
	for r := 1; r < len(capacity); r++ {
		ans += int64(cnt[pair{capacity[r], sum}])
		cnt[pair{capacity[r-1], capacity[r-1] + sum}]++
		sum += capacity[r]
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{capacity}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 思考题

把最小长度 $3$ 改成 $4$，代码应该如何修改？改成 $k$ 呢？

## 专题训练

见下面数据结构题单的「**§1.2 前缀和与哈希表**」。

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

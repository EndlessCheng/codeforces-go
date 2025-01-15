首先，把两个数组中都有的数去掉，那么每个剩余数字的出现次数必须为偶数。这可以用哈希表来统计。

设处理后的剩余数组分别 $a$ 和 $b$。

贪心地想，如果要交换 $a$ 中最小的数，那么找一个 $b$ 中最大的数是最合适的；对于 $b$ 中最小的数也同理。

那么把 $a$ **从小到大**排序，$b$ **从大到小**排序，两两匹配。

但是，还有一种方案。

把 $\textit{basket}_1$ 和 $\textit{basket}_2$ 中的最小值 $\textit{mn}$ 当作「工具人」，对于 $a[i]$ 和 $b[i]$ 的交换，可以分别和 $\textit{mn}$ 交换一次，就相当于 $a[i]$ 和 $b[i]$ 交换了。

因此每次交换的代价为

$$
\min(a[i], b[i], 2\cdot\textit{mn})
$$

累加代价，即为答案。

上式也表明，如果工具人也在需要交换的数字中，那么它的最小代价必然是和其他数交换，不会发生工具人和工具人交换的情况。

设 $m$ 为 $a$ 的长度。代码实现时，由于 $\min(a[i], b[i])$ 的数字都在 $a$ 的某个前缀与 $b$ 某个后缀中，而剩下没有选的数（$a$ 的后缀和 $b$ 的前缀）不比这 $m$ 个数小，所以取出的数一定是这 $2m$ 个数中最小的 $m$ 个数。

> 更详细的证明：设选了 $a[0],\cdots,a[i]$ 和 $b[i+1],\cdots,b[m-1]$，由于 $a[i+1]\ge b[i+1]$ 且 $a[i+1] \ge a[i]$，所以 $a[i+1]$ 大于等于任意已选数字，进而推出 $a[i+2],\cdots,a[m-1]$ 都是大于等于任意已选数字的；对于 $b[i]$ 和 $b[i-1],\cdots,b[0]$ 同理。所以剩下没有选的数字都比已选数字大，进而说明已选数字是这 $2m$ 个数中最小的 $m$ 个数。

那么可以直接把 $a$ 和 $b$ 拼起来，从小到大排序后，遍历前一半的数即可（排序可以用快速选择代替，见 C++）。

[视频讲解](https://www.bilibili.com/video/BV1sG4y1T7oc/)

```py [sol-Python3]
class Solution:
    def minCost(self, basket1: List[int], basket2: List[int]) -> int:
        cnt = defaultdict(int)
        for x, y in zip(basket1, basket2):
            cnt[x] += 1
            cnt[y] -= 1

        mn = min(cnt)
        a = []
        for x, c in cnt.items():
            if c % 2:
                return -1
            a.extend([x] * (abs(c) // 2))

        a.sort()  # 也可以用快速选择
        return sum(min(x, mn * 2) for x in a[:len(a) // 2])
```

```java [sol-Java]
class Solution {
    public long minCost(int[] basket1, int[] basket2) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int i = 0; i < basket1.length; i++) {
            cnt.merge(basket1[i], 1, Integer::sum);
            cnt.merge(basket2[i], -1, Integer::sum);
        }

        int mn = Integer.MAX_VALUE;
        List<Integer> a = new ArrayList<>();
        for (var e : cnt.entrySet()) {
            int x = e.getKey();
            int c = e.getValue();
            if (c % 2 != 0) {
                return -1;
            }
            mn = Math.min(mn, x);
            for (c = Math.abs(c) / 2; c > 0; c--) {
                a.add(x);
            }
        }

        long ans = 0;
        a.sort((x, y) -> x - y); // 也可以用快速选择
        for (int i = 0; i < a.size() / 2; i++) {
            ans += Math.min(a.get(i), mn * 2);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minCost(vector<int>& basket1, vector<int>& basket2) {
        unordered_map<int, int> cnt;
        for (int i = 0; i < basket1.size(); i++) {
            cnt[basket1[i]]++;
            cnt[basket2[i]]--;
        }

        int mn = INT_MAX;
        vector<int> a;
        for (auto& [x, c] : cnt) {
            if (c % 2) {
                return -1;
            }
            mn = min(mn, x);
            for (c = abs(c) / 2; c > 0; c--) {
                a.push_back(x);
            }
        }

        long long ans = 0;
        ranges::nth_element(a, a.begin() + a.size() / 2); // 快速选择
        for (int i = 0; i < a.size() / 2; i++) {
            ans += min(a[i], mn * 2);
        }
        return ans;
    }
};
```

```go [sol-Go]
func minCost(basket1, basket2 []int) (ans int64) {
	cnt := map[int]int{}
	for i, x := range basket1 {
		cnt[x]++
		cnt[basket2[i]]--
	}

	mn := math.MaxInt
	a := []int{}
	for x, c := range cnt {
		if c%2 != 0 {
			return -1
		}
		mn = min(mn, x)
		for c = abs(c) / 2; c > 0; c-- {
			a = append(a, x)
		}
	}

	slices.Sort(a) // 也可以用快速选择
	for _, x := range a[:len(a)/2] {
		ans += int64(min(x, mn*2))
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 为 $\textit{basket}_1$ 的长度。用快速选择可以做到 $\mathcal{O}(n)$（见 C++）。
- 空间复杂度：$\mathcal{O}(n)$。

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

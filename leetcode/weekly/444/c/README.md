## 错误思路

错误思路：DP，把子序列乘积的**最大值**作为 $\textit{dfs}$ 的返回值。

错误原因：类似背包问题，最大乘积不一定最优，可能较小的乘积与其他数相乘，能够更接近 $\textit{limit}$。

## 分析

注意到，如果乘积不为 $0$，那么在乘积不超过 $\textit{limit}$ 的前提下，子序列至多包含 $L = \left\lfloor\log_2 \textit{limit}\right\rfloor$ 个大于 $1$ 的数，以及零个或多个 $1$。

在本题的数据范围下，$L\le 12$。

在两个大于 $1$ 的数之间的连续的 $1$，其交错和 $1-1+1-1+\cdots$ 的绝对值 $\le 1$。所以子序列中的所有 $1$ 的交错和的绝对值 $\le L+1$。

- 如果大于 $1$ 的数都是 $2$，交错和的绝对值 $\le 2L + (L+1) = 3L+1\le 37$。
- 如果大于 $1$ 的数都是 $3$，那么 $L\le 7$，交错和的绝对值 $\le 3L + (L+1) = 4L+1\le 29$。

可以看出，**在乘积不为 $0$ 且不超过 $\textit{limit}$ 的情况下，交错和的绝对值其实远远小于** $\textit{nums}$ 的元素和！

此外，$150$ 个 $[1,12]$ 中的数相乘，只有 $394$ 个 $\le 5000$ 的不同乘积。（计算代码见文末）

> 为什么这么少？因为大于 $12$ 的质数（及其倍数）是无法得到的。想一想，你能得到 $13$ 吗？能得到 $26$ 吗？

如果乘积为 $0$ 呢？继续向后（递归），乘积仍然为 $0$，**此时只需关注交错和**，不同交错和的个数 $\le 150\cdot 12 = 1800$ 也很小。

所以，状态个数比预期的少，直接暴力搜索即可（不用写 DP）。

## 思路

写一个爆搜 $\textit{dfs}(i,s,m,\textit{odd},\textit{empty})$，各个参数分别表示：

- 当前要考虑 $x=\textit{nums}[i]$ 选或不选。
- 子序列的交错和为 $s$。
- 子序列的元素积为 $m$。
- 如果选 $x$，是加 $x$ 还是减 $x$。
- 子序列是否为空。

分类讨论：

- 不选 $x$，递归到 $\textit{dfs}(i+1,s,m,\textit{odd},\textit{empty})$。
- 选 $x$，递归到 $\textit{dfs}(i+1,s',\min(m\cdot x,\textit{limit}+1),\texttt{not}\textit{odd},\texttt{false})$。其中 $s'$ 是 $s+x$ 或者 $s-x$，如果 $\textit{odd}=\texttt{false}$ 则加，否则减。这里超过 $\textit{limit}$ 的乘积一律视作 $\textit{limit}+1$，减少状态个数。

**递归终点**：如果 $i=n$，并且 $\textit{empty}=\texttt{false}$ 且 $s=k$ 且 $m\le \textit{limit}$，那么用 $m$ 更新答案的最大值。

**递归入口**：$\textit{dfs}(0, 0, 1, \texttt{false}, \texttt{true})$。加法单位元是 $0$，乘法单位元是 $1$。

递归过程中，用 $\textit{vis}$ 哈希表记录访问过的状态，避免重复访问。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ezRvYiE27/?t=16m16s)，欢迎点赞关注~

> 注：`sum(nums) < abs(k)` 可以优化成 `sum(sorted(nums)[-((n + 1) // 2):]) < abs(k)`。

```py [sol-Python3]
class Solution:
    def maxProduct(self, nums: List[int], k: int, limit: int) -> int:
        if sum(nums) < abs(k):  # |k| 太大
            return -1

        n = len(nums)
        ans = -1

        @cache  # 当 vis 哈希表用
        def dfs(i: int, s: int, m: int, odd: bool, empty: bool) -> None:
            nonlocal ans
            if ans == limit:  # 已经达到最大值
                return

            if i == n:
                if not empty and s == k and m <= limit:  # 合法子序列
                    ans = max(ans, m)  # 用合法子序列的元素积更新答案的最大值
                return

            # 不选 x
            dfs(i + 1, s, m, odd, empty)

            # 选 x
            x = nums[i]
            dfs(i + 1, s + (-x if odd else x), min(m * x, limit + 1), not odd, False)

        dfs(0, 0, 1, False, True)
        return ans
```

```java [sol-Java]
class Solution {
    private int ans = -1;

    public int maxProduct(int[] nums, int k, int limit) {
        int sum = Arrays.stream(nums).sum();
        if (sum < Math.abs(k)) { // |k| 太大
            return -1;
        }

        Set<Long> vis = new HashSet<>();
        dfs(0, 0, 1, false, true, nums, k, limit, sum, vis);
        return ans;
    }

    private void dfs(int i, int s, int m, boolean odd, boolean empty, int[] nums, int k, int limit, int bias, Set<Long> vis) {
        if (ans == limit) { // 已经达到最大值
            return;
        }
    
        if (i == nums.length) {
            if (!empty && s == k && m <= limit) { // 合法子序列
                ans = Math.max(ans, m); // 用合法子序列的元素积更新答案的最大值
            }
            return;
        }

        long mask = (long) i << 32 | (s + bias) << 14 | m << 2 | (odd ? 1 : 0) << 1 | (empty ? 1 : 0);
        if (!vis.add(mask)) {
            return;
        }

        // 不选 x
        dfs(i + 1, s, m, odd, empty, nums, k, limit, bias, vis);

        // 选 x
        int x = nums[i];
        dfs(i + 1, s + (odd ? -x : x), Math.min(m * x, limit + 1), !odd, false, nums, k, limit, bias, vis);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxProduct(vector<int>& nums, int k, int limit) {
        int total = reduce(nums.begin(), nums.end());
        if (total < abs(k)) { // |k| 太大
            return -1;
        }

        int n = nums.size(), ans = -1;
        unordered_set<long long> vis;
        auto dfs = [&](this auto&& dfs, int i, int s, int m, bool odd, bool empty) -> void {
            if (ans == limit) { // 已经达到最大值
                return;
            }
        
            if (i == n) {
                if (!empty && s == k && m <= limit) { // 合法子序列
                    ans = max(ans, m); // 用合法子序列的元素积更新答案的最大值
                }
                return;
            }

            long long mask = (long long) i << 32 | (s + total) << 14 | m << 2 | odd << 1 | empty;
            if (!vis.insert(mask).second) {
                return;
            }

            // 不选 x
            dfs(i + 1, s, m, odd, empty);

            // 选 x
            int x = nums[i];
            dfs(i + 1, s + (odd ? -x : x), min(m * x, limit + 1), !odd, false);
        };
        dfs(0, 0, 1, false, true);
        return ans;
    }
};
```

```go [sol-Go]
func maxProduct(nums []int, k, limit int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	if total < abs(k) { // |k| 太大
		return -1
	}

	ans := -1
	type args struct {
		i, s, m    int
		odd, empty bool
	}
	vis := map[args]bool{}
	var dfs func(int, int, int, bool, bool)
	dfs = func(i, s, m int, odd, empty bool) {
		if ans == limit { // 已经达到最大值
			return
		}

		if i == len(nums) {
			if !empty && s == k && m <= limit { // 合法子序列
				ans = max(ans, m) // 用合法子序列的元素积更新答案的最大值
			}
			return
		}

		t := args{i, s, m, odd, empty}
		if vis[t] {
			return
		}
		vis[t] = true

		// 不选 x
		dfs(i+1, s, m, odd, empty)

		// 选 x
		x := nums[i]
		if odd {
			s -= x
		} else {
			s += x
		}
		dfs(i+1, s, min(m*x, limit+1), !odd, false)
	}
	dfs(0, 0, 1, false, true)
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

请注意，$150$ 个 $[1,12]$ 中的数相乘，只有 $M=394$ 个 $\le 5000$ 的不同乘积。

- 时间复杂度：$\mathcal{O}(n(nU + M\log \textit{limit}))$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})\le 12$。其中 $nU$ 对应 $m=0$ 或者 $m=\textit{limit}+1$ 的情况，$M\log \textit{limit}$ 对应 $1\le m\le \textit{limit}$ 的情况。
- 空间复杂度：$\mathcal{O}(n(nU + M\log \textit{limit}))$。

## 附：递推写法

```py [sol-Python3]
class Solution:
    def maxProduct(self, nums: List[int], k: int, limit: int) -> int:
        if sum(nums) < abs(k):  # |k| 太大
            return -1

        odd_s = defaultdict(set)  # s -> {m}
        even_s = defaultdict(set)  # s -> {m}

        for x in nums:
            # 长为偶数的子序列的计算结果 new_even_s
            new_even_s = defaultdict(set)
            for s, st in odd_s.items():
                new_even_s[s - x] = {m * x for m in st if m * x <= limit}

            # 长为奇数的子序列的计算结果 odd_s
            for s, st in even_s.items():
                odd_s[s + x] |= {m * x for m in st if m * x <= limit}
                if x == 0:
                    odd_s[s].add(0)

            # 用 new_even_s 更新 even_s
            for s in new_even_s:
                even_s[s] |= new_even_s[s]
                if x == 0:
                    even_s[s].add(0)

            # 子序列只有一个数的情况
            if x <= limit:
                odd_s[x].add(x)  

            if k in odd_s and limit in odd_s[k] or k in even_s and limit in even_s[k]:
                return limit  # 提前返回

        return max(max(odd_s[k], default=-1), max(even_s[k], default=-1))
```

```go [sol-Go]
func maxProduct(nums []int, k, limit int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	if total < abs(k) { // |k| 太大
		return -1
	}

	// s -> {m}
	oddS := map[int]map[int]struct{}{}
	evenS := map[int]map[int]struct{}{}
	add := func(m map[int]map[int]struct{}, key, val int) {
		if _, ok := m[key]; !ok {
			m[key] = map[int]struct{}{}
		}
		m[key][val] = struct{}{}
	}

	for _, x := range nums {
		// 长为偶数的子序列的计算结果 newEvenS
		newEvenS := map[int]map[int]struct{}{}
		for s, set := range oddS {
			newEvenS[s-x] = map[int]struct{}{}
			for m := range set {
				if m*x <= limit {
					newEvenS[s-x][m*x] = struct{}{}
				}
			}
		}

		// 长为奇数的子序列的计算结果 oddS
		for s, set := range evenS {
			if _, ok := oddS[s+x]; !ok {
				oddS[s+x] = map[int]struct{}{}
			}
			for m := range set {
				if m*x <= limit {
					oddS[s+x][m*x] = struct{}{}
				}
			}
			if x == 0 {
				add(oddS, s, 0)
			}
		}

		// 用 newEvenS 更新 evenS
		for s, set := range newEvenS {
			if eSet, ok := evenS[s]; ok {
				for m := range set {
					eSet[m] = struct{}{}
				}
			} else {
				evenS[s] = set
			}
			if x == 0 {
				add(evenS, s, 0)
			}
		}

		// 子序列只有一个数的情况
		if x <= limit {
			add(oddS, x, x)
		}

		if set, ok := oddS[k]; ok {
			if _, ok := set[limit]; ok {
				return limit // 提前返回
			}
		}
		if set, ok := evenS[k]; ok {
			if _, ok := set[limit]; ok {
				return limit // 提前返回
			}
		}
	}

	calcMax := func(m map[int]struct{}) int {
		maxVal := -1
		if m != nil {
			for v := range m {
				maxVal = max(maxVal, v)
			}
		}
		return maxVal
	}
	return max(calcMax(oddS[k]), calcMax(evenS[k]))
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

同上。

## 附：如何计算不同乘积个数

```py
st = {1}
for _ in range(150):
    new_st = set()
    for m in st:
        for x in range(1, 13):
            if m * x <= 5000:
                new_st.add(m * x)
    st = new_st
print(len(st))  # 394
```

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

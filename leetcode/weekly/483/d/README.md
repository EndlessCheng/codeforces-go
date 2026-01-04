## 状态定义与状态转移方程

想一想，**最后一步发生了什么**？

我们合并了两个有序数组，分别来自 $\textit{lists}$ 中的一部分数组，以及 $\textit{lists}$ 中的另一部分数组。

这两部分数组都是 $\textit{lists}$ 的**子集**。子集又可以用同样的方法，拆分成两个更小的子集。这是和原问题相似的子问题。

根据 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/) §9.4 子集状压 DP 的套路，定义 $f[S]$ 表示已选数组（的下标）集合为 $S$ 时，把 $S$ 合并为一个有序列表所需的最小总成本。

枚举 $S$ 的非空真子集 $T$，问题变成得到 $T$ 和 $\complement_ST$ 的最小总成本，即 $f[T] + f[\complement_ST]$，再加上合并 $S$ 和 $T$ 的代价 $\text{len}(S) + \text{len}(T) + |\text{med}(S) - \text{med}(T)|$。

所有情况取最小值，得

$$
f[S] = \min_{\varnothing \neq T \subset S} f[T] + f[\complement_ST] + \text{len}(S) + \text{len}(T) + |\text{med}(S) - \text{med}(T)|
$$

初始值：如果 $S$ 是单元素集合，则 $f[S] = 0$，无需操作。

答案为 $f[U]$，其中 $U=\{0,1,2,\ldots,n-1\}$，$n$ 是 $\textit{lists}$ 的长度。

## 预处理写法一：合并有序数组

由于上述 DP 过程会反复计算同一个 $\text{len}(S)$ 和 $\text{med}(S)$，我们可以预处理所有 $\text{len}(S)$ 和 $\text{med}(S)$，这也是一个 DP。

考虑刷表法，枚举不在 $S$ 中的下标 $i$，设 $T = S\cup\{i\}$，那么

$$
\text{len}(T) = \text{len}(S) + \text{len}(\textit{lists}[i])
$$

为了计算 $\text{med}(S)$，我们还需要预处理 $\text{sorted}(S)$，即 $S$ 中的所有列表合并后的结果。同上，我们有

$$
\text{sorted}(T) = \texttt{merge}(\text{sorted}(S),\textit{lists}[i])
$$

其中 $\texttt{merge}$ 为 [88. 合并两个有序数组](https://leetcode.cn/problems/merge-sorted-array/)（创建一个新的数组再合并）。

根据题目中位数的定义，设 $m = \left\lfloor\dfrac{\text{len}(T)-1}{2}\right\rfloor$，我们有

$$
\text{med}(T) = \text{sorted}(T)[m]
$$

代码实现时，用二进制表示集合，用位运算实现集合操作，具体请看 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

此外，根据对称性，枚举子集时可以只枚举 $T > \complement_ST$ 的子集 $T$，其余子集是重复枚举。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
# 手写 min 更快
min = lambda a, b: b if b < a else a

class Solution:
    def minMergeCost(self, lists: List[List[int]]) -> int:
        u = 1 << len(lists)
        sum_len = [0] * u
        sorted_ = [[] for _ in range(u)]
        median = [0] * u

        for i, a in enumerate(lists):  # 枚举不在 s 中的下标 i
            high_bit = 1 << i
            for s in range(high_bit):
                t = high_bit | s
                sum_len[t] = sum_len[s] + len(a)
                b = sorted_[s] + a
                b.sort()  # 线性合并的写法见另一份代码【Python3 写法二】
                sorted_[t] = b
                median[t] = b[(len(b) - 1) // 2]

        f = [inf] * u
        for i in range(u):
            if i & (i - 1) == 0:  # i 只包含一个元素，无法分解成两个非空子集
                f[i] = 0
                continue
            # 枚举 i 的非空真子集 j
            j = i & (i - 1)
            while j > (i ^ j):
                k = i ^ j  # j 关于 i 的补集是 k
                f[i] = min(f[i], f[j] + f[k] + sum_len[j] + sum_len[k] + abs(median[j] - median[k]))
                j = (j - 1) & i

        return f[-1]
```

```py [sol-Python3 写法二]
# 手写 min 更快
min = lambda a, b: b if b < a else a

class Solution:
    def minMergeCost(self, lists: List[List[int]]) -> int:
        # 88. 合并两个有序数组（创建一个新数组）
        def merge(a: List[int], b: List[int]) -> List[int]:
            i, n = 0, len(a)
            j, m = 0, len(b)
            res = []
            while i < n or j < m:
                if j == m or i < n and a[i] < b[j]:
                    res.append(a[i])
                    i += 1
                else:
                    res.append(b[j])
                    j += 1
            return res

        u = 1 << len(lists)
        sum_len = [0] * u
        sorted_ = [[] for _ in range(u)]
        median = [0] * u

        for i, a in enumerate(lists):  # 枚举不在 s 中的下标 i
            high_bit = 1 << i
            for s in range(high_bit):
                t = high_bit | s
                sum_len[t] = sum_len[s] + len(a)
                b = merge(sorted_[s], a)
                sorted_[t] = b
                median[t] = b[(len(b) - 1) // 2]

        f = [inf] * u
        for i in range(u):
            if i & (i - 1) == 0:  # i 只包含一个元素，无法分解成两个非空子集
                f[i] = 0
                continue
            # 枚举 i 的非空真子集 j
            j = i & (i - 1)
            while j > (i ^ j):
                k = i ^ j  # j 关于 i 的补集是 k
                f[i] = min(f[i], f[j] + f[k] + sum_len[j] + sum_len[k] + abs(median[j] - median[k]))
                j = (j - 1) & i

        return f[-1]
```

```java [sol-Java]
class Solution {
    public long minMergeCost(int[][] lists) {
        int n = lists.length;
        int u = 1 << n;
        int[] sumLen = new int[u];
        int[][] sorted = new int[u][];
        int[] median = new int[u];

        for (int i = 0; i < n; i++) { // 枚举不在 s 中的下标 i
            int highBit = 1 << i;
            for (int s = 0; s < highBit; s++) {
                int t = highBit | s;
                sumLen[t] = sumLen[s] + lists[i].length;
                int[] b = merge(sorted[s], lists[i]);
                sorted[t] = b;
                median[t] = b[(b.length - 1) / 2];
            }
        }

        long[] f = new long[u];
        for (int i = 0; i < u; i++) {
            if ((i & (i - 1)) == 0) { // i 只包含一个元素，无法分解成两个非空子集
                continue; // f[i] = 0
            }
            f[i] = Long.MAX_VALUE;
            // 枚举 i 的非空真子集 j
            for (int j = i & (i - 1); j > (i ^ j); j = (j - 1) & i) {
                int k = i ^ j; // j 关于 i 的补集是 k
                f[i] = Math.min(f[i], f[j] + f[k] + sumLen[j] + sumLen[k] + Math.abs(median[j] - median[k]));
            }
        }

        return f[u - 1];
    }

    // 88. 合并两个有序数组（创建一个新数组）
    private int[] merge(int[] a, int[] b) {
        if (a == null) {
            return b;
        }
        int n = a.length;
        int m = b.length;
        int[] res = new int[n + m];
        int i = 0;
        int j = 0;
        int idx = 0;
        while (i < n || j < m) {
            if (j == m || i < n && a[i] < b[j]) {
                res[idx++] = a[i++];
            } else {
                res[idx++] = b[j++];
            }
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 88. 合并两个有序数组（创建一个新数组）
    vector<int> merge(const vector<int>& a, const vector<int>& b) {
        int n = a.size(), m = b.size();
        vector<int> res;
        res.reserve(n + m);
        int i = 0, j = 0;
        while (i < n || j < m) {
            if (j == m || i < n && a[i] < b[j]) {
                res.push_back(a[i++]);
            } else {
                res.push_back(b[j++]);
            }
        }
        return res;
    }

public:
    long long minMergeCost(vector<vector<int>>& lists) {
        int n = lists.size();
        int u = 1 << n;
        vector<int> sum_len(u);
        vector<int> median(u);
        vector<vector<int>> sorted(u);

        for (int i = 0; i < n; i++) { // 枚举不在 s 中的下标 i
            int high_bit = 1 << i;
            for (int s = 0; s < high_bit; s++) {
                int t = high_bit | s;
                sum_len[t] = sum_len[s] + lists[i].size();
                vector<int> b = merge(sorted[s], lists[i]);
                median[t] = b[(b.size() - 1) / 2];
                sorted[t] = move(b);
            }
        }

        vector<long long> f(u);
        for (int i = 0; i < u; i++) {
            if ((i & (i - 1)) == 0) { // i 只包含一个元素，无法分解成两个非空子集
                continue; // f[i] = 0
            }
            f[i] = LLONG_MAX;
            // 枚举 i 的非空真子集 j
            for (int j = i & (i - 1); j > (i ^ j); j = (j - 1) & i) {
                int k = i ^ j; // j 关于 i 的补集是 k
                f[i] = min(f[i], f[j] + f[k] + sum_len[j] + sum_len[k] + abs(median[j] - median[k]));
            }
        }

        return f[u - 1];
    }
};
```

```go [sol-Go]
func minMergeCost(lists [][]int) int64 {
	u := 1 << len(lists)
	sumLen := make([]int, u)
	sorted := make([][]int, u)
	median := make([]int, u)
	for i, a := range lists { // 枚举不在 s 中的下标 i
		highBit := 1 << i
		for s, sl := range sumLen[:highBit] {
			t := highBit | s
			sumLen[t] = sl + len(a)
			b := merge(sorted[s], a)
			sorted[t] = b
			median[t] = b[(len(b)-1)/2]
		}
	}

	f := make([]int, u)
	for i := range f {
		if i&(i-1) == 0 { // i 只包含一个元素，无法分解成两个非空子集
			continue // f[i] = 0
		}
		f[i] = math.MaxInt
		// 枚举 i 的非空真子集 j
		for j := i & (i - 1); j > i^j; j = (j - 1) & i {
			k := i ^ j // j 关于 i 的补集是 k
			f[i] = min(f[i], f[j]+f[k]+sumLen[j]+sumLen[k]+abs(median[j]-median[k]))
		}
	}
	return int64(f[u-1])
}

// 88. 合并两个有序数组（创建一个新数组）
func merge(a, b []int) []int {
	i, n := 0, len(a)
	j, m := 0, len(b)
	res := make([]int, 0, n+m)
	for {
		if i == n {
			return append(res, b[j:]...)
		}
		if j == m {
			return append(res, a[i:]...)
		}
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L\cdot 2^n + 3^n)$，其中 $n$ 是 $\textit{lists}$ 的长度，$L\le 2000$ 是所有 $\textit{list}[i]$ 的长度之和。对于预处理，瓶颈在 $\texttt{merge}$ 上，或者说所有 $\text{sorted}(S)$ 的长度之和。考虑每个元素的贡献，在 $2^n$ 个子集中，每个 $\textit{lists}[i][j]$ 恰好出现在其中的 $2^{n-1}$ 个子集中（选或不选），所以每个元素对 $\text{sorted}(S)$ 的长度之和的贡献是 $\mathcal{O}(2^n)$，所以预处理的时间复杂度为 $\mathcal{O}(L\cdot 2^n)$。子集状压 DP 的时间复杂度为 $\mathcal{O}(3^n)$，证明见动态规划题单的 §9.4 子集状压 DP。
- 空间复杂度：$\mathcal{O}(L\cdot 2^n)$。所有 $\text{sorted}(S)$ 的长度之和为 $\mathcal{O}(L\cdot 2^n)$。

## 预处理写法二：二分中位数

$S$ 的中位数即 $S$ 的第 $\left\lceil\dfrac{\text{len}(S)}{2}\right\rceil$ 小。

**套路**：第 $k$ 小等价于，求**最小**的 $x$，满足 $\le x$ 的数**至少**有 $k$ 个。

转为二分答案，在 $S$ 的每个数组中二分查找 $\le x$ 的元素个数，如下表：

| **需求**  | **写法**  |
|---|---|
| $< x$ 的元素个数  | $\texttt{lowerBound}(\textit{nums},x)$  | 
| $\le x$ 的元素个数 | $\texttt{lowerBound}(\textit{nums},x+1)$  | 
| $\ge x$ 的元素个数  | $n - \texttt{lowerBound}(\textit{nums},x)$  | 
| $> x$ 的元素个数  | $n - \texttt{lowerBound}(\textit{nums},x+1)$  | 

关于二分查找的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

```py [sol-Python3]
# 手写 min 更快
min = lambda a, b: b if b < a else a

class Solution:
    def minMergeCost(self, lists: List[List[int]]) -> int:
        u = 1 << len(lists)
        sum_len = [0] * u
        for i, len_a in enumerate(map(len, lists)):  # 枚举不在 s 中的下标 i
            high_bit = 1 << i
            for s in range(high_bit):
                sum_len[high_bit | s] = sum_len[s] + len_a

        all_elements = sorted({x for a in lists for x in a})
        median = [0] * u

        for mask in range(1, u):
            k = (sum_len[mask] + 1) // 2

            def check(x: int) -> bool:
                cnt = 0
                for i, a in enumerate(lists):
                    if mask >> i & 1 == 0:
                        continue
                    cnt += bisect_right(a, x)
                    if cnt >= k:
                        return True
                return False

            i = bisect_left(all_elements, True, key=check)
            median[mask] = all_elements[i]

        f = [inf] * u
        for i in range(u):
            if i & (i - 1) == 0:  # i 只包含一个元素，无法分解成两个非空子集
                f[i] = 0
                continue
            # 枚举 i 的非空真子集 j
            j = i & (i - 1)
            while j > (i ^ j):
                k = i ^ j  # j 关于 i 的补集是 k
                f[i] = min(f[i], f[j] + f[k] + sum_len[j] + sum_len[k] + abs(median[j] - median[k]))
                j = (j - 1) & i

        return f[-1]
```

```java [sol-Java]
class Solution {
    public long minMergeCost(int[][] lists) {
        int n = lists.length;
        int u = 1 << n;
        int[] sumLen = new int[u];
        for (int i = 0; i < n; i++) { // 枚举不在 s 中的下标 i
            int highBit = 1 << i;
            for (int s = 0; s < highBit; s++) {
                sumLen[highBit | s] = sumLen[s] + lists[i].length;
            }
        }

        int[] median = new int[u];
        for (int mask = 0; mask < u; mask++) {
            int k = (sumLen[mask] + 1) / 2;
            int left = (int) -1e9 - 1;
            int right = (int) 1e9;
            while (left + 1 < right) {
                int mid = left + (right - left) / 2;
                if (check(lists, mask, k, mid)) {
                    right = mid;
                } else {
                    left = mid;
                }
            }
            median[mask] = right;
        }

        long[] f = new long[u];
        for (int i = 0; i < u; i++) {
            if ((i & (i - 1)) == 0) { // i 只包含一个元素，无法分解成两个非空子集
                continue; // f[i] = 0
            }
            f[i] = Long.MAX_VALUE;
            // 枚举 i 的非空真子集 j
            for (int j = i & (i - 1); j > (i ^ j); j = (j - 1) & i) {
                int k = i ^ j; // j 关于 i 的补集是 k
                f[i] = Math.min(f[i], f[j] + f[k] + sumLen[j] + sumLen[k] + Math.abs(median[j] - median[k]));
            }
        }

        return f[u - 1];
    }

    private boolean check(int[][] lists, int mask, int k, int x) {
        int cnt = 0;
        for (int i = 0; i < lists.length; i++) {
            if ((mask >> i & 1) == 0) {
                continue;
            }
            cnt += upperBound(lists[i], x);
            if (cnt >= k) {
                return true;
            }
        }
        return false;
    }

    // 开区间写法
    // https://www.bilibili.com/video/BV1AP41137w7/
    private int upperBound(int[] nums, int target) {
        int left = -1;
        int right = nums.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] <= target
            // nums[right] > target
            int mid = left + (right - left) / 2;
            if (nums[mid] > target) {
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
public:
    long long minMergeCost(vector<vector<int>>& lists) {
        int n = lists.size();
        int u = 1 << n;
        vector<int> sum_len(u);
        for (int i = 0; i < n; i++) { // 枚举不在 s 中的下标 i
            int high_bit = 1 << i;
            for (int s = 0; s < high_bit; s++) {
                sum_len[high_bit | s] = sum_len[s] + lists[i].size();
            }
        }

        vector<int> median(u); // 每个子集的中位数
        for (int mask = 0; mask < u; mask++) {
            int sl = sum_len[mask];

            auto check = [&](int med) -> bool {
                int cnt = 0;
                for (int i = 0; i < n; i++) {
                    if ((mask >> i & 1) == 0) {
                        continue;
                    }
                    cnt += ranges::upper_bound(lists[i], med) - lists[i].begin();
                    if (cnt >= (sl + 1) / 2) {
                        return true;
                    }
                }
                return false;
            };

            int left = -1e9 - 1, right = 1e9;
            while (left + 1 < right) {
                int mid = left + (right - left) / 2;
                (check(mid) ? right : left) = mid;
            }
            median[mask] = right;
        }

        vector<long long> f(u);
        for (int i = 0; i < u; i++) {
            if ((i & (i - 1)) == 0) { // i 只包含一个元素，无法分解成两个非空子集
                continue; // f[i] = 0
            }
            f[i] = LLONG_MAX;
            // 枚举 i 的非空真子集 j
            for (int j = i & (i - 1); j > (i ^ j); j = (j - 1) & i) {
                int k = i ^ j; // j 关于 i 的补集是 k
                f[i] = min(f[i], f[j] + f[k] + sum_len[j] + sum_len[k] + abs(median[j] - median[k]));
            }
        }

        return f[u - 1];
    }
};
```

```go [sol-Go]
func minMergeCost(lists [][]int) int64 {
	u := 1 << len(lists)
	sumLen := make([]int, u)
	for i, a := range lists { // 枚举不在 s 中的下标 i
		highBit := 1 << i
		for s, sl := range sumLen[:highBit] {
			sumLen[highBit|s] = sl + len(a)
		}
	}

	median := make([]int, u)
	for mask, sl := range sumLen {
		left, right := int(-1e9), int(1e9)
		median[mask] = left + sort.Search(right-left, func(med int) bool {
			med += left
			cnt := 0
			for s := uint32(mask); s > 0; s &= s - 1 {
				i := bits.TrailingZeros32(s)
				cnt += sort.SearchInts(lists[i], med+1)
				if cnt >= (sl+1)/2 {
					return true
				}
			}
			return false
		})
	}

	f := make([]int, u)
	for i := range f {
		if i&(i-1) == 0 {
			continue // f[i] = 0
		}
		f[i] = math.MaxInt
		// 枚举 i 的非空真子集 j
		for j := i & (i - 1); j > i^j; j = (j - 1) & i {
			k := i ^ j // j 关于 i 的补集是 k
			f[i] = min(f[i], f[j]+f[k]+sumLen[j]+sumLen[k]+abs(median[j]-median[k]))
		}
	}
	return int64(f[u-1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(2^n\cdot n\log (n\ell)\log \ell + 3^n)$ 或 $\mathcal{O}(2^n\cdot n\log U\log \ell + 3^n)$，取决于实现，其中 $n$ 是 $\textit{lists}$ 的长度，$U\le 2\times 10^9$ 是二分答案的范围大小，$\ell$ 是 $\textit{lists}[i]$ 的平均长度。
- 空间复杂度：$\mathcal{O}(2^n)$。

## 专题训练

1. 动态规划题单的「**§9.4 子集状压 DP**」。
2. 二分题单的「**§2.6 第 K 小/大**」。

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

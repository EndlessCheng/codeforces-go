## 题意

从 $\textit{nums}$ 中选一个长为 $2k$ 的子序列，计算其前一半的 OR，后一半的 OR，这两个 OR 再计算 XOR。

问：计算出的 XOR 最大能是多少？

## 核心思路

- 想象有一根分割线，把 $\textit{nums}$ 分成左右两部分，左和右分别计算所有长为 $k$ 的子序列的 OR 都**有哪些值**。比如左边计算出的 OR 有 $2,3,5$，右边计算出的 OR 有 $1,3,6$，那么两两组合计算 XOR，其中最大值即为答案。
- 枚举分割线的位置，把 $\textit{nums}$ 分割成一个前缀和一个后缀，问题变成：从前缀/后缀中选一个长为 $k$ 的子序列，这个子序列 OR 的结果能否等于 $x$？

把 OR 理解成一个类似加法的东西，转换成二维 0-1 背包。如果你不了解 0-1 背包，或者不理解为什么下面代码 $j$ 要倒序枚举，请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

> **二维**：指背包有两个约束，一个是所选元素的个数是 $k$，另一个是所选元素的 OR 是 $x$。

## 具体算法

计算后缀。对于 0-1 背包问题，我们定义 $f[i][j][x]$ 表示从 $\textit{nums}[i]$ 到 $\textit{nums}[n-1]$ 中选 $j$ 个数，这些数的 OR 能否等于 $x$。

设 $v=\textit{nums}[i]$，用刷表法转移：

- 不选 $v$，那么 $f[i][j][x] = f[i+1][j][x]$。
- 选 $v$，如果 $f[i+1][j][x]=\texttt{true}$，那么 $f[i][j+1][x\ |\ v]=\texttt{true}$。

> **刷表法**：本题计算 $x = v\ |\ ?$ 中的 $?$ 是困难的，但计算 $x\ |\ v$ 是很简单的。也就是说，对于状态 $f[i][j][x]$ 而言，其转移来源是谁不好计算，但从 $f[i][j][x]$ 转移到的目标状态 $f[i][j+1][x\ |\ v]$ 是好计算的。在动态规划中，根据转移来源计算状态叫查表法，用当前状态更新其他状态叫刷表法。

初始值 $f[n][0][0]=\texttt{true}$。什么也不选，OR 等于 $0$。

对于每个 $i$，由于我们只需要 $f[i][k]$ 中的数据，把 $f[i][k]$ 复制到 $\textit{suf}[i]$ 中。这样做无需创建三维数组，空间复杂度更小。

代码实现时，$f$ 的第一个维度可以优化掉。

对于前缀 $\textit{pre}$ 的计算也同理。

最后，枚举 $i=k-1,k,k+1,\ldots,n-k-1$，两两组合 $\textit{pre}[i]$ 和 $\textit{suf}[i+1]$ 中的数计算 XOR，其中最大值即为答案。

小优化：如果在循环中，发现答案 $\textit{ans}$ 达到了理论最大值 $2^7-1$（或者所有元素的 OR），则立刻返回答案。

> 也可以用哈希集合代替布尔数组，见下面的 Python 优化代码。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Ub4mekE1x/) 第三题，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def maxValue(self, nums: List[int], k: int) -> int:
        mx = reduce(or_, nums)
        n = len(nums)
        suf = [None] * (n - k + 1)
        f = [[False] * (mx + 1) for _ in range(k + 1)]
        f[0][0] = True
        for i in range(n - 1, k - 1, -1):
            v = nums[i]
            # 注意当 i 比较大的时候，循环次数应和 i 有关，因为更大的 j，对应的 f[j] 全为 False
            for j in range(min(k - 1, n - 1 - i), -1, -1):
                for x, has_x in enumerate(f[j]):
                    if has_x:
                        f[j + 1][x | v] = True
            if i <= n - k:
                suf[i] = f[k].copy()

        ans = 0
        f = [[False] * (mx + 1) for _ in range(k + 1)]
        f[0][0] = True
        for i, v in enumerate(nums[:-k]):
            for j in range(min(k - 1, i), -1, -1):
                for x, has_x in enumerate(f[j]):
                    if has_x:
                        f[j + 1][x | v] = True
            if i < k - 1:
                continue
            # 这里 f[k] 就是 pre[i]
            for x, has_x in enumerate(f[k]):
                if has_x:
                    for y, has_y in enumerate(suf[i + 1]):
                        if has_y and x ^ y > ans:  # 手写 if
                            ans = x ^ y
            if ans == mx:
                return ans
        return ans
```

```py [sol-Python3 优化]
# 使用 set 代替 bool list
class Solution:
    def maxValue(self, nums: List[int], k: int) -> int:
        n = len(nums)
        suf = [None] * (n - k + 1)
        f = [set() for _ in range(k + 1)]
        f[0].add(0)
        for i in range(n - 1, k - 1, -1):
            v = nums[i]
            for j in range(min(k - 1, n - 1 - i), -1, -1):
                f[j + 1].update(x | v for x in f[j])
            if i <= n - k:
                suf[i] = f[k].copy()

        mx = reduce(or_, nums)
        ans = 0
        f = [set() for _ in range(k + 1)]
        f[0].add(0)
        for i, v in enumerate(nums[:-k]):
            for j in range(min(k - 1, i), -1, -1):
                f[j + 1].update(x | v for x in f[j])
            if i < k - 1:
                continue
            # 这里 f[k] 就是 pre[i]
            ans = max(ans, max(x ^ y for x in f[k] for y in suf[i + 1]))
            if ans == mx:
                return ans
        return ans
```

```java [sol-Java]
class Solution {
    public int maxValue(int[] nums, int k) {
        final int MX = 1 << 7;
        int n = nums.length;
        boolean[][] suf = new boolean[n - k + 1][];
        boolean[][] f = new boolean[k + 1][MX];
        f[0][0] = true;
        for (int i = n - 1; i >= k; i--) {
            int v = nums[i];
            // 注意当 i 比较大的时候，循环次数应和 i 有关，因为更大的 j，对应的 f[j] 全为 false
            for (int j = Math.min(k - 1, n - 1 - i); j >= 0; j--) {
                for (int x = 0; x < MX; x++) {
                    if (f[j][x]) {
                        f[j + 1][x | v] = true;
                    }
                }
            }
            if (i <= n - k) {
                suf[i] = f[k].clone();
            }
        }

        int ans = 0;
        f = new boolean[k + 1][MX];
        f[0][0] = true;
        for (int i = 0; i < n - k; i++) {
            int v = nums[i];
            for (int j = Math.min(k - 1, i); j >= 0; j--) {
                for (int x = 0; x < MX; x++) {
                    if (f[j][x]) {
                        f[j + 1][x | v] = true;
                    }
                }
            }
            if (i < k - 1) {
                continue;
            }
            // 这里 f[k] 就是 pre[i]
            for (int x = 0; x < MX; x++) {
                if (f[k][x]) {
                    for (int y = 0; y < MX; y++) {
                        if (suf[i + 1][y]) {
                            ans = Math.max(ans, x ^ y);
                        }
                    }
                }
            }
            if (ans == MX - 1) {
                return ans;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxValue(vector<int>& nums, int k) {
        const int MX = 1 << 7;
        int n = nums.size();
        vector<array<int, MX>> suf(n - k + 1);
        vector<array<int, MX>> f(k + 1);
        f[0][0] = true;
        for (int i = n - 1; i >= k; i--) {
            int v = nums[i];
            // 注意当 i 比较大的时候，循环次数应和 i 有关，因为更大的 j，对应的 f[j] 全为 false
            for (int j = min(k - 1, n - 1 - i); j >= 0; j--) {
                for (int x = 0; x < MX; x++) {
                    if (f[j][x]) {
                        f[j + 1][x | v] = true;
                    }
                }
            }
            if (i <= n - k) {
                suf[i] = f[k];
            }
        }

        int ans = 0;
        f = vector<array<int, MX>>(k + 1);
        f[0][0] = true;
        for (int i = 0; i < n - k; i++) {
            int v = nums[i];
            for (int j = min(k - 1, i); j >= 0; j--) {
                for (int x = 0; x < MX; x++) {
                    if (f[j][x]) {
                        f[j + 1][x | v] = true;
                    }
                }
            }
            if (i < k - 1) {
                continue;
            }
            // 这里 f[k] 就是 pre[i]
            for (int x = 0; x < MX; x++) {
                if (f[k][x]) {
                    for (int y = 0; y < MX; y++) {
                        if (suf[i + 1][y]) {
                            ans = max(ans, x ^ y);
                        }
                    }
                }
            }
            if (ans == MX - 1) {
                return ans;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxValue(nums []int, k int) (ans int) {
	const mx = 1 << 7
	n := len(nums)
	suf := make([][mx]bool, n-k+1)
	f := make([][mx]bool, k+1)
	f[0][0] = true
	for i := n - 1; i >= k; i-- {
		v := nums[i]
		// 注意当 i 比较大的时候，循环次数应和 i 有关，因为更大的 j，对应的 f[j] 全为 false
		for j := min(k-1, n-1-i); j >= 0; j-- {
			for x, hasX := range f[j] {
				if hasX {
					f[j+1][x|v] = true
				}
			}
		}
		if i <= n-k {
			suf[i] = f[k]
		}
	}

	clear(f)
	f[0][0] = true
	for i, v := range nums[:n-k] {
		for j := min(k-1, i); j >= 0; j-- {
			for x, hasX := range f[j] {
				if hasX {
					f[j+1][x|v] = true
				}
			}
		}
		if i < k-1 {
			continue
		}
		// 这里 f[k] 就是 pre[i]
		for x, hasX := range f[k] {
			if hasX {
				for y, hasY := range suf[i+1] {
					if hasY {
						ans = max(ans, x^y)
					}
				}
			}
		}
		if ans == mx-1 {
			return
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nkU + nU^2)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U$ 是 $\textit{nums}$ 所有元素的 OR，本题至多为 $2^7-1$。DP 是 $\mathcal{O}(nkU)$ 的，计算 XOR 最大值是 $\mathcal{O}(nU^2)$ 的。
- 空间复杂度：$\mathcal{O}(nU)$。

## 优化

例如 $x=1101_{(2)}$，我们**至多**要选几个 $\textit{nums}[i]$，就能 OR 得到 $x$？（前提是可以得到 $x$）

答案是 $3$ 个。考虑 $x$ 中的每个比特 $1$，它来自某个 $\textit{nums}[i]$。

设 $\textit{nums}$ 所有元素 OR 的二进制中的 $1$ 的个数为 $\textit{ones}$（本题数据范围保证 $\textit{ones}\le 7$）。一般地，我们至多选 $\textit{ones}$ 个 $\textit{nums}[i]$，就能 OR 得到 $x$。

但是，本题要求（前缀/后缀）**恰好**选 $k$ 个元素。选的元素越多 OR 越大，那么某些比较小的 $x$ 可能无法 OR 出来。

为了判断（前缀/后缀）恰好选 $k$ 个元素能否 OR 出整数 $x$，定义：

- $\textit{minI}[x]$，表示从 $0$ 开始遍历，至少要遍历到 $i$ 才有可能找到 $k$ 个数 OR 等于 $x$。如果无法得到 $x$ 那么 $\textit{minI}[x] = \infty$。
- $\textit{maxI}[x]$，表示从 $n-1$ 开始遍历，至少要遍历到 $i$ 才有可能找到 $k$ 个数 OR 等于 $x$。如果无法得到 $x$ 那么 $\textit{maxI}[x] = 0$。

根据 [从集合论到位运算](https://leetcode.cn/circle/discuss/CaOJ45/)，如果能 OR 得到 $x$，那么参与 OR 运算的元素都是 $x$ 的子集。换句话说，$x$ 是参与 OR 运算的元素的**超集**（superset）。

对于 $\textit{minI}[x]$ 的计算，我们可以在遍历 $\textit{nums}$ 的同时，用一个数组 $\textit{cnt}$ 维护 $\textit{nums}$ 元素的超集的出现次数。如果发现 $\textit{cnt}[x]=k$，说明至少要遍历到 $i$ 才有可能找到 $k$ 个数 OR 等于 $x$，记录 $\textit{minI}[x]=i$。对于 $\textit{maxI}[x]$ 的计算也同理。

对于两数异或最大值的计算，可以用**试填法**解决，原理请看[【图解】421. 数组中两个数的最大异或值](https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/solutions/2511644/tu-jie-jian-ji-gao-xiao-yi-tu-miao-dong-1427d/)。

```py [sol-Python3]
class Solution:
    def maxValue(self, nums: List[int], k: int) -> int:
        n = len(nums)
        mx = reduce(or_, nums)
        k2 = min(k, mx.bit_count())  # 至多选 k2 个数

        suf = [None] * (n - k + 1)
        f = [set() for _ in range(k2 + 1)]
        f[0].add(0)
        max_i = [0] * (mx + 1)
        cnt = [0] * (mx + 1)
        for i in range(n - 1, k - 1, -1):
            v = nums[i]
            for j in range(min(k2 - 1, n - 1 - i), -1, -1):
                f[j + 1].update(x | v for x in f[j])
            if i <= n - k:
                suf[i] = f[k2].copy()
            # 枚举 v 的超集
            s = v
            while s <= mx:
                cnt[s] += 1
                if cnt[s] == k:
                    # 从 n-1 开始遍历，至少要遍历到 i 才有可能找到 k 个数 OR 等于 s
                    max_i[s] = i
                s = (s + 1) | v

        ans = 0
        pre = [set() for _ in range(k2 + 1)]
        pre[0].add(0)
        min_i = [inf] * (mx + 1)
        cnt = [0] * (mx + 1)
        w = mx.bit_length()  # 用于 findMaximumXOR
        for i, v in enumerate(nums[:-k]):
            for j in range(min(k2 - 1, i), -1, -1):
                pre[j + 1].update(x | v for x in pre[j])
            # 枚举 v 的超集
            s = v
            while s <= mx:
                cnt[s] += 1
                if cnt[s] == k:
                    # 从 0 开始遍历，至少要遍历到 i 才有可能找到 k 个数 OR 等于 s
                    min_i[s] = i
                s = (s + 1) | v
            if i < k - 1:
                continue
            a = [x for x in pre[k2] if min_i[x] <= i]
            b = [x for x in suf[i + 1] if max_i[x] > i]
            ans = max(ans, self.findMaximumXOR(a, b, w))
            if ans == mx:
                return ans
        return ans

    # 421. 数组中两个数的最大异或值
    # 改成两个数组的最大异或值，做法是类似的，仍然可以用【试填法】解决
    def findMaximumXOR(self, a: List[int], b: List[int], w: int) -> int:
        ans = mask = 0
        for i in range(w - 1, -1, -1):  # 从最高位开始枚举
            mask |= 1 << i
            new_ans = ans | (1 << i)  # 这个比特位可以是 1 吗？
            set_a = set(x & mask for x in a)  # 低于 i 的比特位置为 0
            for x in b:
                x &= mask  # 低于 i 的比特位置为 0
                if new_ans ^ x in set_a:
                    ans = new_ans  # 这个比特位可以是 1
                    break
        return ans
```

```java [sol-Java]
class Solution {
    private static final int BIT_WIDTH = 7;

    public int maxValue(int[] nums, int k) {
        final int MX = 1 << BIT_WIDTH;
        int n = nums.length;
        int k2 = Math.min(k, BIT_WIDTH); // 至多选 k2 个数

        boolean[][] suf = new boolean[n - k + 1][];
        boolean[][] f = new boolean[k2 + 1][MX];
        f[0][0] = true;
        int[] maxI = new int[MX];
        int[] cnt = new int[MX];
        for (int i = n - 1; i >= k; i--) {
            int v = nums[i];
            for (int j = Math.min(k2 - 1, n - 1 - i); j >= 0; j--) {
                for (int x = 0; x < MX; x++) {
                    if (f[j][x]) {
                        f[j + 1][x | v] = true;
                    }
                }
            }
            if (i <= n - k) {
                suf[i] = f[k2].clone();
            }
            // 枚举 v 的超集
            for (int s = v; s < MX; s = (s + 1) | v) {
                if (++cnt[s] == k) {
                    // 从 n-1 开始遍历，至少要遍历到 i 才有可能找到 k 个数 OR 等于 s
                    maxI[s] = i;
                }
            }
        }

        int ans = 0;
        boolean[][] pre = new boolean[k2 + 1][MX];
        pre[0][0] = true;
        int[] minI = new int[MX];
        Arrays.fill(minI, Integer.MAX_VALUE);
        Arrays.fill(cnt, 0);
        int[] a = new int[MX];
        int[] b = new int[MX];
        for (int i = 0; i < n - k; i++) {
            int v = nums[i];
            for (int j = Math.min(k2 - 1, i); j >= 0; j--) {
                for (int x = 0; x < MX; x++) {
                    if (pre[j][x]) {
                        pre[j + 1][x | v] = true;
                    }
                }
            }
            // 枚举 v 的超集
            for (int s = v; s < MX; s = (s + 1) | v) {
                if (++cnt[s] == k) {
                    // 从 0 开始遍历，至少要遍历到 i 才有可能找到 k 个数 OR 等于 s
                    minI[s] = i;
                }
            }
            if (i < k - 1) {
                continue;
            }
            int na = 0;
            int nb = 0;
            for (int x = 0; x < MX; x++) {
                if (pre[k2][x] && minI[x] <= i) {
                    a[na++] = x;
                }
                if (suf[i + 1][x] && maxI[x] > i) {
                    b[nb++] = x;
                }
            }
            ans = Math.max(ans, findMaximumXOR(a, na, b, nb));
            if (ans == MX - 1) {
                return ans;
            }
        }
        return ans;
    }

    // 421. 数组中两个数的最大异或值
    // 改成两个数组的最大异或值，做法是类似的，仍然可以用【试填法】解决
    private int findMaximumXOR(int[] a, int n, int[] b, int m) {
        int ans = 0;
        int mask = 0;
        boolean[] seen = new boolean[1 << BIT_WIDTH];
        for (int i = BIT_WIDTH - 1; i >= 0; i--) { // 从最高位开始枚举
            mask |= 1 << i;
            int newAns = ans | (1 << i); // 这个比特位可以是 1 吗？
            Arrays.fill(seen, false);
            for (int j = 0; j < n; j++) {
                seen[a[j] & mask] = true; // 低于 i 的比特位置为 0
            }
            for (int j = 0; j < m; j++) {
                int x = b[j] & mask; // 低于 i 的比特位置为 0
                if (seen[newAns ^ x]) {
                    ans = newAns; // 这个比特位可以是 1
                    break;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int BIT_WIDTH = 7;

    // 421. 数组中两个数的最大异或值
    // 改成两个数组的最大异或值，做法是类似的，仍然可以用【试填法】解决
    int findMaximumXOR(vector<int>& a, vector<int>& b) {
        int ans = 0, mask = 0;
        vector<int> seen(1 << BIT_WIDTH);
        for (int i = BIT_WIDTH - 1; i >= 0; i--) { // 从最高位开始枚举
            mask |= 1 << i;
            int new_ans = ans | (1 << i); // 这个比特位可以是 1 吗？
            ranges::fill(seen, false);
            for (int x : a) {
                seen[x & mask] = true; // 低于 i 的比特位置为 0
            }
            for (int x : b) {
                x &= mask; // 低于 i 的比特位置为 0
                if (seen[new_ans ^ x]) {
                    ans = new_ans; // 这个比特位可以是 1
                    break;
                }
            }
        }
        return ans;
    }

public:
    int maxValue(vector<int>& nums, int k) {
        const int MX = 1 << BIT_WIDTH;
        int n = nums.size();
        int k2 = min(k, BIT_WIDTH); // 至多选 k2 个数

        vector<array<int, MX>> suf(n - k + 1);
        vector<array<int, MX>> f(k2 + 1);
        f[0][0] = true;
        int max_i[MX]{}, cnt[MX]{};
        for (int i = n - 1; i >= k; i--) {
            int v = nums[i];
            // 注意当 i 比较大的时候，循环次数应和 i 有关，因为更大的 j，对应的 f[j] 全为 false
            for (int j = min(k2 - 1, n - 1 - i); j >= 0; j--) {
                for (int x = 0; x < MX; x++) {
                    if (f[j][x]) {
                        f[j + 1][x | v] = true;
                    }
                }
            }
            if (i <= n - k) {
                suf[i] = f[k2];
            }
            // 枚举 v 的超集
            for (int s = v; s < MX; s = (s + 1) | v) {
                if (++cnt[s] == k) {
                    // 从 n-1 开始遍历，至少要遍历到 i 才有可能找到 k 个数 OR 等于 s
                    max_i[s] = i;
                }
            }
        }

        int ans = 0;
        vector<array<int, MX>> pre(k2 + 1);
        pre[0][0] = true;
        int min_i[MX];
        ranges::fill(min_i, INT_MAX);
        ranges::fill(cnt, 0);
        for (int i = 0; i < n - k; i++) {
            int v = nums[i];
            for (int j = min(k2 - 1, i); j >= 0; j--) {
                for (int x = 0; x < MX; x++) {
                    if (pre[j][x]) {
                        pre[j + 1][x | v] = true;
                    }
                }
            }
            // 枚举 v 的超集
            for (int s = v; s < MX; s = (s + 1) | v) {
                if (++cnt[s] == k) {
                    // 从 0 开始遍历，至少要遍历到 i 才有可能找到 k 个数 OR 等于 s
                    min_i[s] = i;
                }
            }
            if (i < k - 1) {
                continue;
            }
            vector<int> a, b;
            for (int x = 0; x < MX; x++) {
                if (pre[k2][x] && min_i[x] <= i) {
                    a.push_back(x);
                }
                if (suf[i + 1][x] && max_i[x] > i) {
                    b.push_back(x);
                }
            }
            ans = max(ans, findMaximumXOR(a, b));
            if (ans == MX - 1) {
                return ans;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
const bitWidth = 7
const mx = 1 << bitWidth

func maxValue(nums []int, k int) (ans int) {
	n := len(nums)
	k2 := min(k, bitWidth) // 至多选 k2 个数
	suf := make([][mx]bool, n-k+1)
	f := make([][mx]bool, k2+1)
	f[0][0] = true
	maxI := [mx]int{}
	cnt := [mx]int{}
	for i := n - 1; i >= k; i-- {
		v := nums[i]
		for j := min(k2-1, n-1-i); j >= 0; j-- {
			for x, hasX := range f[j] {
				if hasX {
					f[j+1][x|v] = true
				}
			}
		}
		if i <= n-k {
			suf[i] = f[k2]
		}
		// 枚举 v 的超集
		for s := v; s < mx; s = (s + 1) | v {
			cnt[s]++
			if cnt[s] == k {
				// 从 n-1 开始遍历，至少要遍历到 i 才有可能找到 k 个数 OR 等于 s
				maxI[s] = i
			}
		}
	}

	pre := make([][mx]bool, k2+1)
	pre[0][0] = true
	minI := [mx]int{}
	for i := range minI {
		minI[i] = math.MaxInt
	}
	cnt = [mx]int{}
	for i, v := range nums[:n-k] {
		for j := min(k2-1, i); j >= 0; j-- {
			for x, hasX := range pre[j] {
				if hasX {
					pre[j+1][x|v] = true
				}
			}
		}
		// 枚举 v 的超集
		for s := v; s < mx; s = (s + 1) | v {
			cnt[s]++
			if cnt[s] == k {
				// 从 0 开始遍历，至少要遍历到 i 才有可能找到 k 个数 OR 等于 s
				minI[s] = i
			}
		}
		if i < k-1 {
			continue
		}
		a := []int{}
		b := []int{}
		for x, has := range pre[k2] {
			if has && minI[x] <= i {
				a = append(a, x)
			}
			if suf[i+1][x] && maxI[x] > i {
				b = append(b, x)
			}
		}
		ans = max(ans, findMaximumXOR(a, b))
		if ans == mx-1 {
			return
		}
	}
	return
}

// 421. 数组中两个数的最大异或值
// 改成两个数组的最大异或值，做法是类似的，仍然可以用【试填法】解决
func findMaximumXOR(a, b []int) (ans int) {
	mask := 0
	for i := bitWidth - 1; i >= 0; i-- { // 从最高位开始枚举
		mask |= 1 << i
		newAns := ans | 1<<i // 这个比特位可以是 1 吗？
		seen := [mx]bool{}
		for _, x := range a {
			seen[x&mask] = true // 低于 i 的比特位置为 0
		}
		for _, x := range b {
			x &= mask // 低于 i 的比特位置为 0
			if seen[newAns^x] {
				ans = newAns
				break
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U$ 是 $\textit{nums}$ 所有元素的 OR，本题至多为 $2^7-1$。
- 空间复杂度：$\mathcal{O}(nU)$。

更多相似题目，见下面动态规划题单中的「**§3.1 0-1 背包**」和「**专题：前后缀分解**」。

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

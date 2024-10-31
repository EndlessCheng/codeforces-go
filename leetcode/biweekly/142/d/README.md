## 总体思路

把一开始可能想要输入字符串叫做初始字符串。注意这里定义的初始字符串**长度没有限制**。

1. 计算不考虑 $k$ 的情况下，有多少个初始字符串。
2. 计算长度小于 $k$ 的初始字符串个数。
3. 二者相减，即为长度大于等于 $k$ 的初始字符串个数。

## 不考虑 k 的初始字符串个数

示例 1 的字符串，可以分为 $4$ 组（每组内的字母都相同）：$\texttt{aa},\texttt{bb},\texttt{cc},\texttt{dd}$，长度分别为 $2,2,2,2$。

在初始字符串中，每组的长度可以从 $1$ 到 $2$ 不等，根据乘法原理，个数为

$$
2\times 2\times 2\times 2 = 16
$$

## 长度小于 k 的初始字符串个数

### 寻找子问题

假设字符串分为 $4$ 组，当前要用这 $4$ 组构造的初始字符串的长度是 $6$。

由于每组的长度至少是 $1$，为方便计算，先从每组各选 $1$ 个字母，问题转换成从 $4$ 组中额外再选 $6-4=2$ 个字母的方案数。

枚举从最后一组中选多少个字母：

- 选 $0$ 个，问题变成从前 $3$ 组中选 $2-0=2$ 个字母的方案数。
- 选 $1$ 个，问题变成从前 $3$ 组中选 $2-1=1$ 个字母的方案数。

### 状态定义与状态转移方程

这是一个**多重背包**问题。相当于有 $m=4$ 种物品，第 $i$ 种物品有「第 $i$ 组的大小减一」个，计算至多选 $k-m-1$ 个物品的方案数。

根据上面的讨论，定义 $f[i+1][j]$ 表示从前 $i$ 种物品中选**恰好** $j$ 个物品的方案数。

初始值 $f[0][0]=1$，不选算一种方案（在示例 1 中，这对应字符串 $\texttt{abcd}$）。

假设第 $i$ 种物品有 $c$ 个，枚举选 $L=0,1,2,\cdots,c$ 个物品，问题变成从前 $i-1$ 种物品中选恰好 $j-L$ 个物品的方案数，即 $f[i][j-L]$。

累加得

$$
f[i+1][j] = \sum_{L=0}^{c} f[i][j-L]
$$

注意要保证 $j-L\ge 0$。上式等价于

$$
f[i+1][j] = \sum_{p=\max(j-c, 0)}^{j} f[i][p]
$$

### 前缀和优化

定义 $f[i]$ 的 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 数组为 $s$，那么上式等价于

$$
f[i+1][j] = s[j+1] - s[\max(j-c, 0)]
$$

设一共有 $m$ 组，那么至多选 $k-m-1$ 个物品的方案总数为 $\sum\limits_{j=0}^{k-m-1}f[m][j]$。

特别地，如果 $n<k$（$n$ 为 $\textit{word}$ 的长度），那么无法满足要求，直接返回 $0$。

特别地，如果 $m\ge k$，那么长度小于 $k$ 的初始字符串个数为 $0$，直接返回各组长度的乘积。

代码中用到了一些取模的细节，原理见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

具体请看 [视频讲解](https://www.bilibili.com/video/BV13J1MYwEGM/?t=21m4s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def possibleStringCount(self, word: str, k: int) -> int:
        n = len(word)
        if n < k:  # 无法满足要求
            return 0

        MOD = 1_000_000_007
        cnts = []
        ans = 1
        cnt = 0
        for i in range(n):
            cnt += 1
            if i == n - 1 or word[i] != word[i + 1]:
                # 如果 cnt = 1，这组字符串必选，无需参与计算
                if cnt > 1:
                    if k > 0:
                        cnts.append(cnt - 1)
                    ans = ans * cnt % MOD
                k -= 1  # 注意这里把 k 减小了
                cnt = 0

        if k <= 0:
            return ans

        f = [[0] * k for _ in range(len(cnts) + 1)]
        f[0][0] = 1
        for i, c in enumerate(cnts):
            # 计算 f[i] 的前缀和数组 s
            s = list(accumulate(f[i], initial=0))
            # 计算子数组和
            for j in range(k):
                f[i + 1][j] = (s[j + 1] - s[max(j - c, 0)]) % MOD
        return (ans - sum(f[-1])) % MOD
```

```java [sol-Java]
class Solution {
    public int possibleStringCount(String word, int k) {
        int n = word.length();
        if (n < k) { // 无法满足要求
            return 0;
        }

        final int MOD = 1_000_000_007;
        List<Integer> cnts = new ArrayList<>();
        long ans = 1;
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            cnt++;
            if (i == n - 1 || word.charAt(i) != word.charAt(i + 1)) {
                // 如果 cnt = 1，这组字符串必选，无需参与计算
                if (cnt > 1) {
                    if (k > 0) {
                        cnts.add(cnt - 1);
                    }
                    ans = ans * cnt % MOD;
                }
                k--; // 注意这里把 k 减小了
                cnt = 0;
            }
        }

        if (k <= 0) {
            return (int) ans;
        }

        int m = cnts.size();
        int[][] f = new int[m + 1][k];
        f[0][0] = 1;
        int[] s = new int[k + 1];
        for (int i = 0; i < m; i++) {
            // 计算 f[i] 的前缀和数组 s
            for (int j = 0; j < k; j++) {
                s[j + 1] = (s[j] + f[i][j]) % MOD;
            }
            int c = cnts.get(i);
            // 计算子数组和
            for (int j = 0; j < k; j++) {
                f[i + 1][j] = (s[j + 1] - s[Math.max(j - c, 0)]) % MOD;
            }
        }

        for (int x : f[m]) {
            ans -= x;
        }
        return (int) ((ans % MOD + MOD) % MOD); // 保证结果非负
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int possibleStringCount(string word, int k) {
        int n = word.length();
        if (n < k) { // 无法满足要求
            return 0;
        }

        const int MOD = 1'000'000'007;
        vector<int> cnts;
        long long ans = 1;
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            cnt++;
            if (i == n - 1 || word[i] != word[i + 1]) {
                // 如果 cnt = 1，这组字符串必选，无需参与计算
                if (cnt > 1) {
                    if (k > 0) {
                        cnts.push_back(cnt - 1);
                    }
                    ans = ans * cnt % MOD;
                }
                k--; // 注意这里把 k 减小了
                cnt = 0;
            }
        }

        if (k <= 0) {
            return ans;
        }

        int m = cnts.size();
        vector<vector<int>> f(m + 1, vector<int>(k));
        f[0][0] = 1;
        vector<int> s(k + 1);
        for (int i = 0; i < m; i++) {
            // 计算 f[i] 的前缀和数组 s
            for (int j = 0; j < k; j++) {
                s[j + 1] = (s[j] + f[i][j]) % MOD;
            }
            // 计算子数组和
            for (int j = 0; j < k; j++) {
                f[i + 1][j] = (s[j + 1] - s[max(j - cnts[i], 0)]) % MOD;
            }
        }

        ans -= reduce(f[m].begin(), f[m].end(), 0LL);
        return (ans % MOD + MOD) % MOD; // 保证结果非负
    }
};
```

```go [sol-Go]
func possibleStringCount(word string, k int) int {
	if len(word) < k { // 无法满足要求
		return 0
	}

	const mod = 1_000_000_007
	cnts := []int{}
	ans := 1
	cnt := 0
	for i := range word {
		cnt++
		if i == len(word)-1 || word[i] != word[i+1] {
			// 如果 cnt = 1，这组字符串必选，无需参与计算
			if cnt > 1 {
				if k > 0 {
					cnts = append(cnts, cnt-1)
				}
				ans = ans * cnt % mod
			}
			k-- // 注意这里把 k 减小了
			cnt = 0
		}
	}

	if k <= 0 {
		return ans
	}

	m := len(cnts)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, k)
	}
	f[0][0] = 1
	s := make([]int, k+1)
	for i, c := range cnts {
		// 计算 f[i] 的前缀和数组 s
		for j, v := range f[i] {
			s[j+1] = (s[j] + v) % mod
		}
		// 计算子数组和
		for j := range f[i+1] {
			f[i+1][j] = s[j+1] - s[max(j-c, 0)]
		}
	}

	for _, v := range f[m] {
		ans -= v
	}
	return (ans%mod + mod) % mod // 保证结果非负
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + k^2)$，其中 $n$ 是 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(k^2)$。

### 空间优化

去掉 $f$ 的第一个维度。

前缀和直接计算到 $f$ 数组中。

然后和 0-1 背包一样，倒序计算 $f[j] = s[j] - s[j-c-1]$。减一是因为原来前缀和中的 $s[0]=0$ 去掉了，$s$ 的长度不是 $k+1$ 而是 $k$。

```py [sol-Python3]
class Solution:
    def possibleStringCount(self, word: str, k: int) -> int:
        n = len(word)
        if n < k:  # 无法满足要求
            return 0

        MOD = 1_000_000_007
        cnts = []
        ans = 1
        cnt = 0
        for i in range(n):
            cnt += 1
            if i == n - 1 or word[i] != word[i + 1]:
                # 如果 cnt = 1，这组字符串必选，无需参与计算
                if cnt > 1:
                    if k > 0:  # 保证空间复杂度为 O(k)
                        cnts.append(cnt - 1)
                    ans = ans * cnt % MOD
                k -= 1  # 注意这里把 k 减小了
                cnt = 0

        if k <= 0:
            return ans

        f = [0] * k
        f[0] = 1
        for c in cnts:
            # 原地计算 f 的前缀和
            for j in range(1, k):
                f[j] = (f[j] + f[j - 1]) % MOD
            # 计算子数组和
            for j in range(k - 1, c, -1):
                f[j] -= f[j - c - 1]
        return (ans - sum(f)) % MOD
```

```java [sol-Java]
class Solution {
    public int possibleStringCount(String word, int k) {
        int n = word.length();
        if (n < k) { // 无法满足要求
            return 0;
        }

        final int MOD = 1_000_000_007;
        List<Integer> cnts = new ArrayList<>();
        long ans = 1;
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            cnt++;
            if (i == n - 1 || word.charAt(i) != word.charAt(i + 1)) {
                // 如果 cnt = 1，这组字符串必选，无需参与计算
                if (cnt > 1) {
                    if (k > 0) { // 保证空间复杂度为 O(k)
                        cnts.add(cnt - 1);
                    }
                    ans = ans * cnt % MOD;
                }
                k--; // 注意这里把 k 减小了
                cnt = 0;
            }
        }

        if (k <= 0) {
            return (int) ans;
        }

        int[] f = new int[k];
        f[0] = 1;
        for (int c : cnts) {
            // 原地计算 f 的前缀和
            for (int j = 1; j < k; j++) {
                f[j] = (f[j] + f[j - 1]) % MOD;
            }
            // 计算子数组和
            for (int j = k - 1; j > c; j--) {
                f[j] = (f[j] - f[j - c - 1]) % MOD;
            }
        }

        for (int x : f) {
            ans -= x;
        }
        return (int) ((ans % MOD + MOD) % MOD); // 保证结果非负
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int possibleStringCount(string word, int k) {
        int n = word.length();
        if (n < k) { // 无法满足要求
            return 0;
        }

        const int MOD = 1'000'000'007;
        vector<int> cnts;
        long long ans = 1;
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            cnt++;
            if (i == n - 1 || word[i] != word[i + 1]) {
                // 如果 cnt = 1，这组字符串必选，无需参与计算
                if (cnt > 1) {
                    if (k > 0) { // 保证空间复杂度为 O(k)
                        cnts.push_back(cnt - 1);
                    }
                    ans = ans * cnt % MOD;
                }
                k--; // 注意这里把 k 减小了
                cnt = 0;
            }
        }

        if (k <= 0) {
            return ans;
        }

        vector<int> f(k);
        f[0] = 1;
        for (int c : cnts) {
            // 原地计算 f 的前缀和
            for (int j = 1; j < k; j++) {
                f[j] = (f[j] + f[j - 1]) % MOD;
            }
            // 计算子数组和
            for (int j = k - 1; j > c; j--) {
                f[j] = (f[j] - f[j - c - 1]) % MOD;
            }
        }

        ans -= reduce(f.begin(), f.end(), 0LL);
        return (ans % MOD + MOD) % MOD; // 保证结果非负
    }
};
```

```go [sol-Go]
func possibleStringCount(word string, k int) int {
	if len(word) < k { // 无法满足要求
		return 0
	}

	const mod = 1_000_000_007
	cnts := []int{}
	ans := 1
	cnt := 0
	for i := range word {
		cnt++
		if i == len(word)-1 || word[i] != word[i+1] {
			// 如果 cnt = 1，这组字符串必选，无需参与计算
			if cnt > 1 {
				if k > 0 { // 保证空间复杂度为 O(k)
					cnts = append(cnts, cnt-1)
				}
				ans = ans * cnt % mod
			}
			k-- // 注意这里把 k 减小了
			cnt = 0
		}
	}

	if k <= 0 {
		return ans
	}

	f := make([]int, k)
	f[0] = 1
	for _, c := range cnts {
		// 原地计算 f 的前缀和
		for j := 1; j < k; j++ {
			f[j] = (f[j] + f[j-1]) % mod
		}
		// 计算子数组和
		for j := k - 1; j > c; j-- {
			f[j] -= f[j-c-1]
		}
	}

	for _, x := range f {
		ans -= x
	}
	return (ans%mod + mod) % mod // 保证结果非负
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + k^2)$，其中 $n$ 是 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。

## 相似题目

- [2902. 和带限制的子多重集合的数目](https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum/)

更多相似题目，见下面动态规划题单中的「**§11.1 前缀和优化 DP**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

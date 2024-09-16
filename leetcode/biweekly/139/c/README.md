想象有一根分割线，把 $\textit{nums}$ 分成左右两部分，左部和右部分别计算所有长为 $k$ 的子序列的 OR 都**有哪些值**。

比如左部计算出的 OR 有 $[2,3,5]$，右部计算出的 OR 有 $[1,3,6]$，两两组合计算 XOR，取其中最大值作为答案。

把 OR 理解成一个类似加法的东西，本题是一个二维的 0-1 背包。如果你不了解 0-1 背包，或者不理解为什么下面代码 $j$ 要倒序枚举，请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

考虑计算右部，定义 $f[i][j][x]$ 表示从 $\textit{nums}[i]$ 到 $\textit{nums}[n-1]$ 中选出 $j$ 个数，这些数的 OR 是否等于 $x$。

设 $v=\textit{nums}[i]$，用**刷表法**转移：

- 不选 $v$，那么 $f[i][j][x] = f[i+1][j][x]$。
- 选 $v$，如果 $f[i+1][j][x]=\texttt{true}$，那么 $f[i][j+1][x|v]=\texttt{true}$。

初始值 $f[n][0][0]=\texttt{true}$。什么也不选，OR 等于 $0$。

对于每个 $i$，由于我们只需要 $f[i][k]$ 中的数据，把 $f[i][k]$ 复制到 $\textit{suf}[i]$ 中。这样做无需创建三维数组，空间复杂度更小。

代码实现时，$f$ 的第一个维度可以优化掉。

对于左部 $\textit{pre}$ 的计算也同理。

最后，枚举 $i=k-1,k,k+1,\cdots,n-k-1$，两两组合 $\textit{pre}[i]$ 和 $\textit{suf}[i+1]$ 中的元素计算 XOR，取其中最大值作为答案。

**小优化**：如果在循环中，发现答案 $\textit{ans}$ 达到了理论最大值 $2^7-1$（或者所有元素的 OR），则立刻返回答案。

也可以用哈希集合代替布尔数组，见下面的 Python 优化代码。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Ub4mekE1x/) 第三题，欢迎点赞关注~

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
        pre = [[False] * (mx + 1) for _ in range(k + 1)]
        pre[0][0] = True
        for i, v in enumerate(nums[:-k]):
            for j in range(min(k - 1, i), -1, -1):
                for x, has_x in enumerate(pre[j]):
                    if has_x:
                        pre[j + 1][x | v] = True
            if i < k - 1:
                continue
            for x, has_x in enumerate(pre[k]):
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
        pre = [set() for _ in range(k + 1)]
        pre[0].add(0)
        for i, v in enumerate(nums[:-k]):
            for j in range(min(k - 1, i), -1, -1):
                pre[j + 1].update(x | v for x in pre[j])
            if i < k - 1:
                continue
            ans = max(ans, max(x ^ y for x in pre[k] for y in suf[i + 1]))
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
        boolean[][] pre = new boolean[k + 1][MX];
        pre[0][0] = true;
        for (int i = 0; i < n - k; i++) {
            int v = nums[i];
            for (int j = Math.min(k - 1, i); j >= 0; j--) {
                for (int x = 0; x < MX; x++) {
                    if (pre[j][x]) {
                        pre[j + 1][x | v] = true;
                    }
                }
            }
            if (i < k - 1) {
                continue;
            }
            for (int x = 0; x < MX; x++) {
                if (pre[k][x]) {
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
        vector<array<int, MX>> pre(k + 1);
        pre[0][0] = true;
        for (int i = 0; i < n - k; i++) {
            int v = nums[i];
            for (int j = min(k - 1, i); j >= 0; j--) {
                for (int x = 0; x < MX; x++) {
                    if (pre[j][x]) {
                        pre[j + 1][x | v] = true;
                    }
                }
            }
            if (i < k - 1) {
                continue;
            }
            for (int x = 0; x < MX; x++) {
                if (pre[k][x]) {
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

	pre := make([][mx]bool, k+1)
	pre[0][0] = true
	for i, v := range nums[:n-k] {
		for j := min(k-1, i); j >= 0; j-- {
			for x, hasX := range pre[j] {
				if hasX {
					pre[j+1][x|v] = true
				}
			}
		}
		if i < k-1 {
			continue
		}
		for x, hasX := range pre[k] {
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

- 时间复杂度：$\mathcal{O}(nkU)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U$ 是 $\textit{nums}$ 所有元素的 OR，本题至多为 $2^7-1$。
- 空间复杂度：$\mathcal{O}(n+k)U)$。

更多相似题目，见下面动态规划题单中的「**§3.1 0-1 背包**」和「**专题：前后缀分解**」。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

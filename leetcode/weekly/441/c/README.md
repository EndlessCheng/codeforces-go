## 思路

由于题目让我们选的是范围 $[l_i, r_i]$ 内的一个下标**子集**，所以每个 $\textit{nums}[i]$ 是**互相独立**的，可以分别计算。

选出包含 $i$ 的询问，设这些询问的 $\textit{val}$ 组成了数组 $\textit{vals}$，问题变成：

- 从 $\textit{vals}$ 的前缀中选一些数，元素和能否恰好等于 $\textit{nums}[i]$？

这是 0-1 背包的标准应用，原理见[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

从左到右遍历 $\textit{queries}$，计算 0-1 背包，如果每个 $\textit{nums}[i]$ 都能通过一些数的相加得到，那么返回此时 $\textit{queries}$ 的下标加一。

注意特判 $\textit{nums}$ 全为 $0$ 的情况，此时无需操作，返回 $0$。

如果遍历完 $\textit{queries}$ 也没有返回答案，那么返回 $-1$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1JYQ8YWEvD/?t=21m27s)，欢迎点赞关注~

## 写法一：布尔数组

```py [sol-Python3]
class Solution:
    def minZeroArray(self, nums: List[int], queries: List[List[int]]) -> int:
        if all(x == 0 for x in nums):
            return 0  # nums 全为 0
        f = [[True] + [False] * x for x in nums]
        for k, (l, r, val) in enumerate(queries):
            for i in range(l, r + 1):
                if f[i][-1]: continue  # 小优化：已经满足要求，不计算
                for j in range(nums[i], val - 1, -1):
                    f[i][j] = f[i][j] or f[i][j - val]
            if all(fi[-1] for fi in f):
                return k + 1
        return -1
```

```java [sol-Java]
class Solution {
    public int minZeroArray(int[] nums, int[][] queries) {
        if (Arrays.stream(nums).allMatch(x -> x == 0)) {
            return 0; // nums 全为 0
        }

        int n = nums.length;
        boolean[][] f = new boolean[n][];
        for (int i = 0; i < n; i++) {
            f[i] = new boolean[nums[i] + 1];
            f[i][0] = true;
        }

        for (int k = 0; k < queries.length; k++) {
            int[] q = queries[k];
            int val = q[2];
            for (int i = q[0]; i <= q[1]; i++) {
                if (f[i][nums[i]]) continue; // 小优化：已经满足要求，不计算
                for (int j = nums[i]; j >= val; j--) {
                    f[i][j] = f[i][j] || f[i][j - val];
                }
            }
            boolean ok = true;
            for (int i = 0; i < n; i++) {
                if (!f[i][nums[i]]) {
                    ok = false;
                    break;
                }
            }
            if (ok) {
                return k + 1;
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        if (ranges::all_of(nums, [](int x) { return x == 0; })) {
            return 0; // nums 全为 0
        }

        int n = nums.size();
        vector<vector<int>> f(n);
        for (int i = 0; i < n; i++) {
            f[i].resize(nums[i] + 1);
            f[i][0] = true;
        }

        for (int k = 0; k < queries.size(); k++) {
            auto& q = queries[k];
            int val = q[2];
            for (int i = q[0]; i <= q[1]; i++) {
                if (f[i][nums[i]]) continue; // 小优化：已经满足要求，不计算
                for (int j = nums[i]; j >= val; j--) {
                    f[i][j] = f[i][j] || f[i][j - val];
                }
            }
            bool ok = true;
            for (auto& fi : f) {
                if (!fi.back()) {
                    ok = false;
                    break;
                }
            }
            if (ok) {
                return k + 1;
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
func minZeroArray(nums []int, queries [][]int) int {
	for _, x := range nums {
		if x > 0 {
			goto normal
		}
	}
	return 0 // nums 全为 0
normal:
	f := make([][]bool, len(nums))
	for i, x := range nums {
		f[i] = make([]bool, x+1)
		f[i][0] = true
	}
next:
	for k, q := range queries {
		val := q[2]
		for i := q[0]; i <= q[1]; i++ {
			if f[i][nums[i]] {
				continue // 小优化：已经满足要求，不计算
			}
			for j := nums[i]; j >= val; j-- {
				f[i][j] = f[i][j] || f[i][j-val]
			}
		}
		for i, x := range nums {
			if !f[i][x] {
				continue next
			}
		}
		return k + 1
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(qnU)$，其中 $q$ 是 $\textit{queries}$ 的长度，$n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(nU)$。

## 写法二：bitset

用 bitset（视作一个二进制数）代替布尔数组。二进制数从低到高第 $i$ 位是 $0$，表示布尔数组的第 $i$ 个数是 $\texttt{false}$；从低到高第 $i$ 位是 $1$，表示布尔数组的第 $i$ 个数是 $\texttt{true}$。

转移方程等价于，设 $s=f[i]$，把 $s$ 中的每个比特位增加 $\textit{val}$，即左移 $\textit{val}$ 位，再与 $f[i]$ 计算 OR。前者对应选择一个值为 $\textit{val}$ 的物品，后者对应不选。

判断 $f[i][x]$ 是否为 $\texttt{true}$，等价于判断 $f[i]$ 的第 $x$ 位是否为 $1$，即 `(f[i] >> x & 1) == 1`。

更多位运算技巧，请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def minZeroArray(self, nums: List[int], queries: List[List[int]]) -> int:
        if all(x == 0 for x in nums):
            return 0  # nums 全为 0
        f = [1] * len(nums)
        for k, (l, r, val) in enumerate(queries):
            for i in range(l, r + 1):
                f[i] |= f[i] << val  # 本题 val 比较小，超出 nums[i] 比特位没有去掉
            if all(fi >> x & 1 for fi, x in zip(f, nums)):
                return k + 1
        return -1
```

```java [sol-Java]
import java.math.BigInteger;

class Solution {
    public int minZeroArray(int[] nums, int[][] queries) {
        if (Arrays.stream(nums).allMatch(x -> x == 0)) {
            return 0; // nums 全为 0
        }

        int n = nums.length;
        BigInteger[] f = new BigInteger[n];
        Arrays.fill(f, BigInteger.ONE);

        for (int k = 0; k < queries.length; k++) {
            int[] q = queries[k];
            for (int i = q[0]; i <= q[1]; i++) {
                if (!f[i].testBit(nums[i])) { // 小优化：已经满足要求，不计算
                    f[i] = f[i].or(f[i].shiftLeft(q[2]));
                }
            }
            boolean ok = true;
            for (int i = 0; i < n; i++) {
                if (!f[i].testBit(nums[i])) {
                    ok = false;
                    break;
                }
            }
            if (ok) {
                return k + 1;
            }
        }
        return -1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        if (ranges::all_of(nums, [](int x) { return x == 0; })) {
            return 0; // nums 全为 0
        }

        int n = nums.size();
        vector<bitset<1001>> f(n);
        for (auto& bs : f) {
            bs[0] = 1;
        }

        for (int k = 0; k < queries.size(); k++) {
            auto& q = queries[k];
            int val = q[2];
            for (int i = q[0]; i <= q[1]; i++) {
                if (!f[i][nums[i]]) { // 小优化：已经满足要求，不计算
                    f[i] |= f[i] << val;
                }
            }
            bool ok = true;
            for (int i = 0; i < n; i++) {
                if (!f[i][nums[i]]) {
                    ok = false;
                    break;
                }
            }
            if (ok) {
                return k + 1;
            }
        }
        return -1;
    }
};
```

```go [sol-Go]
func minZeroArray(nums []int, queries [][]int) int {
	for _, x := range nums {
		if x > 0 {
			goto normal
		}
	}
	return 0
normal:
	f := make([]*big.Int, len(nums))
	for i := range f {
		f[i] = big.NewInt(1)
	}
	p := new(big.Int)
next:
	for k, q := range queries {
		val := uint(q[2])
		for i := q[0]; i <= q[1]; i++ {
			if f[i].Bit(nums[i]) == 0 { // 小优化：已经满足要求，不计算
				f[i].Or(f[i], p.Lsh(f[i], val))
			}
		}
		for i, x := range nums {
			if f[i].Bit(x) == 0 {
				continue next
			}
		}
		return k + 1
	}
	return -1
}
```

#### 复杂度分析

以下分析，不考虑超出 $\textit{nums}[i]$ 的比特位。

- 时间复杂度：$\mathcal{O}(qnU / w)$，其中 $q$ 是 $\textit{queries}$ 的长度，$n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$，$w=32$ 或 $64$。
- 空间复杂度：$\mathcal{O}(nU / w)$。

更多相似题目，见下面动态规划题单中的「**§3.1 0-1 背包**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. 【本题相关】[动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

根据 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)，把每个 $\textit{nums}_i$ 视作一个集合 $S_i$。

用集合思考，两个数的按位与等于 $0$，等价于两个集合 $A$ 和 $B$ 的交集为空，即 

$$
A\cap B = \varnothing
$$

这等价于 $B$ 是 $A$ 关于全集 $U$ 的补集的子集，即

$$
B \subseteq \complement_UA
$$

枚举 $\textit{nums}_i$，为了让乘积最大，我们需要知道 $\textit{nums}_j$ 的最大值，其中 $j$ 需要满足 $S_j \subseteq \complement_US_i$。

这是**高维前缀和**（Sum Over Subsets DP，SOS DP）的标准应用。

> 注：Sum Over Subsets DP 直译过来是**子集和 DP**，国内算法竞赛圈一般叫高维前缀和。

定义 $f[S]$ 表示 $S$ 的所有子集的对应元素值的最大值。

$S$ 可以分解为如下子集的**并集**。

- 枚举 $S$ 中的元素 $i$。
- 移除 $i$，得到 $S$ 的子集 $T_i = S \setminus \{i\}$。
- 所有 $T_i$ 的并集即为 $S$。

由于 $T_i$ 是一个更小的集合，是规模更小的子问题，所以可以用 DP 解决。

计算所有 $f[T_i] = f[S \setminus \{i\}]$ 的最大值

$$
\max_{i\in S} f[S \setminus \{i\}]
$$

去更新 $f[S]$ 的最大值。

> 注意上式只包含 $S$ 的真子集，没有包含 $S$ 自身，即 $f[S]$ 的初始值。所以不能直接说 $f[S] = \max\limits_{i\in S} f[S \setminus \{i\}]$。

初始值：$f[S_i] = \textit{nums}_i$。

答案为

$$
\max_{i=0}^{n-1} \textit{nums}_i\cdot f[\complement_US_i]
$$

集合用二进制数表示，用位运算实现集合的运算，具体请看 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

[本题视频讲解](https://www.bilibili.com/video/BV1SMaGz7EXe/)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        w = max(nums).bit_length()
        u = 1 << w
        f = [0] * u
        for x in nums:
            f[x] = x  # 初始值

        for s in range(3, u):  # 从小到大枚举集合 s（至少有两个数）
            for i in range(w):  # 枚举 s 中的元素 i
                if s >> i & 1:  # i 属于集合 s
                    v = f[s ^ (1 << i)]  # 从 s 的子集 s \ {i} 转移过来
                    if v > f[s]:
                        f[s] = v  # 手写 max 更快

        return max(x * f[(u - 1) ^ x] for x in nums)
```

```java [sol-Java]
class Solution {
    public long maxProduct(int[] nums) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }
        
        int w = 32 - Integer.numberOfLeadingZeros(mx); // mx 的二进制长度
        int u = 1 << w;
        int[] f = new int[u];
        for (int x : nums) {
            f[x] = x; // 初始值
        }

        for (int s = 3; s < u; s++) { // 从小到大枚举集合 s（至少有两个数）
            for (int i = 0; i < w; i++) { // 枚举 s 中的元素 i
                if ((s >> i & 1) > 0) { // i 属于集合 s
                    f[s] = Math.max(f[s], f[s ^ (1 << i)]); // 从 s 的子集 s \ {i} 转移过来
                }
            }
        }

        long ans = 0;
        for (int x : nums) {
            ans = Math.max(ans, (long) x * f[(u - 1) ^ x]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxProduct(vector<int>& nums) {
        int w = bit_width((uint32_t) ranges::max(nums));
        int u = 1 << w;
        vector<int> f(u);
        for (int x : nums) {
            f[x] = x; // 初始值
        }

        for (int s = 3; s < u; s++) { // 从小到大枚举集合 s（至少有两个数）
            for (int i = 0; i < w; i++) { // 枚举 s 中的元素 i
                if (s >> i & 1) { // i 属于集合 s
                    f[s] = max(f[s], f[s ^ (1 << i)]); // 从 s 的子集 s \ {i} 转移过来
                }
            }
        }

        long long ans = 0;
        for (int x : nums) {
            ans = max(ans, 1LL * x * f[(u - 1) ^ x]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxProduct(nums []int) int64 {
	w := bits.Len(uint(slices.Max(nums)))
	u := 1 << w
	f := make([]int, u)
	for _, x := range nums {
		f[x] = x // 初始值
	}

	for s := 3; s < u; s++ { // 从小到大枚举集合 s（至少有两个数）
		for i := range w { // 枚举 s 中的元素 i
			if s>>i&1 > 0 { // i 属于集合 s
				f[s] = max(f[s], f[s^1<<i]) // 从 s 的子集 s \ {i} 转移过来
			}
		}
	}

	ans := 0
	for _, x := range nums {
		ans = max(ans, x*f[u-1^x])
	}
	return int64(ans)
}
```

## 优化

先枚举 $i$，再枚举 $s$。

通过把 $s$ 或上 $2^i$，快速跳到第 $i$ 位是 $1$ 的 $s$，从而减少无效枚举。相当于跳过了所有 `(s >> i & 1) == 0` 的情况。

```py [sol-Python3]
class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        w = max(nums).bit_length()
        u = 1 << w
        f = [0] * u
        for x in nums:
            f[x] = x

        for i in range(w):
            bit = 1 << i  # 避免在循环中反复计算 1 << i
            s = 3
            while s < u:
                s |= bit  # 快速跳到第 i 位是 1 的 s
                v = f[s ^ bit]
                if v > f[s]:
                    f[s] = v
                s += 1

        return max(x * f[(u - 1) ^ x] for x in nums)
```

```java [sol-Java]
class Solution {
    public long maxProduct(int[] nums) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int w = 32 - Integer.numberOfLeadingZeros(mx);
        int u = 1 << w;
        int[] f = new int[u];
        for (int x : nums) {
            f[x] = x;
        }

        for (int i = 0; i < w; i++) {
            for (int s = 3; s < u; s++) {
                s |= 1 << i; // 快速跳到第 i 位是 1 的 s
                f[s] = Math.max(f[s], f[s ^ (1 << i)]);
            }
        }

        long ans = 0;
        for (int x : nums) {
            ans = Math.max(ans, (long) x * f[(u - 1) ^ x]);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxProduct(vector<int>& nums) {
        int w = bit_width((uint32_t) ranges::max(nums));
        int u = 1 << w;
        vector<int> f(u);
        for (int x : nums) {
            f[x] = x;
        }

        for (int i = 0; i < w; i++) {
            for (int s = 0; s < u; s++) {
                s |= 1 << i; // 快速跳到第 i 位是 1 的 s
                f[s] = max(f[s], f[s ^ (1 << i)]);
            }
        }

        long long ans = 0;
        for (int x : nums) {
            ans = max(ans, 1LL * x * f[(u - 1) ^ x]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxProduct(nums []int) int64 {
	w := bits.Len(uint(slices.Max(nums)))
	u := 1 << w
	f := make([]int, u)
	for _, x := range nums {
		f[x] = x
	}

	for i := range w {
		for s := 3; s < u; s++ {
			s |= 1 << i // 快速跳到第 i 位是 1 的 s
			f[s] = max(f[s], f[s^1<<i])
		}
	}

	ans := 0
	for _, x := range nums {
		ans = max(ans, x*f[u-1^x])
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + U\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(U)$。

## 相似题目

[2732. 找到矩阵中的好子集](https://leetcode.cn/problems/find-a-good-subset-of-the-matrix/)

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

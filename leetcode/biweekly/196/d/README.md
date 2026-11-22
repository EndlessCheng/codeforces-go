## 核心思路

1. 用定长滑动窗口算出每个 $x_i$。**注**：这可以用栈优化到 $\mathcal{O}(n)$。
2. 用 SOS DP，求出每个集合 $S$ 是否有一个子集为 $x_i$。

## 定长滑动窗口

前置知识：[定长滑动窗口原理讲解](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/)。

如果窗口每个数的二进制第 $i$ 位都是 $1$，那么按位与的第 $i$ 位是 $1$，否则是 $0$。

设 $\max(\textit{nums})$ 的二进制长度为 $w$。创建一个长为 $w$ 的数组 $\textit{cnt}$，维护对应的 $1$ 的个数。如果 $\textit{cnt}[i] = k$，那么 $x$ 二进制的 $i$ 位是 $1$，否则是 $0$。

> 小优化：如果发现 $x = 0$，由于 $0$ AND 任何数都是 $0$，所以无解，直接返回 $-1$。

## SOS DP

枚举 $y=1,2,\ldots,2^w-1$，如何判断是否存在 $x_i$，使得 $y\ \&\ x_i = 0$？

下文用到的术语来自 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/discuss/post/3571304/cong-ji-he-lun-dao-wei-yun-suan-chang-ji-enve/)。

$y\ \&\ x_i = 0$ 从集合角度来看，即集合 $y$ 与 $x_i$ 不相交。或者说，$x_i$ 是 $y$ 的补集的子集。

具体地，设 $y$ 关于全集 $U = 2^w-1$ 的补集为 $z = U\oplus y$，那么 $y\ \&\ x_i = 0$ 等价于 $x_i$ 是 $z$ 的子集。

用 **SOS DP** 计算，即可对每个 $z$ 都知道是否存在 $x_i$ 是 $z$ 的子集。

[SOS DP 原理讲解（方法二）](https://leetcode.cn/problems/maximum-product-of-two-integers-with-no-common-bits/solutions/3768219/mo-ban-gao-wei-qian-zhui-he-sos-dppython-78fz/)。

> 注：本题计算的内容在**组合优化**中叫做**最小命中集合**（Minimum Hitting Set）。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

## 优化前

```py [sol-Python3]
class Solution:
    def minSetBits(self, nums: list[int], k: int) -> int:
        w = max(nums).bit_length()
        cnt = [0] * w
        u = 1 << w
        f = [False] * u

        for r, v in enumerate(nums):
            # 1. 入
            for i in range(w):
                cnt[i] += v >> i & 1

            l = r + 1 - k
            if l < 0:  # 尚未形成第一个窗口
                continue

            # 2. 计算 x
            x = 0
            for i, c in enumerate(cnt):
                if c == k:
                    x |= 1 << i
            f[x] = True

            # 3. 出
            v = nums[l]
            for i in range(w):
                cnt[i] -= v >> i & 1

        for i in range(w):
            bit = 1 << i
            s = 0
            while s < u:
                s |= bit  # 优化：快速跳到 i 位是 1 的 s
                f[s] = f[s] or f[s ^ bit]
                s += 1

        mx = -1
        # 先计算 y 的补集的二进制 1 的个数的最大值
        for z, has_subset in enumerate(f):
            if not has_subset:
                mx = max(mx, z.bit_count())

        if mx < 0:
            return -1
        # 最后用 w 减去 mx，即为 y 的二进制 1 的个数的最小值
        return w - mx
```

```java [sol-Java]
class Solution {
    public int minSetBits(int[] nums, int k) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int w = 32 - Integer.numberOfLeadingZeros(mx);
        int[] cnt = new int[w];
        boolean[] f = new boolean[1 << w];

        for (int r = 0; r < nums.length; r++) {
            int v = nums[r];
            // 1. 入
            for (int i = 0; i < w; i++) {
                cnt[i] += v >> i & 1;
            }

            int l = r + 1 - k;
            if (l < 0) { // 尚未形成第一个窗口
                continue;
            }

            // 2. 计算 x
            int x = 0;
            for (int i = 0; i < w; i++) {
                if (cnt[i] == k) {
                    x |= 1 << i;
                }
            }
            f[x] = true;

            // 3. 出
            v = nums[l];
            for (int i = 0; i < w; i++) {
                cnt[i] -= v >> i & 1;
            }
        }

        for (int i = 0; i < w; i++) {
            for (int s = 0; s < 1 << w; s++) {
                s |= 1 << i; // 优化：快速跳到 i 位是 1 的 s
                f[s] = f[s] || f[s ^ (1 << i)];
            }
        }

        int maxCnt = -1;
        // 先计算 y 的补集的二进制 1 的个数的最大值
        for (int z = 0; z < 1 << w; z++) {
            if (!f[z]) {
                maxCnt = Math.max(maxCnt, Integer.bitCount(z));
            }
        }

        if (maxCnt < 0) {
            return -1;
        }
        // 最后用 w 减去 maxCnt，即为 y 的二进制 1 的个数的最小值
        return w - maxCnt;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minSetBits(vector<int>& nums, int k) {
        int w = bit_width(1u * ranges::max(nums));
        vector<int> cnt(w);
        vector<int8_t> f(1 << w);

        for (int r = 0; r < nums.size(); r++) {
            int v = nums[r];
            // 1. 入
            for (int i = 0; i < w; i++) {
                cnt[i] += v >> i & 1;
            }

            int l = r + 1 - k;
            if (l < 0) { // 尚未形成第一个窗口
                continue;
            }

            // 2. 计算 x
            int x = 0;
            for (int i = 0; i < w; i++) {
                if (cnt[i] == k) {
                    x |= 1 << i;
                }
            }
            f[x] = true;

            // 3. 出
            v = nums[l];
            for (int i = 0; i < w; i++) {
                cnt[i] -= v >> i & 1;
            }
        }

        for (int i = 0; i < w; i++) {
            for (int s = 0; s < 1 << w; s++) {
                s |= 1 << i; // 优化：快速跳到 i 位是 1 的 s
                f[s] = f[s] || f[s ^ (1 << i)];
            }
        }

        int mx = -1;
        // 先计算 y 的补集的二进制 1 的个数的最大值
        for (uint32_t z = 0; z < 1 << w; z++) {
            if (!f[z]) {
                mx = max(mx, popcount(z));
            }
        }

        if (mx < 0) {
            return -1;
        }
        // 最后用 w 减去 mx，即为 y 的二进制 1 的个数的最小值
        return w - mx;
    }
};
```

```go [sol-Go]
func minSetBits(nums []int, k int) int {
	w := bits.Len(uint(slices.Max(nums)))
	cnt := make([]int, w)
	f := make([]bool, 1<<w)
	for r, v := range nums {
		// 1. 入
		for ; v > 0; v &= v - 1 {
			cnt[bits.TrailingZeros(uint(v))]++
		}

		l := r + 1 - k
		if l < 0 { // 尚未形成第一个窗口
			continue
		}

		// 2. 计算 x
		x := 0
		for i, c := range cnt {
			if c == k {
				x |= 1 << i
			}
		}
		f[x] = true

		// 3. 出
		for v = nums[l]; v > 0; v &= v - 1 {
			cnt[bits.TrailingZeros(uint(v))]--
		}
	}

	for i := range w {
		for s := 0; s < 1<<w; s++ {
			s |= 1 << i // 优化：快速跳到 i 位是 1 的 s
			f[s] = f[s] || f[s^1<<i]
		}
	}

	mx := -1
	// 先计算 y 的补集的二进制 1 的个数的最大值
	for z, hasSubset := range f[:1<<w] {
		if !hasSubset {
			mx = max(mx, bits.OnesCount(uint(z)))
		}
	}

	if mx < 0 {
		return -1
	}
	// 最后用 w 减去 mx，即为 y 的二进制 1 的个数的最小值
	return w - mx
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+U)\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。设 $W = \log U$，定长滑窗需要 $\mathcal{O}(nW)$ 时间，SOS DP 需要 $\mathcal{O}(W2^W)$ 时间。
- 空间复杂度：$\mathcal{O}(U)$。

## 优化

定长滑窗可以用栈优化到 $\mathcal{O}(n)$，见 [原理讲解（方法二）](https://leetcode.cn/problems/find-subarray-with-bitwise-or-closest-to-k/solutions/2798206/li-yong-and-de-xing-zhi-pythonjavacgo-by-gg4d/)。

```py [sol-Python3]
class Solution:
    def minSetBits(self, nums: list[int], k: int) -> int:
        w = max(nums).bit_length()
        u = 1 << w
        f = [False] * u
        bottom = 0
        right_and = -1

        for r, v in enumerate(nums):
            right_and &= v
            l = r + 1 - k
            if l < 0:
                continue
            if bottom < l:
                # 重新构建一个栈
                for i in range(r - 1, l - 1, -1):
                    nums[i] &= nums[i + 1]
                bottom = r
                right_and = -1
            f[nums[l] & right_and] = True

        for i in range(w):
            bit = 1 << i
            s = 0
            while s < u:
                s |= bit  # 优化：快速跳到 i 位是 1 的 s
                f[s] = f[s] or f[s ^ bit]
                s += 1

        mx = -1
        # 先计算 y 的补集的二进制 1 的个数的最大值
        for z, has_subset in enumerate(f):
            if not has_subset:
                mx = max(mx, z.bit_count())

        if mx < 0:
            return -1
        # 最后用 w 减去 mx，即为 y 的二进制 1 的个数的最小值
        return w - mx
```

```java [sol-Java]
class Solution {
    public int minSetBits(int[] nums, int k) {
        int mx = 0;
        for (int x : nums) {
            mx = Math.max(mx, x);
        }

        int w = 32 - Integer.numberOfLeadingZeros(mx);
        boolean[] f = new boolean[1 << w];
        int bottom = 0;
        int rightAnd = -1;

        for (int r = 0; r < nums.length; r++) {
            int v = nums[r];
            rightAnd &= v;
            int l = r + 1 - k;
            if (l < 0) {
                continue;
            }
            if (bottom < l) {
                // 重新构建一个栈
                for (int i = r - 1; i >= l; i--) {
                    nums[i] &= nums[i + 1];
                }
                bottom = r;
                rightAnd = -1;
            }
            f[nums[l] & rightAnd] = true;
        }

        for (int i = 0; i < w; i++) {
            for (int s = 0; s < 1 << w; s++) {
                s |= 1 << i; // 优化：快速跳到 i 位是 1 的 s
                f[s] = f[s] || f[s ^ (1 << i)];
            }
        }

        int maxCnt = -1;
        // 先计算 y 的补集的二进制 1 的个数的最大值
        for (int z = 0; z < 1 << w; z++) {
            if (!f[z]) {
                maxCnt = Math.max(maxCnt, Integer.bitCount(z));
            }
        }

        if (maxCnt < 0) {
            return -1;
        }
        // 最后用 w 减去 maxCnt，即为 y 的二进制 1 的个数的最小值
        return w - maxCnt;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minSetBits(vector<int>& nums, int k) {
        int w = bit_width(1u * ranges::max(nums));
        vector<int8_t> f(1 << w);
        int bottom = 0;
        int right_and = -1;

        for (int r = 0; r < nums.size(); r++) {
            int v = nums[r];
            right_and &= v;
            int l = r + 1 - k;
            if (l < 0) {
                continue;
            }
            if (bottom < l) {
                // 重新构建一个栈
                for (int i = r - 1; i >= l; i--) {
                    nums[i] &= nums[i + 1];
                }
                bottom = r;
                right_and = -1;
            }
            f[nums[l] & right_and] = true;
        }

        for (int i = 0; i < w; i++) {
            for (int s = 0; s < 1 << w; s++) {
                s |= 1 << i; // 优化：快速跳到 i 位是 1 的 s
                f[s] = f[s] || f[s ^ (1 << i)];
            }
        }

        int mx = -1;
        // 先计算 y 的补集的二进制 1 的个数的最大值
        for (uint32_t z = 0; z < 1 << w; z++) {
            if (!f[z]) {
                mx = max(mx, popcount(z));
            }
        }

        if (mx < 0) {
            return -1;
        }
        // 最后用 w 减去 mx，即为 y 的二进制 1 的个数的最小值
        return w - mx;
    }
};
```

```go [sol-Go]
func minSetBits(nums []int, k int) int {
	w := bits.Len(uint(slices.Max(nums)))
	f := make([]bool, 1<<w)
	bottom := 0
	rightAnd := -1
	for r, v := range nums {
		rightAnd &= v
		l := r + 1 - k
		if l < 0 {
			continue
		}
		if bottom < l {
			// 重新构建一个栈
			for i := r - 1; i >= l; i-- {
				nums[i] &= nums[i+1]
			}
			bottom = r
			rightAnd = -1
		}
		f[nums[l]&rightAnd] = true
	}

	for i := range w {
		for s := 0; s < 1<<w; s++ {
			s |= 1 << i // 优化：快速跳到 i 位是 1 的 s
			f[s] = f[s] || f[s^1<<i]
		}
	}

	mx := -1
	// 先计算 y 的补集的二进制 1 的个数的最大值
	for z, hasSubset := range f[:1<<w] {
		if !hasSubset {
			mx = max(mx, bits.OnesCount(uint(z)))
		}
	}

	if mx < 0 {
		return -1
	}
	// 最后用 w 减去 mx，即为 y 的二进制 1 的个数的最小值
	return w - mx
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + U\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。设 $W = \log U$，定长滑窗需要 $\mathcal{O}(n)$ 时间，SOS DP 需要 $\mathcal{O}(W2^W)$ 时间。
- 空间复杂度：$\mathcal{O}(U)$。

## 专题训练

1. 滑动窗口题单的「**一、定长滑动窗口**」。
2. 位运算题单的「**AND/OR LogTrick**」。
3. 动态规划题单的「**§9.6 SOS DP**」

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

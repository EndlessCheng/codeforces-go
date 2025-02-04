### 情况一

先来看 $a$ 和 $b$ 均小于 $2^n$ 的情况。

举例说明：

```py
a = 11101
b = 00001
```

其中 $a$ 和 $b$ 同为 $0$ 的比特位，$x$ 可以取 $1$，这样异或后都是 $1$。同为 $1$ 的比特位，$x$ 可以取 $0$，这样异或后也都是 $1$。

所以重点讨论 $a$ 和 $b$ 的同一个比特位一个是 $1$ 另一个是 $0$ 的情况。

可以发现，无论 $x$ 取 $0$ 还是取 $1$，总有一个是 $1$ 另一个是 $0$，换句话说，$1$ 的个数是不变的。我把这样的比特位叫做「可分配」的位。

设 $\textit{ax} = a \oplus x,\ \textit{bx} = b \oplus x$，其中 $\oplus$ 表示异或运算。

如果 $x=00010$，则有

```py
ax = 11111
bx = 00011
```

乘积为 $31\cdot 3 = 93$，不够大。如何分配这些 $1$，使得 $\textit{ax}\cdot \textit{bx}$ 尽量大呢？

- $\textit{ax}$ 和 $\textit{bx}$ 的低两位一定是 $1$。
- $\textit{ax}$ 和 $\textit{bx}$ 的其余位，无论 $x$ 怎么取，一定满足总有一个是 $1$ 另一个是 $0$。

也就是说，$\textit{ax}+\textit{bx}$ 是一个**定值**！

根据**基本不等式（均值定理）**，在和为定值的前提下，要让乘积尽量大，应当让 $\textit{ax}$ 与 $\textit{bx}$ 尽量接近。

所以，最优的分配方式是把剩下「可分配」的位中的最高位分给 $\textit{ax}$，其余位分给 $\textit{bx}$，即

```py
ax = 10011
bx = 01111
```

此时乘积最大，为 $19\cdot 15 = 285$。

### 情况二

然后来看 $a$ 和 $b$ 至少有一个 $\ge 2^n$ 的情况。

不妨设 $a\ge b$。

对于高于或等于 $n$ 的比特位，我们是无法修改的，这会对「分配」产生什么影响呢？

分类讨论：

- 如果高于或等于 $n$ 的比特位，$a$ 和 $b$ 是一样的，那么问题转换成情况一。
- 否则高于或等于 $n$ 的比特位，满足 $a>b$，由于这些比特位无法修改，所以无论怎么分配，永远有 $\textit{ax} > \textit{bx}$，所以对于可修改的部分（低于 $n$ 的比特位），应当把「可分配」的位**全部**分给 $\textit{bx}$，让 $\textit{bx}$ 尽量大，从而使 $\textit{ax}$ 和 $\textit{bx}$ 尽量接近，得到最大的 $\textit{ax}\cdot \textit{bx}$。

代码中用到了大量位运算的技巧，请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

[本题视频讲解](https://www.bilibili.com/video/BV1pC4y1j7Pw/)，欢迎点赞投币~

```py [sol-Python3]
class Solution:
    def maximumXorProduct(self, a: int, b: int, n: int) -> int:
        if a < b:
            a, b = b, a  # 保证 a >= b

        mask = (1 << n) - 1
        ax = a & ~mask  # 第 n 位及其左边，无法被 x 影响，先算出来
        bx = b & ~mask
        a &= mask  # 低于第 n 位，能被 x 影响
        b &= mask

        left = a ^ b  # 可分配：a XOR x 和 b XOR x 一个是 1 另一个是 0
        one = mask ^ left  # 无需分配：a XOR x 和 b XOR x 均为 1
        ax |= one  # 先加到异或结果中
        bx |= one

        # 现在要把 left 分配到 ax 和 bx 中
        # 根据基本不等式（均值定理），分配后应当使 ax 和 bx 尽量接近，乘积才能尽量大
        if left > 0 and ax == bx:
            # 尽量均匀分配，例如把 1111 分成 1000 和 0111
            high_bit = 1 << (left.bit_length() - 1)
            ax |= high_bit
            left ^= high_bit
        # 如果 a & ~mask 更大，则应当全部分给 bx（注意最上面保证了 a>=b）
        bx |= left

        return ax * bx % 1_000_000_007
```

```java [sol-Java]
class Solution {
    public int maximumXorProduct(long a, long b, int n) {
        if (a < b) {
            // 保证 a >= b
            long temp = a;
            a = b;
            b = temp;
        }

        long mask = (1L << n) - 1;
        long ax = a & ~mask; // 第 n 位及其左边，无法被 x 影响，先算出来
        long bx = b & ~mask;
        a &= mask; // 低于第 n 位，能被 x 影响
        b &= mask;

        long left = a ^ b; // 可分配：a XOR x 和 b XOR x 一个是 1 另一个是 0
        long one = mask ^ left; // 无需分配：a XOR x 和 b XOR x 均为 1
        ax |= one; // 先加到异或结果中
        bx |= one;

        // 现在要把 left 分配到 ax 和 bx 中
        // 根据基本不等式（均值定理），分配后应当使 ax 和 bx 尽量接近，乘积才能尽量大
        if (left > 0 && ax == bx) {
            // 尽量均匀分配，例如把 1111 分成 1000 和 0111
            long highBit = 1L << (63 - Long.numberOfLeadingZeros(left));
            ax |= highBit;
            left ^= highBit;
        }
        // 如果 a & ~mask 更大，则应当全部分给 bx（注意最上面保证了 a>=b）
        bx |= left;

        final long MOD = 1_000_000_007;
        return (int) (ax % MOD * (bx % MOD) % MOD); // 注意不能直接 long * long，否则溢出
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumXorProduct(long long a, long long b, int n) {
        if (a < b) {
            swap(a, b); // 保证 a >= b
        }

        long long mask = (1LL << n) - 1;
        long long ax = a & ~mask; // 第 n 位及其左边，无法被 x 影响，先算出来
        long long bx = b & ~mask;
        a &= mask; // 低于第 n 位，能被 x 影响
        b &= mask;

        long long left = a ^ b; // 可分配：a XOR x 和 b XOR x 一个是 1 另一个是 0
        long long one = mask ^ left; // 无需分配：a XOR x 和 b XOR x 均为 1
        ax |= one; // 先加到异或结果中
        bx |= one;

        // 现在要把 left 分配到 ax 和 bx 中
        // 根据基本不等式（均值定理），分配后应当使 ax 和 bx 尽量接近，乘积才能尽量大
        if (left > 0 && ax == bx) {
            // 尽量均匀分配，例如把 1111 分成 1000 和 0111
            long long high_bit = 1LL << (63 - __builtin_clzll(left));
            ax |= high_bit;
            left ^= high_bit;
        }
        // 如果 a & ~mask 更大，则应当全部分给 bx（注意最上面保证了 a>=b）
        bx |= left;

        const long long MOD = 1'000'000'007;
        return ax % MOD * (bx % MOD) % MOD; // 注意不能直接 LL * LL，否则溢出
    }
};
```

```go [sol-Go]
func maximumXorProduct(A, B int64, n int) int {
	const mod = 1_000_000_007
	a, b := int(A), int(B)
	if a < b {
		a, b = b, a // 保证 a >= b
	}

	mask := 1<<n - 1
	ax := a &^ mask // 第 n 位及其左边，无法被 x 影响，先算出来
	bx := b &^ mask
	a &= mask // 低于第 n 位，能被 x 影响
	b &= mask

	left := a ^ b      // 可分配：a XOR x 和 b XOR x 一个是 1 另一个是 0
	one := mask ^ left // 无需分配：a XOR x 和 b XOR x 均为 1
	ax |= one          // 先加到异或结果中
	bx |= one

	// 现在要把 left 分配到 ax 和 bx 中
	// 根据基本不等式（均值定理），分配后应当使 ax 和 bx 尽量接近，乘积才能尽量大
	if left > 0 && ax == bx { // a &^ mask = b &^ mask
		// 尽量均匀分配，例如把 1111 分成 1000 和 0111
		highBit := 1 << (bits.Len(uint(left)) - 1)
		ax |= highBit
		left ^= highBit
	}
	// 如果 a &^ mask 更大，则应当全部分给 bx（注意最上面保证了 a>=b）
	bx |= left

	return ax % mod * (bx % mod) % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

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

## 分析

由于任意两个相邻的元素都不相等，所以前两个数要么递增，要么递减。

第三个数呢？由于任意三个连续的元素不能构成一个严格递增或严格递减的序列：

- 如果前两个数递增，那么第三个数一定比第二个数小（变成递减）。
- 如果前两个数递减，那么第三个数一定比第二个数大（变成递增）。

所以锯齿形数组是一个增减增减交替的数组。

## 寻找子问题

由于我们只关心递增递减，可以把值域 $[l,r]$ 调整为 $[0,r-l]$，方便计算。

比如第 $n$ 个数是 $4$，且最后两个数是递增的，那么第 $n-1$ 个数可以是 $0,1,2,3$，且第 $n-2$ 个数和第 $n-1$ 个数是递减的。枚举第 $n-1$ 个数为 $k$，问题变成在第 $n-1$ 个数为 $k$，且第 $n-2$ 个数和第 $n-1$ 个数是递减的情况下，包含 $n-1$ 个数的锯齿形数组个数。

比如第 $n$ 个数是 $4$，且最后两个数是递减的，那么第 $n-1$ 个数可以是 $5,6,\dots,r-l$，且第 $n-2$ 个数和第 $n-1$ 个数是递增的。枚举第 $n-1$ 个数为 $k$，问题变成在第 $n-1$ 个数为 $k$，且第 $n-2$ 个数和第 $n-1$ 个数是递增的情况下，包含 $n-1$ 个数的锯齿形数组个数。

这些问题都是**和原问题相似的、规模更小的子问题**。

## 状态定义与状态转移方程

定义 $f_0[i][j]$ 表示在第 $i$ 个数为 $j$，且第 $i-1$ 个数和第 $i$ 个数是**递增**的情况下，包含 $i$ 个数的锯齿形数组个数。

定义 $f_1[i][j]$ 表示在第 $i$ 个数为 $j$，且第 $i-1$ 个数和第 $i$ 个数是**递减**的情况下，包含 $i$ 个数的锯齿形数组个数。

对于 $f_0[i][j]$，枚举第 $i-1$ 个数为 $k=0,1,2,\dots,j-1$，问题变成在第 $i-1$ 个数为 $k$，且第 $i-2$ 个数和第 $i-1$ 个数是递减的情况下，包含 $i-1$ 个数的锯齿形数组个数，即 $f_1[i-1][k]$。

累加得

$$
f_0[i][j] = \sum_{k=0}^{j-1} f_1[i-1][k]
$$

对于 $f_1[i][j]$，同理有

$$
f_1[i][j] = \sum_{k=j+1}^{r-l} f_0[i-1][k]
$$

初始值 $f_0[1][j] = f_1[1][j] = 1$。

答案为 $f_0[n]$ 和 $f_1[n]$ 的元素和。

注意转移方程是个子数组的和，可以用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 优化（注意左闭右开）。

代码实现时，$f_0$ 和 $f_1$ 的第一个维度可以优化掉。

注意取模。为什么可以在计算中途取模？请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV156n9z7E9o/?t=14m42s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        MOD = 1_000_000_007
        k = r - l + 1

        f0 = [1] * k  # 后两个数递增
        f1 = [1] * k  # 后两个数递减
        for _ in range(n - 1):
            s0 = list(accumulate(f0, initial=0))
            s1 = list(accumulate(f1, initial=0))
            for j in range(k):
                f0[j] = s1[j] % MOD
                f1[j] = (s0[k] - s0[j + 1]) % MOD

        return (sum(f0) + sum(f1)) % MOD
```

```java [sol-Java]
class Solution {
    public int zigZagArrays(int n, int l, int r) {
        final int MOD = 1_000_000_007;
        int k = r - l + 1;

        int[] f0 = new int[k]; // 后两个数递增
        int[] f1 = new int[k]; // 后两个数递减
        Arrays.fill(f0, 1);
        Arrays.fill(f1, 1);
        long[] s0 = new long[k + 1];
        long[] s1 = new long[k + 1];

        for (int i = 2; i <= n; i++) {
            for (int j = 0; j < k; j++) {
                s0[j + 1] = s0[j] + f0[j];
                s1[j + 1] = s1[j] + f1[j];
            }
            for (int j = 0; j < k; j++) {
                f0[j] = (int) (s1[j] % MOD);
                f1[j] = (int) ((s0[k] - s0[j + 1]) % MOD);
            }
        }

        long ans = 0;
        for (int j = 0; j < k; j++) {
            ans += f0[j] + f1[j];
        }
        return (int) (ans % MOD);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int zigZagArrays(int n, int l, int r) {
        constexpr int MOD = 1'000'000'007;
        int k = r - l + 1;

        vector<int> f0(k, 1); // 后两个数递增
        vector<int> f1(k, 1); // 后两个数递减
        vector<long long> s0(k + 1);
        vector<long long> s1(k + 1);

        for (int i = 2; i <= n; i++) {
            for (int j = 0; j < k; j++) {
                s0[j + 1] = s0[j] + f0[j];
                s1[j + 1] = s1[j] + f1[j];
            }
            for (int j = 0; j < k; j++) {
                f0[j] = s1[j] % MOD;
                f1[j] = (s0[k] - s0[j + 1]) % MOD;
            }
        }

        long long res0 = reduce(f0.begin(), f0.end(), 0LL);
        long long res1 = reduce(f1.begin(), f1.end(), 0LL);
        return (res0 + res1) % MOD;
    }
};
```

```go [sol-Go]
func zigZagArrays(n, l, r int) (ans int) {
	const mod = 1_000_000_007
	k := r - l + 1
	f0 := make([]int, k) // 后两个数递增
	f1 := make([]int, k) // 后两个数递减
	for i := range f0 {
		f0[i] = 1
		f1[i] = 1
	}

	s0 := make([]int, k+1)
	s1 := make([]int, k+1)
	for range n - 1 {
		for j, v := range f0 {
			s0[j+1] = s0[j] + v
			s1[j+1] = s1[j] + f1[j]
		}
		for j := range f0 {
			f0[j] = s1[j] % mod
			f1[j] = (s0[k] - s0[j+1]) % mod
		}
	}

	for j, v := range f0 {
		ans += v + f1[j]
	}
	return ans % mod
}
```

## 优化一：原地计算

根据上面的代码，我们可以原地计算 $f_1$ 的**前缀和**，以及 $f_0$ 的**后缀和**。

计算后，交换 $f_0$ 和 $f_1$。

```py [sol-Python3]
class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        MOD = 1_000_000_007
        k = r - l + 1

        f0 = [1] * k  # 后两个数递增
        f1 = [1] * k  # 后两个数递减
        for _ in range(n - 1):
            pre = 0
            for i, v in enumerate(f1):
                f1[i] = pre % MOD
                pre += v

            suf = 0
            for i in range(k - 1, -1, -1):
                v = f0[i]
                f0[i] = suf % MOD
                suf += v

            f0, f1 = f1, f0

        return (sum(f0) + sum(f1)) % MOD
```

```java [sol-Java]
class Solution {
    public int zigZagArrays(int n, int l, int r) {
        final int MOD = 1_000_000_007;
        int k = r - l + 1;

        int[] f0 = new int[k]; // 后两个数递增
        int[] f1 = new int[k]; // 后两个数递减
        Arrays.fill(f0, 1);
        Arrays.fill(f1, 1);

        for (int i = 2; i <= n; i++) {
            long pre = 0;
            for (int j = 0; j < k; j++) {
                int v = f1[j];
                f1[j] = (int) (pre % MOD);
                pre += v;
            }

            long suf = 0;
            for (int j = k - 1; j >= 0; j--) {
                int v = f0[j];
                f0[j] = (int) (suf % MOD);
                suf += v;
            }

            int[] tmp = f0;
            f0 = f1;
            f1 = tmp;
        }

        long ans = 0;
        for (int j = 0; j < k; j++) {
            ans += f0[j] + f1[j];
        }
        return (int) (ans % MOD);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int zigZagArrays(int n, int l, int r) {
        constexpr int MOD = 1'000'000'007;
        int k = r - l + 1;

        vector<int> f0(k, 1); // 后两个数递增
        vector<int> f1(k, 1); // 后两个数递减

        for (int i = 2; i <= n; i++) {
            long long pre = 0;
            for (int j = 0; j < k; j++) {
                int v = f1[j];
                f1[j] = pre % MOD;
                pre += v;
            }

            long long suf = 0;
            for (int j = k - 1; j >= 0; j--) {
                int v = f0[j];
                f0[j] = suf % MOD;
                suf += v;
            }

            swap(f0, f1);
        }

        long long res0 = reduce(f0.begin(), f0.end(), 0LL);
        long long res1 = reduce(f1.begin(), f1.end(), 0LL);
        return (res0 + res1) % MOD;
    }
};
```

```go [sol-Go]
func zigZagArrays(n, l, r int) (ans int) {
	const mod = 1_000_000_007
	k := r - l + 1
	f0 := make([]int, k) // 后两个数递增
	f1 := make([]int, k) // 后两个数递减
	for i := range f0 {
		f0[i] = 1
		f1[i] = 1
	}

	for range n - 1 {
		pre := 0
		for i, v := range f1 {
			f1[i] = pre % mod
			pre += v
		}

		suf := 0
		for i := k - 1; i >= 0; i-- {
			v := f0[i]
			f0[i] = suf % mod
			suf += v
		}

		f0, f1 = f1, f0
	}

	for j, v := range f0 {
		ans += v + f1[j]
	}
	return ans % mod
}
```

## 优化二：利用对称性

根据对称性，增减增减的方案数与减增减增的方案数相等，所以只需计算增减增减的方案数，然后乘以 $2$，即为答案。

```py [sol-Python3]
class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        MOD = 1_000_000_007
        k = r - l + 1
        f = [1] * k

        for i in range(1, n):
            if i % 2:  # 增
                pre = 0
                for j, v in enumerate(f):
                    f[j] = pre % MOD
                    pre += v
            else:  # 减
                suf = 0
                for j in range(k - 1, -1, -1):
                    v = f[j]
                    f[j] = suf % MOD
                    suf += v

        return sum(f) * 2 % MOD
```

```java [sol-Java]
class Solution {
    public int zigZagArrays(int n, int l, int r) {
        final int MOD = 1_000_000_007;
        int k = r - l + 1;
        int[] f = new int[k];
        Arrays.fill(f, 1);

        for (int i = 1; i < n; i++) {
            if (i % 2 > 0) { // 增
                long pre = 0;
                for (int j = 0; j < k; j++) {
                    int v = f[j];
                    f[j] = (int) (pre % MOD);
                    pre += v;
                }
            } else { // 减
                long suf = 0;
                for (int j = k - 1; j >= 0; j--) {
                    int v = f[j];
                    f[j] = (int) (suf % MOD);
                    suf += v;
                }
            }
        }

        long ans = 0;
        for (int v : f) {
            ans += v;
        }
        return (int) (ans * 2 % MOD);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int zigZagArrays(int n, int l, int r) {
        constexpr int MOD = 1'000'000'007;
        int k = r - l + 1;
        vector<int> f(k, 1);

        for (int i = 1; i < n; i++) {
            if (i % 2) { // 增
                long long pre = 0;
                for (int j = 0; j < k; j++) {
                    int v = f[j];
                    f[j] = pre % MOD;
                    pre += v;
                }
            } else { // 减
                long long suf = 0;
                for (int j = k - 1; j >= 0; j--) {
                    int v = f[j];
                    f[j] = suf % MOD;
                    suf += v;
                }
            }
        }

        return reduce(f.begin(), f.end(), 0LL) * 2 % MOD;
    }
};
```

```go [sol-Go]
func zigZagArrays(n, l, r int) (ans int) {
	const mod = 1_000_000_007
	k := r - l + 1
	f := make([]int, k)
	for i := range f {
		f[i] = 1
	}

	for i := 1; i < n; i++ {
		if i%2 > 0 { // 增
			pre := 0
			for j, v := range f {
				f[j] = pre % mod
				pre += v
			}
		} else { // 减
			suf := 0
			for j := k - 1; j >= 0; j-- {
				v := f[j]
				f[j] = suf % mod
				suf += v
			}
		}
	}

	for _, v := range f {
		ans += v
	}
	return ans * 2 % mod
}
```

## 另一种写法

注：这个写法只是为下一道题 [3700. 锯齿形数组的总数 II](https://leetcode.cn/problems/number-of-zigzag-arrays-ii/) 做铺垫，效率可能不如上面的写法。

我们可以把要计算的数据倒着保存在 $f[k-1-j]$ 中，这样就可以每次都正序遍历了。

但我们只有一个数组，不能提前覆盖，怎么办？

可以用滚动数组。也可以先保存在 $f[j]$ 中，最后反转一下。

```py [sol-Python3]
class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        MOD = 1_000_000_007
        k = r - l + 1
        f = [1] * k

        for i in range(1, n):
            pre = 0
            for j, v in enumerate(f):
                f[j] = pre % MOD
                pre += v
            f.reverse()

        return sum(f) * 2 % MOD
```

```java [sol-Java]
class Solution {
    public int zigZagArrays(int n, int l, int r) {
        final int MOD = 1_000_000_007;
        int k = r - l + 1;
        int[] f = new int[k];
        Arrays.fill(f, 1);

        for (int i = 1; i < n; i++) {
            long pre = 0;
            for (int j = 0; j < k; j++) {
                int v = f[j];
                f[j] = (int) (pre % MOD);
                pre += v;
            }
            reverse(f);
        }

        long ans = 0;
        for (int v : f) {
            ans += v;
        }
        return (int) (ans * 2 % MOD);
    }

    private void reverse(int[] a) {
        for (int i = 0, j = a.length - 1; i < j; i++, j--) {
            int tmp = a[i];
            a[i] = a[j];
            a[j] = tmp;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int zigZagArrays(int n, int l, int r) {
        constexpr int MOD = 1'000'000'007;
        int k = r - l + 1;
        vector<int> f(k, 1);

        for (int i = 1; i < n; i++) {
            long long pre = 0;
            for (int j = 0; j < f.size(); j++) {
                int v = f[j];
                f[j] = pre % MOD;
                pre += v;
            }
            ranges::reverse(f);
        }

        return reduce(f.begin(), f.end(), 0LL) * 2 % MOD;
    }
};
```

```go [sol-Go]
func zigZagArrays(n, l, r int) (ans int) {
	const mod = 1_000_000_007
	k := r - l + 1
	f := make([]int, k)
	for i := range f {
		f[i] = 1
	}

	for i := 1; i < n; i++ {
		pre := 0
		for j, v := range f {
			f[j] = pre % mod
			pre += v
		}
		slices.Reverse(f)
	}

	for _, v := range f {
		ans += v
	}
	return ans * 2 % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(r-l))$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(r-l)$。

## 专题训练

见下面动态规划题单的「**§11.1 前缀和优化 DP**」。

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

[本题视频讲解](https://www.bilibili.com/video/BV19N411j7Dj/)

## 前置知识

如果你之前没有做过前缀和+哈希表的题目，请先做这题：[前缀和+哈希表，消除分支的技巧](https://leetcode.cn/problems/find-longest-subarray-lcci/solution/tao-lu-qian-zhui-he-ha-xi-biao-xiao-chu-3mb11/)

## 转化

把元音视作 $1$，辅音视作 $-1$。

「元音字母和辅音字母的数量相等」就等价于：找到一个和为 $0$ 的连续子数组。注意这种子数组的长度一定是偶数，因为元音辅音数量相等。

设子数组的长度为 $L$。由于元音辅音数量相等，所以元音辅音数量都等于 $\dfrac{L}{2}$，所以「元音字母和辅音字母的数量的乘积能被 $k$ 整除」等价于

$$
\left(\dfrac{L}{2}\right)^2 \bmod k = 0
$$

这等价于

$$
L^2 \bmod (4k) = 0
$$

这个平方很烦人，如果能去掉平方就好做了。

## 来点数学

我们来研究下，如果一个数 $L$ 的平方能被 $n$ 整除，意味着什么？

假设 $n$ 是一个质数，例如 $3$，那么 $L$ 必须包含质因子 $3$，此时题目约束就变成了：$L$ 是 $3$ 的倍数。我们把平方去掉了！

如果 $n$ 是一个质数 $p$ 的 $e$ 次幂呢？分类讨论：

- 如果 $e$ 是偶数，比如 $n=3^4$，那么 $L$ 必须包含因子 $3^2$，才能使得 $L^2$ 能被 $n$ 整除。此时题目约束就变成了：$L$ 是 $3^2$ 的倍数。
- 如果 $e$ 是奇数，比如 $n=3^5$，那么 $L$ 必须包含因子 $3^3$，才能使得 $L^2$ 能被 $n$ 整除。此时题目约束就变成了：$L$ 是 $3^3$ 的倍数。

所以 $L$ 必须包含因子 $p^r$，其中 $r=\left\lceil\dfrac{e}{2}\right\rceil = \left\lfloor\dfrac{e+1}{2}\right\rfloor$。

如果 $n$ 可以分解出多个质因子，只需要把每个质因子及其幂次按照上面的方法处理，把结果相乘，就得到 $L$ 必须是什么数的倍数了。

这样就把平方去掉了。

## 套路：前缀和+哈希表

把 $4k$ 按照上述方法计算，设 $L$ 必须是 $k'$ 的倍数。

现在问题变成，有多少个和为 $0$ 的子数组，其长度是 $k'$ 的倍数？

设子数组的下标范围为 $[j,i)$，那么其长度 $L=i-j$，则有

$$
(i-j)\bmod k' = 0
$$

等价于

$$
i \bmod k' = j\bmod k'
$$

对于前缀和来说（定义见最上面贴的题解），子数组和为 $0$ 相当于 $s[i]-s[j] = 0$，即

$$
s[i] = s[j]
$$

我们需要**同时满足**这两个条件（下标模 $k'$ 相等，$s$ 值相等），这可以一并用哈希表解决。

哈希表的 key 是一个 pair：$(i\bmod k', s[i])$，哈希表的 value 是这个 pair 的出现次数。

代码实现时，前缀和数组可以用一个变量表示。

代码实现时，可以把 aeiou 压缩成一个二进制数，从而快速判断字母是否为元音，原理请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

#### 答疑

**问**：为什么哈希表要在一开始插入一个 $(k'-1, 0)$？

**答**：前缀和的第一项是 $0$，由于代码中是从下标 $0$ 开始算第二个前缀和，所以相当于 $s[-1] = 0$，而 $k'-1$ 和 $-1$ 关于 $k'$ 同余，所以插入 $(k'-1, 0)$。

```py [sol-Python3]
class Solution:
    def beautifulSubstrings(self, s: str, k: int) -> int:
        k = self.sqrt(k * 4)
        cnt = Counter([(k - 1, 0)])  # k-1 和 -1 同余
        ans = pre_sum = 0
        for i, c in enumerate(s):
            pre_sum += 1 if c in "aeiou" else -1
            p = (i % k, pre_sum)
            ans += cnt[p]
            cnt[p] += 1
        return ans

    def sqrt(self, n: int) -> int:
        res = 1
        i = 2
        while i * i <= n:
            i2 = i * i
            while n % i2 == 0:
                res *= i
                n //= i2
            if n % i == 0:
                res *= i
                n //= i
            i += 1
        if n > 1:
            res *= n
        return res
```

```java [sol-Java]
class Solution {
    private static final int AEIOU_MASK = 1065233;

    public int beautifulSubstrings(String s, int k) {
        k = pSqrt(k * 4);
        Map<Integer, Integer> cnt = new HashMap<>();
        int n = s.length();
        int sum = n; // 保证非负
        cnt.put((k - 1) << 16 | sum, 1); // 添加 (k-1, sum)
        int ans = 0;
        for (int i = 0; i < n; i++) {
            int bit = (AEIOU_MASK >> (s.charAt(i) - 'a')) & 1;
            sum += bit * 2 - 1; // 1 -> 1    0 -> -1
            ans += cnt.merge((i % k) << 16 | sum, 1, Integer::sum) - 1; // ans += cnt[(i%k,sum)]++
        }
        return ans;
    }

    private int pSqrt(int n) {
        int res = 1;
        for (int i = 2; i * i <= n; i++) {
            int i2 = i * i;
            while (n % i2 == 0) {
                res *= i;
                n /= i2;
            }
            if (n % i == 0) {
                res *= i;
                n /= i;
            }
        }
        if (n > 1) {
            res *= n;
        }
        return res;
    }
}
```

```cpp [sol-C++ map]
class Solution {
    int p_sqrt(int n) {
        int res = 1;
        for (int i = 2; i * i <= n; i++) {
            int i2 = i * i;
            while (n % i2 == 0) {
                res *= i;
                n /= i2;
            }
            if (n % i == 0) {
                res *= i;
                n /= i;
            }
        }
        if (n > 1) {
            res *= n;
        }
        return res;
    }

    const int AEIOU_MASK = 1065233;

public:
    int beautifulSubstrings(string s, int k) {
        k = p_sqrt(k * 4);
        // 把 pair 压缩成 long long（或者 int）就可以用 umap 了
        map<pair<int, int>, int> cnt;
        cnt[{k - 1, 0}]++; // 添加 (k-1, sum)
        int ans = 0;
        int sum = 0;
        for (int i = 0; i < s.length(); i++) {
            int bit = (AEIOU_MASK >> (s[i] - 'a')) & 1;
            sum += bit * 2 - 1; // 1 -> 1    0 -> -1
            ans += cnt[{i % k, sum}]++;
        }
        return ans;
    }
};
```

```cpp [sol-C++ umap]
class Solution {
    int p_sqrt(int n) {
        int res = 1;
        for (int i = 2; i * i <= n; i++) {
            int i2 = i * i;
            while (n % i2 == 0) {
                res *= i;
                n /= i2;
            }
            if (n % i == 0) {
                res *= i;
                n /= i;
            }
        }
        if (n > 1) {
            res *= n;
        }
        return res;
    }

    const int AEIOU_MASK = 1065233;

public:
    int beautifulSubstrings(string s, int k) {
        k = p_sqrt(k * 4);
        unordered_map<int, int> cnt;
        int n = s.length();
        int sum = n; // 保证非负
        cnt[(k - 1) << 16 | sum]++; // 添加 (k-1, sum)
        int ans = 0;
        for (int i = 0; i < n; i++) {
            int bit = (AEIOU_MASK >> (s[i] - 'a')) & 1;
            sum += bit * 2 - 1; // 1 -> 1    0 -> -1
            ans += cnt[(i % k) << 16 | sum]++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func pSqrt(n int) int {
	res := 1
	for i := 2; i*i <= n; i++ {
		i2 := i * i
		for n%i2 == 0 {
			res *= i
			n /= i2
		}
		if n%i == 0 {
			res *= i
			n /= i
		}
	}
	if n > 1 {
		res *= n
	}
	return res
}

func beautifulSubstrings(s string, k int) (ans int) {
	k = pSqrt(k * 4)

	type pair struct{ i, sum int }
	cnt := map[pair]int{{k - 1, 0}: 1} // k-1 和 -1 同余
	sum := 0
	const aeiouMask = 1065233
	for i, c := range s {
		bit := aeiouMask >> (c - 'a') & 1
		sum += bit*2 - 1 // 1 -> 1    0 -> -1
		p := pair{i % k, sum}
		ans += cnt[p]
		cnt[p]++
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + \sqrt k)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 强化训练：前缀和+哈希表

推荐按顺序做。

- [560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/)
- [974. 和可被 K 整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/)
- [1590. 使数组和能被 P 整除](https://leetcode.cn/problems/make-sum-divisible-by-p/)
- [523. 连续的子数组和](https://leetcode.cn/problems/continuous-subarray-sum/)
- [525. 连续数组](https://leetcode.cn/problems/contiguous-array/)
- [面试题 17.05. 字母与数字](https://leetcode.cn/problems/find-longest-subarray-lcci/)
- [1915. 最美子字符串的数目](https://leetcode.cn/problems/number-of-wonderful-substrings/)
- [930. 和相同的二元子数组](https://leetcode-cn.com/problems/binary-subarrays-with-sum/)
- [1371. 每个元音包含偶数次的最长子字符串](https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/)
- [1542. 找出最长的超赞子字符串](https://leetcode-cn.com/problems/find-longest-awesome-substring/)

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

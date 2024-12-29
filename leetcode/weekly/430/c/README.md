## 方法一：前后缀分解

为方便描述，下文把 $\textit{nums}[p],\textit{nums}[q],\textit{nums}[r],\textit{nums}[s]$ 分别简称为 $a,b,c,d$。

题目要求

$$
a\cdot c = b\cdot d
$$

将其变形为

$$
\dfrac{a}{b} = \dfrac{d}{c}
$$

这样 $a$ 和 $b$ 都在 $c$ 和 $d$ 的左边，从而方便用**前后缀分解**解决。

首先统计下标 $[4,n-1]$ 中的（间隔至少一个数的）数对 $(c,d)$ 的**最简分数**的个数，记录到一个哈希表 $\textit{suf}$ 中。

> **最简分数**：分子和分母互质的分数。如果分子和分母不互质，可以除以二者的最大公约数（GCD）。比如分数 $4/6$，分子和分母都除以二者的最大公约数 $2$，得到最简分数 $2/3$。

然后枚举 $b$，以及 $b$ 左边的（间隔至少一个数的）$a$。计算最简分数 $a'/b'$，去 $\textit{suf}$ 中查找 $a'/b'$ 的个数，加入答案。

从左向右枚举 $b$ 的过程中，维护（撤销）$\textit{suf}$ 中的最简分数个数。

具体请看 [视频讲解](https://www.bilibili.com/video/BV13f68YjE7o/?t=17m10s)（先讲的方法二），欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def numberOfSubsequences(self, nums: List[int]) -> int:
        n = len(nums)
        suf = defaultdict(int)
        # 枚举 c
        for i in range(4, n - 2):
            c = nums[i]
            # 枚举 d
            for d in nums[i + 2:]:
                g = gcd(c, d)
                suf[d // g, c // g] += 1

        ans = 0
        # 枚举 b
        for i in range(2, n - 4):
            b = nums[i]
            # 枚举 a
            for a in nums[:i - 1]:
                g = gcd(a, b)
                ans += suf[a // g, b // g]
            # 撤销之前统计的 d'/c'
            c = nums[i + 2]
            for d in nums[i + 4:]:
                g = gcd(c, d)
                suf[d // g, c // g] -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public long numberOfSubsequences(int[] nums) {
        int n = nums.length;
        Map<Integer, Integer> suf = new HashMap<>();
        // 枚举 c
        for (int i = 4; i < n - 2; i++) {
            int c = nums[i];
            // 枚举 d
            for (int j = i + 2; j < n; j++) {
                int d = nums[j];
                int g = gcd(c, d);
                // 把分子和分母（两个 short）压缩成一个 int
                suf.merge((d / g) << 16 | (c / g), 1, Integer::sum);
            }
        }

        long ans = 0;
        // 枚举 b
        for (int i = 2; i < n - 4; i++) {
            int b = nums[i];
            // 枚举 a
            for (int j = 0; j < i - 1; j++) {
                int a = nums[j];
                int g = gcd(a, b);
                ans += suf.getOrDefault((a / g) << 16 | (b / g), 0);
            }
            // 撤销之前统计的 d'/c'
            int c = nums[i + 2];
            for (int j = i + 4; j < n; j++) {
                int d = nums[j];
                int g = gcd(c, d);
                suf.merge((d / g) << 16 | (c / g), -1, Integer::sum);
            }
        }
        return ans;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```java [sol-Java 写法二]
class Solution {
    private record Pair(int x, int y) {}

    public long numberOfSubsequences(int[] nums) {
        int n = nums.length;
        Map<Pair, Integer> suf = new HashMap<>();
        // 枚举 c
        for (int i = 4; i < n - 2; i++) {
            int c = nums[i];
            // 枚举 d
            for (int j = i + 2; j < n; j++) {
                int d = nums[j];
                int g = gcd(c, d);
                suf.merge(new Pair(d / g, c / g), 1, Integer::sum);
            }
        }

        long ans = 0;
        // 枚举 b
        for (int i = 2; i < n - 4; i++) {
            int b = nums[i];
            // 枚举 a
            for (int j = 0; j < i - 1; j++) {
                int a = nums[j];
                int g = gcd(a, b);
                ans += suf.getOrDefault(new Pair(a / g, b / g), 0);
            }
            // 撤销之前统计的 d'/c'
            int c = nums[i + 2];
            for (int j = i + 4; j < n; j++) {
                int d = nums[j];
                int g = gcd(c, d);
                suf.merge(new Pair(d / g, c / g), -1, Integer::sum);
            }
        }
        return ans;
    }

    private int gcd(int a, int b) {
        while (a != 0) {
            int tmp = a;
            a = b % a;
            b = tmp;
        }
        return b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long numberOfSubsequences(vector<int>& nums) {
        int n = nums.size();
        unordered_map<int, int> suf;
        // 枚举 c
        for (int i = 4; i < n - 2; i++) {
            int c = nums[i];
            // 枚举 d
            for (int j = i + 2; j < n; j++) {
                int d = nums[j];
                int g = gcd(c, d);
                // 把分子和分母（两个 short）压缩成一个 int
                suf[(d / g) << 16 | (c / g)]++;
            }
        }

        long long ans = 0;
        // 枚举 b
        for (int i = 2; i < n - 4; i++) {
            int b = nums[i];
            // 枚举 a
            for (int j = 0; j < i - 1; j++) {
                int a = nums[j];
                int g = gcd(a, b);
                ans += suf[(a / g) << 16 | (b / g)];
            }
            // 撤销之前统计的 d'/c'
            int c = nums[i + 2];
            for (int j = i + 4; j < n; j++) {
                int d = nums[j];
                int g = gcd(c, d);
                suf[(d / g) << 16 | (c / g)]--;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfSubsequences(nums []int) (ans int64) {
	n := len(nums)
	type pair struct{ x, y int }
	suf := map[pair]int{}
	// 枚举 c
	for i := 4; i < n-2; i++ {
		c := nums[i]
		// 枚举 d
		for _, d := range nums[i+2:] {
			g := gcd(c, d)
			suf[pair{d / g, c / g}]++
		}
	}

	// 枚举 b
	for i := 2; i < n-4; i++ {
		b := nums[i]
		// 枚举 a
		for _, a := range nums[:i-1] {
			g := gcd(a, b)
			ans += int64(suf[pair{a / g, b / g}])
		}
		// 撤销之前统计的 d'/c'
		c := nums[i+2]
		for _, d := range nums[i+4:] {
			g := gcd(c, d)
			suf[pair{d / g, c / g}]--
		}
	}
	return
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。计算一次 GCD 需要 $\mathcal{O}(\log U)$ 的时间。
- 空间复杂度：$\mathcal{O}(\min(n^2,U^2))$。由**欧拉函数**的平均值 $\dfrac{6}{\pi^2}$ 可知，互质数对是很常见的。在 $[1,U]$ 中任选两个数，可以组成 $\mathcal{O}(U^2)$ 个不同的最简分数。所以最简分数的个数并不少，总体来说是 $\mathcal{O}(\min(n^2,U^2))$ 的，这决定了哈希表的大小。

## 方法二：枚举右，维护左

推荐先完成两个数的版本：[1512. 好数对的数目](https://leetcode.cn/problems/number-of-good-pairs/)。

在枚举右边的 $c$ 和 $d$ 的同时，用哈希表维护左边的 $\dfrac{a}{b}$ 的个数。

本题由于数据范围只有 $[1,1000]$，在这个范围内比较分数是否相等，是无误的，所以也可以直接用浮点数计算。

### 什么情况下用浮点数是错的？

取两个接近 $1$ 但不相同的分数 $\dfrac{a}{a+1}$ 和 $\dfrac{a-1}{a}$，根据 IEEE 754，在使用双精度浮点数的情况下，如果这两个数的绝对差 $\dfrac{1}{a(a+1)}$ 比 $2^{-52}$ 还小，那么计算机可能会把这两个数舍入到同一个附近的浮点数上。所以当 $a$ 达到 $2^{26}$ 的时候，算法就可能有问题了。本题只有 $1000$，可以放心地使用浮点数除法。

如果用单精度浮点数，就是当 $a$ 达到 $2^{11.5}$ 时才会有问题。所以本题用单精度浮点数也可以。

读者可以做做 [2280. 表示一个折线图的最少线段数](https://leetcode.cn/problems/minimum-lines-to-represent-a-line-chart/) 感受一下。

```py [sol-Python3]
class Solution:
    def numberOfSubsequences(self, nums: List[int]) -> int:
        ans = 0
        cnt = defaultdict(int)
        # 枚举 b 和 c
        for i in range(4, len(nums) - 2):
            # 增量式更新，本轮循环只需枚举 b=nums[i-2] 这一个数
            # 至于更前面的 b，已经在前面的循环中添加到 cnt 中了，不能重复添加
            b = nums[i - 2]
            # 枚举 a
            for a in nums[:i - 3]:
                cnt[a / b] += 1

            c = nums[i]
            # 枚举 d
            for d in nums[i + 2:]:
                ans += cnt[d / c]
        return ans
```

```java [sol-Java]
class Solution {
    public long numberOfSubsequences(int[] nums) {
        int n = nums.length;
        long ans = 0;
        Map<Float, Integer> cnt = new HashMap<>();
        // 枚举 b 和 c
        for (int i = 4; i < n - 2; i++) {
            // 增量式更新，本轮循环只需枚举 b=nums[i-2] 这一个数
            // 至于更前面的 b，已经在前面的循环中添加到 cnt 中了，不能重复添加
            float b = nums[i - 2];
            // 枚举 a
            for (int j = 0; j < i - 3; j++) {
                cnt.merge(nums[j] / b, 1, Integer::sum);
            }

            float c = nums[i];
            // 枚举 d
            for (int j = i + 2; j < n; j++) {
                ans += cnt.getOrDefault(nums[j] / c, 0);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long numberOfSubsequences(vector<int>& nums) {
        int n = nums.size();
        long long ans = 0;
        unordered_map<float, int> cnt;
        // 枚举 b 和 c
        for (int i = 4; i < n - 2; i++) {
            // 增量式更新，本轮循环只需枚举 b=nums[i-2] 这一个数
            // 至于更前面的 b，已经在前面的循环中添加到 cnt 中了，不能重复添加
            float b = nums[i - 2];
            // 枚举 a
            for (int j = 0; j < i - 3; j++) {
                cnt[nums[j] / b]++;
            }

            float c = nums[i];
            // 枚举 d
            for (int j = i + 2; j < n; j++) {
                ans += cnt[nums[j] / c];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func numberOfSubsequences(nums []int) (ans int64) {
	n := len(nums)
	cnt := map[float32]int{}
	// 枚举 b 和 c
	for i := 4; i < n-2; i++ {
		// 增量式更新，本轮循环只需枚举 b=nums[i-2] 这一个数
		// 至于更前面的 b，已经在前面的循环中添加到 cnt 中了，不能重复添加
		b := float32(nums[i-2])
		// 枚举 a
		for _, a := range nums[:i-3] {
			cnt[float32(a)/b]++
		}

		c := float32(nums[i])
		// 枚举 d
		for _, d := range nums[i+2:] {
			ans += int64(cnt[float32(d)/c])
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(\min(n^2,U^2))$。见方法一的复杂度分析。

更多相似题目，见下面动态规划题单中的「**专题：前后缀分解**」，以及数据结构题单中的「**§0.1 枚举右，维护左**」。

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

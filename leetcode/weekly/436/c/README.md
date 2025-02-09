## 前言

如何把一个字符串转成 $\texttt{int}$ 数字？

比如字符串 $s=\texttt{123}$，我们可以从左到右遍历 $s$，初始化 $x=0$，利用

$$
x = x\cdot 10 + \texttt{int}(s_i)
$$

不断迭代 $x$，得到 $0\to 1\to 12\to 123$。

## 公式推导

下文把 $s_i$ 当作数字 $0$ 到 $9$。

把子串按照末尾数字 $s_i$ 分组，那么：

- 以 $s_i$ 结尾的子串，相当于在所有以 $s_{i-1}$ 结尾的子串的末尾添加 $s_i$。以及 $s_i$ 单独算作一个子串。

题目让我们计算是否整除，也就是模 $s_i$ 的余数是否为 $0$。用数学公式描述，就是

$$
\textit{val}_i \bmod s_i = 0
$$

其中 $\textit{val}_{i}$ 表示以 $s_i$ 结尾的子串的数值。（可能有多个，这里只是表示其中任意一个。）

根据前言中的公式，我们有

$$
\textit{val}_i = \textit{val}_{i-1}\cdot 10 + s_i
$$

代入得

$$
(\textit{val}_{i-1}\cdot 10 + s_i) \bmod s_i = 0
$$

根据 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/) 中的恒等式，上式等价于

$$
((\textit{val}_{i-1}\bmod s_i)\cdot 10 + s_i) \bmod s_i = 0
$$

这意味着，我们只需要知道以 $s_{i-1}$ 结尾的子串的数值模 $s_i$ 的结果 $\textit{rem}_{i-1}$，并不需要知道 $\textit{val}_{i-1}$ 具体是多少。所以将上式变为

$$
(\textit{rem}_{i-1}\cdot 10 + s_i) \bmod s_i = 0
$$

## 动态规划

根据上面的分析，定义 $f[i+1][m][\textit{rem}]$ 表示以 $s_i$ 结尾的、模 $m$ 余数为 $\textit{rem}$ 的子串个数。其中 $m=1,2,3\ldots,9$。

考虑用刷表法转移，在以 $s_{i-1}$ 结尾的、模 $m$ 余数为 $\textit{rem}$ 的子串末尾添加 $s_i$，可以得到以 $s_i$ 结尾的、模 $m$ 余数为 $(\textit{rem}\cdot 10 + s_i)\bmod m$ 的子串，所以有

$$
f[i+1][m][(\textit{rem}\cdot 10 + s_i)\bmod m] += f[i][m][\textit{rem}]
$$

> 注：在动态规划中，用转移来源更新当前状态叫**查表法**，用当前状态更新其他状态叫**刷表法**。本题在已知余数的情况下，并不好计算从哪些状态转移过来，但是从当前状态转移到哪些状态是很好计算的，所以用刷表法。

初始值 $f[i+1][m][s_i\bmod m] = 1$，对应着 $s_i$ 单独形成的子串。

累加以 $s_i$ 结尾的、模 $s_i$ 余数为 $0$ 的子串个数，即为答案：

$$
\sum_{i=0}^{n-1} f[i+1][s_i][0]
$$

代码实现时，$f$ 第一个维度可以去掉，用滚动数组优化。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ekN2ebEHx/?t=27m05s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countSubstrings(self, s: str) -> int:
        ans = 0
        f = [[0] * 9 for _ in range(10)]
        for d in map(int, s):
            for m in range(1, 10):  # 枚举模数 m
                # 滚动数组计算 f
                nf = [0] * m
                nf[d % m] = 1
                for rem in range(m):  # 枚举模 m 的余数 rem
                    nf[(rem * 10 + d) % m] += f[m][rem]  # 刷表法
                f[m] = nf
            # 以 s[i] 结尾的，模 s[i] 余数为 0 的子串个数
            ans += f[d][0]
        return ans
```

```java [sol-Java]
class Solution {
    public long countSubstrings(String s) {
        long ans = 0;
        int[][] f = new int[10][9];
        for (char d : s.toCharArray()) {
            d -= '0';
            for (int m = 1; m < 10; m++) { // 枚举模数 m
                // 滚动数组计算 f
                int[] nf = new int[m];
                nf[d % m] = 1;
                for (int rem = 0; rem < m; rem++) { // 枚举模 m 的余数 rem
                    nf[(rem * 10 + d) % m] += f[m][rem]; // 刷表法
                }
                f[m] = nf;
            }
            // 以 s[i] 结尾的，模 s[i] 余数为 0 的子串个数
            ans += f[d][0];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countSubstrings(string s) {
        long long ans = 0;
        array<int, 9> f[10]{};
        for (char d : s) {
            d -= '0';
            for (int m = 1; m < 10; m++) { // 枚举模数 m
                // 滚动数组计算 f
                array<int, 9> nf{};
                nf[d % m] = 1;
                for (int rem = 0; rem < m; rem++) { // 枚举模 m 的余数 rem
                    nf[(rem * 10 + d) % m] += f[m][rem]; // 刷表法
                }
                f[m] = nf;
            }
            // 以 s[i] 结尾的，模 s[i] 余数为 0 的子串个数
            ans += f[d][0];
        }
        return ans;
    }
};
```

```go [sol-Go]
func countSubstrings(s string) (ans int64) {
	f := [10][9]int{}
	for _, c := range s {
		d := int(c - '0')
		for m := 1; m < 10; m++ { // 枚举模数 m
			// 滚动数组计算 f
			nf := [9]int{}
			nf[d%m] = 1
			for rem, fv := range f[m][:m] { // 枚举模 m 的余数 rem
				nf[(rem*10+d)%m] += fv // 刷表法
			}
			f[m] = nf
		}
		// 以 s[i] 结尾的，模 s[i] 余数为 0 的子串个数
		ans += int64(f[d][0])
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nD^2)$，其中 $n$ 是 $s$ 的长度，$D=9$。
- 空间复杂度：$\mathcal{O}(D^2)$。

更多相似题目，见下面动态规划题单中的「**§7.5 多维 DP**」。

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

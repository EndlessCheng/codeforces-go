## 方法一：贡献法

考虑每个 $s[i]$ 的贡献。

对于重复字符，人为规定，**只考虑第一个字符的贡献**。比如子串中有多个字母 $\texttt{a}$，我们只计算其中第一个字母 $\texttt{a}$ 的贡献。

设 $s[i]$ 上一次出现的下标为 $\textit{last}[s[i]]$。如果不存在则为 $-1$。

那么子串的左端点可以从 $\textit{last}[s[i]]+1$ 到 $i$，子串的右端点可以从 $i$ 到 $n-1$。

所以 $s[i]$ 的贡献是

$$
(i-\textit{last}[s[i]]) \cdot (n-i)。
$$

累加即为答案。

```py [sol-Python3]
class Solution:
    def appealSum(self, s: str) -> int:
        last = {}
        n = len(s)
        ans = 0
        for i, c in enumerate(s):
            ans += (i - last.get(c, -1)) * (n - i)
            last[c] = i
        return ans
```

```java [sol-Java]
class Solution {
    public long appealSum(String s) {
        int[] last = new int[26];
        Arrays.fill(last, -1);
        int n = s.length();
        long ans = 0;
        for (int i = 0; i < n; i++) {
            int c = s.charAt(i) - 'a';
            ans += (long) (i - last[c]) * (n - i);
            last[c] = i;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long appealSum(string s) {
        vector<int> last(26, -1);
        int n = s.size();
        long long ans = 0;
        for (int i = 0; i < n; i++) {
            int c = s[i] - 'a';
            ans += 1LL * (i - last[c]) * (n - i);
            last[c] = i;
        }
        return ans;
    }
};
```

```go [sol-Go]
func appealSum(s string) (ans int64) {
	last := [26]int{}
	for i := range last {
		last[i] = -1 // 初始化成 -1 可以让提示 2-2 中的两种情况合并成一个公式
	}
	n := len(s)
	for i, c := range s {
		c -= 'a'
		ans += int64(i-last[c]) * int64(n-i)
		last[c] = i
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$，其中 $|\Sigma|=26$ 是字符集合的大小。

## 方法二：增量法

#### 提示 1-1

将**所有子串**按照其末尾字符的下标分组。

#### 提示 1-2

考虑两组**相邻**的子串：以 $s[i-1]$ 结尾的子串、以 $s[i]$ 结尾的子串。

#### 提示 1-3

以 $s[i]$ 结尾的子串，可以看成是以 $s[i-1]$ 结尾的子串，在末尾添加上 $s[i]$ 组成。

**上面这一串提示是思考子串统计类问题的通用技巧之一。**

#### 提示 2-1

从左往右遍历 $s$，考虑将 $s[i]$ 添加到以 $s[i-1]$ 结尾的子串的末尾。添加后，这些以 $s[i-1]$ 结尾的子串的引力值会增加多少？

#### 提示 2-2

分类讨论：

- 如果 $s[i]$ 之前没有遇到过，那么这些子串的引力值都会增加 $1$，这些子串的引力值之和会增加 $i$，再加上 $1$，即 $s[i]$ 单独组成的子串的引力值；
- 如果 $s[i]$ 之前遇到过，设其上次出现的下标为 $j$，那么向子串 $s[0..i-1],\ s[1..i-1],\ s[2..i-1],\cdots,s[j..i-1]$ 的末尾添加 $s[i]$ 后，这些子串的引力值是不会变化的，因为 $s[i]$ 已经在 $s[j]$ 处出现过了；而子串 $s[j+1..i-1],\ s[j+2..i-1],\cdots,s[i-1..i-1]$ 由于不包含字符 $s[i]$，这些子串的引力值都会增加 $1$，因此有 $i-j-1$ 个子串的引力值会增加 $1$，这些子串的引力值之和会增加 $i-j-1$，再加上 $1$，即 $s[i]$ 单独组成的子串的引力值。

#### 提示 2-3

模拟上述过程，遍历 $s$ 的过程中用一个变量 $\textit{sumG}$ 维护以 $s[i]$ 结尾的子串的引力值之和，同时用一个数组或哈希表 $\textit{last}$ 记录每个字符上次出现的下标。

累加遍历中的 $\textit{sumG}$ 即为答案（根据提示 1-1）。

```py [sol-Python3]
class Solution:
    def appealSum(self, s: str) -> int:
        last = {}
        ans = sum_g = 0
        for i, c in enumerate(s):
            sum_g += i - last.get(c, -1)  # 将提示 2-2 中的两种情况合并成一个公式
            ans += sum_g
            last[c] = i
        return ans
```

```java [sol-Java]
class Solution {
    public long appealSum(String s) {
        int[] last = new int[26];
        Arrays.fill(last, -1); // 初始化成 -1 可以让提示 2-2 中的两种情况合并成一个公式
        long ans = 0;
        int sumG = 0;
        for (int i = 0; i < s.length(); i++) {
            int c = s.charAt(i) - 'a';
            sumG += i - last[c];
            ans += sumG;
            last[c] = i;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long appealSum(string s) {
        vector<int> last(26, -1); // 初始化成 -1 可以让提示 2-2 中的两种情况合并成一个公式
        long long ans = 0;
        int sum_g = 0;
        for (int i = 0; i < s.size(); i++) {
            int c = s[i] - 'a';
            sum_g += i - last[c];
            ans += sum_g;
            last[c] = i;
        }
        return ans;
    }
};
```

```go [sol-Go]
func appealSum(s string) (ans int64) {
	last := [26]int{}
	for i := range last {
		last[i] = -1 // 初始化成 -1 可以让提示 2-2 中的两种情况合并成一个公式
	}
	sumG := 0
	for i, c := range s {
		c -= 'a'
		sumG += i - last[c]
		ans += int64(sumG)
		last[c] = i
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$，其中 $|\Sigma|=26$ 是字符集合的大小。

## 专题训练

见下面动态规划题单的「**§7.3 子数组 DP**」的**思维扩展**。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

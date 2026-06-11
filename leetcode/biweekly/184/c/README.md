## 方法一：贪心

由于 $\texttt{1}$ 只能向左移且只能移动一次，所以每个 $\texttt{0111}\ldots \texttt{11}$ 可以视作一段。例如 $\texttt{01111}$ 的某种移动方式为 $\texttt{11101}$，左边三个 $\texttt{1}$ 移动一次，最右边的 $\texttt{1}$ 不动。

每一段恰好有一个 $\textit{nums}[i]$ 不能选，其余 $\textit{nums}[i]$ 全选，所以贪心地，不选最小的 $\textit{nums}[i]$。具体地，遍历每个 $\texttt{0111}\ldots \texttt{11}$ 段，累加其中的 $\textit{nums}[i]$，在遍历到段的末尾时，减去其中最小的 $\textit{nums}[i]$。

特别地，如果 $s[0] = \texttt{1}$，那么第一段可以全选，最小值可以视作 $0$（题目保证 $\textit{nums}[i]$ 非负）。

[本题视频讲解](https://www.bilibili.com/video/BV1mREx6KEp1/?t=5m55s)，欢迎点赞关注~

### 写法一：正序遍历

```py [sol-Python3]
class Solution:
    def maxTotal(self, nums: list[int], s: str) -> int:
        n = len(s)
        ans = mn = 0
        for i, x in enumerate(nums):
            if s[i] == '1':
                ans += x
                mn = min(mn, x)  # 维护这一段的最小值
                if i == n - 1 or s[i + 1] == '0':  # 遍历到了段的末尾
                    ans -= mn  # 段的最小值不选
            elif i < n - 1 and s[i + 1] == '1':  # 0111..11 段的开头
                ans += x
                mn = x
        return ans
```

```java [sol-Java]
class Solution {
    public long maxTotal(int[] nums, String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        long ans = 0;
        int mn = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (s[i] == '1') {
                ans += x;
                mn = Math.min(mn, x); // 维护这一段的最小值
                if (i == n - 1 || s[i + 1] == '0') { // 遍历到了段的末尾
                    ans -= mn; // 段的最小值不选
                }
            } else if (i < n - 1 && s[i + 1] == '1') { // 0111..11 段的开头
                ans += x;
                mn = x;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxTotal(vector<int>& nums, string s) {
        int n = s.size();
        long long ans = 0;
        int mn = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i];
            if (s[i] == '1') {
                ans += x;
                mn = min(mn, x); // 维护这一段的最小值
                if (i == n - 1 || s[i + 1] == '0') { // 遍历到了段的末尾
                    ans -= mn; // 段的最小值不选
                }
            } else if (i < n - 1 && s[i + 1] == '1') { // 0111..11 段的开头
                ans += x;
                mn = x;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxTotal(nums []int, s string) int64 {
	n := len(s)
	ans, mn := 0, 0
	for i, x := range nums {
		if s[i] == '1' {
			ans += x
			mn = min(mn, x) // 维护这一段的最小值
			if i == n-1 || s[i+1] == '0' { // 遍历到了段的末尾
				ans -= mn // 段的最小值不选
			}
		} else if i < n-1 && s[i+1] == '1' { // 0111..11 段的开头
			ans += x
			mn = x
		}
	}
	return int64(ans)
}
```

### 写法二：倒序遍历

```py [sol-Python3]
class Solution:
    def maxTotal(self, nums: list[int], s: str) -> int:
        ans = 0
        mn = inf
        for i in range(len(nums) - 1, -1, -1):
            mn = min(mn, nums[i])
            ans += nums[i]
            if s[i] == '0':
                ans -= mn
                mn = inf
        return ans
```

```java [sol-Java]
class Solution {
    public long maxTotal(int[] nums, String s) {
        long ans = 0;
        int mn = Integer.MAX_VALUE;
        for (int i = nums.length - 1; i >= 0; i--) {
            mn = Math.min(mn, nums[i]);
            ans += nums[i];
            if (s.charAt(i) == '0') {
                ans -= mn;
                mn = Integer.MAX_VALUE;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxTotal(vector<int>& nums, string s) {
        long long ans = 0;
        int mn = INT_MAX;
        for (int i = nums.size() - 1; i >= 0; i--) {
            mn = min(mn, nums[i]);
            ans += nums[i];
            if (s[i] == '0') {
                ans -= mn;
                mn = INT_MAX;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxTotal(nums []int, s string) int64 {
	ans, mn := 0, math.MaxInt
	for i := len(nums) - 1; i >= 0; i-- {
		mn = min(mn, nums[i])
		ans += nums[i]
		if s[i] == '0' {
			ans -= mn
			mn = math.MaxInt
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：状态机 DP

定义 $f[i+1][0/1]$ 表示在前缀 $[0,i]$ 中选数字，能选到的元素和的最大值。第二个维度 $j=0$ 表示没有标记从 $i+1$ 移动到 $i$，$j=1$ 表示有标记从 $i+1$ 移动到 $i$。

分类讨论：

- 如果 $s[i] = \texttt{0}$。
    - 如果没有标记从 $i+1$ 移动到 $i$，那么 $\textit{nums}[i]$ 不能选，问题变成在前缀 $[0,i-1]$ 中选数字，能选到的元素和的最大值，且没有标记从 $i$ 移动到 $i-1$，即 $f[i][0]$。转移方程为 $f[i+1][0] = f[i][0]$。
    - 如果有标记从 $i+1$ 移动到 $i$，那么 $\textit{nums}[i]$ 一定选，问题变成在前缀 $[0,i-1]$ 中选数字，能选到的元素和的最大值，且没有标记从 $i$ 移动到 $i-1$，即 $f[i][0]$。转移方程为 $f[i+1][1] = f[i][0] + \textit{nums}[i]$。
- 如果 $s[i] = \texttt{1}$。
    - 如果没有标记从 $i+1$ 移动到 $i$。
        - 选 $\textit{nums}[i]$（不移动标记 $i$），问题变成在前缀 $[0,i-1]$ 中选数字，能选到的元素和的最大值，且没有标记从 $i$ 移动到 $i-1$，即 $f[i][0]$，加上 $\textit{nums}[i]$。
        - 不选 $\textit{nums}[i]$（移动标记 $i$），问题变成在前缀 $[0,i-1]$ 中选数字，能选到的元素和的最大值，且有标记从 $i$ 移动到 $i-1$，即 $f[i][1]$。
        - 二者取最大值，转移方程为 $f[i+1][0] = \max(f[i][0] + \textit{nums}[i], f[i][1])$。
    - 如果有标记从 $i+1$ 移动到 $i$，那么 $\textit{nums}[i]$ 一定选，且标记 $i$ 移动一定更优，问题变成在前缀 $[0,i-1]$ 中选数字，能选到的元素和的最大值，且有标记从 $i$ 移动到 $i-1$，即 $f[i][1]$。转移方程为 $f[i+1][1] = f[i][1] + \textit{nums}[i]$。

初始值 $f[0][0] = f[0][1] = 0$。

答案为 $f[n][0]$。

代码实现时，$f$ 的第一个维度可以优化掉。

```py [sol-Python3]
class Solution:
    def maxTotal(self, nums: list[int], s: str) -> int:
        f0 = f1 = 0
        for x, ch in zip(nums, s):
            if ch == '0':
                f1 = f0 + x
            else:
                f0 = max(f0 + x, f1)
                f1 += x
        return f0
```

```java [sol-Java]
class Solution {
    public long maxTotal(int[] nums, String s) {
        long f0 = 0, f1 = 0;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (s.charAt(i) == '0') {
                f1 = f0 + x;
            } else {
                f0 = Math.max(f0 + x, f1);
                f1 += x;
            }
        }
        return f0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxTotal(vector<int>& nums, string s) {
        long long f0 = 0, f1 = 0;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            if (s[i] == '0') {
                f1 = f0 + x;
            } else {
                f0 = max(f0 + x, f1);
                f1 += x;
            }
        }
        return f0;
    }
};
```

```go [sol-Go]
func maxTotal(nums []int, s string) int64 {
	f0, f1 := 0, 0
	for i, x := range nums {
		if s[i] == '0' {
			f1 = f0 + x
		} else {
			f0 = max(f0+x, f1)
			f1 += x
		}
	}
	return int64(f0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

如果一个标记可以移动不止一次呢？

这题是 [3781. 二进制交换后的最大分数](https://leetcode.cn/problems/maximum-score-after-binary-swaps/)。

## 专题训练

1. 双指针题单的「**六、分组循环**」。
2. 动态规划题单的「**六、状态机 DP**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

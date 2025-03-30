## 初步分类

设 $s$ 中选的子串为 $x$，$t$ 中选的子串为 $y$。字符串 $A$ 的长度记作 $|A|$。

答案有如下三种情况：

- $|x|=|y|$。
- $|x|>|y|$。这意味着 $x$ 的长为 $|x|-|y|$ 后缀是一个回文串。
- $|x|<|y|$。

我们只需要计算 $|x|\ge |y|$ 的情况，对于 $|x| < |y|$ 的情况，可以计算 $\text{reverse}(t)$ 和 $\text{reverse}(s)$，复用同一份代码。

## |x| = |y|

回想一下 [1143. 最长公共子序列](https://leetcode.cn/problems/longest-common-subsequence/)。

定义 $f[i+1][j]$ 表示以 $s[i]$ 结尾的子串（倒序）与以 $t[j]$ 开头的子串的最长匹配长度。

- 如果 $s[i]\ne t[j]$，那么 $f[i+1][j] = 0$。
- 如果 $s[i]= t[j]$，那么问题变成以 $s[i-1]$ 结尾的子串（倒序）与以 $t[j+1]$ 开头的子串的最长匹配长度，在这个基础上加一，即 $f[i+1][j] = f[i][j+1] + 1$。

初始值 $f[0][j] = f[i][|t|] = 0$。

$|x|=|y|$ 时，答案为所有 $f[i][j]$ 的最大值乘以 $2$。

## |x| > |y|

回想一下 [5. 最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/) 的**中心扩展法**。

枚举 $s$ 的回文中心，贪心地，这个回文子串越长越好，因为 $x$ 去掉的回文后缀越长，越容易找到一个与之匹配的 $y$。

假设 $x$ 的回文后缀是 $[l,r]$，再加上以 $s[l-1]$ 结尾的子串（倒序）与以 $t[j]$ 开头的子串的最长匹配长度（乘以 $2$），可以得到一个长为

$$
r-l+1 + 2\cdot \max_{j=0}^{|t|-1} f[l][j]
$$

的回文串，更新答案的最大值。

代码实现时，用一个数组 $\textit{mx}[i]$ 记录 $\max\limits_{j=0}^{|t|-1} f[i][j]$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
class Solution:
    def calc(self, s: str, t: str) -> int:
        n, m = len(s), len(t)
        f = [[0] * (m + 1) for _ in range(n + 1)]
        for i, x in enumerate(s):
            for j, y in enumerate(t):
                if x == y:
                    f[i + 1][j] = f[i][j + 1] + 1
        mx = list(map(max, f))
        ans = max(mx) * 2  # |x| = |y| 的情况

        # 计算 |x| > |y| 的情况，中心扩展法
        for i in range(2 * n - 1):
            l, r = i // 2, (i + 1) // 2
            while l >= 0 and r < n and s[l] == s[r]:
                l -= 1
                r += 1
            if l + 1 <= r - 1:  # s[l+1] 到 s[r-1] 是非空回文串
                ans = max(ans, r - l - 1 + mx[l + 1] * 2)
        return ans

    def longestPalindrome(self, s: str, t: str) -> int:
        return max(self.calc(s, t), self.calc(t[::-1], s[::-1]))
```

```java [sol-Java]
class Solution {
    private int calc(String S, String T) {
        int ans = 0;
        char[] s = S.toCharArray();
        char[] t = T.toCharArray();
        int n = s.length;
        int m = t.length;
        int[] mx = new int[n + 1];
        int[][] f = new int[n + 1][m + 1];

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (s[i] == t[j]) {
                    f[i + 1][j] = f[i][j + 1] + 1;
                    mx[i + 1] = Math.max(mx[i + 1], f[i + 1][j]);
                }
            }
            ans = Math.max(ans, mx[i + 1] * 2); // |x| = |y| 的情况
        }

        // 计算 |x| > |y| 的情况，中心扩展法
        for (int i = 0; i < 2 * n - 1; i++) {
            int l = i / 2, r = (i + 1) / 2;
            while (l >= 0 && r < n && s[l] == s[r]) {
                l--;
                r++;
            }
            if (l + 1 <= r - 1) { // s[l+1] 到 s[r-1] 是非空回文串
                ans = Math.max(ans, r - l - 1 + mx[l + 1] * 2);
            }
        }
        return ans;
    }

    public int longestPalindrome(String s, String t) {
        String rs = new StringBuilder(s).reverse().toString();
        String rt = new StringBuilder(t).reverse().toString();
        return Math.max(calc(s, t), calc(rt, rs));
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int calc(string& s, string& t) {
        int n = s.size(), m = t.size();
        vector<int> mx(n + 1);
        vector f(n + 1, vector<int>(m + 1));
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (s[i] == t[j]) {
                    f[i + 1][j] = f[i][j + 1] + 1;
                }
            }
            mx[i + 1] = ranges::max(f[i + 1]);
        }
        int ans = ranges::max(mx) * 2; // |x| = |y| 的情况

        // 计算 |x| > |y| 的情况，中心扩展法
        for (int i = 0; i < 2 * n - 1; i++) {
            int l = i / 2, r = (i + 1) / 2;
            while (l >= 0 && r < n && s[l] == s[r]) {
                l--;
                r++;
            }
            if (l + 1 <= r - 1) { // s[l+1] 到 s[r-1] 是非空回文串
                ans = max(ans, r - l - 1 + mx[l + 1] * 2);
            }
        }
        return ans;
    }

    int longestPalindrome(string s, string t) {
        string rs = s, rt = t;
        ranges::reverse(rs);
        ranges::reverse(rt);
        return max(calc(s, t), calc(rt, rs));
    }
};
```

```go [sol-Go]
func calc(s, t string) int {
	n, m := len(s), len(t)
	mx := make([]int, n+1)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i, x := range s {
		for j, y := range t {
			if x == y {
				f[i+1][j] = f[i][j+1] + 1
			}
		}
		mx[i+1] = slices.Max(f[i+1])
	}
	ans := slices.Max(mx) * 2 // |x| = |y| 的情况

	// 计算 |x| > |y| 的情况，中心扩展法
	for i := range 2*n - 1 {
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if l+1 <= r-1 { // s[l+1] 到 s[r-1] 是非空回文串
			ans = max(ans, r-l-1+mx[l+1]*2)
		}
	}
	return ans
}

func longestPalindrome(s, t string) int {
	return max(calc(s, t), calc(reverse(t), reverse(s)))
}

func reverse(s string) string {
	t := []byte(s)
	slices.Reverse(t)
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm+n^2+m^2)$，其中 $n$ 是 $s$ 的长度，$m$ 是 $t$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2+m^2)$。用滚动数组可以优化至 $\mathcal{O}(n+m)$。

**注**：利用后缀数组 + 高度数组 + Manacher 算法，可以把时间复杂度优化到 $\mathcal{O}(n+m)$。后面补充。

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

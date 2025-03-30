## 分类

设 $s$ 中选的子串为 $x$，$t$ 中选的子串为 $y$。字符串 $A$ 的长度记作 $|A|$。

有如下三种情况：

- $|x|=|y|$。这类似**最长公共子串**问题，见 [718. 最长重复子数组](https://leetcode.cn/problems/maximum-length-of-repeated-subarray/)。
- $|x|>|y|$。这意味着 $x$ 的长为 $|x|-|y|$ 后缀是一个回文串。
- $|x|<|y|$。我们只需要计算 $|x|\ge |y|$ 的情况，对于 $|x| < |y|$ 的情况，等同于 $s=\text{reverse}(t)$，$t=\text{reverse}(s)$，可以复用同一个逻辑。

## |x| = |y|

回想一下 [1143. 最长公共子序列](https://leetcode.cn/problems/longest-common-subsequence/)。这里是连续的子串。

定义 $f[i+1][j]$ 表示以 $s[i]$ 结尾的子串（倒序）与以 $t[j]$ 开头的子串的**最长匹配长度**。

- 如果 $s[i]\ne t[j]$，那么 $f[i+1][j] = 0$。
- 如果 $s[i]= t[j]$，那么问题变成以 $s[i-1]$ 结尾的子串（倒序）与以 $t[j+1]$ 开头的子串的最长匹配长度，在这个基础上加一，即 $f[i+1][j] = f[i][j+1] + 1$。

初始值 $f[0][j] = f[i][|t|] = 0$。

$|x|=|y|$ 时，答案为所有 $f[i][j]$ 的最大值（即全局最长匹配长度）乘以 $2$（在 $s$ 和 $t$ 中各选一个一样长的子串）。

## |x| > |y|

回想一下 [5. 最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/) 的**中心扩展法**。

枚举 $s$ 的回文中心，贪心地，这个回文子串越长越好，因为 $x$ 去掉的回文后缀越长，越容易找到一个与之匹配的 $y$。

假设 $x$ 的回文后缀是 $[l,r]$，再加上以 $s[l-1]$ 结尾的子串（倒序）与以 $t[j]$ 开头的子串的最长匹配长度（乘以 $2$），可以得到一个长为

$$
r-l+1 + 2\cdot \max_{j=0}^{|t|-1} f[l][j]
$$

的回文串，更新答案的最大值。

代码实现时，用一个数组 $\textit{mx}[i]$ 记录 $\max\limits_{j=0}^{|t|-1} f[i][j]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV17yZzYbEP8/?t=4m28s)，欢迎点赞关注~

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
        String revS = new StringBuilder(s).reverse().toString();
        String revT = new StringBuilder(t).reverse().toString();
        return Math.max(calc(s, t), calc(revT, revS));
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
        string rev_s = s, rev_t = t;
        ranges::reverse(rev_s);
        ranges::reverse(rev_t);
        return max(calc(s, t), calc(rev_t, rev_s));
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

## 附：后缀数组 + 高度数组 + Manacher 算法

### 前置知识

1. [后缀数组](https://oi-wiki.org/string/sa/)。
2. [Manacher 算法【视频讲解】](https://www.bilibili.com/video/BV1UcyYY4EnQ/)。

利用后缀数组的高度数组，可以求出方法一中的 $\textit{mx}$ 数组，具体见代码。

计算每个回文中心的最长回文子串，可以用 Manacher 算法。

```go [sol-Go]
func calc(s, t string) int {
	// ts = t + "#" + s
	ts := append([]byte(t), '#')
	tmp := []byte(s)
	slices.Reverse(tmp)
	ts = append(ts, tmp...)
	sa := (*struct {
		_  []byte
		sa []int32
	})(unsafe.Pointer(suffixarray.New(ts))).sa

	// 后缀名次数组 rank
	// 后缀 ts[i:] 位于后缀字典序中的第 rank[i] 个
	// 特别地，rank[0] 即 ts 在后缀字典序中的排名，rank[n-1] 即 ts[n-1:] 在字典序中的排名
	rank := make([]int, len(sa))
	for i, p := range sa {
		rank[p] = i
	}

	// 高度数组 height
	// sa 中相邻后缀的最长公共前缀 LCP
	// height[0] = height[len(sa)] = 0（哨兵）
	// height[i] = LCP(ts[sa[i]:], ts[sa[i-1]:])
	height := make([]int, len(sa)+1)
	h := 0
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < len(ts) && j+h < len(ts) && ts[i+h] == ts[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	mx := make([]int, len(s)+1)
	lcp := 0
	// sa[0] 对应 '#' 开头的后缀，不遍历
	for i := 1; i < len(sa); i++ {
		if int(sa[i]) < len(t) {
			lcp = math.MaxInt // 找到了 t 中的后缀，可以开始计算 LCP
		} else {
			lcp = min(lcp, height[i])
			mx[int(sa[i])-len(t)-1] = lcp
		}
	}
	lcp = 0
	for i := len(sa) - 1; i > 0; i-- { // 反着再来一遍
		if int(sa[i]) < len(t) {
			lcp = math.MaxInt
		} else {
			lcp = min(lcp, height[i+1])
			j := int(sa[i]) - len(t) - 1
			mx[j] = max(mx[j], lcp)
		}
	}
	slices.Reverse(mx)
	ans := slices.Max(mx) * 2 // |x| = |y| 的情况

	// 计算 |x| > |y| 的情况
	s2 := append(make([]byte, 0, len(s)*2+3), '^')
	for _, c := range s {
		s2 = append(s2, '#', byte(c))
	}
	s2 = append(s2, '#', '$')
	halfLen := make([]int, len(s2)-2)
	halfLen[1] = 1
	boxM, boxR := 0, 0
	for i := 2; i < len(halfLen); i++ {
		hl := 1
		if i < boxR {
			hl = min(halfLen[boxM*2-i], boxR-i)
		}
		for s2[i-hl] == s2[i+hl] {
			hl++
			boxM, boxR = i, i+hl
		}
		halfLen[i] = hl

		if hl > 1 { // 回文子串不为空
			l := (i - hl) / 2 // 回文子串左端点
			ans = max(ans, hl-1+mx[l]*2)
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

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 是 $s$ 的长度，$m$ 是 $t$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

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

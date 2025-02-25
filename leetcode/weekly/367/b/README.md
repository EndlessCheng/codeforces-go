## 方法一：枚举

首先，如果 $s$ 中 $1$ 的个数不足 $k$，直接返回空串。

否则一定有解。

从 $k$ 开始枚举答案的长度 $\textit{size}$，然后在 $s$ 中枚举所有长为 $\textit{size}$ 的子串，同时维护字典序最小的子串。如果存在一个子串，其中 $1$ 的个数等于 $k$，则返回字典序最小的子串。

[本题视频讲解](https://www.bilibili.com/video/BV1aC4y1G7dB/)

```py [sol-Python3]
class Solution:
    def shortestBeautifulSubstring(self, s: str, k: int) -> str:
        # 1 的个数不足 k
        if s.count('1') < k:
            return ''
        # 否则一定有解
        for size in count(k):  # 从 k 开始枚举
            ans = ''
            for i in range(size, len(s) + 1):
                t = s[i - size: i]
                if (ans == '' or t < ans) and t.count('1') == k:
                    ans = t
            if ans: return ans
```

```java [sol-Java]
class Solution {
    public String shortestBeautifulSubstring(String s, int k) {
        // 1 的个数不足 k
        if (s.replace("0", "").length() < k) {
            return "";
        }
        // 否则一定有解
        for (int size = k; ; size++) {
            String ans = "";
            for (int i = size; i <= s.length(); i++) {
                String t = s.substring(i - size, i);
                if ((ans.isEmpty() || t.compareTo(ans) < 0) && t.replace("0", "").length() == k) {
                    ans = t;
                }
            }
            if (!ans.isEmpty()) {
                return ans;
            }
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string shortestBeautifulSubstring(string s, int k) {
        // 1 的个数不足 k
        if (ranges::count(s, '1') < k) {
            return "";
        }
        // 否则一定有解
        for (int size = k; ; size++) {
            string ans = "";
            for (int i = size; i <= s.length(); i++) {
                string t = s.substr(i - size, size);
                if ((ans.empty() || t < ans) && ranges::count(t, '1') == k) {
                    ans = t;
                }
            }
            if (!ans.empty()) {
                return ans;
            }
        }
    }
};
```

```go [sol-Go]
func shortestBeautifulSubstring(s string, k int) string {
	// 1 的个数不足 k
	if strings.Count(s, "1") < k {
		return ""
	}
	// 否则一定有解
	for size := k; ; size++ {
		ans := ""
		for i := size; i <= len(s); i++ {
			t := s[i-size : i]
			if (ans == "" || t < ans) && strings.Count(t, "1") == k {
				ans = t
			}
		}
		if ans != "" {
			return ans
		}
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^3)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。字符串切片需要 $\mathcal{O}(n)$ 的空间，Go 除外。

## 方法二：滑动窗口

原理请看 [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)

由于答案中恰好有 $k$ 个 $1$，我们也可以用滑动窗口找最短的答案。

如果窗口内的 $1$ 的个数超过 $k$，或者窗口端点是 $0$，就可以缩小窗口。

> 注：利用字符串哈希（或者后缀数组等），可以把比较字典序的时间降至 $\mathcal{O}(n\log n)$，这样可以做到 $\mathcal{O}(n\log n)$ 的时间复杂度。

```py [sol-Python3]
class Solution:
    def shortestBeautifulSubstring(self, s: str, k: int) -> str:
        if s.count('1') < k:
            return ''
        ans = s
        cnt1 = left = 0
        for right, c in enumerate(s):
            cnt1 += int(c)
            while cnt1 > k or s[left] == '0':
                cnt1 -= int(s[left])
                left += 1
            if cnt1 == k:
                t = s[left: right + 1]
                if len(t) < len(ans) or len(t) == len(ans) and t < ans:
                    ans = t
        return ans
```

```java [sol-Java]
class Solution {
    public String shortestBeautifulSubstring(String S, int k) {
        if (S.replace("0", "").length() < k) {
            return "";
        }
        char[] s = S.toCharArray();
        String ans = S;
        int cnt1 = 0, left = 0;
        for (int right = 0; right < s.length; right++) {
            cnt1 += s[right] - '0';
            while (cnt1 > k || s[left] == '0') {
                cnt1 -= s[left++] - '0';
            }
            if (cnt1 == k) {
                String t = S.substring(left, right + 1);
                if (t.length() < ans.length() || t.length() == ans.length() && t.compareTo(ans) < 0) {
                    ans = t;
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string shortestBeautifulSubstring(string s, int k) {
        if (ranges::count(s, '1') < k) {
            return "";
        }
        string ans = s;
        int cnt1 = 0, left = 0;
        for (int right = 0; right < s.length(); right++) {
            cnt1 += s[right] - '0';
            while (cnt1 > k || s[left] == '0') {
                cnt1 -= s[left++] - '0';
            }
            if (cnt1 == k) {
                string t = s.substr(left, right - left + 1);
                if (t.length() < ans.length() || t.length() == ans.length() && t < ans) {
                    ans = move(t);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func shortestBeautifulSubstring(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}
	ans := s
	cnt1 := 0
	left := 0
	for right, b := range s {
		cnt1 += int(b & 1)
		for cnt1 > k || s[left] == '0' {
			cnt1 -= int(s[left] & 1)
			left++
		}
		if cnt1 == k {
			t := s[left : right+1]
			if len(t) < len(ans) || len(t) == len(ans) && t < ans {
				ans = t
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。字符串切片需要 $\mathcal{O}(n)$ 的空间，Go 除外。

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

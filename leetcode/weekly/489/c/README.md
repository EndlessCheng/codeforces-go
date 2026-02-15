**前置题目**：[5. 最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/)，[我的题解](https://leetcode.cn/problems/longest-palindromic-substring/solutions/2958179/mo-ban-on-manacher-suan-fa-pythonjavacgo-t6cx/)。

用中心扩展法，枚举最终答案的回文中心，向外扩展。

如果发现 $s[l]\ne s[r]$，那么必须删除字母，才能继续扩展：

- 选择删除 $s[l]$，从 $l-1$ 和 $r$ 开始，继续向左向右扩展。
- 选择删除 $s[r]$，从 $l$ 和 $r+1$ 开始，继续向左向右扩展。
- 两种情况取最大值。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def almostPalindromic(self, s: str) -> int:
        n = len(s)
        ans = 0

        def expand(l: int, r: int) -> None:
            while l >= 0 and r < n and s[l] == s[r]:
                l -= 1
                r += 1
            nonlocal ans
            ans = max(ans, r - l - 1)  # [l+1, r-1] 是回文串

        for i in range(2 * n - 1):
            l, r = i // 2, (i + 1) // 2
            while l >= 0 and r < n and s[l] == s[r]:
                l -= 1
                r += 1
            expand(l - 1, r)  # 删除 s[l]，继续扩展
            expand(l, r + 1)  # 删除 s[r]，继续扩展
            if ans >= n:  # 优化：提前返回答案
                return n
        return ans
```

```java [sol-Java]
class Solution {
    public int almostPalindromic(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int ans = 0;

        for (int i = 0; i < 2 * n - 1 && ans < n; i++) {
            int l = i / 2;
            int r = (i + 1) / 2;
            while (l >= 0 && r < n && s[l] == s[r]) {
                l--;
                r++;
            }
            ans = Math.max(ans, expand(s, l - 1, r)); // 删除 s[l]，继续扩展
            ans = Math.max(ans, expand(s, l, r + 1)); // 删除 s[r]，继续扩展
        }

        return Math.min(ans, n);
    }

    private int expand(char[] s, int l, int r) {
        while (l >= 0 && r < s.length && s[l] == s[r]) {
            l--;
            r++;
        }
        return r - l - 1;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int almostPalindromic(string s) {
        int n = s.size();
        int ans = 0;

        auto expand = [&](int l, int r) -> void {
            while (l >= 0 && r < n && s[l] == s[r]) {
                l--;
                r++;
            }
            ans = max(ans, r - l - 1); // [l+1, r-1] 是回文串
        };

        for (int i = 0; i < 2 * n - 1 && ans < n; i++) {
            int l = i / 2, r = (i + 1) / 2;
            while (l >= 0 && r < n && s[l] == s[r]) {
                l--;
                r++;
            }
            expand(l - 1, r); // 删除 s[l]，继续扩展
            expand(l, r + 1); // 删除 s[r]，继续扩展
        }
        return min(ans, n);
    }
};
```

```go [sol-Go]
func almostPalindromic(s string) (ans int) {
	n := len(s)
	expand := func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		ans = max(ans, r-l-1) // [l+1, r-1] 是回文串
	}

	for i := range 2*n - 1 {
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		expand(l-1, r) // 删除 s[l]，继续扩展
		expand(l, r+1) // 删除 s[r]，继续扩展
		if ans >= n { // 优化：提前返回答案
			return n
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 更快的做法

本题还可以用 Manacher 算法 + 后缀数组做到 $\mathcal{O}(n\log n)$（或者字符串哈希做到 $\mathcal{O}(n)$），直播结束后补充。

## 专题训练

见下面字符串题单的「**三、Manacher 算法**」和「**八、后缀数组**」。

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

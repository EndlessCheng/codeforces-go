**前置知识**：[滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

本题维护的是窗口中的相邻字符相等的次数 $\textit{cnt}$。

1. $s[\textit{right}]$ 进入窗口，检查 $s[\textit{right}]$ 是否等于 $s[\textit{right}-1]$。如果相等，$\textit{cnt}$ 增加 $1$。
2. 如果 $\textit{cnt}>1$，则不断右移左指针 $\textit{left}$ 直到 $\textit{cnt}=1$。$s[\textit{left}]$ 离开窗口前，如果其等于 $s[\textit{left}+1]$，则把 $\textit{cnt}$ 减少 $1$。
3. 统计子串长度 $\textit{right}-\textit{left}+1$ 的最大值。

```py [sol-Python3]
class Solution:
    def longestSemiRepetitiveSubstring(self, s: str) -> int:
        ans = 1
        cnt = left = 0
        for right in range(1, len(s)):
            if s[right] == s[right - 1]:
                cnt += 1
            while cnt > 1:
                if s[left] == s[left + 1]:
                    cnt -= 1  # left 离开窗口后，和 left+1 断开
                left += 1
            ans = max(ans, right - left + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestSemiRepetitiveSubstring(String S) {
        char[] s = S.toCharArray();
        int ans = 1;
        int cnt = 0;
        int left = 0;
        for (int right = 1; right < s.length; right++) {
            if (s[right] == s[right - 1]) {
                cnt++;
            }
            while (cnt > 1) {
                if (s[left] == s[left + 1]) {
                    cnt--; // left 离开窗口后，和 left+1 断开
                }
                left++;
            }
            ans = Math.max(ans, right - left + 1);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestSemiRepetitiveSubstring(string s) {
        int ans = 1, cnt = 0, left = 0;
        for (int right = 1; right < s.size(); right++) {
            if (s[right] == s[right - 1]) {
                cnt++;
            }
            while (cnt > 1) {
                if (s[left] == s[left + 1]) {
                    cnt--; // left 离开窗口后，和 left+1 断开
                }
                left++;
            }
            ans = max(ans, right - left + 1);
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestSemiRepetitiveSubstring(s string) int {
	ans, cnt, left := 1, 0, 0
	for right := 1; right < len(s); right++ {
		if s[right] == s[right-1] {
			cnt++
		}
		for cnt > 1 {
			if s[left] == s[left+1] {
				cnt-- // left 离开窗口后，和 left+1 断开
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。注意 $\textit{left}$ 只会增加不会减少，所以二重循环的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。

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

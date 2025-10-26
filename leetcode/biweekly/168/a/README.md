枚举 $k=1,2,\ldots,n$，反转长为 $k$ 的前缀/后缀，用得到的字符串更新答案的最小值。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def lexSmallest(self, s: str) -> str:
        ans = s  # k = 1 时，操作不改变 s
        for k in range(2, len(s) + 1):
            ans = min(ans, s[:k][::-1] + s[k:], s[:-k] + s[-k:][::-1])
        return ans
```

```java [sol-Java]
class Solution {
    public String lexSmallest(String s) {
        int n = s.length();
        String ans = s; // k = 1 时，操作不改变 s
        for (int k = 2; k <= n; k++) {
            StringBuilder t = new StringBuilder(s.substring(0, k)).reverse();
            ans = min(ans, t + s.substring(k));

            t = new StringBuilder(s.substring(n - k)).reverse();
            ans = min(ans, s.substring(0, n - k) + t);
        }
        return ans;
    }

    private String min(String a, String b) {
        return a.compareTo(b) <= 0 ? a : b;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string lexSmallest(string s) {
        int n = s.size();
        string ans = s; // k = 1 时，操作不改变 s
        for (int k = 2; k <= n; k++) {
            string t = s.substr(0, k);
            ranges::reverse(t);
            ans = min(ans, t + s.substr(k));

            t = s.substr(n - k);
            ranges::reverse(t);
            ans = min(ans, s.substr(0, n - k) + t);
        }
        return ans;
    }
};
```

```go [sol-Go]
func lexSmallest(s string) string {
	n := len(s)
	ans := s // k = 1 时，操作不改变 s
	for k := 2; k <= n; k++ {
		t := []byte(s[:k])
		slices.Reverse(t)
		ans = min(ans, string(t)+s[k:])

		t = []byte(s[n-k:])
		slices.Reverse(t)
		ans = min(ans, s[:n-k]+string(t))
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 更快的做法

可以用后缀数组（或者字符串哈希）快速比较两个子串的字典序大小，时间复杂度 $\mathcal{O}(n\log n)$。

代码稍后补充。

## 专题训练

见下面字符串题单的「**八、后缀数组/后缀自动机**」。

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

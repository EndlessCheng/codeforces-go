**题意**：找一个连续相同子串 $s[i]$ 到 $s[j]$，要求长度恰好等于 $k$，且 $s[i-1]\ne s[i]$ 且 $s[j]\ne s[j+1]$（如果有）。

遍历 $s$ 的同时，维护连续相同子串长度 $\textit{cnt}$，遇到子串末尾就看看 $\textit{cnt}=k$ 是否成立，成立就立刻返回 $\texttt{true}$，否则重置 $\textit{cnt}=0$。

[本题视频讲解](https://www.bilibili.com/video/BV1pmAGegEcw/)，欢迎点赞关注~

```py [sol-Py3]
class Solution:
    def hasSpecialSubstring(self, s: str, k: int) -> bool:
        cnt = 0
        for i, c in enumerate(s):
            cnt += 1
            if i == len(s) - 1 or c != s[i + 1]:
                if cnt == k:
                    return True
                cnt = 0
        return False
```

```py [sol-Py3 写法二]
class Solution:
    def hasSpecialSubstring(self, s: str, k: int) -> bool:
        for _, it in groupby(s):
            if len(list(it)) == k:
                return True
        return False
```

```py [sol-Py3 写法三]
class Solution:
    def hasSpecialSubstring(self, s: str, k: int) -> bool:
        return any(len(list(it)) == k for _, it in groupby(s))
```

```java [sol-Java]
class Solution {
    public boolean hasSpecialSubstring(String S, int k) {
        char[] s = S.toCharArray();
        int n = s.length;
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            cnt++;
            if (i == n - 1 || s[i] != s[i + 1]) {
                if (cnt == k) {
                    return true;
                }
                cnt = 0;
            }
        }
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool hasSpecialSubstring(string s, int k) {
        int n = s.size();
        int cnt = 0;
        for (int i = 0; i < n; i++) {
            cnt++;
            if (i == n - 1 || s[i] != s[i + 1]) {
                if (cnt == k) {
                    return true;
                }
                cnt = 0;
            }
        }
        return false;
    }
};
```

```go [sol-Go]
func hasSpecialSubstring(s string, k int) bool {
	cnt := 0
	for i := range s {
		cnt++
		if i == len(s)-1 || s[i] != s[i+1] {
			if cnt == k {
				return true
			}
			cnt = 0
		}
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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

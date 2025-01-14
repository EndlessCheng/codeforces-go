## 算法一：暴力

本题思路和 [20. 有效的括号](https://leetcode.cn/problems/valid-parentheses/) 是一样的。把 AB 看成一对括号 `()`，CD 看成另一对括号 `{}`。

暴力做法是不断把 AB 和 CD 去掉，直到 $s$ 中没有 AB 和 CD 为止。

请看[【周赛 346】](https://www.bilibili.com/video/BV1Qm4y1t7cx/)第一题，欢迎点赞关注~

```py [sol1-Python3]
class Solution:
    def minLength(self, s: str) -> int:
        while "AB" in s or "CD" in s:
            s = s.replace("AB", "").replace("CD", "")
        return len(s)
```

```java [sol1-Java]
class Solution {
    public int minLength(String s) {
        while (s.contains("AB") || s.contains("CD"))
            s = s.replace("AB", "").replace("CD", "");
        return s.length();
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int minLength(string s) {
        regex ab("AB"), cd("CD");
        while (s.find("AB") != string::npos || s.find("CD") != string::npos) {
            s = regex_replace(s, ab, "");
            s = regex_replace(s, cd, "");
        }
        return s.length();
    }
};
```

```go [sol1-Go]
func minLength(s string) int {
	for strings.Contains(s, "AB") || strings.Contains(s, "CD") {
		s = strings.ReplaceAll(s, "AB", "")
		s = strings.ReplaceAll(s, "CD", "")
	}
	return len(s)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $s$ 的长度。对于 AAA...BBB 这样的字符串，需要循环 $\mathcal{O}(n)$ 次，每次需要 $\mathcal{O}(n)$ 的时间。
- 空间复杂度：$\mathcal{O}(n)$。替换过程中生成的字符串需要 $\mathcal{O}(n)$ 的空间。

## 算法二：栈

用栈记录遍历过的，没有删除的字母。

如果当前字母是 B，且栈顶为 A，那么这两个字母都可以删除。同理，如果当前字母是 D，且栈顶为 C，那么这两个字母都可以删除。

否则，把当前字母入栈。

```py [sol-Python3]
class Solution:
    def minLength(self, s: str) -> int:
        st = []
        for c in s:
            if st and (c == 'B' and st[-1] == 'A' or c == 'D' and st[-1] == 'C'):
                st.pop()
            else:
                st.append(c)
        return len(st)
```

```java [sol-Java]
class Solution {
    public int minLength(String s) {
        var st = new ArrayDeque<Character>();
        for (var c : s.toCharArray()) {
            if (!st.isEmpty() && (c == 'B' && st.peek() == 'A' || c == 'D' && st.peek() == 'C'))
                st.pop();
            else
                st.push(c);
        }
        return st.size();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minLength(string s) {
        stack<char> st;
        for (char c: s) {
            if (!st.empty() && (c == 'B' && st.top() == 'A' || c == 'D' && st.top() == 'C'))
                st.pop();
            else
                st.push(c);
        }
        return st.size();
    }
};
```

```go [sol-Go]
func minLength(s string) int {
	st := []rune{}
	for _, c := range s {
		if len(st) > 0 && (c == 'B' && st[len(st)-1] == 'A' || c == 'D' && st[len(st)-1] == 'C') {
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return len(st)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，请看下面数据结构题单中的「**§3.3 邻项消除**」。

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

注意字符串为空的时候，不能删除最后一个字母（否则会报错）。

```py [sol-Python3]
class Solution:
    def processStr(self, s: str) -> str:
        ans = []
        for c in s:
            if c == '*':
                if ans:
                    ans.pop()
            elif c == '#':
                ans += ans
            elif c == '%':
                ans.reverse()
            else:
                ans.append(c)
        return ''.join(ans)
```

```java [sol-Java]
class Solution {
    String processStr(String s) {
        StringBuilder ans = new StringBuilder();
        for (char c : s.toCharArray()) {
            if (c == '*') {
                if (!ans.isEmpty()) {
                    ans.setLength(ans.length() - 1);
                }
            } else if (c == '#') {
                ans.append(ans);
            } else if (c == '%') {
                ans.reverse();
            } else {
                ans.append(c);
            }
        }
        return ans.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string processStr(string s) {
        string ans;
        for (char c : s) {
            if (c == '*') {
                if (!ans.empty()) {
                    ans.pop_back();
                }
            } else if (c == '#') {
                ans += ans;
            } else if (c == '%') {
                ranges::reverse(ans);
            } else {
                ans += c;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func processStr(s string) string {
	ans := []byte{}
	for _, c := range s {
		if c == '*' {
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
		} else if c == '#' {
			ans = append(ans, ans...)
		} else if c == '%' {
			slices.Reverse(ans)
		} else {
			ans = append(ans, byte(c))
		}
	}
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(2^n)$，其中 $n$ 是 $s$ 的长度。最坏情况下把一个字母复制 $n-1$ 次，得到长为 $2^{n-1}$ 的字符串。这个过程的时间为等比数列之和 $\mathcal{O}(1+2+4+8+\cdots) = \mathcal{O}(2^n)$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

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

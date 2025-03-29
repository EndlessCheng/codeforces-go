## 方法一：双指针

初始化 $j=0$。遍历 $s$，把 $s[i]$ 加到答案中。如果 $\textit{spaces}[j]$ 等于当前下标 $i$，那么先加空格，再加 $s[i]$，然后把 $j$ 加一。

```py [sol-Python3]
class Solution:
    def addSpaces(self, s: str, spaces: List[int]) -> str:
        ans = []
        j = 0
        for i, c in enumerate(s):
            if j < len(spaces) and spaces[j] == i:
                ans.append(' ')
                j += 1
            ans.append(c)
        return ''.join(ans)
```

```java [sol-Java]
class Solution {
    public String addSpaces(String s, int[] spaces) {
        StringBuilder ans = new StringBuilder(s.length() + spaces.length);
        int j = 0;
        for (int i = 0; i < s.length(); i++) {
            if (j < spaces.length && spaces[j] == i) {
                ans.append(' ');
                j++;
            }
            ans.append(s.charAt(i));
        }
        return ans.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string addSpaces(string s, vector<int>& spaces) {
        string ans;
        int j = 0;
        for (int i = 0; i < s.size(); i++) {
            if (j < spaces.size() && spaces[j] == i) {
                ans += ' ';
                j++;
            }
            ans += s[i];
        }
        return ans;
    }
};
```

```go [sol-Go]
func addSpaces(s string, spaces []int) string {
	ans := make([]byte, 0, len(s)+len(spaces))
	j := 0
	for i, c := range s {
		if j < len(spaces) && spaces[j] == i {
			ans = append(ans, ' ')
			j++
		}
		ans = append(ans, byte(c))
	}
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。注意 $\textit{spaces}$ 的长度不超过 $n$。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。

## 方法二：按照空格分割

```py [sol-Python3]
class Solution:
    def addSpaces(self, s: str, spaces: List[int]) -> str:
        spaces.append(len(s))  # 这样可以在循环中处理最后一段
        ans = [s[:spaces[0]]]
        for p, q in pairwise(spaces):
            ans.append(' ')
            ans.append(s[p: q])
        return ''.join(ans)
```

```java [sol-Java]
class Solution {
    public String addSpaces(String s, int[] spaces) {
        StringBuilder ans = new StringBuilder(s.length() + spaces.length);
        ans.append(s, 0, spaces[0]);
        for (int i = 1; i < spaces.length; i++) {
            ans.append(' ');
            ans.append(s, spaces[i - 1], spaces[i]);
        }
        ans.append(' ');
        ans.append(s, spaces[spaces.length - 1], s.length());
        return ans.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string addSpaces(string s, vector<int>& spaces) {
        spaces.push_back(s.length()); // 这样可以在循环中处理最后一段
        string ans(s, 0, spaces[0]);
        for (int i = 1; i < spaces.size(); i++) {
            ans += ' ';
            ans.append(s, spaces[i - 1], spaces[i] - spaces[i - 1]);
        }
        return ans;
    }
};
```

```go [sol-Go]
func addSpaces(s string, spaces []int) string {
	spaces = append(spaces, len(s)) // 这样可以在循环中处理最后一段
	ans := make([]byte, 0, len(s)+len(spaces))
	ans = append(ans, s[:spaces[0]]...)
	for i := 1; i < len(spaces); i++ {
		ans = append(ans, ' ')
		ans = append(ans, s[spaces[i-1]:spaces[i]]...)
	}
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。

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

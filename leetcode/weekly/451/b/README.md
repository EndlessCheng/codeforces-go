注意相邻字符消除后，原本不相邻的字符会变成相邻，可以继续消除。

由于每次都是消除最左边的字符，我们从左到右遍历 $s$，同时用栈记录字符：

- 如果栈不为空，且 $s[i]$ 与栈顶是「连续」的，那么立刻消除，弹出栈顶。
- 否则把 $s[i]$ 入栈。

最后答案就是栈中剩余字符，即栈底到栈顶。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1o1jgzJE51/?t=3m47s)，欢迎点赞关注~

```py [sol-Python3]
# 更快的写法见【Python3 写法二】
def is_consecutive(x: str, y: str) -> bool:
    d = abs(ord(x) - ord(y))
    return d == 1 or d == 25

class Solution:
    def resultingString(self, s: str) -> str:
        st = []
        for b in s:
            if st and is_consecutive(b, st[-1]):
                st.pop()
            else:
                st.append(b)
        return ''.join(st)
```

```py [sol-Python3 写法二]
# "a" -> "zb", "b" -> "ac", "c" -> "bd", ..., "z" -> "ya"
lc = ascii_lowercase
mp = {b: lc[i - 1] + lc[(i + 1) % 26] for i, b in enumerate(lc)}

class Solution:
    def resultingString(self, s: str) -> str:
        st = []
        for b in s:
            if st and st[-1] in mp[b]:
                st.pop()
            else:
                st.append(b)
        return ''.join(st)
```

```java [sol-Java]
class Solution {
    public String resultingString(String s) {
        StringBuilder st = new StringBuilder(); // 更快的写法见【Java 数组】
        for (char b : s.toCharArray()) {
            if (!st.isEmpty() && isConsecutive(b, st.charAt(st.length() - 1))) {
                st.setLength(st.length() - 1);
            } else {
                st.append(b);
            }
        }
        return st.toString();
    }

    private boolean isConsecutive(char x, char y) {
        int d = Math.abs(x - y);
        return d == 1 || d == 25;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public String resultingString(String s) {
        char[] st = new char[s.length()];
        int top = -1; // 栈顶下标
        for (char b : s.toCharArray()) {
            if (top >= 0 && isConsecutive(b, st[top])) {
                top--; // 出栈
            } else {
                st[++top] = b; // 入栈
            }
        }
        return new String(st, 0, top + 1);
    }

    private boolean isConsecutive(char x, char y) {
        int d = Math.abs(x - y);
        return d == 1 || d == 25;
    }
}
```

```cpp [sol-C++]
class Solution {
    bool is_consecutive(char x, char y) {
        int d = abs(x - y);
        return d == 1 || d == 25;
    }

public:
    string resultingString(string s) {
        string st;
        for (char b : s) {
            if (!st.empty() && is_consecutive(b, st.back())) {
                st.pop_back();
            } else {
                st += b;
            }
        }
        return st;
    }
};
```

```cpp [sol-C++ 原地]
class Solution {
    bool is_consecutive(char x, char y) {
        int d = abs(x - y);
        return d == 1 || d == 25;
    }

public:
    string resultingString(string s) {
        int top = -1; // 栈顶下标
        for (char b : s) {
            if (top >= 0 && is_consecutive(b, s[top])) {
                top--; // 出栈
            } else {
                s[++top] = b; // 入栈
            }
        }
        s.resize(top + 1);
        return s;
    }
};
```

```go [sol-Go]
func isConsecutive(x, y byte) bool {
	d := abs(int(x) - int(y))
	return d == 1 || d == 25
}

func resultingString(s string) string {
	st := []byte{}
	for _, b := range s {
		if len(st) > 0 && isConsecutive(byte(b), st[len(st)-1]) {
			st = st[:len(st)-1]
		} else {
			st = append(st, byte(b))
		}
	}
	return string(st)
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度，每个字符至多入栈出栈各一次。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。C++ 原地做法可以 $\mathcal{O}(1)$。

更多相似题目，见下面数据结构题单的「**§3.3 邻项消除**」。

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

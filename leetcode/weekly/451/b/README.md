注意相邻字符消除后，原本不相邻的字符会变成相邻，可以继续消除。

由于每次都是消除最左边的字符，我们从左到右遍历 $s$，同时用栈记录遍历过的字符。

如果栈顶的两个字符是「连续」的，那么立刻消除，即弹出栈顶的两个字符。

最后答案就是栈中剩余字符，即栈底到栈顶。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注！

```py [sol-Python3]
def is_consecutive(x: str, y: str) -> bool:
    d = abs(ord(x) - ord(y))
    return d == 1 or d == 25

class Solution:
    def resultingString(self, s: str) -> str:
        st = []
        for b in s:
            st.append(b)
            if len(st) > 1 and is_consecutive(st[-2], st[-1]):
                del st[-2:]
        return ''.join(st)
```

```java [sol-Java]
class Solution {
    public String resultingString(String s) {
        StringBuilder st = new StringBuilder();
        for (char b : s.toCharArray()) {
            st.append(b);
            if (st.length() > 1 && isConsecutive(st.charAt(st.length() - 2), st.charAt(st.length() - 1))) {
                st.setLength(st.length() - 2);
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
            st += b;
            if (st.size() > 1 && is_consecutive(st[st.size() - 2], st.back())) {
                st.resize(st.size() - 2);
            }
        }
        return st;
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
		st = append(st, byte(b))
		if len(st) > 1 && isConsecutive(st[len(st)-2], st[len(st)-1]) {
			st = st[:len(st)-2]
		}
	}
	return string(st)
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度，每个字符至多入栈出栈各一次。
- 空间复杂度：$\mathcal{O}(n)$。

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

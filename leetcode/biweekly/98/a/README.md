示例 1 的 $\textit{num}=11891$。要想得到最大的数，替换前面的数更好，还是后面的数更好？

- 把 $1$ 替换成 $9$，得到 $99899$。
- 把 $8$ 替换成 $9$，得到 $11991$。由于前面的 $1$ 没有变大，替换后面的数不是最优的。

一般地，把 $\textit{num}$ 转成十进制字符串 $s$，从左到右找第一个不等于 $9$ 的字符 $c$，把 $s$ 中的 $c$ 都替换成 $9$，得到最大数。

同理，找第一个不等于 $0$ 的字符 $c$，把 $s$ 中的 $c$ 都替换成 $0$，得到最小数。由于 $s[0]$ 一定不等于 $0$，直接替换 $s[0]$ 就行。

```py [sol-Python3]
class Solution:
    def minMaxDifference(self, num: int) -> int:
        s = str(num)
        mx = num
        for c in s:
            if c != '9':  # 第一个不等于 9 的字符
                mx = int(s.replace(c, '9'))  # 替换成 9
                break
        mn = int(s.replace(s[0], '0'))  # 第一个不等于 0 的字符，替换成 0
        return mx - mn
```

```java [sol-Java]
class Solution {
    public int minMaxDifference(int num) {
        String s = String.valueOf(num);
        int mx = num;
        for (char c : s.toCharArray()) {
            if (c != '9') { // 第一个不等于 9 的字符
                mx = Integer.parseInt(s.replace(c, '9')); // 替换成 9
                break;
            }
        }
        // 第一个不等于 0 的字符，替换成 0
        int mn = Integer.parseInt(s.replace(s.charAt(0), '0'));
        return mx - mn;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minMaxDifference(int num) {
        string s = to_string(num);
        int mx = num;
        for (char c : s) {
            if (c != '9') { // 第一个不等于 9 的字符
                string tmp = s;
                for (char& ch : tmp) {
                    if (ch == c) {
                        ch = '9'; // 替换成 9
                    }
                }
                mx = stoi(tmp);
                break;
            }
        }

        char s0 = s[0]; // 第一个不等于 0 的字符
        for (char& ch : s) {
            if (ch == s0) {
                ch = '0'; // 替换成 0
            }
        }
        int mn = stoi(s);

        return mx - mn;
    }
};
```

```go [sol-Go]
func minMaxDifference(num int) int {
	s := strconv.Itoa(num)
	mx := num
	for _, c := range s {
		if c != '9' { // 第一个不等于 9 的字符，替换成 9
			mx, _ = strconv.Atoi(strings.ReplaceAll(s, string(c), "9"))
			break
		}
	}
	// 第一个不等于 0 的字符，替换成 0
	mn, _ := strconv.Atoi(strings.ReplaceAll(s, s[:1], "0"))
	return mx - mn
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log \textit{num})$。
- 空间复杂度：$\mathcal{O}(\log \textit{num})$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

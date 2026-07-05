为方便描述，把 $s_1$ 简记为 $s$，把 $s_2$ 简记为 $t$。

从左到右思考：

- 如果 $s[i]=t[i]$，那么不改，继续向右遍历。
- 如果 $s[i]=\texttt{0}$，$t[i]=\texttt{1}$，那么用操作一即可让 $s[i]=t[i] = \texttt{1}$，操作 $1$ 次。
- 如果 $s[i]=\texttt{1}$，$t[i]=\texttt{0}$。
    - 如果 $i<n-1$ 且 $s[i+1]=\texttt{1}$，那么用操作二即可让 $s[i]=t[i] = \texttt{0}$（$s[i+1]$ 也变成了 $\texttt{0}$），操作 $1$ 次。至于 $s[i+1]$ 怎么操作，在下一轮循环中讨论。为什么不和 $s[i-1]$ 一起操作？理由见后文。
    - 如果 $i<n-1$ 且 $s[i+1]=\texttt{0}$，那么先用操作一把 $s[i+1]$ 变成 $\texttt{1}$，转化成上面的情况。一共操作 $2$ 次。
    - 如果 $i=n-1$，那么只能改 $s[i-1]$。
        - 如果 $s[i-1] = \texttt{0}$，那么先用操作一把 $s[i-1]$ 变成 $\texttt{1}$，再用操作二把 $s[i-1]$ 恢复成 $\texttt{0}$，也是一共操作 $2$ 次。
        - 如果 $s[i-1] = \texttt{1}$，那么先用操作二（$s[i-1]$ 变成 $\texttt{0}$），再用操作一把 $s[i-1]$ 恢复成 $\texttt{1}$，也是一共操作 $2$ 次。

上述分类讨论的最后一段表明，**操作左边**（$s[i-1]$）**不如操作右边**（$s[i+1]$）：

- 操作左边，必须操作 $2$ 次，而且 $s[i+1]$ 还没改。
- 操作右边，如果顺带把 $s[i+1]$ 改成等于 $t[i+1]$，一共只需操作 $1$ 次。如果操作后 $s[i+1]\ne t[i+1]$，这说明此时 $s[i+1]=\texttt{0}$，$t[i+1]=\texttt{1}$，那也只需再用操作一 $1$ 次。
- 所以操作右边不比操作左边差，无脑操作右边即可。

特判 $n=1$ 的情况，如果 $s=\texttt{1}$ 且 $t=\texttt{0}$（示例 3），那么无法操作，返回 $-1$。

[本题视频讲解](https://www.bilibili.com/video/BV1qXTC63EQa/?t=4m39s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minOperations(self, s1: str, t: str) -> int:
        n = len(s1)
        if n == 1 and s1 == "1" and t == "0":
            return -1

        # 也可以用一个布尔变量表示 s1[i] 是否操作过，从而做到 O(1) 空间，见 Python3 写法二
        s = list(s1)
        ans = 0
        for i in range(n):
            if s[i] == t[i]:
                continue
            if s[i] == '0':
                ans += 1
            elif i < n - 1 and s[i + 1] == '1':
                ans += 1
                s[i + 1] = '0'
            else:
                ans += 2
        return ans
```

```py [sol-Python3 写法二]
class Solution:
    def minOperations(self, s: str, t: str) -> int:
        n = len(s)
        if n == 1 and s == "1" and t == "0":
            return -1

        changed = False
        ans = 0
        for i in range(n):
            ch = '0' if changed else s[i]
            changed = False
            if ch == t[i]:
                continue
            if ch == '0':
                ans += 1
            elif i < n - 1 and s[i + 1] == '1':
                ans += 1
                changed = True
            else:
                ans += 2
        return ans
```

```java [sol-Java]
class Solution {
    public int minOperations(String s1, String t) {
        int n = s1.length();
        if (n == 1 && s1.charAt(0) == '1' && t.charAt(0) == '0') {
            return -1;
        }

        // 也可以用一个布尔变量表示 s1[i] 是否操作过，从而做到 O(1) 空间，见 Python3 写法二
        char[] s = s1.toCharArray();
        int ans = 0;
        for (int i = 0; i < n; i++) {
            if (s[i] == t.charAt(i)) {
                continue;
            }
            if (s[i] == '0') {
                ans++;
            } else if (i < n - 1 && s[i + 1] == '1') {
                ans++;
                s[i + 1] = '0';
            } else {
                ans += 2;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(string s, string t) {
        int n = s.size();
        if (n == 1 && s == "1" && t == "0") {
            return -1;
        }

        int ans = 0;
        for (int i = 0; i < n; i++) {
            if (s[i] == t[i]) {
                continue;
            }
            if (s[i] == '0') {
                ans++;
            } else if (i < n - 1 && s[i + 1] == '1') {
                ans++;
                s[i + 1] = '0';
            } else {
                ans += 2;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minOperations(s1, t string) (ans int) {
	n := len(s1)
	if n == 1 && s1 == "1" && t == "0" {
		return -1
	}

	// 也可以用一个布尔变量表示 s1[i] 是否操作过，从而做到 O(1) 空间，见 Python3 写法二
	s := []byte(s1)
	for i := range n {
		if s[i] == t[i] {
			continue
		}
		if s[i] == '0' {
			ans++
		} else if i < n-1 && s[i+1] == '1' {
			ans++
			s[i+1] = '0'
		} else {
			ans += 2
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。可以用一个布尔变量表示是否修改了当前字符，从而做到 $\mathcal{O}(1)$ 空间，见 Python3 写法二。

## 专题训练

见下面贪心题单的「**§1.4 从最左/最右开始贪心**」。

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

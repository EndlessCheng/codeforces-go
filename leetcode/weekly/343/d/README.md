### 提示 1

由于回文串去掉首尾字母后，仍然是回文串，所以长为 $m$ 的回文串必然包含长为 $m-2$ 的回文串。这等价于（逆否命题）如果没有长为 $m-2$ 的回文串，那么也不会有长为 $m$ 的回文串。

根据这一性质，「不包含任何长度为 $2$ 或更长的回文串」等价于「不包含长度为 $2$ 和长度为 $3$ 的回文串」。换句话说，不能出现 $s[i]=s[i-1]$ 以及 $s[i]=s[i-2]$。

这个性质十分重要，意味着我们**只需判断** $s[i]$ **及其左侧的两个字母**。

> $s[i]$ 及其右侧的两个字母 $s[i+1]$ 和 $s[i+2]$ 呢？交给 $i+1$ 和 $i+2$ 来判断。

### 提示 2

既然要字典序最小，那么修改的位置越靠右越好。

> 下面把 $s$ 视作一个 $k$ 进制数。也就是说，把字母 $\text{a}$ 视作 $0$，字母 $\text{b}$ 视作 $1$，依此类推。

例如 $s=\text{dacd},\ k=4$，由于答案要比 $s$ 大，先把末尾的 $s[3]=\text{d}$ 加一，进位得到 $\text{dada}$。但这样前三个字母和后三个字母都形成了回文串，我们先来解决前面的，也就是把 $s[2]=\text{d}$ 加一，进位得到 $\text{dbaa}$，这样前面就没有回文串了。反过来处理后面的回文串 $\text{aa}$，把 $s[3]=\text{a}$ 加一，得到 $\text{dbab}$，还是有回文串，再把 $s[3]=\text{b}$ 加一，得到 $\text{dbac}$，这样就没有回文串了。

请注意，题目已经保证输入的 $s$ 是美丽字符串了（不包含回文串），所以一旦发现左侧和右侧都没有回文串，那么就可以返回答案了。

如果计算过程中出现把 $s[0]$ 加一后不在前 $k$ 个字母中的情况，说明答案不存在，返回空字符串。

> 注：上面的描述中并没有用到 $k\ge 4$ 这个条件。更多相关讨论见视频讲解。

[视频讲解](https://www.bilibili.com/video/BV1QX4y1m71X/) 第四题。

```py [sol-Python3]
class Solution:
    def smallestBeautifulString(self, s: str, k: int) -> str:
        a = ord('a')
        k += a
        s = list(map(ord, s))
        n = len(s)
        i = n - 1  # 从最后一个字母开始
        s[i] += 1  # 先加一
        while i < n:
            if s[i] == k:  # 需要进位
                if i == 0:  # 无法进位
                    return ""
                # 进位
                s[i] = a
                i -= 1
                s[i] += 1
            elif i and s[i] == s[i - 1] or i > 1 and s[i] == s[i - 2]:
                s[i] += 1  # 如果 s[i] 和左侧的字符形成回文串，就继续增加 s[i]
            else:
                i += 1  # 反过来检查后面是否有回文串
        return ''.join(map(chr, s))
```

```java [sol-Java]
class Solution {
    public String smallestBeautifulString(String S, int k) {
        k += 'a';
        char[] s = S.toCharArray();
        int n = s.length;
        int i = n - 1; // 从最后一个字母开始
        s[i]++; // 先加一
        while (i < n) {
            if (s[i] == k) { // 需要进位
                if (i == 0) { // 无法进位
                    return "";
                }
                // 进位
                s[i] = 'a';
                s[--i]++;
            } else if (i > 0 && s[i] == s[i - 1] || i > 1 && s[i] == s[i - 2]) {
                s[i]++; // 如果 s[i] 和左侧的字符形成回文串，就继续增加 s[i]
            } else {
                i++; // 反过来检查后面是否有回文串
            }
        }
        return new String(s);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string smallestBeautifulString(string s, int k) {
        k += 'a';
        int n = s.length();
        int i = n - 1; // 从最后一个字母开始
        s[i]++; // 先加一
        while (i < n) {
            if (s[i] == k) { // 需要进位
                if (i == 0) { // 无法进位
                    return "";
                }
                // 进位
                s[i] = 'a';
                s[--i]++;
            } else if (i && s[i] == s[i - 1] || i > 1 && s[i] == s[i - 2]) {
                s[i]++; // 如果 s[i] 和左侧的字符形成回文串，就继续增加 s[i]
            } else {
                i++; // 反过来检查后面是否有回文串
            }
        }
        return s;
    }
};
```

```go [sol-Go]
func smallestBeautifulString(S string, k int) string {
	limit := 'a' + byte(k)
	s := []byte(S)
	n := len(s)
	i := n - 1 // 从最后一个字母开始
	s[i]++ // 先加一
	for i < n {
		if s[i] == limit { // 需要进位
			if i == 0 { // 无法进位
				return ""
			}
			// 进位
			s[i] = 'a'
			i--
			s[i]++
		} else if i > 0 && s[i] == s[i-1] || i > 1 && s[i] == s[i-2] {
			s[i]++ // 如果 s[i] 和左侧的字符形成回文串，就继续增加 s[i]
		} else {
			i++ // 反过来检查后面是否有回文串
		}
	}
	return string(s)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。注意不是 $\mathcal{O}(nk)$，因为只考虑相邻或者相隔一个字母的情况，$s[i]$ 不会增加很多次。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。如果可以直接修改字符串（例如 C++）就只需要 $\mathcal{O}(1)$ 的额外空间。

## 思考题

如果把「不包含任何长度为 $2$」中的 $2$ 改成 $3$，改成 $4$，改成一个输入的数字 $m$，要怎么做？

欢迎在评论区发表你的思路/代码。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

如果你从未做过类似题目，可以先做一些简单的，比如 [3561. 移除相邻字符](https://leetcode.cn/problems/resulting-string-after-adjacent-removals/)。

本题也是邻项消除，但消除的不是相邻的一对字符，而是一串字符。

为方便知道能否消除，栈中不能只存字符，而是存字符连续出现的次数（相当于**压缩**了连续出现的字符）。具体来说，栈中保存二元组 $(b,\textit{cnt})$，其中 $b$ 为括号字符，$\textit{cnt}$ 为该括号的连续出现次数。

具体算法如下：

1. 创建一个空栈。
2. 从左到右遍历 $s$。
3. **维护字符连续出现次数**：如果栈不为空，且 $s[i]$ 与栈顶字符相同，那么把栈顶的 $\textit{cnt}$ 加一，否则把二元组 $(s[i],1)$ 入栈。
4. **检查是否可以消除**：如果栈的大小 $\ge 2$，栈顶是右括号且个数为 $k$，并且栈顶下面（倒数第二项）左括号的个数 $\ge k$，那么就找到了一个 $k$-平衡子串，将其消除。做法是弹出栈顶，然后把新的栈顶的 $\textit{cnt}$ 减少 $k$，如果减少 $k$ 后 $\textit{cnt}=0$，那么弹出新的栈顶。

[本题视频讲解](https://www.bilibili.com/video/BV1ESxKzeEt5/?t=7m48s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def removeSubstring(self, s: str, k: int) -> str:
        st = []  # 栈中保存 [字符, 连续出现次数]
        for b in s:
            if st and st[-1][0] == b:
                st[-1][1] += 1  # 连续相同括号个数 +1
            else:
                st.append([b, 1])  # 新的括号

            # 栈顶的 k 个右括号与栈顶下面的 k 个左括号抵消
            if b == ')' and len(st) > 1 and st[-1][1] == k and st[-2][1] >= k:
                st.pop()
                st[-1][1] -= k
                if st[-1][1] == 0:
                    st.pop()

        return ''.join(b * cnt for b, cnt in st)
```

```java [sol-Java]
class Solution {
    public String removeSubstring(String s, int k) {
        List<int[]> st = new ArrayList<>(); // 栈中保存 [字符, 连续出现次数]
        for (char b : s.toCharArray()) {
            if (!st.isEmpty() && st.getLast()[0] == b) {
                st.getLast()[1]++; // 连续相同括号个数 +1
            } else {
                st.add(new int[]{b, 1}); // 新的括号
            }

            // 栈顶的 k 个右括号与栈顶下面的 k 个左括号抵消
            if (b == ')' && st.size() > 1 && st.getLast()[1] == k && st.get(st.size() - 2)[1] >= k) {
                st.removeLast();
                st.getLast()[1] -= k;
                if (st.getLast()[1] == 0) {
                    st.removeLast();
                }
            }
        }

        StringBuilder ans = new StringBuilder();
        for (int[] p : st) {
            ans.repeat(p[0], p[1]);
        }
        return ans.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string removeSubstring(string s, int k) {
        vector<pair<char, int>> st; // 栈中保存 pair{字符, 连续出现次数}
        for (char b : s) {
            if (!st.empty() && st.back().first == b) {
                st.back().second++; // 连续相同括号个数 +1
            } else {
                st.emplace_back(b, 1); // 新的括号
            }

            // 栈顶的 k 个右括号与栈顶下面的 k 个左括号抵消
            if (b == ')' && st.size() > 1 && st.back().second == k && st[st.size() - 2].second >= k) {
                st.pop_back();
                st.back().second -= k;
                if (st.back().second == 0) {
                    st.pop_back();
                }
            }
        }

        string ans;
        for (auto& p : st) {
            ans += string(p.second, p.first);
        }
        return ans;
    }
};
```

```go [sol-Go]
func removeSubstring(s string, k int) string {
	type pair struct {
		b   rune
		cnt int
	}
	st := []pair{} // 栈中保存 pair{字符, 连续出现次数}
	for _, b := range s {
		if len(st) > 0 && st[len(st)-1].b == b {
			st[len(st)-1].cnt++ // 连续相同括号个数 +1
		} else {
			st = append(st, pair{b, 1}) // 新的括号
		}

		// 栈顶的 k 个右括号与栈顶下面的 k 个左括号抵消
		if b == ')' && len(st) > 1 && st[len(st)-1].cnt == k && st[len(st)-2].cnt >= k {
			st = st[:len(st)-1]
			st[len(st)-1].cnt -= k
			if st[len(st)-1].cnt == 0 {
				st = st[:len(st)-1]
			}
		}
	}

	ans := []byte{}
	for _, p := range st {
		ans = append(ans, strings.Repeat(string(p.b), p.cnt)...)
	}
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§3.3 邻项消除**」。

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

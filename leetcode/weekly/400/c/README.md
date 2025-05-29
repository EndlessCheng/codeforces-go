每次操作，我们要去掉 $\texttt{*}$ 左边最小的字母。

如果有多个相同字母，去掉哪个？去掉更靠左的，还是更靠右的？

比如示例 1，$s=\texttt{aaba*}$：

- 去掉 $s[0]$ 或者 $s[1]$，结果为 $\texttt{aba}$。
- 去掉 $s[3]$，结果为 $\texttt{aab}$。

可以发现，如果去掉的字母靠左，那么会让更大的字母往左移，字典序更大；如果去掉的字母靠右，那么前面会保留更多更小的字母，字典序更小。

**核心思路**：每次操作应当去掉最小字母中靠后的字母。

从左到右遍历 $s$，用 $26$ 个栈记录遍历过的每种字母的下标。

遇到 $\texttt{*}$，弹出最小字母栈（第一个非空栈）的栈顶。

最后把剩余下标对应的字母按顺序串起来，即为答案。

## 写法一

```py [sol-Python3]
class Solution:
    def clearStars(self, s: str) -> str:
        stacks = [[] for _ in range(26)]
        for i, c in enumerate(s):
            if c != '*':
                stacks[ord(c) - ord('a')].append(i)
                continue
            # 找第一个非空栈，即为最小字母
            for st in stacks:
                if st:
                    st.pop()
                    break
        return ''.join(s[i] for i in sorted(chain.from_iterable(stacks)))
```

```java [sol-Java]
class Solution {
    public String clearStars(String S) {
        char[] s = S.toCharArray();
        List<Integer>[] stacks = new ArrayList[26];
        Arrays.setAll(stacks, i -> new ArrayList<>());
        for (int i = 0; i < s.length; i++) {
            if (s[i] != '*') {
                stacks[s[i] - 'a'].add(i);
                continue;
            }
            // 找第一个非空栈，即为最小字母
            for (List<Integer> st : stacks) {
                if (!st.isEmpty()) {
                    st.removeLast();
                    break;
                }
            }
        }

        List<Integer> idx = new ArrayList<>();
        for (List<Integer> st : stacks) {
            idx.addAll(st);
        }
        Collections.sort(idx);

        StringBuilder ans = new StringBuilder(idx.size());
        for (int i : idx) {
            ans.append(s[i]);
        }
        return ans.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string clearStars(string s) {
        vector<int> stacks[26];
        for (int i = 0; i < s.size(); i++) {
            if (s[i] != '*') {
                stacks[s[i] - 'a'].push_back(i);
                continue;
            }
            // 找第一个非空栈，即为最小字母
            for (auto& st : stacks) {
                if (!st.empty()) {
                    st.pop_back();
                    break;
                }
            }
        }

        vector<int> idx;
        for (auto& st : stacks) {
            idx.insert(idx.end(), st.begin(), st.end());
        }
        ranges::sort(idx);

        string ans(idx.size(), 0);
        for (int i = 0; i < idx.size(); i++) {
            ans[i] = s[idx[i]];
        }
        return ans;
    }
};
```

```go [sol-Go]
func clearStars(s string) string {
	stacks := make([][]int, 26)
	for i, c := range s {
		if c != '*' {
			stacks[c-'a'] = append(stacks[c-'a'], i)
			continue
		}
		for j, st := range stacks {
			if len(st) > 0 {
				stacks[j] = st[:len(st)-1]
				break
			}
		}
	}

	idx := []int{}
	for _, p := range stacks {
		idx = append(idx, p...)
	}
	slices.Sort(idx)

	ans := make([]byte, len(idx))
	for i, j := range idx {
		ans[i] = s[j]
	}
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma| + n\log n)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n+|\Sigma|)$。

## 写法二

把要删除的字母**标记**为 $\texttt{*}$ 号，最后去掉所有 $\texttt{*}$ 号。

这样可以避免排序。

```py [sol-Python3]
class Solution:
    def clearStars(self, s: str) -> str:
        s = list(s)
        stacks = [[] for _ in range(26)]
        for i, c in enumerate(s):
            if c != '*':
                stacks[ord(c) - ord('a')].append(i)
                continue
            for st in stacks:
                if st:
                    s[st.pop()] = '*'
                    break
        return ''.join(c for c in s if c != '*')
```

```java [sol-Java]
class Solution {
    public String clearStars(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        List<Integer>[] stacks = new ArrayList[26];
        Arrays.setAll(stacks, i -> new ArrayList<>());
        for (int i = 0; i < n; i++) {
            if (s[i] != '*') {
                stacks[s[i] - 'a'].add(i);
                continue;
            }
            for (List<Integer> st : stacks) {
                if (!st.isEmpty()) {
                    s[st.removeLast()] = '*';
                    break;
                }
            }
        }

        // 原地修改，去掉 s 中的 '*'
        int idx = 0;
        for (int i = 0; i < n; i++) {
            if (s[i] != '*') {
                s[idx++] = s[i];
            }
        }
        return new String(s, 0, idx);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string clearStars(string s) {
        stack<int> stacks[26];
        for (int i = 0; i < s.size(); i++) {
            if (s[i] != '*') {
                stacks[s[i] - 'a'].push(i);
                continue;
            }
            for (auto& st : stacks) {
                if (!st.empty()) {
                    s[st.top()] = '*';
                    st.pop();
                    break;
                }
            }
        }

        s.erase(ranges::remove(s, '*').begin(), s.end());
        return s;
    }
};
```

```go [sol-Go]
func clearStars(S string) string {
	s := []byte(S)
	stacks := make([][]int, 26)
	for i, c := range s {
		if c != '*' {
			stacks[c-'a'] = append(stacks[c-'a'], i)
			continue
		}
		for j, st := range stacks {
			if m := len(st); m > 0 {
				s[st[m-1]] = '*'
				stacks[j] = st[:m-1]
				break
			}
		}
	}

	ans := s[:0]
	for _, c := range s {
		if c != '*' {
			ans = append(ans, c)
		}
	}
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n+|\Sigma|)$。

## 写法三：位运算优化

用一个二进制数 $\textit{mask}$ 记录非空栈：

- $\textit{mask}$ 从右往左第 $i$ 位是 $0$，表示第 $i$ 个栈是空的。
- $\textit{mask}$ 从右往左第 $i$ 位是 $1$，表示第 $i$ 个栈不是空的。

那么 $\textit{mask}$ 尾零的个数，即为最小的字母。这可以 $\mathcal{O}(1)$ 求出。

详见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

```py [sol-Python3]
class Solution:
    def clearStars(self, s: str) -> str:
        s = list(s)
        stacks = [[] for _ in range(26)]
        mask = 0
        for i, c in enumerate(s):
            if c != '*':
                c = ord(c) - ord('a')
                stacks[c].append(i)
                mask |= 1 << c  # 标记第 c 个栈为非空
            else:
                lb = mask & -mask
                st = stacks[lb.bit_length() - 1]
                s[st.pop()] = '*'
                if not st:
                    mask ^= lb  # 标记栈为空
        return ''.join(c for c in s if c != '*')
```

```java [sol-Java]
class Solution {
    public String clearStars(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        List<Integer>[] stacks = new ArrayList[26];
        Arrays.setAll(stacks, i -> new ArrayList<>());
        int mask = 0;
        for (int i = 0; i < n; i++) {
            if (s[i] != '*') {
                stacks[s[i] - 'a'].add(i);
                mask |= 1 << (s[i] - 'a'); // 标记第 s[i]-'a' 个栈为非空
            } else {
                int k = Integer.numberOfTrailingZeros(mask);
                List<Integer> st = stacks[k];
                s[st.removeLast()] = '*';
                if (st.isEmpty()) {
                    mask ^= 1 << k; // 标记第 k 个栈为空
                }
            }
        }

        int idx = 0;
        for (int i = 0; i < n; i++) {
            if (s[i] != '*') {
                s[idx++] = s[i];
            }
        }
        return new String(s, 0, idx);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string clearStars(string s) {
        stack<int> stacks[26];
        uint32_t mask = 0;
        for (int i = 0; i < s.size(); i++) {
            if (s[i] != '*') {
                stacks[s[i] - 'a'].push(i);
                mask |= 1 << (s[i] - 'a'); // 标记第 s[i]-'a' 个栈为非空
            } else {
                int k = countr_zero(mask);
                auto& st = stacks[k];
                s[st.top()] = '*';
                st.pop();
                if (st.empty()) {
                    mask ^= 1 << k; // 标记第 k 个栈为空
                }
            }
        }

        s.erase(ranges::remove(s, '*').begin(), s.end());
        return s;
    }
};
```

```go [sol-Go]
func clearStars(S string) string {
	s := []byte(S)
	stacks := [26][]int{}
	mask := 0
	for i, c := range s {
		if c != '*' {
			c -= 'a'
			stacks[c] = append(stacks[c], i)
			mask |= 1 << c // 标记第 c 个栈为非空
		} else {
			k := bits.TrailingZeros(uint(mask))
			st := stacks[k]
			s[st[len(st)-1]] = '*'
			stacks[k] = st[:len(st)-1]
			if len(stacks[k]) == 0 {
				mask ^= 1 << k // 标记第 k 个栈为空
			}
		}
	}

	t := s[:0]
	for _, c := range s {
		if c != '*' {
			t = append(t, c)
		}
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n+|\Sigma|)$。

## 思考题

改成每次操作删除最大的字母（其余要求不变），要怎么做？

欢迎在评论区分享你的思路/代码。

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

做法类似 [316. 去除重复字母](https://leetcode.cn/problems/remove-duplicate-letters/)，请看 [我的题解](https://leetcode.cn/problems/remove-duplicate-letters/solutions/2381483/gen-zhao-wo-guo-yi-bian-shi-li-2ni-jiu-m-zd6u/)。

[本题视频讲解](https://www.bilibili.com/video/BV1PskxBnEP7/?t=34m13s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def lexSmallestAfterDeletion(self, s: str) -> str:
        left = Counter(s)

        st = []
        for ch in s:
            # 如果 ch 比栈顶小，移除栈顶，可以让字典序更小
            while st and ch < st[-1] and left[st[-1]] > 1:
                left[st.pop()] -= 1
            st.append(ch)

        # 最后，移除末尾的重复字母，可以让字典序更小
        while left[st[-1]] > 1:
            left[st.pop()] -= 1

        return ''.join(st)
```

```java [sol-Java]
class Solution {
    public String lexSmallestAfterDeletion(String S) {
        char[] s = S.toCharArray();
        int[] left = new int[26];
        for (char ch : s) {
            left[ch - 'a']++;
        }

        StringBuilder st = new StringBuilder();
        for (char ch : s) {
            // 如果 ch 比栈顶小，移除栈顶，可以让字典序更小
            while (!st.isEmpty() && ch < st.charAt(st.length() - 1) && left[st.charAt(st.length() - 1) - 'a'] > 1) {
                left[st.charAt(st.length() - 1) - 'a']--;
                st.setLength(st.length() - 1);
            }
            st.append(ch);
        }

        // 最后，移除末尾的重复字母，可以让字典序更小
        while (left[st.charAt(st.length() - 1) - 'a'] > 1) {
            left[st.charAt(st.length() - 1) - 'a']--;
            st.setLength(st.length() - 1);
        }

        return st.toString();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string lexSmallestAfterDeletion(string s) {
        int left[26]{};
        for (char ch : s) {
            left[ch - 'a']++;
        }

        string st;
        for (char ch : s) {
            // 如果 ch 比栈顶小，移除栈顶，可以让字典序更小
            while (!st.empty() && ch < st.back() && left[st.back() - 'a'] > 1) {
                left[st.back() - 'a']--;
                st.pop_back();
            }
            st.push_back(ch);
        }

        // 最后，移除末尾的重复字母，可以让字典序更小
        while (left[st.back() - 'a'] > 1) {
            left[st.back() - 'a']--;
            st.pop_back();
        }

        return st;
    }
};
```

```cpp [sol-C++ 原地]
class Solution {
public:
    string lexSmallestAfterDeletion(string s) {
        int left[26]{};
        for (char ch : s) {
            left[ch - 'a']++;
        }

        int top = -1; // 把 s 当作栈
        for (char ch : s) {
            // 如果 ch 比栈顶小，移除栈顶，可以让字典序更小
            while (top >= 0 && ch < s[top] && left[s[top] - 'a'] > 1) {
                left[s[top] - 'a']--;
                top--;
            }
            s[++top] = ch;
        }

        // 最后，移除末尾的重复字母，可以让字典序更小
        while (left[s[top] - 'a'] > 1) {
            left[s[top] - 'a']--;
            top--;
        }

        s.resize(top + 1);
        return s;
    }
};
```

```go [sol-Go]
func lexSmallestAfterDeletion(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}

	st := []rune{}
	for _, ch := range s {
		// 如果 ch 比栈顶小，移除栈顶，可以让字典序更小
		for len(st) > 0 && ch < st[len(st)-1] && left[st[len(st)-1]-'a'] > 1 {
			left[st[len(st)-1]-'a']--
			st = st[:len(st)-1]
		}
		st = append(st, ch)
	}

	// 最后，移除末尾的重复字母，可以让字典序更小
	for left[st[len(st)-1]-'a'] > 1 {
		left[st[len(st)-1]-'a']--
		st = st[:len(st)-1]
	}

	return string(st)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n + |\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n + |\Sigma|)$ 或 $\mathcal{O}(|\Sigma|)$。C++ 不计入返回值，空间复杂度为 $\mathcal{O}(|\Sigma|)$。

## 专题训练

见下面单调栈题单的「**四、最小字典序**」。

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

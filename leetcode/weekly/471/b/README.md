枚举子串左端点 $i$，然后枚举子串右端点 $j=i,i+1,i+2,\ldots,n-1$。

在枚举右端点 $j$ 的过程中，统计子串 $[i,j]$ 每种字母的出现次数 $\textit{cnt}$。

遍历 $\textit{cnt}$，如果所有字母的出现次数均相同，用子串长度 $j-i+1$ 更新答案的最大值。

[本题视频讲解](https://www.bilibili.com/video/BV1FJ4uz1EkN/?t=1m20s)，欢迎点赞关注~

## 优化前

```py [sol-Python3]
class Solution:
    def longestBalanced(self, s: str) -> int:
        ans = 0
        n = len(s)
        for i in range(n):
            cnt = defaultdict(int)
            for j in range(i, n):
                cnt[s[j]] += 1
                if len(set(cnt.values())) == 1:
                    ans = max(ans, j - i + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestBalanced(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int ans = 0;
        for (int i = 0; i < n; i++) {
            int[] cnt = new int[26];
            next:
            for (int j = i; j < n; j++) {
                int base = ++cnt[s[j] - 'a'];
                for (int c : cnt) {
                    if (c > 0 && c != base) {
                        continue next;
                    }
                }
                ans = Math.max(ans, j - i + 1);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestBalanced(string s) {
        int n = s.size();
        int ans = 0;
        for (int i = 0; i < n; i++) {
            int cnt[26]{};
            for (int j = i; j < n; j++) {
                int base = ++cnt[s[j] - 'a'];
                for (int c : cnt) {
                    if (c && c != base) {
                        base = -1;
                        break;
                    }
                }
                if (base != -1) {
                    ans = max(ans, j - i + 1);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestBalanced(s string) (ans int) {
	for i := range s {
		cnt := make([]int, 26)
	next:
		for j := i; j < len(s); j++ {
			cnt[s[j]-'a']++
			base := cnt[s[j]-'a']
			for _, c := range cnt {
				if c > 0 && c != base {
					continue next
				}
			}
			ans = max(ans, j-i+1)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 优化

设 $\textit{mx} = \max(\textit{cnt})$，设 $\textit{kinds}$ 为子串中的不同字母个数。

如果 $\textit{mx}\cdot \textit{kinds} = j-i+1$，说明子串所有字母的出现次数均为 $\textit{mx}$，均相等。

```py [sol-Python3]
# 手写 max 更快
max = lambda a, b: b if b > a else a

class Solution:
    def longestBalanced(self, s: str) -> int:
        ans = 0
        for i in range(len(s)):
            cnt = defaultdict(int)
            mx = 0
            for j in range(i, len(s)):
                cnt[s[j]] += 1
                mx = max(mx, cnt[s[j]])
                if mx * len(cnt) == j - i + 1:
                    ans = max(ans, j - i + 1)
        return ans
```

```java [sol-Java]
class Solution {
    public int longestBalanced(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int ans = 0;
        for (int i = 0; i < n; i++) {
            int[] cnt = new int[26];
            int mx = 0, kinds = 0;
            for (int j = i; j < n; j++) {
                int b = s[j] - 'a';
                if (cnt[b] == 0) {
                    kinds++;
                }
                mx = Math.max(mx, ++cnt[b]);
                if (mx * kinds == j - i + 1) {
                    ans = Math.max(ans, j - i + 1);
                }
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestBalanced(string s) {
        int n = s.size();
        int ans = 0;
        for (int i = 0; i < n; i++) {
            int cnt[26]{};
            int mx = 0, kinds = 0;
            for (int j = i; j < n; j++) {
                int b = s[j] - 'a';
                if (cnt[b] == 0) {
                    kinds++;
                }
                mx = max(mx, ++cnt[b]);
                if (mx * kinds == j - i + 1) {
                    ans = max(ans, j - i + 1);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestBalanced(s string) (ans int) {
	for i := range s {
		cnt := [26]int{}
		mx, kinds := 0, 0
		for j := i; j < len(s); j++ {
			b := s[j] - 'a'
			if cnt[b] == 0 {
				kinds++
			}
			cnt[b]++
			mx = max(mx, cnt[b])
			if mx*kinds == j-i+1 {
				ans = max(ans, j-i+1)
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$，其中 $|\Sigma|=26$ 是字符集合的大小。

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

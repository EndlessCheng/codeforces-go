枚举子串左端点 $i$，然后枚举子串右端点 $j=i,i+1,i+2,\ldots,n-1$。

在枚举右端点 $j$ 的过程中，统计子串 $[i,j]$ 每种字母的出现次数 $\textit{cnt}$。

遍历 $\textit{cnt}$，如果所有字母的出现次数均相同，用子串长度 $j-i+1$ 更新答案的最大值。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

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
                cnt[s[j] - 'a']++;
                int base = 0;
                for (int c : cnt) {
                    if (c == 0) {
                        continue;
                    }
                    if (base == 0) {
                        base = c;
                    } else if (c != base) {
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
                cnt[s[j] - 'a']++;
                int base = 0;
                for (int c : cnt) {
                    if (c == 0) {
                        continue;
                    }
                    if (base == 0) {
                        base = c;
                    } else if (c != base) {
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
			base := 0
			for _, c := range cnt {
				if c == 0 {
					continue
				}
				if base == 0 {
					base = c
				} else if c != base {
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

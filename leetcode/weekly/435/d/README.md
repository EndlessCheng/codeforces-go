## 前置知识

1. [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/)，推荐先完成 [1371. 每个元音包含偶数次的最长子字符串](https://leetcode.cn/problems/find-the-longest-substring-containing-vowels-in-even-counts/)，对本题做法有一定启发。
2. [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

## 引入前缀和

枚举 $a$ 和 $b$ 分别是哪两个数字。

定义 $\textit{sum}[i+1][j]$ 表示 $s[0]$ 到 $s[i]$ 中的 $j$ 的出现次数。

我们要计算的是

$$
\begin{aligned}
    & (\textit{sum}[r][x] - \textit{sum}[l][x]) - (\textit{sum}[r][y] - \textit{sum}[l][y])      \\
={} & (\textit{sum}[r][x] - \textit{sum}[r][y]) - (\textit{sum}[l][x] - \textit{sum}[l][y])      \\
\end{aligned}
$$

的最大值，其中：

- 子串对应的下标区间为 $[l,r)$。
- $r-l\ge k$
- 子串必须包含奇数个 $x$，那么至少要满足 $\textit{sum}[r][x] > \textit{sum}[l][x]$。
- 子串必须包含正偶数个 $y$，那么至少要满足 $\textit{sum}[r][y] > \textit{sum}[l][y]$。

枚举 $r$（枚举右），问题变成计算满足上述约束的

$$
\textit{sum}[l][x] - \textit{sum}[l][y]
$$

的最小值（维护左）。

## 处理奇偶性

题目奇偶性的要求，等价于：

- $\textit{sum}[r][x]$ 的奇偶性必须与 $\textit{sum}[l][x]$ **不同**。
- $\textit{sum}[r][y]$ 的奇偶性必须与 $\textit{sum}[l][y]$ **相同**。

$x$ 个数的奇偶和 $y$ 个数的奇偶两两组合，有 $4$ 种情况，我们需要维护 $4$ 种最小前缀和。

定义 $\textit{minS}[p][q]$ 表示最小的 $\textit{sum}[l][x] - \textit{sum}[l][y]$，其中

- $p=0$ 表示偶数，$p=1$ 表示奇数，$q$ 同理。
- $x$ 在下标 $[0,l)$ 中的出现次数的奇偶性为 $p$。
- $y$ 在下标 $[0,l)$ 中的出现次数的奇偶性为 $q$。

并满足以下条件（把 $r$ 当作定值）：  

- $r-l\ge k$
- $\textit{sum}[r][x] > \textit{sum}[l][x]$
- $\textit{sum}[r][y] > \textit{sum}[l][y]$

由于子串越长，越能满足上述要求，有单调性，所以用**滑动窗口**维护 $l+1$ 的最大值 $\textit{left}$，同时维护相应的 $\textit{minS}[p][q]$。

⚠**注意**：我们要维护的是窗口左边的 $4$ 种最小前缀和，并不关心窗口内的东西。

内层循环结束后，在 $[0,\textit{left}-1]$ 中的左端点 $l$ 都是符合要求的，并且相应的最小前缀和也已保存到 $\textit{minS}[p][q]$ 中。

此时用

$$
(\textit{sum}[r][x] - \textit{sum}[r][y]) - \textit{minS}[p][q]
$$

更新答案的最大值，其中 $p=1-(\textit{sum}[r][x]\bmod 2)$，$q = \textit{sum}[r][y]\bmod 2$。

代码实现时，可以一边滑窗，一边计算前缀和。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1D5F6eRECp/?t=49m08s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxDifference(self, s: str, k: int) -> int:
        s = list(map(int, s))
        ans = -inf
        for x in range(5):
            for y in range(5):
                if y == x:
                    continue
                cur_s = [0] * 5
                pre_s = [0] * 5
                min_s = [[inf, inf], [inf, inf]]
                left = 0
                for i, v in enumerate(s):
                    cur_s[v] += 1
                    r = i + 1
                    while r - left >= k and cur_s[x] > pre_s[x] and cur_s[y] > pre_s[y]:
                        p, q = pre_s[x] & 1, pre_s[y] & 1
                        min_s[p][q] = min(min_s[p][q], pre_s[x] - pre_s[y])
                        pre_s[s[left]] += 1
                        left += 1
                    if r >= k:
                        ans = max(ans, cur_s[x] - cur_s[y] - min_s[cur_s[x] & 1 ^ 1][cur_s[y] & 1])
        return ans
```

```java [sol-Java]
class Solution {
    public int maxDifference(String S, int k) {
        final int INF = Integer.MAX_VALUE / 2;
        char[] s = S.toCharArray();
        int ans = -INF;
        for (int x = 0; x < 5; x++) {
            for (int y = 0; y < 5; y++) {
                if (y == x) {
                    continue;
                }
                int[] curS = new int[5];
                int[] preS = new int[5];
                int[][] minS = {{INF, INF}, {INF, INF}};
                int left = 0;
                for (int i = 0; i < s.length; i++) {
                    curS[s[i] - '0']++;
                    int r = i + 1;
                    while (r - left >= k && curS[x] > preS[x] && curS[y] > preS[y]) {
                        int p = preS[x] & 1;
                        int q = preS[y] & 1;
                        minS[p][q] = Math.min(minS[p][q], preS[x] - preS[y]);
                        preS[s[left] - '0']++;
                        left++;
                    }
                    ans = Math.max(ans, curS[x] - curS[y] - minS[curS[x] & 1 ^ 1][curS[y] & 1]);
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
    int maxDifference(string s, int k) {
        const int inf = INT_MAX / 2;
        int ans = -inf;
        for (int x = 0; x < 5; x++) {
            for (int y = 0; y < 5; y++) {
                if (y == x) {
                    continue;
                }
                int cur_s[5]{}, pre_s[5]{};
                int min_s[2][2] = {{inf, inf}, {inf, inf}};
                int left = 0;
                for (int i = 0; i < s.size(); i++) {
                    cur_s[s[i] - '0']++;
                    int r = i + 1;
                    while (r - left >= k && cur_s[x] > pre_s[x] && cur_s[y] > pre_s[y]) {
                        int& p = min_s[pre_s[x] & 1][pre_s[y] & 1];
                        p = min(p, pre_s[x] - pre_s[y]);
                        pre_s[s[left] - '0']++;
                        left++;
                    }
                    ans = max(ans, cur_s[x] - cur_s[y] - min_s[cur_s[x] & 1 ^ 1][cur_s[y] & 1]);
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func maxDifference(s string, k int) int {
	const inf = math.MaxInt / 2
	ans := -inf
	for x := range 5 {
		for y := range 5 {
			if y == x {
				continue
			}
			curS := [5]int{}
			preS := [5]int{}
			minS := [2][2]int{{inf, inf}, {inf, inf}}
			left := 0
			for i, b := range s {
				curS[b-'0']++
				r := i + 1
				for r-left >= k && curS[x] > preS[x] && curS[y] > preS[y] {
					p := &minS[preS[x]&1][preS[y]&1]
					*p = min(*p, preS[x]-preS[y])
					preS[s[left]-'0']++
					left++
				}
                ans = max(ans, curS[x]-curS[y]-minS[curS[x]&1^1][curS[y]&1])
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|^2)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=5$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

更多相似题目，见下面数据结构题单中的「**§1.2 前缀和与哈希表**」和「**§1.4 前缀异或和**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. 【本题相关】[常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

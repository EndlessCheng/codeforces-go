如果不允许交换，本题就是 [525. 连续数组](https://leetcode.cn/problems/contiguous-array/)。下面接着 [我的题解](https://leetcode.cn/problems/contiguous-array/solutions/3805089/shi-zi-bian-xing-mei-ju-you-wei-hu-zuo-p-x9q2/) 继续讲。

如果把子串中的一个 $\texttt{1}$ 和外面的一个 $\texttt{0}$ 交换后，子串变成平衡的，那么交换前，子串中的 $\texttt{1}$ 的个数恰好比 $\texttt{0}$ 的个数多 $2$。用 525 题的做法来说，就是子串的元素和恰好等于 $2$。

用前缀和表示，就是 $\textit{sum}[i] - \textit{sum}[j] = 2$。枚举 $i$，为了让子串长度 $i-j$ 尽量长，我们需要找到 $\textit{sum}[i] - 2$ 的最小的出现位置 $j$。所以要用哈希表（或者数组）记录每个前缀和首次出现的位置。

但如果子串外面没有 $\texttt{0}$ 呢？设 $s$ 一共有 $\textit{total}_0$ 个 $\texttt{0}$，分类讨论：

- 子串长度为 $i-j$，其中有 $\dfrac{i-j-2}{2}$ 个 $\texttt{0}$。
- 如果 $\dfrac{i-j-2}{2} < \textit{total}_0$，那么子串外面还有 $\texttt{0}$，可以交换。
- 如果 $\dfrac{i-j-2}{2} = \textit{total}_0$，那么子串外面没有 $\texttt{0}$，无法交换。我们可以额外记录每个前缀和**第二次**出现的位置 $j'$。由于同一个前缀和两次出现的位置之间的元素和为 $0$，所以两个位置之间的 $\texttt{0}$ 和 $\texttt{1}$ 的个数相等且都大于 $0$，所以长为 $i-j'$ 的这个子串的外面一定有 $\texttt{0}$，可以完成交换。所以只需维护每个前缀和前两次出现的位置。

对于把子串中的一个 $\texttt{0}$ 和外面的一个 $\texttt{1}$ 交换的情况，做法是类似的，枚举 $i$，找 $\textit{sum}[i] + 2$ 前两次出现的位置。

> **注**：本题允许子串为空，可以先更新哈希表，再计算子串长度。

[本题视频讲解](https://www.bilibili.com/video/BV1JNDQBBE7n/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def longestBalanced(self, s: str) -> int:
        total0 = s.count('0')
        total1 = len(s) - total0

        pos = defaultdict(list)
        pos[0] = [-1]  # 见 525 题
        ans = 0
        pre = 0  # 前缀和

        for i, ch in enumerate(s):
            pre += 1 if ch == '1' else -1

            if len(pos[pre]) < 2:
                pos[pre].append(i)

            # 不交换
            ans = max(ans, i - pos[pre][0])

            # 交换子串内的一个 1 和子串外的一个 0
            if pre - 2 in pos:
                p = pos[pre - 2]
                if (i - p[0] - 2) // 2 < total0:
                    ans = max(ans, i - p[0])
                elif len(p) > 1:
                    ans = max(ans, i - p[1])

            # 交换子串内的一个 0 和子串外的一个 1
            if pre + 2 in pos:
                p = pos[pre + 2]
                if (i - p[0] - 2) // 2 < total1:
                    ans = max(ans, i - p[0])
                elif len(p) > 1:
                    ans = max(ans, i - p[1])

        return ans
```

```java [sol-Java]
class Solution {
    public int longestBalanced(String S) {
        char[] s = S.toCharArray();
        int total0 = 0;
        for (char c : s) {
            if (c == '0') {
                total0++;
            }
        }
        int total1 = s.length - total0;

        Map<Integer, List<Integer>> pos = new HashMap<>();
        pos.computeIfAbsent(0, _ -> new ArrayList<>()).add(-1); // 见 525 题

        int ans = 0;
        int sum = 0; // 前缀和

        for (int i = 0; i < s.length; i++) {
            sum += (s[i] - '0') * 2 - 1;

            List<Integer> p = pos.computeIfAbsent(sum, _ -> new ArrayList<>());
            if (p.size() < 2) {
                p.add(i);
            }

            // 不交换
            ans = Math.max(ans, i - p.get(0));

            // 交换子串内的一个 1 和子串外的一个 0
            p = pos.get(sum - 2);
            if (p != null) {
                if ((i - p.get(0) - 2) / 2 < total0) {
                    ans = Math.max(ans, i - p.get(0));
                } else if (p.size() > 1) {
                    ans = Math.max(ans, i - p.get(1));
                }
            }

            // 交换子串内的一个 0 和子串外的一个 1
            p = pos.get(sum + 2);
            if (p != null) {
                if ((i - p.get(0) - 2) / 2 < total1) {
                    ans = Math.max(ans, i - p.get(0));
                } else if (p.size() > 1) {
                    ans = Math.max(ans, i - p.get(1));
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
        int total0 = ranges::count(s, '0');
        int total1 = s.size() - total0;

        unordered_map<int, vector<int>> pos = {{0, {-1}}}; // 见 525 题
        int ans = 0;
        int sum = 0; // 前缀和

        for (int i = 0; i < s.size(); i++) {
            sum += (s[i] - '0') * 2 - 1;

            if (pos[sum].size() < 2) {
                pos[sum].push_back(i);
            }

            // 不交换
            ans = max(ans, i - pos[sum][0]);

            // 交换子串内的一个 1 和子串外的一个 0
            auto it = pos.find(sum - 2);
            if (it != pos.end()) {
                auto& p = it->second;
                if ((i - p[0] - 2) / 2 < total0) {
                    ans = max(ans, i - p[0]);
                } else if (p.size() > 1) {
                    ans = max(ans, i - p[1]);
                }
            }

            // 交换子串内的一个 0 和子串外的一个 1
            it = pos.find(sum + 2);
            if (it != pos.end()) {
                auto& p = it->second;
                if ((i - p[0] - 2) / 2 < total1) {
                    ans = max(ans, i - p[0]);
                } else if (p.size() > 1) {
                    ans = max(ans, i - p[1]);
                }
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func longestBalanced(s string) (ans int) {
	total0 := strings.Count(s, "0")
	total1 := len(s) - total0

	pos := map[int][]int{0: {-1}} // 见 525 题
	sum := 0 // 前缀和
	for i, ch := range s {
		sum += int(ch-'0')*2 - 1

		if p := pos[sum]; len(p) < 2 {
			pos[sum] = append(p, i)
		}

		// 不交换
		ans = max(ans, i-pos[sum][0])

		// 交换子串内的一个 1 和子串外的一个 0
		if p, ok := pos[sum-2]; ok {
			if (i-p[0]-2)/2 < total0 {
				ans = max(ans, i-p[0])
			} else if len(p) > 1 {
				ans = max(ans, i-p[1])
			}
		}

		// 交换子串内的一个 0 和子串外的一个 1
		if p, ok := pos[sum+2]; ok {
			if (i-p[0]-2)/2 < total1 {
				ans = max(ans, i-p[0])
			} else if len(p) > 1 {
				ans = max(ans, i-p[1])
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

**注**：把哈希表换成数组会更快一些。

## 专题训练

见下面数据结构题单的「**§1.2 前缀和与哈希表**」。

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

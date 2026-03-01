要想让答案尽量小，那么答案二进制的高位是 $0$ 比是 $1$ 更好，所以优先判断答案的高位能否是 $0$，即**从高到低**依次判断答案的第 $i$ 位能不能是 $0$。

如果在每一行的**能选的数字**中，都存在二进制第 $i$ 位是 $0$ 的数，那么答案的第 $i$ 位可以是 $0$，否则必须是 $1$。

怎么判断一个数能不能选？

比如答案现在是 $101\texttt{\_}\texttt{\_}$，那么 $\textit{grid}[i][j]$ 绝对不能是 $\texttt{\_}1\texttt{\_}\texttt{\_}\texttt{\_}$，这会导致 OR 的结果是 $\texttt{\_}1\texttt{\_}\texttt{\_}\texttt{\_}$，不符合目前的答案。

一般地，对于答案中的是 $0$ 的比特位，如果 $\textit{grid}[i][j]$ 二进制同一位上是 $1$，那么 $\textit{grid}[i][j]$ 不能选，否则可以选。从集合的角度理解，对于目前已经确定的比特位，$\textit{grid}[i][j]$ 必须是答案的**子集**。具体请看 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

[本题视频讲解](https://www.bilibili.com/video/BV1V4PMzrEYG/?t=9m50s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minimumOR(self, grid: List[List[int]]) -> int:
        mx = max(map(max, grid))
        ans = 0
        # 试填法：ans 的第 i 位能不能是 0？
        # 如果在每一行的能选的数字中，都存在第 i 位是 0 的数，那么 ans 的第 i 位可以是 0，否则必须是 1
        for i in range(mx.bit_length() - 1, -1, -1):
            mask = ans | ((1 << i) - 1)  # mask 低于 i 的比特位全是 1，表示 grid[i][j] 的低位是 0 还是 1 无所谓
            for row in grid:
                for x in row:
                    # x 的高于 i 的比特位，如果 ans 是 0，那么 x 的这一位必须也是 0
                    # x 的低于 i 的比特位，随意
                    # x 的第 i 个比特位，我们期望它是 0
                    if (x | mask) == mask:  # x 可以选，且第 i 位是 0
                        break
                else:  # 这一行的可选数字中，第 i 位全是 1
                    ans |= 1 << i  # ans 第 i 位必须是 1
                    break  # 填下一位
        return ans
```

```java [sol-Java]
class Solution {
    public int minimumOR(int[][] grid) {
        int mx = 0;
        for (int[] row : grid) {
            for (int x : row) {
                mx = Math.max(mx, x);
            }
        }
        int bitLength = 32 - Integer.numberOfLeadingZeros(mx);

        int ans = 0;
        // 试填法：ans 的第 i 位能不能是 0？
        // 如果在每一行的能选的数字中，都存在第 i 位是 0 的数，那么 ans 的第 i 位可以是 0，否则必须是 1
        for (int i = bitLength - 1; i >= 0; i--) {
            int mask = ans | ((1 << i) - 1); // mask 低于 i 的比特位全是 1，表示 grid[i][j] 的低位是 0 还是 1 无所谓
            next:
            for (int[] row : grid) {
                for (int x : row) {
                    // x 的高于 i 的比特位，如果 ans 是 0，那么 x 的这一位必须也是 0
                    // x 的低于 i 的比特位，随意
                    // x 的第 i 个比特位，我们期望它是 0
                    if ((x | mask) == mask) { // x 可以选，且第 i 位是 0
                        continue next;
                    }
                }
                // 这一行的可选数字中，第 i 位全是 1
                ans |= 1 << i; // ans 第 i 位必须是 1
                break; // 填下一位
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumOR(vector<vector<int>>& grid) {
        int mx = 0;
        for (auto& row : grid) {
            mx = max(mx, ranges::max(row));
        }

        int ans = 0;
        // 试填法：ans 的第 i 位能不能是 0？
        // 如果在每一行的能选的数字中，都存在第 i 位是 0 的数，那么 ans 的第 i 位可以是 0，否则必须是 1
        for (int i = bit_width((uint32_t) mx) - 1; i >= 0; i--) {
            int mask = ans | ((1 << i) - 1); // mask 低于 i 的比特位全是 1，表示 grid[i][j] 的低位是 0 还是 1 无所谓
            for (auto& row : grid) {
                bool found0 = false;
                for (int x : row) {
                    // x 的高于 i 的比特位，如果 ans 是 0，那么 x 的这一位必须也是 0
                    // x 的低于 i 的比特位，随意
                    // x 的第 i 个比特位，我们期望它是 0
                    if ((x | mask) == mask) { // x 可以选，且第 i 位是 0
                        found0 = true;
                        break;
                    }
                }
                if (!found0) { // 这一行的可选数字中，第 i 位全是 1
                    ans |= 1 << i; // ans 第 i 位必须是 1
                    break; // 填下一位
                }
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumOR(grid [][]int) (ans int) {
	mx := 0
	for _, row := range grid {
		mx = max(mx, slices.Max(row))
	}

	// 试填法：ans 的第 i 位能不能是 0？
	// 如果在每一行的能选的数字中，都存在第 i 位是 0 的数，那么 ans 的第 i 位可以是 0，否则必须是 1
	for i := bits.Len(uint(mx)) - 1; i >= 0; i-- {
		mask := ans | (1<<i - 1) // mask 低于 i 的比特位全是 1，表示 grid[i][j] 的低位是 0 还是 1 无所谓
	next:
		for _, row := range grid {
			for _, x := range row {
				// x 的高于 i 的比特位，如果 ans 是 0，那么 x 的这一位必须也是 0
				// x 的低于 i 的比特位，随意
				// x 的第 i 个比特位，我们期望它是 0
				if x|mask == mask { // x 可以选，且第 i 位是 0
					continue next
				}
			}
			// 这一行的可选数字中，第 i 位全是 1
			ans |= 1 << i // ans 第 i 位必须是 1
			break // 填下一位
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\log U)$，其中 $m$ 和 $n$ 分别是 $\textit{grid}$ 的行数和列数，$U$ 是所有 $\textit{grid}[i][j]$ 中的最大值。注意题目保证 $m\cdot n\le 10^5$。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

改成求 OR 的最大值，怎么做？

欢迎在评论区分享你的思路/代码。

## 专题训练

见下面位运算题单的「**五、试填法**」。

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

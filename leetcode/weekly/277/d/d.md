由于至多有 $n=15$ 个人，我们可以枚举这 $n$ 个人谁是好人，谁是坏人，这一共有 $2^n$ 种不同的情况。

根据 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/) 中的「枚举子集」的技巧，枚举好人子集。

用二进制数表示这 $n$ 个人中谁好谁坏，其中 $1$ 表示好人，$0$ 表示坏人。这样就可以枚举 $i\in [0, 2^n-1]$ 中的所有数字，然后判断 $i$ 中好人的陈述是否与实际情况矛盾，若不矛盾则 $i$ 为一种合法的情况。所有合法情况中的好人个数的最大值即为答案。

代码实现时，可以从 $i=1$ 开始枚举。

```py [sol-Python3]
class Solution:
    def maximumGood(self, statements: List[List[int]]) -> int:
        def check(i: int) -> int:
            cnt = 0  # i 中好人个数
            for j, row in enumerate(statements):  # 枚举 i 中的好人 j
                if i >> j & 1:
                    if any(st < 2 and st != i >> k & 1 for k, st in enumerate(row)):
                        return 0  # 好人 j 的某个陈述 st 与实际情况矛盾
                    cnt += 1
            return cnt

        return max(check(i) for i in range(1, 1 << len(statements)))
```

```java [sol-Java]
class Solution {
    public int maximumGood(int[][] statements) {
        int n = statements.length;
        int ans = 0;
        next:
        for (int i = 1; i < 1 << n; i++) {
            int cnt = 0; // i 中好人个数
            for (int j = 0; j < n; j++) {
                if ((i >> j & 1) > 0) { // 枚举 i 中的好人 j
                    for (int k = 0; k < n; k++) { // 枚举 j 的所有陈述
                        if (statements[j][k] < 2 && statements[j][k] != (i >> k & 1)) { // 该陈述与实际情况矛盾
                            continue next;
                        }
                    }
                    cnt++;
                }
            }
            ans = Math.max(ans, cnt);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumGood(vector<vector<int>>& statements) {
        int ans = 0, n = statements.size();
        for (int i = 1; i < 1 << n; i++) {
            int cnt = 0; // i 中好人个数
            for (int j = 0; j < n; j++) {
                if (i >> j & 1) { // 枚举 i 中的好人 j
                    for (int k = 0; k < n; k++) { // 枚举 j 的所有陈述
                        if (statements[j][k] < 2 && statements[j][k] != (i >> k & 1)) { // 该陈述与实际情况矛盾
                            goto next;
                        }
                    }
                    ++cnt;
                }
            }
            ans = max(ans, cnt);
            next:;
        }
        return ans;
    }
};
```

```go [sol-Go]
func maximumGood(statements [][]int) (ans int) {
next:
	for i := 1; i < 1<<len(statements); i++ {
		cnt := 0 // i 中好人个数
		for j, row := range statements {
			if i>>j&1 > 0 { // 枚举 i 中的好人 j
				for k, st := range row { // 枚举 j 的所有陈述 st
					if st < 2 && st != i>>k&1 { // 该陈述与实际情况矛盾
						continue next
					}
				}
				cnt++
			}
		}
		ans = max(ans, cnt)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^22^n)$，其中 $n$ 是 $\textit{statements }$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

如果坏人一定说假话呢？

这题是 [CF1594D](https://codeforces.com/problemset/problem/1594/D)。

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

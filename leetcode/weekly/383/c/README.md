1. 遍历所有 $3 \times 3$ 的子网格。
2. 遍历网格内的所有左右相邻格子和上下相邻格子，如果存在差值超过 $\textit{threshold}$ 的情况，则枚举下一个子网格。
3. 如果合法，计算子网格的平均值 $\textit{avg}$，等于子网格的元素和除以 $9$ 下取整。
4. 更新子网格内的 $\textit{result}[i][j]$，由于需要计算平均值，我们先把 $\textit{avg}$ 加到 $\textit{result}[i][j]$ 中，同时用一个 $\textit{cnt}$ 矩阵统计 $(i,j)$ 在多少个合法子网格内。
5. 最后返回答案。如果 $\textit{cnt}[i][j]=0$ 则 $\textit{result}[i][j] = \textit{image}[i][j]$，否则 $\textit{result}[i][j] = \left\lfloor\dfrac{\textit{result}[i][j]}{\textit{cnt}[i][j]}\right\rfloor$。

```py [sol-Python3]
class Solution:
    def resultGrid(self, a: List[List[int]], threshold: int) -> List[List[int]]:
        m, n = len(a), len(a[0])
        result = [[0] * n for _ in range(m)]
        cnt = [[0] * n for _ in range(m)]
        for i in range(2, m):
            for j in range(2, n):
                # 检查左右相邻格子
                ok = True
                for row in a[i - 2: i + 1]:
                    if abs(row[j - 2] - row[j - 1]) > threshold or abs(row[j - 1] - row[j]) > threshold:
                        ok = False
                        break  # 不合法，下一个
                if not ok: continue

                # 检查上下相邻格子
                for y in range(j - 2, j + 1):
                    if abs(a[i - 2][y] - a[i - 1][y]) > threshold or abs(a[i - 1][y] - a[i][y]) > threshold:
                        ok = False
                        break  # 不合法，下一个
                if not ok: continue

                # 合法，计算 3x3 子网格的平均值
                avg = sum(a[x][y] for x in range(i - 2, i + 1) for y in range(j - 2, j + 1)) // 9

                # 更新 3x3 子网格内的 result
                for x in range(i - 2, i + 1):
                    for y in range(j - 2, j + 1):
                        result[x][y] += avg  # 先累加，最后再求平均值
                        cnt[x][y] += 1

        for i, row in enumerate(cnt):
            for j, c in enumerate(row):
                if c == 0:  # (i,j) 不属于任何子网格
                    result[i][j] = a[i][j]
                else:
                    result[i][j] //= c  # 求平均值
        return result
```

```java [sol-Java]
class Solution {
    public int[][] resultGrid(int[][] a, int threshold) {
        int m = a.length;
        int n = a[0].length;
        int[][] result = new int[m][n];
        int[][] cnt = new int[m][n];
        for (int i = 2; i < m; i++) {
            next:
            for (int j = 2; j < n; j++) {
                // 检查左右相邻格子
                for (int x = i - 2; x <= i; x++) {
                    if (Math.abs(a[x][j - 2] - a[x][j - 1]) > threshold || Math.abs(a[x][j - 1] - a[x][j]) > threshold) {
                        continue next; // 不合法，下一个
                    }
                }

                // 检查上下相邻格子
                for (int y = j - 2; y <= j; y++) {
                    if (Math.abs(a[i - 2][y] - a[i - 1][y]) > threshold || Math.abs(a[i - 1][y] - a[i][y]) > threshold) {
                        continue next; // 不合法，下一个
                    }
                }

                // 合法，计算 3x3 子网格的平均值
                int avg = 0;
                for (int x = i - 2; x <= i; x++) {
                    for (int y = j - 2; y <= j; y++) {
                        avg += a[x][y];
                    }
                }
                avg /= 9;

                // 更新 3x3 子网格内的 result
                for (int x = i - 2; x <= i; x++) {
                    for (int y = j - 2; y <= j; y++) {
                        result[x][y] += avg; // 先累加，最后再求平均值
                        cnt[x][y]++;
                    }
                }
            }
        }

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (cnt[i][j] == 0) { // (i,j) 不属于任何子网格
                    result[i][j] = a[i][j];
                } else {
                    result[i][j] /= cnt[i][j]; // 求平均值
                }
            }
        }
        return result;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> resultGrid(vector<vector<int>>& a, int threshold) {
        int m = a.size(), n = a[0].size();
        vector<vector<int>> result(m, vector<int>(n));
        vector<vector<int>> cnt(m, vector<int>(n));
        for (int i = 2; i < m; i++) {
            for (int j = 2; j < n; j++) {
                // 检查左右相邻格子
                bool ok = true;
                for (int x = i - 2; x <= i; x++) {
                    if (abs(a[x][j - 2] - a[x][j - 1]) > threshold || abs(a[x][j - 1] - a[x][j]) > threshold) {
                        ok = false;
                        break; // 不合法，下一个
                    }
                }
                if (!ok) continue;

                // 检查上下相邻格子
                for (int y = j - 2; y <= j; y++) {
                    if (abs(a[i - 2][y] - a[i - 1][y]) > threshold || abs(a[i - 1][y] - a[i][y]) > threshold) {
                        ok = false;
                        break; // 不合法，下一个
                    }
                }
                if (!ok) continue;

                // 合法，计算 3x3 子网格的平均值
                int avg = 0;
                for (int x = i - 2; x <= i; x++) {
                    for (int y = j - 2; y <= j; y++) {
                        avg += a[x][y];
                    }
                }
                avg /= 9;

                // 更新 3x3 子网格内的 result
                for (int x = i - 2; x <= i; x++) {
                    for (int y = j - 2; y <= j; y++) {
                        result[x][y] += avg; // 先累加，最后再求平均值
                        cnt[x][y]++;
                    }
                }
            }
        }

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (cnt[i][j] == 0) { // (i,j) 不属于任何子网格
                    result[i][j] = a[i][j];
                } else {
                    result[i][j] /= cnt[i][j]; // 求平均值
                }
            }
        }
        return result;
    }
};
```

```go [sol-Go]
func resultGrid(a [][]int, threshold int) [][]int {
	m, n := len(a), len(a[0])
	result := make([][]int, m)
	cnt := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		cnt[i] = make([]int, n)
	}
	for i := 2; i < m; i++ {
	next:
		for j := 2; j < n; j++ {
			// 检查左右相邻格子
			for _, row := range a[i-2 : i+1] {
				if abs(row[j-2]-row[j-1]) > threshold || abs(row[j-1]-row[j]) > threshold {
					continue next // 不合法，下一个
				}
			}

			// 检查上下相邻格子
			for y := j - 2; y <= j; y++ {
				if abs(a[i-2][y]-a[i-1][y]) > threshold || abs(a[i-1][y]-a[i][y]) > threshold {
					continue next // 不合法，下一个
				}
			}

			// 合法，计算 3x3 子网格的平均值
			avg := 0
			for x := i - 2; x <= i; x++ {
				for y := j - 2; y <= j; y++ {
					avg += a[x][y]
				}
			}
			avg /= 9

			// 更新 3x3 子网格内的 result
			for x := i - 2; x <= i; x++ {
				for y := j - 2; y <= j; y++ {
					result[x][y] += avg // 先累加，最后再求平均值
					cnt[x][y]++
				}
			}
		}
	}

	for i, row := range cnt {
		for j, c := range row {
			if c == 0 { // (i,j) 不属于任何子网格
				result[i][j] = a[i][j]
			} else {
				result[i][j] /= c // 求平均值
			}
		}
	}
	return result
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(Amn)$，其中 $m$ 和 $n$ 分别为 $\textit{a}$ 的行数和列数，$A=9$ 表示子网格大小。
- 空间复杂度：$\mathcal{O}(mn)$。

## 思考题

如果额外输入两个数 $w$ 和 $h$，把题目中 $3\times 3$ 改成 $w\times h$，要怎么做？你能做到 $\mathcal{O}(mn)$ 的时间复杂度吗？

欢迎在评论区分享你的思路/代码。

相关题目：[2132. 用邮票贴满网格图](https://leetcode.cn/problems/stamping-the-grid/)，[我的题解](https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/)，包含了解决思考题需要掌握的算法。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

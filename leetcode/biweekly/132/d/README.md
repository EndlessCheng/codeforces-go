遍历到 $x=\textit{nums}[i]$ 时，我们需要维护以 $x$ 结尾的、至多包含 $j$ 个不同相邻元素的子序列的最大长度，定义为 $f[x][j]$，初始全为 $0$。

对于 $x$，有三种决策：

1. 不选：$f[x][j]$ 不变。
2. 选，且和子序列的前一个数一样，或者作为子序列的第一个数：$f[x][j]$ 增加 $1$。
3. 选，且和子序列的前一个数不一样：设前一个数为 $y$，我们需要知道最大的 $f[y][j-1]$。

对于第三种决策，暴力枚举 $y$ 就太慢了（可以通过第三题但无法通过本题）。我们可以维护 $f[\cdot][j-1]$ 中的最大值 $\textit{mx}$、最大值对应的数字 $\textit{num}$，以及 $f[\textit{num}_2][j-1]$ 中的最大值 $\textit{mx}_2$，其中 $\textit{num}_2\ne \textit{num}$。

于是：

- 如果 $x\ne \textit{num}$，那么最大的 $f[y][j-1]$ 就是 $\textit{mx}$。
- 如果 $x = \textit{num}$，那么最大的 $f[y][j-1]$ 就是 $\textit{mx}_2$。

把最大的 $f[y][j-1]$ 记作 $m$，则 $f[x][j]$ 更新为

$$
\max(f[x][j] + 1, m + 1)
$$

对于不同的 $j$，我们需要维护对应的 $\textit{mx},\textit{mx}_2,\textit{num}$。用一个长为 $k+1$ 的数组 $\textit{records}$ 记录。

由于在计算 $f[x][j]$ 时会用到 $\textit{records}[j-1]$，然后更新 $\textit{records}[j]$，可以倒序枚举 $j$，以避免使用覆盖后的数据。

## 优化前

```py [sol-Python3]
class Solution:
    def maximumLength(self, nums: List[int], k: int) -> int:
        fs = {}
        records = [[0] * 3 for _ in range(k + 1)]
        for x in nums:
            if x not in fs:
                fs[x] = [0] * (k + 1)
            f = fs[x]
            for j in range(k, -1, -1):
                f[j] += 1
                if j > 0:
                    mx, mx2, num = records[j - 1]
                    f[j] = max(f[j], (mx if x != num else mx2) + 1)

                # records[j] 维护 fs[.][j] 的 mx, mx2, num
                v = f[j]
                p = records[j]
                if v > p[0]:
                    if x != p[2]:
                        p[2] = x
                        p[1] = p[0]
                    p[0] = v
                elif x != p[2] and v > p[1]:
                    p[1] = v
        return records[k][0]
```

```java [sol-Java]
class Solution {
    public int maximumLength(int[] nums, int k) {
        Map<Integer, int[]> fs = new HashMap<>();
        int[][] records = new int[k + 1][3];
        for (int x : nums) {
            int[] f = fs.computeIfAbsent(x, i -> new int[k + 1]);
            for (int j = k; j >= 0; j--) {
                f[j]++;
                if (j > 0) {
                    int mx = records[j - 1][0], mx2 = records[j - 1][1], num = records[j - 1][2];
                    f[j] = Math.max(f[j], (x != num ? mx : mx2) + 1);
                }

                // records[j] 维护 fs[.][j] 的 mx, mx2, num
                int v = f[j];
                int[] p = records[j];
                if (v > p[0]) {
                    if (x != p[2]) {
                        p[2] = x;
                        p[1] = p[0];
                    }
                    p[0] = v;
                } else if (x != p[2] && v > p[1]) {
                    p[1] = v;
                }
            }
        }
        return records[k][0];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumLength(vector<int>& nums, int k) {
        unordered_map<int, vector<int>> fs;
        vector<array<int, 3>> records(k + 1);
        for (int x : nums) {
            if (!fs.contains(x)) {
                fs[x] = vector<int>(k + 1);
            }
            auto& f = fs[x];
            for (int j = k; j >= 0; j--) {
                f[j]++;
                if (j) {
                    auto& r = records[j - 1];
                    int mx = r[0], mx2 = r[1], num = r[2];
                    f[j] = max(f[j], (x != num ? mx : mx2) + 1);
                }

                // records[j] 维护 fs[.][j] 的 mx, mx2, num
                int v = f[j];
                auto& p = records[j];
                if (v > p[0]) {
                    if (x != p[2]) {
                        p[2] = x;
                        p[1] = p[0];
                    }
                    p[0] = v;
                } else if (x != p[2] && v > p[1]) {
                    p[1] = v;
                }
            }
        }
        return records[k][0];
    }
};
```

```go [sol-Go]
func maximumLength(nums []int, k int) int {
	fs := map[int][]int{}
	records := make([]struct{ mx, mx2, num int }, k+1)
	for _, x := range nums {
		if fs[x] == nil {
			fs[x] = make([]int, k+1)
		}
		f := fs[x]
		for j := k; j >= 0; j-- {
			f[j]++
			if j > 0 {
				p := records[j-1]
				m := p.mx
				if x == p.num {
					m = p.mx2
				}
				f[j] = max(f[j], m+1)
			}

			// records[j] 维护 fs[.][j] 的 mx,mx2,num
			v := f[j]
			p := &records[j]
			if v > p.mx {
				if x != p.num {
					p.num = x
					p.mx2 = p.mx
				}
				p.mx = v
			} else if x != p.num && v > p.mx2 {
				p.mx2 = v
			}
		}
	}
	return records[k].mx
}
```

## 优化

其实只需要维护 $\textit{mx}$，因为：

- 如果 $x\ne \textit{num}$，那么最大的 $f[y][j-1]$ 就是 $\textit{mx}$。
- 如果 $x = \textit{num}$，相当于把 $x$ 加到以 $x$ 结尾的子序列的末尾，也就是用 $f[x][j-1] + 1$ 更新 $f[x][j]$ 的最大值。注意这个转移方程是不符合状态定义的，但由于 $j$ 越大，能选的数越多，所以 $f[x][j]\ge f[x][j-1]$，这样更新其实不影响结果，因为第二种决策会用 $f[x][j] + 1$ 更新 $f[x][j]$，这不会低于 $f[x][j-1] + 1$。

所以直接用 $\textit{mx}+1$ 更新 $f[x][j]$ 的最大值即可。

此外，为了避免判断 $i=0$ 的情况，可以往 $\textit{mx}$ 数组的最左边插入一个 $0$，把 $\textit{mx}$ 的下标加一。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Tx4y1b7wk/) 第四题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def maximumLength(self, nums: List[int], k: int) -> int:
        fs = {}
        mx = [0] * (k + 2)
        for x in nums:
            if x not in fs:
                fs[x] = [0] * (k + 1)
            f = fs[x]
            for j in range(k, -1, -1):
                f[j] = max(f[j], mx[j]) + 1
                mx[j + 1] = max(mx[j + 1], f[j])
        return mx[-1]
```

```java [sol-Java]
class Solution {
    public int maximumLength(int[] nums, int k) {
        Map<Integer, int[]> fs = new HashMap<>();
        int[] mx = new int[k + 2];
        for (int x : nums) {
            int[] f = fs.computeIfAbsent(x, i -> new int[k + 1]);
            for (int j = k; j >= 0; j--) {
                f[j] = Math.max(f[j], mx[j]) + 1;
                mx[j + 1] = Math.max(mx[j + 1], f[j]);
            }
        }
        return mx[k + 1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumLength(vector<int>& nums, int k) {
        unordered_map<int, vector<int>> fs;
        vector<int> mx(k + 2);
        for (int x : nums) {
            if (!fs.contains(x)) {
                fs[x] = vector<int>(k + 1);
            }
            auto& f = fs[x];
            for (int j = k; j >= 0; j--) {
                f[j] = max(f[j], mx[j]) + 1;
                mx[j + 1] = max(mx[j + 1], f[j]);
            }
        }
        return mx[k + 1];
    }
};
```

```go [sol-Go]
func maximumLength(nums []int, k int) int {
	fs := map[int][]int{}
	mx := make([]int, k+2)
	for _, x := range nums {
		if fs[x] == nil {
			fs[x] = make([]int, k+1)
		}
		f := fs[x]
		for j := k; j >= 0; j-- {
			f[j] = max(f[j], mx[j]) + 1
			mx[j+1] = max(mx[j+1], f[j])
		}
	}
	return mx[k+1]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(nk)$。

#### 相似题目

- [CF264C. Choosing Balls](https://codeforces.com/problemset/problem/264/C)

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

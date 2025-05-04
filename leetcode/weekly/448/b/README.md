初始化 $2^n\times 2^n$ 的答案矩阵 $a$，初始化 $\textit{val}=0$。

定义 $\textit{dfs}(u,d,l,r)$ 表示填充行号在 $[u,d)$，列号在 $[l,r)$ 中的子矩阵。

设 $m = \dfrac{d-u}{2}$。按照从小到大的顺序，依次递归填充四个象限（四等分矩阵）：

- 填充右上角象限 $\textit{dfs}(u,u+m,l+m,r)$。
- 填充右下角象限 $\textit{dfs}(u+m,d,l+m,r)$。
- 填充左下角象限 $\textit{dfs}(u+m,d,l,l+m)$。
- 填充左上角象限 $\textit{dfs}(u,u+m,l,l+m)$。

递归边界：如果 $d-u=1$，只有一个格子，填充 $a[u][l]=\textit{val}$，然后把 $\textit{val}$ 加一。

递归入口：$\textit{dfs}(0,2^n,0,2^n)$。

**注**：由于矩阵的长宽都是 $2$ 的幂，所以每一步四等分，矩阵的长宽都可以被 $2$ 整除，不会出现无法四等分的情况。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1avVwz5EbY/?t=3m13s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def specialGrid(self, n: int) -> List[List[int]]:
        a = [[0] * (1 << n) for _ in range(1 << n)]
        val = 0

        def dfs(u: int, d: int, l: int, r: int) -> None:
            if d - u == 1:
                nonlocal val
                a[u][l] = val
                val += 1
                return
            m = (d - u) // 2
            dfs(u, u + m, l + m, r)
            dfs(u + m, d, l + m, r)
            dfs(u + m, d, l, l + m)
            dfs(u, u + m, l, l + m)

        dfs(0, 1 << n, 0, 1 << n)
        return a
```

```java [sol-Java]
class Solution {
    public int[][] specialGrid(int n) {
        int[][] a = new int[1 << n][1 << n];
        dfs(a, 0, 1 << n, 0, 1 << n);
        return a;
    }

    private int val = 0;

    private void dfs(int[][] a, int u, int d, int l, int r) {
        if (d - u == 1) {
            a[u][l] = val++;
            return;
        }
        int m = (d - u) / 2;
        dfs(a, u, u + m, l + m, r);
        dfs(a, u + m, d, l + m, r);
        dfs(a, u + m, d, l, l + m);
        dfs(a, u, u + m, l, l + m);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> specialGrid(int n) {
        vector a(1 << n, vector<int>(1 << n));
        int val = 0;
        auto dfs = [&](this auto&& dfs, int u, int d, int l, int r) -> void {
            if (d - u == 1) {
                a[u][l] = val++;
                return;
            }
            int m = (d - u) / 2;
            dfs(u, u + m, l + m, r);
            dfs(u + m, d, l + m, r);
            dfs(u + m, d, l, l + m);
            dfs(u, u + m, l, l + m);
        };
        dfs(0, 1 << n, 0, 1 << n);
        return a;
    }
};
```

```go [sol-Go]
func specialGrid(n int) [][]int {
	val := 0
	var dfs func([][]int, int, int)
	dfs = func(a [][]int, l, r int) {
		if len(a) == 1 {
			a[0][l] = val
			val++
			return
		}
		m := len(a) / 2
		dfs(a[:m], l+m, r)
		dfs(a[m:], l+m, r)
		dfs(a[m:], l, l+m)
		dfs(a[:m], l, l+m)
	}

	a := make([][]int, 1<<n)
	for i := range a {
		a[i] = make([]int, 1<<n)
	}
	dfs(a, 0, 1<<n)
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(4^n)$。每个格子恰好访问一次。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。递归需要 $\mathcal{O}(n)$ 的栈空间。

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

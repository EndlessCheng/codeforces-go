设矩形长宽分别为 $x$ 和 $y$。

根据勾股定理，矩形对角线长度的平方为

$$
x^2+y^2
$$

根据题意，我们需要用双关键字比大小，找到最大的面积。

- 第一关键字是对角线长度，直接用其平方值。如果遍历到更大的长度，则覆盖矩形面积。
- 第二关键字是矩形面积，即 $x\cdot y$。如果遇到了和最长长度一样长的矩形，那么更新面积的最大值。

```py [sol-Python3]
class Solution:
    def areaOfMaxDiagonal(self, dimensions: List[List[int]]) -> int:
        return max((x * x + y * y, x * y) for x, y in dimensions)[1]
```

```java [sol-Java]
class Solution {
    public int areaOfMaxDiagonal(int[][] dimensions) {
        int ans = 0, maxL = 0;
        for (int[] d : dimensions) {
            int x = d[0], y = d[1];
            int l = x * x + y * y;
            if (l > maxL || l == maxL && x * y > ans) {
                maxL = l;
                ans = x * y;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int areaOfMaxDiagonal(vector<vector<int>>& dimensions) {
        pair<int, int> mx{};
        for (auto& d : dimensions) {
            int x = d[0], y = d[1];
            mx = max(mx, pair(x * x + y * y, x * y));
        }
        return mx.second;
    }
};
```

```c [sol-C]
int areaOfMaxDiagonal(int** dimensions, int dimensionsSize, int* dimensionsColSize) {
    int ans = 0, max_l = 0;
    for (int i = 0; i < dimensionsSize; i++) {
        int x = dimensions[i][0], y = dimensions[i][1];
        int l = x * x + y * y;
        if (l > max_l || l == max_l && x * y > ans) {
            max_l = l;
            ans = x * y;
        }
    }
    return ans;
}
```

```go [sol-Go]
func areaOfMaxDiagonal(dimensions [][]int) (ans int) {
	maxL := 0
	for _, d := range dimensions {
		x, y := d[0], d[1]
		l := x*x + y*y
		if l > maxL || l == maxL && x*y > ans {
			maxL = l
			ans = x * y
		}
	}
	return
}
```

```js [sol-JavaScript]
var areaOfMaxDiagonal = function(dimensions) {
    let ans = 0, maxL = 0;
    for (const [x, y] of dimensions) {
        const l = x * x + y * y;
        if (l > maxL || l === maxL && x * y > ans) {
            maxL = l;
            ans = x * y;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn area_of_max_diagonal(dimensions: Vec<Vec<i32>>) -> i32 {
        let mut mx = (0, 0);
        for d in dimensions {
            let x = d[0];
            let y = d[1];
            mx = mx.max((x * x + y * y, x * y));
        }
        mx.1
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{dimensions}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

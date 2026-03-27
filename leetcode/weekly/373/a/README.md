实际上，只需把每行右移 $k$ 次。无需判断奇偶，无需考虑左移。

为什么？如果一行左移 $k$ 次等于自己，那么这个过程的逆过程，就是把自己右移 $k$ 次，得到自己。

判断 $\textit{row}$ 右移 $k$ 次是否等于 $\textit{row}$，可以比较 $\textit{row}[j]$ 与右移 $k$ 次后的位置 $\textit{row}[(j+k)\bmod n]$ 是否相等。

```py [sol-Python3]
class Solution:
    def areSimilar(self, mat: List[List[int]], k: int) -> bool:
        k %= len(mat[0])  # 右移 n 次等价于右移 0 次，右移 n+1 次等价于右移 1 次，依此类推，先模个 n
        return k == 0 or all(row == row[k:] + row[:k] for row in mat)
```

```py [sol-Python3 写法二]
class Solution:
    def areSimilar(self, mat: List[List[int]], k: int) -> bool:
        n = len(mat[0])
        for row in mat:
            for j in range(n):
                if row[j] != row[(j + k) % n]:
                    return False
        return True
```

```java [sol-Java]
class Solution {
    public boolean areSimilar(int[][] mat, int k) {
        int n = mat[0].length;
        for (int[] row : mat) {
            for (int j = 0; j < n; j++) {
                if (row[j] != row[(j + k) % n]) {
                    return false;
                }
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool areSimilar(vector<vector<int>>& mat, int k) {
        int n = mat[0].size();
        for (auto& row : mat) {
            for (int j = 0; j < n; j++) {
                if (row[j] != row[(j + k) % n]) {
                    return false;
                }
            }
        }
        return true;
    }
};
```

```c [sol-C]
bool areSimilar(int** mat, int matSize, int* matColSize, int k) {
    int n = matColSize[0];
    for (int i = 0; i < matSize; i++) {
        int* row = mat[i];
        for (int j = 0; j < n; j++) {
            if (row[j] != row[(j + k) % n]) {
                return false;
            }
        }
    }
    return true;
}
```

```go [sol-Go]
func areSimilar(mat [][]int, k int) bool {
	n := len(mat[0])
	for _, row := range mat {
		for j, x := range row {
			if x != row[(j+k)%n] {
				return false
			}
		}
	}
	return true
}
```

```go [sol-Go 写法二]
func areSimilar(mat [][]int, k int) bool {
	k %= len(mat[0]) // 右移 n 次等价于右移 0 次，右移 n+1 次等价于右移 1 次，依此类推，先模个 n
	for _, row := range mat {
		if !slices.Equal(row, append(row[k:], row[:k]...)) {
			return false
		}
	}
	return true
}
```

```js [sol-JS]
var areSimilar = function(mat, k) {
    const n = mat[0].length;
    for (const row of mat) {
        for (let j = 0; j < n; j++) {
            if (row[j] !== row[(j + k) % n]) {
                return false;
            }
        }
    }
    return true;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn are_similar(mat: Vec<Vec<i32>>, k: i32) -> bool {
        let n = mat[0].len();
        let k = k as usize;
        for row in mat {
            for j in 0..n {
                if row[j] != row[(j + k) % n] {
                    return false;
                }
            }
        }
        true
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别是 $\textit{mat}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。

## 详细解释

可能有同学脑筋没转过来，这里详细解释下。

如果给你两个数组 $a$ 和 $b$，要判断 $a$ 左移/右移后是否等于 $b$，那么 $a$ 左移 $k$ 次和右移 $k$ 次是不一样的。

但本题这两个数组都是 $a$，要判断 $a$ 左移/右移后是否等于 $a$ 自己。

由于 $a$ 左移 $k$ 次后和 $b$ 比较，等价于 $b$ 右移 $k$ 次后和 $a$ 比较。在 $b$ 就是 $a$ 的情况下，等价于 $a$ 自己右移 $k$ 次和 $a$ 比较。

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

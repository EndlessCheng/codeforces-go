## 方法一：统计元素出现次数

用数组统计每个数的出现次数，然后遍历 $[1,n^2]$，寻找出现两次的数，和出现零次的数。

```py [sol-Python3]
class Solution:
    def findMissingAndRepeatedValues(self, grid: List[List[int]]) -> List[int]:
        n = len(grid)
        cnt = [0] * (n * n + 1)
        for row in grid:
            for x in row:
                cnt[x] += 1

        ans = [0, 0]
        for i in range(1, n * n + 1):
            if cnt[i] == 2:
                ans[0] = i  # 出现两次的数
            elif cnt[i] == 0:
                ans[1] = i  # 出现零次的数
        return ans
```

```java [sol-Java]
public class Solution {
    public int[] findMissingAndRepeatedValues(int[][] grid) {
        int n = grid.length;
        int[] cnt = new int[n * n + 1];
        for (int[] row : grid) {
            for (int x : row) {
                cnt[x]++;
            }
        }

        int[] ans = new int[2];
        for (int i = 1; i <= n * n; i++) {
            if (cnt[i] == 2) {
                ans[0] = i; // 出现两次的数
            } else if (cnt[i] == 0) {
                ans[1] = i; // 出现零次的数
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findMissingAndRepeatedValues(vector<vector<int>>& grid) {
        int n = grid.size();
        vector<int> cnt(n * n + 1);
        for (auto& row : grid) {
            for (int x : row) {
                cnt[x]++;
            }
        }

        vector<int> ans(2);
        for (int i = 1; i <= n * n; i++) {
            if (cnt[i] == 2) {
                ans[0] = i; // 出现两次的数
            } else if (cnt[i] == 0) {
                ans[1] = i; // 出现零次的数
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	cnt := make([]int, n*n+1)
	for _, row := range grid {
		for _, x := range row {
			cnt[x]++
		}
	}

	ans := make([]int, 2)
	for i := 1; i <= n*n; i++ {
		if cnt[i] == 2 {
			ans[0] = i // 出现两次的数
		} else if cnt[i] == 0 {
			ans[1] = i // 出现零次的数
		}
	}
	return ans
}
```

```js [sol-JavaScript]
var findMissingAndRepeatedValues = function(grid) {
    const n = grid.length;
    const cnt = Array(n * n + 1).fill(0);
    for (const row of grid) {
        for (const x of row) {
            cnt[x]++;
        }
    }

    const ans = Array(2);
    for (let i = 1; i <= n * n; i++) {
        if (cnt[i] === 2) {
            ans[0] = i; // 出现两次的数
        } else if (cnt[i] === 0) {
            ans[1] = i; // 出现零次的数
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_missing_and_repeated_values(grid: Vec<Vec<i32>>) -> Vec<i32> {
        let n = grid.len();
        let mut cnt = vec![0; n * n + 1];
        for row in grid {
            for x in row {
                cnt[x as usize] += 1;
            }
        }

        let mut ans = vec![0; 2];
        for i in 1..=n * n {
            if cnt[i] == 2 {
                ans[0] = i as i32; // 出现两次的数
            } else if cnt[i] == 0 {
                ans[1] = i as i32; // 出现零次的数
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 方法二：位运算

能否做到 $\mathcal{O}(1)$ 额外空间？

额外添加 $1,2,3,\cdots,n^2$，那么就有一个数出现一次，一个数出现三次，其余数均出现两次。

在异或操作下，一个数出现三次和出现一次是一样的，于是问题变成 [260. 只出现一次的数字 III](https://leetcode.cn/problems/single-number-iii/)，具体请看我的[【图解】](https://leetcode.cn/problems/single-number-iii/solution/tu-jie-yi-zhang-tu-miao-dong-zhuan-huan-np9d2/)。

以下代码基于我 260 题解的第二份代码修改：

1. 计算 $\textit{xorAll}$ 时，额外计算 $1$ 到 $n^2$ 的异或和。有 $\mathcal{O}(1)$ 公式，见 [1486. 数组异或操作的题解](https://leetcode.cn/problems/xor-operation-in-an-array/solution/o1-gong-shi-tui-dao-pythonjavaccgojsrust-le23/)。注意，$(4k)^2=16k^2$ 是 $4$ 的倍数，$(4k+2)^2=16k^2+16k+4$ 也是 $4$ 的倍数，$(4k+1)^2=16k^2+8k+1$ 模 $4$ 等于 $1$，$(4k+3)^2=16k^2+24k+9$ 模 $4$ 也等于 $1$，所以 $n^2$ 模 $4$ 要么是 $0$ 要么是 $1$，取决于 $n$ 是偶数还是奇数。如果 $n$ 是偶数，那么 $1$ 到 $n^2$ 的异或和等于 $n^2$，否则等于 $1$。
2. 计算 $\textit{ans}$ 时，额外遍历 $1$ 到 $n^2$ 的每个数。

```py [sol-Python3]
class Solution:
    def findMissingAndRepeatedValues(self, grid: List[List[int]]) -> List[int]:
        n = len(grid)
        xor_all = reduce(xor, (x for row in grid for x in row)) ^ (1 if n % 2 else n * n)
        shift = xor_all.bit_length() - 1

        ans = [0, 0]
        for x in range(1, n * n + 1):
            ans[x >> shift & 1] ^= x
        for row in grid:
            for x in row:
                ans[x >> shift & 1] ^= x

        return ans if ans[0] in (x for row in grid for x in row) else ans[::-1]
```

```java [sol-Java]
class Solution {
    public int[] findMissingAndRepeatedValues(int[][] grid) {
        int n = grid.length;
        int xorAll = 0;
        for (int[] row : grid) {
            for (int x : row) {
                xorAll ^= x;
            }
        }
        xorAll ^= n % 2 > 0 ? 1 : n * n;
        int shift = Integer.numberOfTrailingZeros(xorAll);

        int[] ans = new int[2];
        for (int x = 1; x <= n * n; x++) {
            ans[x >> shift & 1] ^= x;
        }
        for (int[] row : grid) {
            for (int x : row) {
                ans[x >> shift & 1] ^= x;
            }
        }

        for (int[] row : grid) {
            for (int x : row) {
                if (x == ans[0]) {
                    return ans;
                }
            }
        }
        return new int[]{ans[1], ans[0]};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findMissingAndRepeatedValues(vector<vector<int>>& grid) {
        int n = grid.size();
        int xor_all = 0;
        for (auto& row : grid) {
            for (int x : row) {
                xor_all ^= x;
            }
        }
        xor_all ^= n % 2 ? 1 : n * n;
        int shift = __builtin_ctz(xor_all);

        vector<int> ans(2);
        for (int x = 1; x <= n * n; x++) {
            ans[x >> shift & 1] ^= x;
        }
        for (auto& row : grid) {
            for (int x : row) {
                ans[x >> shift & 1] ^= x;
            }
        }

        for (auto& row : grid) {
            if (ranges::find(row, ans[0]) != row.end()) {
                return ans;
            }
        }
        return {ans[1], ans[0]};
    }
};
```

```go [sol-Go]
func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    xorAll := 0
    for _, row := range grid {
        for _, x := range row {
            xorAll ^= x
        }
    }
    if n%2 > 0 {
        xorAll ^= 1
    } else {
        xorAll ^= n * n
    }
    shift := bits.TrailingZeros(uint(xorAll))

    ans := make([]int, 2)
    for x := 1; x <= n*n; x++ {
        ans[x>>shift&1] ^= x
    }
    for _, row := range grid {
        for _, x := range row {
            ans[x>>shift&1] ^= x
        }
    }

    for _, row := range grid {
        if slices.Contains(row, ans[0]) {
            return ans
        }
    }
    return []int{ans[1], ans[0]}
}
```

```js [sol-JavaScript]
var findMissingAndRepeatedValues = function(grid) {
    const n = grid.length;
    let xorAll = 0;
    for (const row of grid) {
        for (const x of row) {
            xorAll ^= x;
        }
    }
    xorAll ^= n % 2 ? 1 : n * n;
    const shift = 31 - Math.clz32(xorAll);

    const ans = [0, 0];
    for (let x = 1; x <= n * n; x++) {
        ans[x >> shift & 1] ^= x;
    }
    for (const row of grid) {
        for (const x of row) {
            ans[x >> shift & 1] ^= x;
        }
    }

    for (const row of grid) {
        if (row.includes(ans[0])) {
            return ans;
        }
    }
    return [ans[1], ans[0]];
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_missing_and_repeated_values(grid: Vec<Vec<i32>>) -> Vec<i32> {
        let n = grid.len() as i32;
        let mut xor_all = 0;
        for row in &grid {
            for &x in row {
                xor_all ^= x;
            }
        }
        xor_all ^= if n % 2 > 0 { 1 } else { n * n };
        let shift = xor_all.trailing_zeros();

        let mut ans = vec![0, 0];
        for x in 1..=n * n {
            ans[(x >> shift & 1) as usize] ^= x;
        }
        for row in &grid {
            for &x in row {
                ans[(x >> shift & 1) as usize] ^= x;
            }
        }

        for row in grid {
            if row.contains(&ans[0]) {
                return ans;
            }
        }
        vec![ans[1], ans[0]]
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。

## 思考题

代码在计算 `ans[x >> shift & 1] ^= x` 时，遍历了 $1$ 到 $n^2$ 的每个数，你能用 $\mathcal{O}(1)$ 的公式解决吗？

欢迎在评论区发表你的思路/代码。

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

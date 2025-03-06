## 思路

遍历两条对角线上的元素，如果是质数则更新答案的最大值。

注意 $1$ 不是质数。

## 如何判断质数

对于整数 $n$，暴力做法是枚举 $[2,n-1]$ 中的整数 $i$，判断 $n\bmod i=0$ 是否成立，如果成立则说明 $i$ 是 $n$ 的一个因子，$n$ 不是质数。

但其实只需枚举 $[2,\left\lfloor\sqrt{n}\right\rfloor]$ 中的整数 $i$。

**证明**：反证法。从小到大枚举，假设我们枚举到了一个超过 $\left\lfloor\sqrt{n}\right\rfloor$ 的整数 $j$，且 $j$ 是 $n$ 的因子，那么 $\dfrac{n}{j}$ 也是 $n$ 的因子。由于 $\dfrac{n}{j} < j$，我们会先枚举到 $i=\dfrac{n}{j}$，再枚举到 $i=j$。但由于 $\dfrac{n}{j}$ 是 $n$ 的因子，我们会在 $i=\dfrac{n}{j}$ 时返回，不可能继续枚举到 $i=j$，矛盾，所以原命题成立。

## 细节

如果元素 $x$ 没有超过答案 $\textit{ans}$，那么无需判断 $x$ 是否为质数，因为它不会让答案变得更大。

```py [sol-Python3]
class Solution:
    def is_prime(self, n: int) -> bool:
        for i in range(2, isqrt(n) + 1):
            if n % i == 0:
                return False
        return n >= 2  # 1 不是质数

    def diagonalPrime(self, nums: List[List[int]]) -> int:
        ans = 0
        for i, row in enumerate(nums):
            for x in row[i], row[-1 - i]:
                if x > ans and self.is_prime(x):
                    ans = x
        return ans
```

```java [sol-Java]
class Solution {
    public int diagonalPrime(int[][] nums) {
        int n = nums.length;
        int ans = 0;
        for (int i = 0; i < n; i++) {
            int x = nums[i][i];
            if (x > ans && isPrime(x)) {
                ans = x;
            }
            x = nums[i][n - 1 - i];
            if (x > ans && isPrime(x)) {
                ans = x;
            }
        }
        return ans;
    }

    private boolean isPrime(int n) {
        for (int i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return n >= 2; // 1 不是质数
    }
}
```

```cpp [sol-C++]
class Solution {
    bool is_prime(int n) {
        for (int i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return n >= 2; // 1 不是质数
    }

public:
    int diagonalPrime(vector<vector<int>>& nums) {
        int n = nums.size(), ans = 0;
        for (int i = 0; i < n; i++) {
            for (int x : {nums[i][i], nums[i][n - 1 - i]}) {
                if (x > ans && is_prime(x)) {
                    ans = x;
                }
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int isPrime(int n) {
    for (int i = 2; i * i <= n; i++) {
        if (n % i == 0) {
            return 0;
        }
    }
    return n >= 2; // 1 不是质数
}

int diagonalPrime(int** nums, int numsSize, int* numsColSize) {
    int ans = 0;
    for (int i = 0; i < numsSize; i++) {
        int x = nums[i][i];
        if (x > ans && isPrime(x)) {
            ans = x;
        }
        x = nums[i][numsSize - 1 - i];
        if (x > ans && isPrime(x)) {
            ans = x;
        }
    }
    return ans;
}
```

```go [sol-Go]
func isPrime(n int) bool {
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return n >= 2 // 1 不是质数
}

func diagonalPrime(nums [][]int) (ans int) {
    for i, row := range nums {
        if x := row[i]; x > ans && isPrime(x) {
            ans = x
        }
        if x := row[len(nums)-1-i]; x > ans && isPrime(x) {
            ans = x
        }
    }
    return
}
```

```js [sol-JavaScript]
var isPrime = function(n) {
    for (let i = 2; i * i <= n; i++) {
        if (n % i === 0) {
            return false;
        }
    }
    return n >= 2; // 1 不是质数
};

var diagonalPrime = function(nums) {
    const n = nums.length;
    let ans = 0;
    for (let i = 0; i < n; i++) {
        for (const x of [nums[i][i], nums[i][n - 1 - i]]) {
            if (x > ans && isPrime(x)) {
                ans = x;
            }
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    fn is_prime(n: i32) -> bool {
        for i in 2..=((n as f64).sqrt() as i32) {
            if n % i == 0 {
                return false;
            }
        }
        n >= 2 // 1 不是质数
    }

    pub fn diagonal_prime(nums: Vec<Vec<i32>>) -> i32 {
        let n = nums.len();
        let mut ans = 0;
        for (i, row) in nums.iter().enumerate() {
            for x in [row[i], row[n - 1 - i]] {
                if x > ans && Self::is_prime(x) {
                    ans = x;
                }
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt{U})$，其中 $n$ 为 $\textit{nums}$ 的长度，$U$ 为两条对角线上的最大值。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

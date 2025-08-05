把 $n$ 看成背包容量，$1^x,2^x,3^x,\dots$ 看成物品，本题是恰好装满型 **0-1 背包**的方案数，做法同 [494. 目标和](https://leetcode.cn/problems/target-sum/)，原理讲解见[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)，包含倒序循环的讲解。

代码实现时，由于最坏情况 $n=300$，$x=1$ 的答案没有超过 $64$ 位整数的最大值，可以在计算结束后再取模。

```py [sol-Python3]
class Solution:
    def numberOfWays(self, n: int, x: int) -> int:
        f = [1] + [0] * n
        for i in range(1, n + 1):
            v = i ** x
            if v > n:
                break
            for s in range(n, v - 1, -1):
                f[s] += f[s - v]
        return f[n] % 1_000_000_007
```

```py [sol-Python3 预处理 0ms]
# 预处理所有答案
MX_N = 300
MX_X = 5
f = [[1] + [0] * MX_N for _ in range(MX_X + 1)]
for x in range(1, MX_X + 1):
    for i in count(1):
        v = i ** x
        if v > MX_N:
            break
        for s in range(MX_N, v - 1, -1):
            f[x][s] += f[x][s - v]

class Solution:
    def numberOfWays(self, n: int, x: int) -> int:
        return f[x][n] % 1_000_000_007
```

```java [sol-Java]
class Solution {
    int numberOfWays(int n, int x) {
        long[] f = new long[n + 1];
        f[0] = 1;
        // 本题数据范围小，Math.pow 的计算结果一定准确
        for (int i = 1; Math.pow(i, x) <= n; i++) {
            int v = (int) Math.pow(i, x);
            for (int s = n; s >= v; s--) {
                f[s] += f[s - v];
            }
        }
        return (int) (f[n] % 1_000_000_007);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int numberOfWays(int n, int x) {
        vector<long long> f(n + 1);
        f[0] = 1;
        // 本题数据范围小，pow 的计算结果一定准确
        for (int i = 1; pow(i, x) <= n; i++) {
            int v = pow(i, x);
            for (int s = n; s >= v; s--) {
                f[s] += f[s - v];
            }
        }
        return f[n] % 1'000'000'007;
    }
};
```

```c [sol-C]
#define MOD 1000000007

int numberOfWays(int n, int x) {
    long long* f = calloc(n + 1, sizeof(long long));
    f[0] = 1;

    // 本题数据范围小，pow 的计算结果一定准确
    for (int i = 1; pow(i, x) <= n; i++) {
        int v = pow(i, x);
        for (int s = n; s >= v; s--) {
            f[s] += f[s - v];
        }
    }

    int ans = f[n] % MOD;
    free(f);
    return ans;
}
```

```go [sol-Go]
func numberOfWays(n, x int) int {
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; pow(i, x) <= n; i++ {
		v := pow(i, x)
		for s := n; s >= v; s-- {
			f[s] += f[s-v]
		}
	}
	return f[n] % 1_000_000_007
}

// 本题数据范围小，math.Pow 的计算结果一定准确
func pow(i, x int) int {
	return int(math.Pow(float64(i), float64(x)))
}
```

```js [sol-JS]
var numberOfWays = function(n, x) {
    const f = new Array(n + 1).fill(0);
    f[0] = 1;
    // 本题数据范围小，pow 的计算结果一定准确
    for (let i = 1; Math.pow(i, x) <= n; i++) {
        const v = Math.pow(i, x);
        for (let s = n; s >= v; s--) {
            f[s] += f[s - v];
        }
    }
    return f[n] % 1_000_000_007;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn number_of_ways(n: i32, x: i32) -> i32 {
        let n = n as usize;
        let mut f = vec![0i64; n + 1];
        f[0] = 1;
        for i in 1usize.. {
            let v = i.pow(x as u32);
            if v > n {
                break;
            }
            for s in (v..=n).rev() {
                f[s] += f[s - v];
            }
        }
        (f[n] % 1_000_000_007) as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt[x]n)$。外层循环的循环次数为 $\mathcal{O}(\sqrt[x]n)$。
- 空间复杂度：$\mathcal{O}(n)$。

**注**：也可以用快速幂计算 $\texttt{pow}$，原理见[【图解】一张图秒懂快速幂](https://leetcode.cn/problems/powx-n/solution/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/)。

## 专题训练

见下面动态规划题单的「**§3.1 0-1 背包**」。

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

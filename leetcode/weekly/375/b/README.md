由于 $b_i$ 和 $c_i$ 都比较小，可以暴力循环计算。

更快的做法是用**快速幂**，请看[【图解】一张图秒懂快速幂](https://leetcode.cn/problems/powx-n/solution/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/)。

本题还需要取模，如果你不知道如何正确地取模，请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

```py [sol-Python3]
class Solution:
    def getGoodIndices(self, variables: List[List[int]], target: int) -> List[int]:
        return [i for i, (a, b, c, m) in enumerate(variables)
                if pow(pow(a, b, 10), c, m) == target]
```

```java [sol-Java]
class Solution {
    public List<Integer> getGoodIndices(int[][] variables, int target) {
        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < variables.length; i++) {
            int[] v = variables[i];
            if (pow(pow(v[0], v[1], 10), v[2], v[3]) == target) {
                ans.add(i);
            }
        }
        return ans;
    }

    // 本题 mod 很小，即使平方也不会超过 int 范围，所以不需要用 long
    private int pow(int x, int n, int mod) {
        int res = 1;
        while (n > 0) {
            if (n % 2 > 0) {
                res = res * x % mod;
            }
            x = x * x % mod;
            n /= 2;
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
    // 本题 mod 很小，即使平方也不会超过 int 范围，所以不需要用 long long
    int pow(int x, int n, int mod) {
        int res = 1;
        while (n) {
            if (n & 1) {
                res = res * x % mod;
            }
            x = x * x % mod;
            n >>= 1;
        }
        return res;
    }

public:
    vector<int> getGoodIndices(vector<vector<int>>& variables, int target) {
        vector<int> ans;
        for (int i = 0; i < variables.size(); i++) {
            auto& v = variables[i];
            if (pow(pow(v[0], v[1], 10), v[2], v[3]) == target) {
                ans.push_back(i);
            }
        }
        return ans;
    }
};
```

```c [sol-C]
// 本题 mod 很小，即使平方也不会超过 int 范围，所以不需要用 long long
int qpow(int x, int n, int mod) {
    int res = 1;
    while (n) {
        if (n & 1) {
            res = res * x % mod;
        }
        x = x * x % mod;
        n >>= 1;
    }
    return res;
}

int* getGoodIndices(int** variables, int variablesSize, int* variablesColSize, int target, int* returnSize) {
    int* ans = malloc(variablesSize * sizeof(int));
    *returnSize = 0;
    for (int i = 0; i < variablesSize; i++) {
        int* v = variables[i];
        if (qpow(qpow(v[0], v[1], 10), v[2], v[3]) == target) {
            ans[(*returnSize)++] = i;
        }
    }
    return ans;
}
```

```go [sol-Go]
func getGoodIndices(variables [][]int, target int) (ans []int) {
	for i, v := range variables {
		if pow(pow(v[0], v[1], 10), v[2], v[3]) == target {
			ans = append(ans, i)
		}
	}
	return
}

func pow(x, n, mod int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

```js [sol-JavaScript]
var getGoodIndices = function(variables, target) {
    const ans = [];
    for (let i = 0; i < variables.length; i++) {
        const [a, b, c, m] = variables[i];
        if (pow(pow(a, b, 10), c, m) === target) {
            ans.push(i);
        }
    }
    return ans;
};

// 本题 mod 很小，即使平方也不会超过 MAX_SAFE_INTEGER 范围，所以不需要用 BigInt
function pow(x, n, mod) {
    let res = 1;
    while (n) {
        if (n % 2) {
            res = res * x % mod;
        }
        x = x * x % mod;
        n = Math.floor(n / 2);
    }
    return res;
}
```

```rust [sol-Rust]
impl Solution {
    pub fn get_good_indices(variables: Vec<Vec<i32>>, target: i32) -> Vec<i32> {
        let pow = |mut x, mut n, m| {
            // 本题 m 很小，即使平方也不会超过 i32 范围，所以不需要用 i64
            let mut res = 1;
            while n > 0 {
                if n % 2 > 0 {
                    res = res * x % m;
                }
                x = x * x % m;
                n /= 2;
            }
            res
        };
        let check = |v: &Vec<_>| pow(pow(v[0], v[1], 10), v[2], v[3]) == target;
        variables.iter()
            .enumerate()
            .filter_map(|(i, v)| check(v).then_some(i as i32))
            .collect()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{variables}$ 的长度，$U$ 为 $b_i$ 和 $c_i$ 的最大值，本题 $U=10^3$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

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

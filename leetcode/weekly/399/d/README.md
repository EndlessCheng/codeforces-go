本题是可以修改数组元素值的 [198. 打家劫舍](https://leetcode.cn/problems/house-robber/)。

为了解决本题，首先来换一个角度，用**分治**的思想解决打家劫舍。

设 $f(A)$ 为数组 $A$ 的打家劫舍的答案。把 $\textit{nums}$ 从中间切开，分成左右两个子数组，分别记作 $a$ 和 $b$。要计算 $f(\textit{nums})$，看上去，我们只需要分别计算 $f(a)$ 和 $f(b)$。但这是不对的，万一同时选了 $a$ 的最后一个数和 $b$ 的第一个数，就不满足题目要求了。

怎么办？要么不选 $a$ 的最后一个数，要么不选 $b$ 的第一个数。所以加个约束，先来（非正式地）讨论一下：

- 约束 $a$ 的最后一个数一定不选，即 $f(\textit{nums}) = f(a') + f(b)$，其中 $a'$ 是 $a$ 去掉最后一个数后的数组。
- 约束 $b$ 的第一个数一定不选，即 $f(\textit{nums}) = f(a) + f('b)$，其中 $'b$ 表示 $b$ 去掉第一个数后的数组。

两种情况取最大值，即

$$
f(\textit{nums}) = \max(f(a') + f(b), f(a) + f('b))
$$

继续计算下去，就需要进一步地把 $a',b,a,\ 'b$ 切开，继续分类讨论。

为方便描述，定义：

- $f_{00}(A)$ 表示在 $A$ 第一个数一定不选，最后一个数也一定不选的情况下，打家劫舍的答案。
- $f_{01}(A)$ 表示在 $A$ 第一个数一定不选，最后一个数可选可不选的情况下，打家劫舍的答案。
- $f_{10}(A)$ 表示在 $A$ 第一个数可选可不选，最后一个数一定不选的情况下，打家劫舍的答案。
- $f_{11}(A)$ 表示在 $A$ 第一个数可选可不选，最后一个数也可选可不选的情况下，打家劫舍的答案，这等于上面定义的 $f(A)$。

前文的分类讨论可以（正式地）表述为

$$
f_{11}(\textit{nums}) = \max(f_{10}(a) + f_{11}(b),\ f_{11}(a) + f_{01}(b))
$$

$f_{10}$ 和 $f_{01}$ 又该如何计算？以 $f_{10}(a)$ 为例，把 $a$ 分成左右两个数组 $p$ 和 $q$，分类讨论：

- 约束 $p$ 的最后一个数一定不选，即 $f_{10}(a) = f_{10}(p) + f_{10}(q)$，注意 $q$ 的最后一个数也不能选，因为我们计算的是 $f_{10}(a)$，$a$ 的最后一个数一定不能选。
- 约束 $q$ 的第一个数一定不选，即 $f_{10}(a) = f_{11}(p) + f_{00}(q)$。

两种情况取最大值，得

$$
f_{10}(a) = \max(f_{10}(p) + f_{10}(q),\ f_{11}(p) + f_{00}(q))
$$

同理可以得到 $f_{00}$ 和 $f_{01}$。

综上所述：

$$
\begin{aligned}
f_{00}(a) &= \max(f_{00}(p) + f_{10}(q),\ f_{01}(p) + f_{00}(q))    \\
f_{01}(a) &= \max(f_{00}(p) + f_{11}(q),\ f_{01}(p) + f_{01}(q))    \\
f_{10}(a) &= \max(f_{10}(p) + f_{10}(q),\ f_{11}(p) + f_{00}(q))    \\
f_{11}(a) &= \max(f_{10}(p) + f_{11}(q),\ f_{11}(p) + f_{01}(q))    \\
\end{aligned}
$$

这样就可以分治计算 $\textit{nums}$ 的打家劫舍了。

递归边界：如果 $a$ 的长度等于 $1$，那么按照定义，$f_{11}(a) = \max(a[0], 0)$，其余 $f_{00},f_{01},f_{10}$ 均为 $0$。

回到本题：

- 对于修改操作，我们可以用线段树的单点修改实现，线段树的每个节点维护对应区间的 $f_{00},f_{01},f_{10},f_{11}$。
- 对于查询操作，由于询问的是整个数组，询问结果就是线段树根节点的 $f_{11}$，累加到答案中。

具体请看 [视频讲解](https://www.bilibili.com/video/BV17t421N7L6/) 第四题，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maximumSumSubsequence(self, nums: List[int], queries: List[List[int]]) -> int:
        n = len(nums)
        # 4 个数分别保存 f00, f01, f10, f11
        t = [[0] * 4 for _ in range(2 << n.bit_length())]

        # 手写 max，效率更高
        def max(a: int, b: int) -> int:
            return b if b > a else a

        # 合并左右儿子
        def maintain(o: int):
            a, b = t[o * 2], t[o * 2 + 1]
            t[o][0] = max(a[0] + b[2], a[1] + b[0])
            t[o][1] = max(a[0] + b[3], a[1] + b[1])
            t[o][2] = max(a[2] + b[2], a[3] + b[0])
            t[o][3] = max(a[2] + b[3], a[3] + b[1])

        # 用 nums 初始化线段树
        def build(o: int, l: int, r: int) -> None:
            if l == r:
                t[o][3] = max(nums[l], 0)
                return
            m = (l + r) // 2
            build(o * 2, l, m)
            build(o * 2 + 1, m + 1, r)
            maintain(o)

        # 把 nums[i] 改成 val
        def update(o: int, l: int, r: int, i: int, val: int) -> None:
            if l == r:
                t[o][3] = max(val, 0)
                return
            m = (l + r) // 2
            if i <= m:
                update(o * 2, l, m, i, val)
            else:
                update(o * 2 + 1, m + 1, r, i, val)
            maintain(o)

        build(1, 0, n - 1)

        ans = 0
        for i, x in queries:
            update(1, 0, n - 1, i, x)
            ans += t[1][3]  # 注意 f11 没有任何限制，也就是整个数组的打家劫舍
        return ans % 1_000_000_007
```

```java [sol-Java]
class Solution {
    public int maximumSumSubsequence(int[] nums, int[][] queries) {
        int n = nums.length;
        // 4 个数分别保存 f00, f01, f10, f11
        long[][] t = new long[2 << (32 - Integer.numberOfLeadingZeros(n))][4];
        build(t, nums, 1, 0, n - 1);

        long ans = 0;
        for (int[] q : queries) {
            update(t, 1, 0, n - 1, q[0], q[1]);
            ans += t[1][3]; // 注意 f11 没有任何限制，也就是整个数组的打家劫舍
        }
        return (int) (ans % 1_000_000_007);
    }

    // 合并左右儿子
    private void maintain(long[][] t, int o) {
        long[] a = t[o * 2];
        long[] b = t[o * 2 + 1];
        t[o][0] = Math.max(a[0] + b[2], a[1] + b[0]);
        t[o][1] = Math.max(a[0] + b[3], a[1] + b[1]);
        t[o][2] = Math.max(a[2] + b[2], a[3] + b[0]);
        t[o][3] = Math.max(a[2] + b[3], a[3] + b[1]);
    }

    // 用 nums 初始化线段树
    private void build(long[][] t, int[] nums, int o, int l, int r) {
        if (l == r) {
            t[o][3] = Math.max(nums[l], 0);
            return;
        }
        int m = (l + r) / 2;
        build(t, nums, o * 2, l, m);
        build(t, nums, o * 2 + 1, m + 1, r);
        maintain(t, o);
    }

    // 把 nums[i] 改成 val
    private void update(long[][] t, int o, int l, int r, int i, int val) {
        if (l == r) {
            t[o][3] = Math.max(val, 0);
            return;
        }
        int m = (l + r) / 2;
        if (i <= m) {
            update(t, o * 2, l, m, i, val);
        } else {
            update(t, o * 2 + 1, m + 1, r, i, val);
        }
        maintain(t, o);
    }
}
```

```cpp [sol-C++]
class Solution {
    // 根据本题的数据范围，unsigned int 足矣，无需 long long
    // 4 个数分别保存 f00, f01, f10, f11
    vector<array<unsigned int, 4>> t;

    // 合并左右儿子
    void maintain(int o) {
        auto& a = t[o * 2];
        auto& b = t[o * 2 + 1];
        t[o] = {
            max(a[0] + b[2], a[1] + b[0]),
            max(a[0] + b[3], a[1] + b[1]),
            max(a[2] + b[2], a[3] + b[0]),
            max(a[2] + b[3], a[3] + b[1]),
        };
    }

    // 用 nums 初始化线段树
    void build(vector<int>& nums, int o, int l, int r) {
        if (l == r) {
            t[o][3] = max(nums[l], 0);
            return;
        }
        int m = (l + r) / 2;
        build(nums, o * 2, l, m);
        build(nums, o * 2 + 1, m + 1, r);
        maintain(o);
    };

    // 把 nums[i] 改成 val
    void update(int o, int l, int r, int i, int val) {
        if (l == r) {
            t[o][3] = max(val, 0);
            return;
        }
        int m = (l + r) / 2;
        if (i <= m) {
            update(o * 2, l, m, i, val);
        } else {
            update(o * 2 + 1, m + 1, r, i, val);
        }
        maintain(o);
    };

public:
    int maximumSumSubsequence(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        t.resize(2 << (32 - __builtin_clz(n)));
        build(nums, 1, 0, n - 1);

        long long ans = 0;
        for (auto& q : queries) {
            update(1, 0, n - 1, q[0], q[1]);
            ans += t[1][3]; // 注意 f11 没有任何限制，也就是整个数组的打家劫舍
        }
        return ans % 1'000'000'007;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

// 根据本题的数据范围，unsigned int 足矣，无需 long long
typedef struct {
    unsigned int f00; // 第一个数一定不选，最后一个数一定不选
    unsigned int f01; // 第一个数一定不选，最后一个数可选可不选
    unsigned int f10; // 第一个数可选可不选，最后一个数一定不选
    unsigned int f11; // 第一个数可选可不选，最后一个数可选可不选，也就是没有任何限制
} Data;

typedef Data* SegmentTree;

// 合并左右儿子
void maintain(SegmentTree t, int o) {
    Data a = t[o * 2], b = t[o * 2 + 1];
    t[o].f00 = MAX(a.f00 + b.f10, a.f01 + b.f00);
    t[o].f01 = MAX(a.f00 + b.f11, a.f01 + b.f01);
    t[o].f10 = MAX(a.f10 + b.f10, a.f11 + b.f00);
    t[o].f11 = MAX(a.f10 + b.f11, a.f11 + b.f01);
}

// 用 nums 初始化线段树
void build(SegmentTree t, int* nums, int o, int l, int r) {
    if (l == r) {
        t[o].f11 = MAX(nums[l], 0);
        return;
    }
    int m = (l + r) / 2;
    build(t, nums, o * 2, l, m);
    build(t, nums, o * 2 + 1, m + 1, r);
    maintain(t, o);
}

// 把 nums[i] 改成 val
void update(SegmentTree t, int o, int l, int r, int i, int val) {
    if (l == r) {
        t[o].f11 = MAX(val, 0);
        return;
    }
    int m = (l + r) / 2;
    if (i <= m) {
        update(t, o * 2, l, m, i, val);
    } else {
        update(t, o * 2 + 1, m + 1, r, i, val);
    }
    maintain(t, o);
}

int maximumSumSubsequence(int* nums, int numsSize, int** queries, int queriesSize, int* queriesColSize) {
    SegmentTree t = calloc(2 << (32 - __builtin_clz(numsSize)), sizeof(Data));
    build(t, nums, 1, 0, numsSize - 1);

    long long ans = 0;
    for (int i = 0; i < queriesSize; i++) {
        update(t, 1, 0, numsSize - 1, queries[i][0], queries[i][1]);
        ans += t[1].f11; // 注意 f11 没有任何限制，也就是整个数组的打家劫舍
    }

    free(t);
    return ans % 1000000007;
}
```

```go [sol-Go]
type data struct {
    f00 int // 第一个数一定不选，最后一个数一定不选
    f01 int // 第一个数一定不选，最后一个数可选可不选
    f10 int // 第一个数可选可不选，最后一个数一定不选
    f11 int // 第一个数可选可不选，最后一个数可选可不选，也就是没有任何限制
}

type seg []data

// 合并左右儿子
func (t seg) maintain(o int) {
    a, b := t[o<<1], t[o<<1|1]
    t[o] = data{
        max(a.f00+b.f10, a.f01+b.f00),
        max(a.f00+b.f11, a.f01+b.f01),
        max(a.f10+b.f10, a.f11+b.f00),
        max(a.f10+b.f11, a.f11+b.f01),
    }
}

// 用 a 初始化线段树
func (t seg) build(a []int, o, l, r int) {
    if l == r {
        t[o].f11 = max(a[l], 0)
        return
    }
    m := (l + r) >> 1
    t.build(a, o<<1, l, m)
    t.build(a, o<<1|1, m+1, r)
    t.maintain(o)
}

// 把 a[i] 改成 val
func (t seg) update(o, l, r, i, val int) {
    if l == r {
        t[o].f11 = max(val, 0)
        return
    }
    m := (l + r) >> 1
    if i <= m {
        t.update(o<<1, l, m, i, val)
    } else {
        t.update(o<<1|1, m+1, r, i, val)
    }
    t.maintain(o)
}

func maximumSumSubsequence(nums []int, queries [][]int) (ans int) {
    n := len(nums)
    t := make(seg, 2<<bits.Len(uint(n-1)))
    t.build(nums, 1, 0, n-1)

    for _, q := range queries {
        t.update(1, 0, n-1, q[0], q[1])
        ans += t[1].f11 // 注意 f11 没有任何限制，也就是整个数组的打家劫舍
    }
    return ans % 1_000_000_007
}
```

```js [sol-JavaScript]
var maximumSumSubsequence = function(nums, queries) {
    const n = nums.length;
    // 4 个数分别保存 f00, f01, f10, f11
    const t = Array.from({length: 2 << (32 - Math.clz32(n))}, () => [0, 0, 0, 0]);

    // 合并左右儿子
    function maintain(o) {
        const a = t[o * 2], b = t[o * 2 + 1];
        t[o][0] = Math.max(a[0] + b[2], a[1] + b[0]);
        t[o][1] = Math.max(a[0] + b[3], a[1] + b[1]);
        t[o][2] = Math.max(a[2] + b[2], a[3] + b[0]);
        t[o][3] = Math.max(a[2] + b[3], a[3] + b[1]);
    }

    // 用 nums 初始化线段树
    function build(o, l, r) {
        if (l === r) {
            t[o][3] = Math.max(nums[l], 0);
            return;
        }
        const m = Math.floor((l + r) / 2);
        build(o * 2, l, m);
        build(o * 2 + 1, m + 1, r);
        maintain(o);
    }

    // 把 nums[i] 改成 val
    function update(o, l, r, i, val) {
        if (l === r) {
            t[o][3] = Math.max(val, 0);
            return;
        }
        const m = Math.floor((l + r) / 2);
        if (i <= m) {
            update(o * 2, l, m, i, val);
        } else {
            update(o * 2 + 1, m + 1, r, i, val);
        }
        maintain(o);
    }

    build(1, 0, n - 1);

    let ans = 0;
    for (const [i, x] of queries) {
        update(1, 0, n - 1, i, x);
        ans += t[1][3]; // 注意 f11 没有任何限制，也就是整个数组的打家劫舍
    }
    return ans % 1_000_000_007;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_sum_subsequence(nums: Vec<i32>, queries: Vec<Vec<i32>>) -> i32 {
        let n = nums.len();
        // 4 个数分别保存 f00, f01, f10, f11
        let mut t = vec![vec![0u32; 4]; 2 << (64 - n.leading_zeros())];
        Self::build(&mut t, &nums, 1, 0, n - 1);

        let mut ans = 0u64;
        for q in queries {
            Self::update(&mut t, 1, 0, n - 1, q[0] as usize, q[1]);
            ans += t[1][3] as u64; // 注意 f11 没有任何限制，也就是整个数组的打家劫舍
        }
        (ans % 1_000_000_007) as _
    }

    // 合并左右儿子
    fn maintain(t: &mut Vec<Vec<u32>>, o: usize) {
        let l = o * 2;
        let r = o * 2 + 1;
        t[o][0] = (t[l][0] + t[r][2]).max(t[l][1] + t[r][0]);
        t[o][1] = (t[l][0] + t[r][3]).max(t[l][1] + t[r][1]);
        t[o][2] = (t[l][2] + t[r][2]).max(t[l][3] + t[r][0]);
        t[o][3] = (t[l][2] + t[r][3]).max(t[l][3] + t[r][1]);
    }

    // 用 nums 初始化线段树
    fn build(t: &mut Vec<Vec<u32>>, nums: &[i32], o: usize, l: usize, r: usize) {
        if l == r {
            t[o][3] = nums[l].max(0) as u32;
            return;
        }
        let m = (l + r) / 2;
        Self::build(t, nums, o * 2, l, m);
        Self::build(t, nums, o * 2 + 1, m + 1, r);
        Self::maintain(t, o);
    }

    // 把 nums[i] 改成 val
    fn update(t: &mut Vec<Vec<u32>>, o: usize, l: usize, r: usize, i: usize, val: i32) {
        if l == r {
            t[o][3] = val.max(0) as u32;
            return;
        }
        let m = (l + r) / 2;
        if i <= m {
            Self::update(t, o * 2, l, m, i, val);
        } else {
            Self::update(t, o * 2 + 1, m + 1, r, i, val);
        }
        Self::maintain(t, o);
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 总结

如果一个题目可以用分治解决，那么这个题目的带修改版本可以用线段树解决。

## 相似题目

- [2213. 由单个字符重复的最长子字符串](https://leetcode.cn/problems/longest-substring-of-one-repeating-character/)
- [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/) 的分治做法，见官方题解的方法二。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

设 $1$ 在数组中的下标为 $p$，$n$ 在数组中的下标为 $q$。

分类讨论：

- 如果 $p<q$，那么 $1$ 和 $n$ 井水不犯河水，$1$ 移动（交换）到数组的最左边，$n$ 移动到数组的最右边，操作次数为 $p + (n-1-q)$。
- 否则 $p>q$，那么 $1$ 和 $n$ 在移动过程中会相遇，互相穿过对方，也就是只花费一次操作，就让两个数都移动了一步（互相穿过），所以操作次数比上面的情况要少 $1$，即 $p + (n-1-q) - 1$。

```py [sol-Python3]
class Solution:
    def semiOrderedPermutation(self, nums: List[int]) -> int:
        n = len(nums)
        p = nums.index(1)
        q = nums.index(n)
        return p + n - 1 - q - (p > q)
```

```java [sol-Java]
class Solution {
    public int semiOrderedPermutation(int[] nums) {
        int n = nums.length;
        int p = 0;
        int q = 0;
        for (int i = 0; i < n; i++) {
            if (nums[i] == 1) {
                p = i;
            } else if (nums[i] == n) {
                q = i;
            }
        }
        return p + n - 1 - q - (p > q ? 1 : 0);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int semiOrderedPermutation(vector<int>& nums) {
        auto [p, q] = ranges::minmax_element(nums);
        return p - q + nums.size() - 1 - (p > q);
    }
};
```

```c [sol-C]
int semiOrderedPermutation(int* nums, int n) {
    int p, q;
    for (int i = 0; i < n; i++) {
        if (nums[i] == 1) {
            p = i;
        } else if (nums[i] == n) {
            q = i;
        }
    }
    return p + n - 1 - q - (p > q);
}
```

```go [sol-Go]
func semiOrderedPermutation(nums []int) int {
    n := len(nums)
    p := slices.Index(nums, 1)
    q := slices.Index(nums, n)
    if p < q {
        return p + n - 1 - q
    }
    return p + n - 2 - q // 1 向左移动的时候和 n 交换了一次
}
```

```js [sol-JavaScript]
var semiOrderedPermutation = function(nums) {
    const n = nums.length;
    const p = nums.indexOf(1);
    const q = nums.indexOf(n);
    return p + n - 1 - q - (p > q ? 1 : 0);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn semi_ordered_permutation(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let p = nums.iter().position(|&x| x == 1).unwrap();
        let q = nums.iter().position(|&x| x == n as i32).unwrap();
        (p + n - 1 - q - (p > q) as usize) as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

由于元素值一定不小于其数位和，所以答案就是元素和减去数位和，无需取绝对值。

```py [sol-Python3]
class Solution:
    def differenceOfSum(self, nums: List[int]) -> int:
        ans = 0
        for x in nums:
            ans += x  # 累加元素和
            while x:
                ans -= x % 10  # 减去数位和
                x //= 10
        return ans
```

```java [sol-Java]
class Solution {
    public int differenceOfSum(int[] nums) {
        int ans = 0;
        for (int x : nums) {
            ans += x; // 累加元素和
            while (x > 0) {
                ans -= x % 10; // 减去数位和
                x /= 10;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int differenceOfSum(vector<int>& nums) {
        int ans = 0;
        for (int x : nums) {
            ans += x; // 累加元素和
            while (x) {
                ans -= x % 10; // 减去数位和
                x /= 10;
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int differenceOfSum(int* nums, int numsSize) {
    int ans = 0;
    for (int i = 0; i < numsSize; i++) {
        int x = nums[i];
        ans += x; // 累加元素和
        while (x) {
            ans -= x % 10; // 减去数位和
            x /= 10;
        }
    }
    return ans;
}
```

```go [sol-Go]
func differenceOfSum(nums []int) (ans int) {
    for _, x := range nums {
        ans += x // 累加元素和
        for x > 0 {
            ans -= x % 10 // 减去数位和
            x /= 10
        }
    }
    return
}
```

```js [sol-JavaScript]
var differenceOfSum = function(nums) {
    let ans = 0;
    for (let x of nums) {
        ans += x; // 累加元素和
        while (x) {
            ans -= x % 10; // 减去数位和
            x = Math.floor(x / 10);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn difference_of_sum(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        for mut x in nums {
            ans += x; // 累加元素和
            while x > 0 {
                ans -= x % 10; // 减去数位和
                x /= 10;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$O(1)$，仅用到若干额外变量。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

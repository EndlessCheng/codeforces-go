## 思路

由于 $\textit{nums}[i]$ 会被发生在 $i$ 左侧的操作影响，我们先从最左边的 $\textit{nums}[0]$ 开始思考。

讨论是否要在 $i=0$ 处操作：

- 如果 $\textit{nums}[0]=1$，不需要操作，问题变成剩下 $n-1$ 个数的子问题。
- 如果 $\textit{nums}[0]=0$，一定要操作，问题变成剩下 $n-1$ 个数（在操作次数是 $1$ 的情况下）的子问题。

对后续元素来说，由于反转偶数次等于没反转，所以只需考虑操作次数的**奇偶性**。

一般地，设遍历到 $x=\textit{nums}[i]$ 时，之前执行了 $k$ 次操作，分类讨论：

- 如果 $x=0$ 且 $k$ 是奇数，或者 $x=1$ 且 $k$ 是偶数，那么这 $k$ 次操作执行完后 $\textit{nums}[i]$ 变成 $1$。所以如果 $x\ne k\bmod 2$，则不需要操作。
- 如果 $x=0$ 且 $k$ 是偶数，或者 $x=1$ 且 $k$ 是奇数，那么这 $k$ 次操作执行完后 $\textit{nums}[i]$ 变成 $0$。所以如果 $x= k\bmod 2$，则一定要操作。

## 正确性

**问**：为什么这样做是对的？

**答**：

1. 先选择 $i$ 操作再选择 $j$ 操作（$i\ne j$），和先选择 $j$ 操作再选择 $i$ 操作的结果是一样的，所以操作顺序不影响答案。既然操作顺序无影响，我们可以从左到右操作。或者说，假设某种操作顺序是最优的，那么总是可以把这个操作顺序重排成从左到右选择下标操作。
2. 对于同一个 $i$，选择 $i$ 操作两次（偶数次）等于没有操作，所以同一个 $i$ 至多选择一次。
3. 结合上述两点，既然同一个 $i$ 至多选择一次，那么从左到右操作的过程中，遇到 $1$ 一定不能操作，遇到 $0$ 一定要操作，所以**从左到右的操作方式有且仅有一种**。
4. 既然操作方式是唯一的，我们只需**模拟**这个过程。

**问**：题目要求的「最少」体现在哪里？

**答**：对同一个 $i$ 至多选择一次，就可以做到最少的操作次数。

具体请看 [视频讲解](https://www.bilibili.com/video/BV17w4m1e7Nw/) 第三题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int]) -> int:
        k = 0
        for x in nums:
            if x == k % 2:  # 必须操作
                k += 1
        return k
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums) {
        int k = 0;
        for (int x : nums) {
            if (x == k % 2) { // 必须操作
                k++;
            }
        }
        return k;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums) {
        int k = 0;
        for (int x : nums) {
            if (x == k % 2) { // 必须操作
                k++;
            }
        }
        return k;
    }
};
```

```c [sol-C]
int minOperations(int* nums, int numsSize) {
    int k = 0;
    for (int i = 0; i < numsSize; i++) {
        if (nums[i] == k % 2) { // 必须操作
            k++;
        }
    }
    return k;
}
```

```go [sol-Go]
func minOperations(nums []int) (k int) {
	for _, x := range nums {
		if x == k%2 { // 必须操作
			k++
		}
	}
	return
}
```

```js [sol-JavaScript]
var minOperations = function(nums) {
    let k = 0;
    for (const x of nums) {
        if (x === k % 2) { // 必须操作
            k++;
        }
    }
    return k;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn min_operations(nums: Vec<i32>) -> i32 {
        let mut k = 0;
        for x in nums {
            if x == k % 2 { // 必须操作
                k += 1;
            }
        }
        k
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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

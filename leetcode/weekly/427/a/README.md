操作相当于从下标 $i$ 移动到下标 $i+\textit{nums}[i]$。

如果 $i+\textit{nums}[i]$ 下标越界呢？

需要把 $i+\textit{nums}[i]$ 调整到 $[0,n-1]$ 范围中。具体来说，把下标 $i+\textit{nums}[i]$ 模 $n$。比如 $n=4$，在循环数组中，正数下标 $5,9,13,\ldots$ 都是下标 $1$，负数下标 $-3,-7,-11,\ldots$ 也都是下标 $1$。

不了解取模的同学，请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

[本题视频讲解](https://www.bilibili.com/video/BV1YeqHYSEhK/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def constructTransformedArray(self, nums: List[int]) -> List[int]:
        n = len(nums)
        return [nums[(i + x) % n] for i, x in enumerate(nums)]
```

```java [sol-Java]
class Solution {
    public int[] constructTransformedArray(int[] nums) {
        int n = nums.length;
        int[] result = new int[n];
        for (int i = 0; i < n; i++) {
            result[i] = nums[((i + nums[i]) % n + n) % n]; // 保证结果在 [0,n-1] 中
        }
        return result;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> constructTransformedArray(vector<int>& nums) {
        int n = nums.size();
        vector<int> result(n);
        for (int i = 0; i < n; i++) {
            result[i] = nums[((i + nums[i]) % n + n) % n]; // 保证结果在 [0,n-1] 中
        }
        return result;
    }
};
```

```c [sol-C]
int* constructTransformedArray(int* nums, int numsSize, int* returnSize) {
    int n = numsSize;
    int* result = malloc(n * sizeof(int));
    for (int i = 0; i < n; i++) {
        result[i] = nums[((i + nums[i]) % n + n) % n]; // 保证结果在 [0,n-1] 中
    }
    *returnSize = n;
    return result;
}
```

```go [sol-Go]
func constructTransformedArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i, x := range nums {
		result[i] = nums[((i+x)%n+n)%n] // 保证结果在 [0,n-1] 中
	}
	return result
}
```

```js [sol-JavaScript]
var constructTransformedArray = function(nums) {
    const n = nums.length;
    const result = new Array(n);
    for (let i = 0; i < n; i++) {
        result[i] = nums[((i + nums[i]) % n + n) % n]; // 保证结果在 [0,n-1] 中
    }
    return result;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn construct_transformed_array(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let m = n as i32;
        let mut result = vec![0; n];
        for i in 0..n {
            let j = ((i as i32 + nums[i]) % m + m) % m; // 保证结果在 [0,n-1] 中
            result[i] = nums[j as usize];
        }
        result
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

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

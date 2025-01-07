由于所有小于 $k$ 的数都要被删除，所以统计 $\textit{nums}$ 中的小于 $k$ 的元素个数即可。

```py [sol-Python3]
class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        return sum(x < k for x in nums)
```

```java [sol-Java]
class Solution {
    public int minOperations(int[] nums, int k) {
        int ans = 0;
        for (int x : nums) {
            if (x < k) {
                ans++;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minOperations(vector<int>& nums, int k) {
        int ans = 0;
        for (int x : nums) {
            ans += x < k;
        }
        return ans;
    }
};
```

```cpp [sol-C++ 写法二]
class Solution {
public:
    int minOperations(vector<int>& nums, int k) {
        return ranges::count_if(nums, [&](int x) {
            return x < k;
        });
    }
};
```

```c [sol-C]
int minOperations(int* nums, int numsSize, int k) {
    int ans = 0;
    for (int i = 0; i < numsSize; i++) {
        ans += nums[i] < k;
    }
    return ans;
}
```

```go [sol-Go]
func minOperations(nums []int, k int) (ans int) {
	for _, x := range nums {
		if x < k {
			ans++
		}
	}
	return
}
```

```js [sol-JS]
var minOperations = function(nums, k) {
    let ans = 0;
    for (const x of nums) {
        if (x < k) {
            ans++;
        }
    }
    return ans;
};
```

```js [sol-JS 写法二]
var minOperations = function(nums, k) {
    return nums.filter(x => x < k).length;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn min_operations(nums: Vec<i32>, k: i32) -> i32 {
        nums.iter().filter(|&&x| x < k).count() as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
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

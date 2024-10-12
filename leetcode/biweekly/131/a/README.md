遍历 $\textit{nums}$，同时用一个 $\textit{vis}$ 集合记录遇到的数字。

- 设 $x=\textit{nums}[i]$。
- 如果 $x$ 不在 $\textit{vis}$ 中，说明是第一次遇到，加入 $\textit{vis}$。
- 如果 $x$ 在 $\textit{vis}$ 中，说明是第二次遇到（注意每个数至多出现两次），加入答案（异或）。

代码实现时，由于 $\textit{nums}[i]\le 50$，可以用二进制数表示集合，具体见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

[本题视频讲解](https://www.bilibili.com/video/BV1SU411d7wj/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def duplicateNumbersXOR(self, nums: List[int]) -> int:
        ans = vis = 0
        for x in nums:
            if vis >> x & 1:  # x 在 vis 中
                ans ^= x
            else:
                vis |= 1 << x  # 把 x 加到 vis 中
        return ans
```

```java [sol-Java]
class Solution {
    public int duplicateNumbersXOR(int[] nums) {
        int ans = 0;
        long vis = 0;
        for (int x : nums) {
            if ((vis >> x & 1) > 0) { // x 在 vis 中
                ans ^= x;
            } else {
                vis |= 1L << x; // 把 x 加到 vis 中
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int duplicateNumbersXOR(vector<int>& nums) {
        int ans = 0;
        long long vis = 0;
        for (int x : nums) {
            if (vis >> x & 1) { // x 在 vis 中
                ans ^= x;
            } else {
                vis |= 1LL << x; // 把 x 加到 vis 中
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int duplicateNumbersXOR(int* nums, int numsSize) {
    int ans = 0;
    long long vis = 0;
    for (int i = 0; i < numsSize; i++) {
        int x = nums[i];
        if (vis >> x & 1) { // x 在 vis 中
            ans ^= x;
        } else {
            vis |= 1LL << x; // 把 x 加到 vis 中
        }
    }
    return ans;
}
```

```go [sol-Go]
func duplicateNumbersXOR(nums []int) (ans int) {
	vis := 0
	for _, x := range nums {
		if vis>>x&1 > 0 { // x 在 vis 中
			ans ^= x
		} else {
			vis |= 1 << x // 把 x 加到 vis 中
		}
	}
	return
}
```

```js [sol-JavaScript]
var duplicateNumbersXOR = function(nums) {
    let ans = 0;
    // JS 的位运算会强转成 32 位整数，需要用 BigInt 处理
    let vis = 0n;
    for (const x of nums) {
        if (vis >> BigInt(x) & 1n) { // x 在 vis 中
            ans ^= x;
        } else {
            vis |= 1n << BigInt(x); // 把 x 加到 vis 中
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn duplicate_numbers_xor(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut vis = 0i64;
        for x in nums {
            if (vis >> x & 1) > 0 { // x 在 vis 中
                ans ^= x;
            } else {
                vis |= 1 << x; // 把 x 加到 vis 中
            }
        }
        ans
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

本题可以视作一个「**小球打砖块**」游戏，具体请看 [本题视频讲解](https://www.bilibili.com/video/BV1yiU6YnEfU/) 画的图。

每打掉一个砖块，小球的移动方向就要反向，所以要想打掉所有砖块，起始位置左右的砖块个数之差不能超过 $1$。

设整个数组的元素和为 $\textit{total}$。遍历数组的同时，维护前缀和 $pre$。

如果 $\textit{nums}[i]=0$，那么这个位置可以是起始位置吗？分类讨论：

- 如果前缀和 $\textit{pre}$ 等于后缀和 $\textit{total}-\textit{pre}$，那么小球初始方向可以向左，也可以向右，答案加 $2$。
- 如果前缀和比后缀和多 $1$，那么小球初始方向必须向左，才能打掉所有砖块，答案加 $1$。
- 如果前缀和比后缀和少 $1$，那么小球初始方向必须向右，才能打掉所有砖块，答案加 $1$。
- 其余情况，不能是起始位置。

```py [sol-Python3]
class Solution:
    def countValidSelections(self, nums: List[int]) -> int:
        total = sum(nums)
        ans = pre = 0
        for x in nums:
            if x:
                pre += x
            elif pre * 2 == total:
                ans += 2
            elif abs(pre * 2 - total) == 1:
                ans += 1
        return ans
```

```java [sol-Java]
class Solution {
    public int countValidSelections(int[] nums) {
        int total = 0;
        for (int x : nums) {
            total += x;
        }

        int ans = 0;
        int pre = 0;
        for (int x : nums) {
            if (x > 0) {
                pre += x;
            } else if (pre * 2 == total) {
                ans += 2;
            } else if (Math.abs(pre * 2 - total) == 1) {
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
    int countValidSelections(vector<int>& nums) {
        int total = reduce(nums.begin(), nums.end());
        int ans = 0, pre = 0;
        for (int x : nums) {
            if (x) {
                pre += x;
            } else if (pre * 2 == total) {
                ans += 2;
            } else if (abs(pre * 2 - total) == 1) {
                ans++;
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int countValidSelections(int* nums, int numsSize) {
    int total = 0;
    for (int i = 0; i < numsSize; i++) {
        total += nums[i];
    }

    int ans = 0, pre = 0;
    for (int i = 0; i < numsSize; i++) {
        if (nums[i]) {
            pre += nums[i];
        } else if (pre * 2 == total) {
            ans += 2;
        } else if (abs(pre * 2 - total) == 1) {
            ans++;
        }
    }
    return ans;
}
```

```go [sol-Go]
func countValidSelections(nums []int) (ans int) {
	total := 0
	for _, x := range nums {
		total += x
	}

	pre := 0
	for _, x := range nums {
		if x > 0 {
			pre += x
		} else if pre*2 == total {
			ans += 2
		} else if abs(pre*2-total) == 1 {
			ans++
		}
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

```js [sol-JavaScript]
var countValidSelections = function(nums) {
    const total = _.sum(nums);
    let ans = 0, pre = 0;
    for (const x of nums) {
        if (x) {
            pre += x;
        } else if (pre * 2 === total) {
            ans += 2;
        } else if (Math.abs(pre * 2 - total) === 1) {
            ans += 1;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_valid_selections(nums: Vec<i32>) -> i32 {
        let total = nums.iter().sum::<i32>();
        let mut ans = 0;
        let mut pre = 0;
        for x in nums {
            if x != 0 {
                pre += x;
            } else if pre * 2 == total {
                ans += 2;
            } else if (pre * 2 - total).abs() == 1 {
                ans += 1;
            }
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面动态规划题单的「**专题：前后缀分解**」。

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

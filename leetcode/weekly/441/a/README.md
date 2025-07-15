为了让元素和尽量大，负数不能留，可以全部去掉。但如果 $\textit{nums}$ 中的元素都是负数，题目规定不能全部去掉，此时答案为 $\textit{nums}$ 中的最大元素（绝对值最小的负数）。

如果有非负数，那么能选则选。题目规定「子数组中的所有元素互不相同」，意味着每个非负数只能选一个，所以答案就是 $\textit{nums}$ 中的非负数去重后的元素和。可以用哈希集合判断元素是否选过。

```py [sol-Python3]
class Solution:
    def maxSum(self, nums: List[int]) -> int:
        st = set(x for x in nums if x >= 0)  # 去掉非负数并去重
        return sum(st) if st else max(nums)
```

```py [sol-Python3 写法二]
class Solution:
    def maxSum(self, nums: List[int]) -> int:
        return sum(set(x for x in nums if x >= 0)) or max(nums)
```

```java [sol-Java]
class Solution {
    public int maxSum(int[] nums) {
        Set<Integer> set = new HashSet<>();
        int s = 0;
        int mx = Integer.MIN_VALUE;
        for (int x : nums) { // 一次遍历
            if (x < 0) {
                mx = Math.max(mx, x); // 计算最大负数
            } else if (set.add(x)) { // x 不在 set 中
                s += x; // 相同元素只留一个，累加元素和
            }
        }
        return set.isEmpty() ? mx : s;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxSum(vector<int>& nums) {
        unordered_set<int> st;
        int s = 0, mx = INT_MIN;
        for (int x : nums) { // 一次遍历
            if (x < 0) {
                mx = max(mx, x); // 计算最大负数
            } else if (st.insert(x).second) { // x 不在 set 中
                s += x; // 相同元素只留一个，累加元素和
            }
        }
        return st.empty() ? mx : s;
    }
};
```

```go [sol-Go]
func maxSum(nums []int) (ans int) {
	set := map[int]struct{}{}
	mx := math.MinInt
	for _, x := range nums { // 一次遍历
		if x < 0 {
			mx = max(mx, x) // 计算最大负数
		} else if _, ok := set[x]; !ok {
			set[x] = struct{}{}
			ans += x // 相同元素只留一个，累加元素和
		}
	}
	if len(set) == 0 {
		return mx
	}
	return
}
```

```js [sol-JavaScript]
var maxSum = function(nums) {
    const set = new Set();
    let s = 0;
    let mx = -Infinity;
    for (const x of nums) { // 一次遍历
        if (x < 0) {
            mx = Math.max(mx, x); // 计算最大负数
        } else if (!set.has(x)) { // x 不在 set 中
            set.add(x);
            s += x; // 相同元素只留一个，累加元素和
        }
    }
    return set.size ? s : mx;
};
```

```rust [sol-Rust]
use std::collections::HashSet;

impl Solution {
    pub fn max_sum(nums: Vec<i32>) -> i32 {
        let mut set = HashSet::new();
        let mut s = 0;
        let mut mx = i32::MIN;
        for x in nums { // 一次遍历
            if x < 0 {
                mx = mx.max(x); // 计算最大负数
            } else if set.insert(x) { // x 不在 set 中
                s += x; // 相同元素只留一个，累加元素和
            }
        }
        if set.is_empty() { mx } else { s }
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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

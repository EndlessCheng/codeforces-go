根据题意，如果 $\textit{nums}[i]$ 的出现次数超过 $2$，则无法分割，否则可以分割。

```py [sol-Python3]
class Solution:
    def isPossibleToSplit(self, nums: List[int]) -> bool:
        return max(Counter(nums).values()) <= 2
```

```py [sol-Python3 写法二]
class Solution:
    def isPossibleToSplit(self, nums: List[int]) -> bool:
        return all(c <= 2 for c in Counter(nums).values())
```

```java [sol-Java]
class Solution {
    public boolean isPossibleToSplit(int[] nums) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums) {
            if (cnt.merge(x, 1, Integer::sum) > 2) { // ++cnt[x] > 2
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool isPossibleToSplit(vector<int>& nums) {
        unordered_map<int, int> cnt;
        for (int x : nums) {
            if (++cnt[x] > 2) {
                return false;
            }
        }
        return true;
    }
};
```

```c [sol-C]
bool isPossibleToSplit(int* nums, int numsSize) {
    int cnt[101] = {};
    for (int i = 0; i < numsSize; i++) {
        if (++cnt[nums[i]] > 2) {
            return false;
        }
    }
    return true;
}
```

```go [sol-Go]
func isPossibleToSplit(nums []int) bool {
    cnt := map[int]int{}
    for _, x := range nums {
        cnt[x]++
        if cnt[x] > 2 {
            return false
        }
    }
    return true
}
```

```js [sol-JS]
var isPossibleToSplit = function(nums) {
    const cnt = new Map();
    for (const x of nums) {
        const c = (cnt.get(x) ?? 0) + 1;
        if (c > 2) {
            return false;
        }
        cnt.set(x, c);
    }
    return true;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn is_possible_to_split(nums: Vec<i32>) -> bool {
        let mut cnt = HashMap::new();
        for x in nums {
            *cnt.entry(x).or_insert(0) += 1;
            if cnt[&x] > 2 {
                return false;
            }
        }
        true
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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

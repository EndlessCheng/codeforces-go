读题：每次操作，将位于 $\textit{moveFrom}[i]$ 的**所有石块**全部移到 $\textit{moveTo}[i]$。最后返回有石块的位置。

注意这个过程并不关心每个位置有多少石块。所以只需用哈希集合来维护这些石头的位置，不需要维护每个位置有多少石块。

1. 把所有 $\textit{nums}[i]$ 加到一个哈希集合中。
2. 遍历 $\textit{moveFrom}$ 和 $\textit{moveTo}$，先把 $\textit{moveFrom}[i]$ 从哈希集合中去掉，然后把 $\textit{moveTo}[i]$ 加入哈希集合。
3. 取出哈希集合中的元素，从小到大排序后返回。

```py [sol-Python3]
class Solution:
    def relocateMarbles(self, nums: List[int], moveFrom: List[int], moveTo: List[int]) -> List[int]:
        st = set(nums)
        for f, t in zip(moveFrom, moveTo):
            st.remove(f)
            st.add(t)
        return sorted(st)
```

```java [sol-Java]
class Solution {
    public List<Integer> relocateMarbles(int[] nums, int[] moveFrom, int[] moveTo) {
        Set<Integer> set = new HashSet<>(nums.length); // 预分配空间，效率更高
        for (int x : nums) {
            set.add(x);
        }

        for (int i = 0; i < moveFrom.length; i++) {
            set.remove(moveFrom[i]);
            set.add(moveTo[i]);
        }

        List<Integer> ans = new ArrayList<>(set);
        Collections.sort(ans);
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> relocateMarbles(vector<int>& nums, vector<int>& moveFrom, vector<int>& moveTo) {
        unordered_set<int> st(nums.begin(), nums.end());
        for (int i = 0; i < moveFrom.size(); i++) {
            st.erase(moveFrom[i]);
            st.insert(moveTo[i]);
        }
        vector<int> ans(st.begin(), st.end());
        ranges::sort(ans);
        return ans;
    }
};
```

```go [sol-Go]
func relocateMarbles(nums, moveFrom, moveTo []int) []int {
    set := map[int]struct{}{}
    for _, x := range nums {
        set[x] = struct{}{}
    }

    for i, x := range moveFrom {
        delete(set, x)
        set[moveTo[i]] = struct{}{}
    }

    ans := make([]int, 0, len(set))
    for x := range set {
        ans = append(ans, x)
    }
    slices.Sort(ans)
    return ans
}
```

```js [sol-JavaScript]
var relocateMarbles = function(nums, moveFrom, moveTo) {
    const set = new Set(nums);
    for (let i = 0; i < moveFrom.length; i++) {
        set.delete(moveFrom[i]);
        set.add(moveTo[i]);
    }
    const ans = Array.from(set);
    ans.sort((a, b) => a - b);
    return ans;
};
```

```rust [sol-Rust]
use std::collections::HashSet;

impl Solution {
    pub fn relocate_marbles(nums: Vec<i32>, move_from: Vec<i32>, move_to: Vec<i32>) -> Vec<i32> {
        let mut set = nums.into_iter().collect::<HashSet<_>>();
        for (f, t) in move_from.into_iter().zip(move_to) {
            set.remove(&f);
            set.insert(t);
        }
        let mut ans = set.into_iter().collect::<Vec<_>>();
        ans.sort_unstable();
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + m)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m$ 为 $\textit{moveFrom}$ 的长度。排序是瓶颈。注意哈希集合中至多有 $\mathcal{O}(n)$ 个元素。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

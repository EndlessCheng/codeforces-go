![lc2610-c.png](https://pic.leetcode.cn/1741320320-YkycVn-lc2610-c.png){:width=350}

```py [sol-Python3]
class Solution:
    def findMatrix(self, nums: List[int]) -> List[List[int]]:
        ans = []
        cnt = Counter(nums)  # 统计每个元素的出现次数
        while cnt:  # 还有剩余元素
            row = list(cnt)
            ans.append(row)
            # cnt 中的每个元素的出现次数都减一
            for x in row:
                cnt[x] -= 1
                if cnt[x] == 0:
                    del cnt[x]  # 去掉出现次数为 0 的元素
        return ans
```

```java [sol-Java]
class Solution {
    public List<List<Integer>> findMatrix(int[] nums) {
        // 统计每个元素的出现次数
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums) {
            cnt.merge(x, 1, Integer::sum); // cnt[x]++
        }

        List<List<Integer>> ans = new ArrayList<>();
        while (!cnt.isEmpty()) { // 还有剩余元素
            List<Integer> row = new ArrayList<>(cnt.keySet());
            ans.add(row);
            // cnt 中的每个元素的出现次数都减一
            for (Integer x : row) {
                int c = cnt.get(x) - 1;
                if (c == 0) {
                    cnt.remove(x); // 去掉出现次数为 0 的元素
                } else {
                    cnt.put(x, c); // 更新出现次数
                }
            }
        }
        return ans;
    }
}
```

```java [sol-Java 写法二]
class Solution {
    public List<List<Integer>> findMatrix(int[] nums) {
        // 统计每个元素的出现次数
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums) {
            cnt.merge(x, 1, Integer::sum); // cnt[x]++
        }

        List<List<Integer>> ans = new ArrayList<>();
        while (!cnt.isEmpty()) { // 还有剩余元素
            List<Integer> row = new ArrayList<>(cnt.size()); // 预分配空间
            // 一边遍历哈希表，一边删除元素
            Iterator<Map.Entry<Integer, Integer>> it = cnt.entrySet().iterator();
            while (it.hasNext()) {
                Map.Entry<Integer, Integer> e = it.next();
                row.add(e.getKey());
                int c = e.getValue() - 1; // 出现次数减一
                if (c == 0) {
                    it.remove(); // 去掉出现次数为 0 的元素
                } else {
                    e.setValue(c); // 更新出现次数
                }
            }
            ans.add(row);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> findMatrix(vector<int>& nums) {
        // 统计每个元素的出现次数
        unordered_map<int, int> cnt;
        for (int x : nums) {
            cnt[x]++;
        }

        vector<vector<int>> ans;
        while (!cnt.empty()) {
            vector<int> row;
            // cnt 中的每个元素的出现次数都减一
            // 一边遍历哈希表，一边删除元素
            for (auto it = cnt.begin(); it != cnt.end();) {
                row.push_back(it->first);
                if (--it->second == 0) {
                    it = cnt.erase(it);
                } else {
                    it++;
                }
            }
            ans.push_back(row);
        }
        return ans;
    }
};
```

```go [sol-Go]
func findMatrix(nums []int) (ans [][]int) {
    // 统计每个元素的出现次数
    cnt := map[int]int{}
    for _, x := range nums {
        cnt[x]++
    }

    for len(cnt) > 0 {
        row := make([]int, 0, len(cnt)) // 预分配空间
        // cnt 中的每个元素的出现次数都减一
        for x := range cnt {
            row = append(row, x)
            cnt[x]--
            if cnt[x] == 0 {
                delete(cnt, x) // 删除当前正在遍历的元素是安全的
            }
        }
        ans = append(ans, row)
    }
    return
}
```

```js [sol-JavaScript]
var findMatrix = function(nums) {
    // 统计每个元素的出现次数
    const cnt = new Map();
    for (const x of nums) {
        cnt.set(x, (cnt.get(x) ?? 0) + 1);
    }

    const ans = [];
    while (cnt.size) { // 还有剩余元素
        const row = [...cnt.keys()];
        ans.push(row);
        // cnt 中的每个元素的出现次数都减一
        for (const x of row) {
            const c = cnt.get(x) - 1;
            if (c === 0) {
                cnt.delete(x); // 去掉出现次数为 0 的元素
            } else {
                cnt.set(x, c); // 更新出现次数
            }
        }
    }
    return ans;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn find_matrix(nums: Vec<i32>) -> Vec<Vec<i32>> {
        // 统计每个元素的出现次数
        let mut cnt = HashMap::new();
        for x in nums {
            *cnt.entry(x).or_insert(0) += 1;
        }

        let mut ans = vec![];
        while !cnt.is_empty() { // 还有剩余元素
            let row = cnt.keys().cloned().collect::<Vec<_>>();
            ans.push(row.clone());
            // cnt 中的每个元素的出现次数都减一
            for x in row {
                let c = cnt.get_mut(&x).unwrap();
                *c -= 1;
                if *c == 0 {
                    cnt.remove(&x); // 去掉出现次数为 0 的元素
                }
            }
        }
        ans
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

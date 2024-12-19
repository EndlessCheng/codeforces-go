用 $\textit{pos}$ 数组记录所有等于 $x$ 的 $\textit{nums}[i]$ 的下标 $i$。

对于每个询问 $q=\textit{queries}[i]$，如果 $q$ 大于 $\textit{pos}$ 的长度，则答案为 $-1$，否则答案为 $\textit{pos}[q-1]$。

```py [sol-Python3]
class Solution:
    def occurrencesOfElement(self, nums: List[int], queries: List[int], x: int) -> List[int]:
        pos = [i for i, num in enumerate(nums) if num == x]
        return [-1 if q > len(pos) else pos[q - 1] for q in queries]
```

```java [sol-Java]
class Solution {
    public int[] occurrencesOfElement(int[] nums, int[] queries, int x) {
        List<Integer> pos = new ArrayList<>();
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == x) {
                pos.add(i);
            }
        }
        for (int i = 0; i < queries.length; i++) {
            queries[i] = queries[i] > pos.size() ? -1 : pos.get(queries[i] - 1);
        }
        return queries;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> occurrencesOfElement(vector<int>& nums, vector<int>& queries, int x) {
        vector<int> pos;
        for (int i = 0; i < nums.size(); i++) {
            if (nums[i] == x) {
                pos.push_back(i);
            }
        }
        for (int& q : queries) {
            q = q > pos.size() ? -1 : pos[q - 1];
        }
        return queries;
    }
};
```

```c [sol-C]
int* occurrencesOfElement(int* nums, int numsSize, int* queries, int queriesSize, int x, int* returnSize) {
    int* pos = malloc(numsSize * sizeof(int));
    int k = 0;
    for (int i = 0; i < numsSize; i++) {
        if (nums[i] == x) {
            pos[k++] = i;
        }
    }

    for (int i = 0; i < queriesSize; i++) {
        queries[i] = queries[i] > k ? -1 : pos[queries[i] - 1];
    }

    free(pos);
    *returnSize = queriesSize;
    return queries;
}
```

```go [sol-Go]
func occurrencesOfElement(nums, queries []int, x int) []int {
	pos := []int{}
	for i, num := range nums {
		if num == x {
			pos = append(pos, i)
		}
	}
	for i, q := range queries {
		if q > len(pos) {
			queries[i] = -1
		} else {
			queries[i] = pos[q-1]
		}
	}
	return queries
}
```

```js [sol-JavaScript]
var occurrencesOfElement = function(nums, queries, x) {
    const pos = [];
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] === x) {
            pos.push(i);
        }
    }
    return queries.map(q => q > pos.length ? -1 : pos[q - 1]);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn occurrences_of_element(nums: Vec<i32>, queries: Vec<i32>, x: i32) -> Vec<i32> {
        let pos = nums.iter()
            .enumerate()
            .filter(|(_, &num)| num == x)
            .map(|(i, _)| i)
            .collect::<Vec<_>>();
        queries.iter()
            .map(|&q| if q as usize > pos.len() { -1 } else { pos[q as usize - 1] as i32 })
            .collect()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

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

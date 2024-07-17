把数组转成哈希集合，就可以 $\mathcal{O}(1)$ 判断元素是否在数组中了。

1. 把 $\textit{nums}_1$ 中的元素加入哈希集合 $\textit{set}_1$ 中。
2. 把 $\textit{nums}_2$ 中的元素加入哈希集合 $\textit{set}_2$ 中。
3. 遍历 $\textit{nums}_1$，统计在 $\textit{set}_2$ 中的元素个数，即为 $\textit{answer}[0]$。
4. 遍历 $\textit{nums}_2$，统计在 $\textit{set}_1$ 中的元素个数，即为 $\textit{answer}[1]$。

```py [sol-Python3]
class Solution:
    def findIntersectionValues(self, nums1: List[int], nums2: List[int]) -> List[int]:
        set1 = set(nums1)
        set2 = set(nums2)
        return [sum(x in set2 for x in nums1),
                sum(x in set1 for x in nums2)]
```

```java [sol-Java]
class Solution {
    public int[] findIntersectionValues(int[] nums1, int[] nums2) {
        HashSet<Integer> set1 = new HashSet<>();
        for (int x : nums1) {
            set1.add(x);
        }
        HashSet<Integer> set2 = new HashSet<>();
        for (int x : nums2) {
            set2.add(x);
        }

        int[] ans = new int[2];
        for (int x : nums1) {
            if (set2.contains(x)) {
                ans[0]++;
            }
        }
        for (int x : nums2) {
            if (set1.contains(x)) {
                ans[1]++;
            }
        }
        return ans;
    }
}
```

```java [sol-Java Stream]
class Solution {
    public int[] findIntersectionValues(int[] nums1, int[] nums2) {
        Set<Integer> set1 = Arrays.stream(nums1).boxed().collect(Collectors.toSet());
        Set<Integer> set2 = Arrays.stream(nums2).boxed().collect(Collectors.toSet());
        int cnt1 = (int) Arrays.stream(nums1).filter(set2::contains).count();
        int cnt2 = (int) Arrays.stream(nums2).filter(set1::contains).count();
        return new int[]{cnt1, cnt2};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> findIntersectionValues(vector<int>& nums1, vector<int>& nums2) {
        unordered_set<int> set1(nums1.begin(), nums1.end());
        unordered_set<int> set2(nums2.begin(), nums2.end());
        vector<int> ans(2);
        for (int x : nums1) ans[0] += set2.count(x);
        for (int x : nums2) ans[1] += set1.count(x);
        return ans;
    }
};
```

```go [sol-Go]
func findIntersectionValues(nums1, nums2 []int) []int {
	set1 := map[int]int{}
	for _, x := range nums1 {
		set1[x] = 1
	}
	set2 := map[int]int{}
	for _, x := range nums2 {
		set2[x] = 1
	}
	
	ans := [2]int{}
	for _, x := range nums1 {
		ans[0] += set2[x]
	}
	for _, x := range nums2 {
		ans[1] += set1[x]
	}
	return ans[:]
}
```

```js [sol-JS]
var findIntersectionValues = function(nums1, nums2) {
    const set1 = new Set(nums1);
    const set2 = new Set(nums2);
    const cnt1 = nums1.reduce((cnt, x) => cnt + (set2.has(x) ? 1 : 0), 0);
    const cnt2 = nums2.reduce((cnt, x) => cnt + (set1.has(x) ? 1 : 0), 0);
    return [cnt1, cnt2];
};
```

```rust [sol-Rust]
use std::collections::HashSet;

impl Solution {
    pub fn find_intersection_values(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        let set1 = nums1.iter().cloned().collect::<HashSet<_>>();
        let set2 = nums2.iter().cloned().collect::<HashSet<_>>();
        let cnt1 = nums1.iter().filter(|&x| set2.contains(x)).count() as i32;
        let cnt2 = nums2.iter().filter(|&x| set1.contains(x)).count() as i32;
        vec![cnt1, cnt2]
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$m$ 为 $\textit{nums}_2$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

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

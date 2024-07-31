## 提示 1

最终所有元素一定变成了一个在 $\textit{nums}$ 中的数。

枚举这个数。

## 提示 2

考虑把数字 $x$「扩散」到其它位置，那么每一秒 $x$ 都可以向左右扩散一位。

多个相同数字 $x$ 同时扩散，那么扩散完整个数组的耗时，就取决于相距**最远**的两个相邻的 $x$。

假设这两个 $x$ 的下标分别为 $i$ 和 $j$，且 $i<j$，那么耗时为：

$$
\left\lfloor\dfrac{j-i}{2}\right\rfloor
$$

枚举不同的 $x$，计算相应的耗时，更新答案的最小值。

## 提示 3

统计所有相同数字的下标，记到一个哈希表 $\textit{pos}$ 中。

设 $\textit{pos}[x]$ 列表第一个下标是 $p$，最后一个下标是 $q$。本题数组可以视作是**环形**的，所以 $p$ 和 $q$ 也是相邻的，耗时为 $\left\lfloor\dfrac{n-(q-p)}{2}\right\rfloor$。

也可以在 $\textit{pos}[x]$ 列表末尾添加一个 $p+n$，就可以转换成非环形数组处理了。

请看 [视频讲解](https://www.bilibili.com/video/BV1bV4y1e72v/) 第三题，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def minimumSeconds(self, nums: List[int]) -> int:
        pos = defaultdict(list)
        for i, x in enumerate(nums):
            pos[x].append(i)

        ans = n = len(nums)
        for a in pos.values():
            a.append(a[0] + n)
            mx = max(j - i for i, j in pairwise(a))
            ans = min(ans, mx)
        return ans // 2  # 最后再除 2
```

```java [sol-Java]
class Solution {
    public int minimumSeconds(List<Integer> nums) {
        int n = nums.size();
        Map<Integer, List<Integer>> pos = new HashMap<>();
        for (int i = 0; i < n; i++) {
            pos.computeIfAbsent(nums.get(i), k -> new ArrayList<>()).add(i);
        }

        int ans = n;
        for (List<Integer> a : pos.values()) {
            int mx = n - a.get(a.size() - 1) + a.get(0);
            for (int i = 1; i < a.size(); i++) {
                mx = Math.max(mx, a.get(i) - a.get(i - 1));
            }
            ans = Math.min(ans, mx);
        }
        return ans / 2; // 最后再除 2
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumSeconds(vector<int> &nums) {
        int n = nums.size();
        unordered_map<int, vector<int>> pos;
        for (int i = 0; i < n; i++) {
            pos[nums[i]].push_back(i);
        }

        int ans = n;
        for (auto &[_, a] : pos) {
            int mx = n - a.back() + a[0];
            for (int i = 1; i < a.size(); ++i) {
                mx = max(mx, a[i] - a[i - 1]);
            }
            ans = min(ans, mx);
        }
        return ans / 2; // 最后再除 2
    }
};
```

```go [sol-Go]
func minimumSeconds(nums []int) int {
	pos := map[int][]int{}
	for i, x := range nums {
		pos[x] = append(pos[x], i)
	}

	n := len(nums)
	ans := n
	for _, a := range pos {
		mx := n - a[len(a)-1] + a[0]
		for i := 1; i < len(a); i++ {
			mx = max(mx, a[i]-a[i-1])
		}
		ans = min(ans, mx)
	}
	return ans / 2 // 最后再除 2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
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

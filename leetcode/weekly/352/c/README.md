**前置知识**：[滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

在遍历数组的同时，维护窗口内的数字。

由于绝对差至多为 $2$，所以用平衡树或者哈希表维护都行，反正至多维护 $3$ 个数，添加删除可以视作是 $\mathcal{O}(1)$ 的。（如果用哈希表，还需记录数字的出现次数。）

如果窗口内的最大值与最小值的差大于 $2$，就不断移动左端点 $\textit{left}$，减少窗口内的数字。

最后

$$
[\textit{left},\textit{right}],[\textit{left}+1,\textit{right}],\cdots,[\textit{right},\textit{right}]
$$

这一共 $\textit{right}-\textit{left}+1$ 个子数组都是合法的，加入答案。

```py [sol-Python3]
class Solution:
    def continuousSubarrays(self, nums: List[int]) -> int:
        ans = left = 0
        cnt = Counter()
        for right, x in enumerate(nums):
            cnt[x] += 1
            while max(cnt) - min(cnt) > 2:
                y = nums[left]
                cnt[y] -= 1
                if cnt[y] == 0:
                    del cnt[y]
                left += 1
            ans += right - left + 1
        return ans
```

```java [sol-Java]
class Solution {
    public long continuousSubarrays(int[] nums) {
        long ans = 0;
        var t = new TreeMap<Integer, Integer>();
        int left = 0;
        for (int right = 0; right < nums.length; right++) {
            t.merge(nums[right], 1, Integer::sum);
            while (t.lastKey() - t.firstKey() > 2) {
                int y = nums[left++];
                if (t.get(y) == 1) t.remove(y);
                else t.merge(y, -1, Integer::sum);
            }
            ans += right - left + 1;
        }
        return ans;
    }
}
```

```cpp [sol-C++ map]
class Solution {
public:
    long long continuousSubarrays(vector<int>& nums) {
        long long ans = 0;
        map<int, int> cnt;
        int left = 0, n = nums.size();
        for (int right = 0; right < n; right++) {
            cnt[nums[right]]++;
            while (cnt.rbegin()->first - cnt.begin()->first > 2) {
                int y = nums[left++];
                if (--cnt[y] == 0) {
                    cnt.erase(y);
                }
            }
            ans += right - left + 1;
        }
        return ans;
    }
};
```

```cpp [sol-C++ multiset]
class Solution {
public:
    long long continuousSubarrays(vector<int>& nums) {
        long long ans = 0;
        multiset<int> s;
        int left = 0, n = nums.size();
        for (int right = 0; right < n; right++) {
            s.insert(nums[right]);
            while (*s.rbegin() - *s.begin() > 2) {
                s.erase(s.find(nums[left++]));
            }
            ans += right - left + 1;
        }
        return ans;
    }
};
```

```go [sol-Go]
func continuousSubarrays(a []int) (ans int64) {
	cnt := map[int]int{}
	left := 0
	for right, x := range a {
		cnt[x]++
		for {
			mx, mn := x, x
			for k := range cnt {
				mx = max(mx, k)
				mn = min(mn, k)
			}
			if mx-mn <= 2 {
				break
			}
			y := a[left]
			if cnt[y]--; cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		ans += int64(right - left + 1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。注意至多维护 $3$ 个数，仅用到常量额外空间。

## 相似题目

- [1438. 绝对差不超过限制的最长连续子数组](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/)

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

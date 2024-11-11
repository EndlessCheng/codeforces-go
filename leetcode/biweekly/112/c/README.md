看到「长度固定的子数组」就要想到滑动窗口！

维护窗口内的元素出现次数 $\textit{cnt}$，以及元素和 $\textit{sum}$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1um4y1M7Rv/) 第三题。

```py [sol-Python3]
class Solution:
    def maxSum(self, nums: List[int], m: int, k: int) -> int:
        ans = 0
        s = sum(nums[:k - 1])  # 先统计 k-1 个数
        cnt = Counter(nums[:k - 1])
        for out, in_ in zip(nums, nums[k - 1:]):
            s += in_  # 再添加一个数就是 k 个数了
            cnt[in_] += 1
            if len(cnt) >= m:
                ans = max(ans, s)

            s -= out  # 下一个子数组不包含 out，移出窗口
            cnt[out] -= 1
            if cnt[out] == 0:
                del cnt[out]
        return ans
```

```java [sol-Java]
class Solution {
    public long maxSum(List<Integer> nums, int m, int k) {
        Integer[] a = nums.toArray(Integer[]::new);
        long ans = 0;
        long sum = 0;
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int i = 0; i < k - 1; i++) { // 先统计 k-1 个数
            sum += a[i];
            cnt.merge(a[i], 1, Integer::sum); // cnt[a[i]]++
        }
        for (int i = k - 1; i < nums.size(); i++) {
            sum += a[i]; // 再添加一个数就是 k 个数了
            cnt.merge(a[i], 1, Integer::sum); // cnt[a[i]]++
            if (cnt.size() >= m) {
                ans = Math.max(ans, sum);
            }

            int out = a[i - k + 1];
            sum -= out; // 下一个子数组不包含 out，移出窗口
            if (cnt.merge(out, -1, Integer::sum) == 0) { // --cnt[out] == 0
                cnt.remove(out);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long maxSum(vector<int> &nums, int m, int k) {
        long long ans = 0, sum = 0;
        unordered_map<int, int> cnt;
        for (int i = 0; i < k - 1; i++) { // 先统计 k-1 个数
            sum += nums[i];
            cnt[nums[i]]++;
        }

        for (int i = k - 1; i < nums.size(); i++) {
            sum += nums[i]; // 再添加一个数就是 k 个数了
            cnt[nums[i]]++;
            if (cnt.size() >= m) {
                ans = max(ans, sum);
            }

            int out = nums[i - k + 1];
            sum -= out; // 下一个子数组不包含 out，移出窗口
            if (--cnt[out] == 0) {
                cnt.erase(out);
            }
        }

        return ans;
    }
};
```

```go [sol-Go]
func maxSum(nums []int, m, k int) (ans int64) {
	sum := int64(0)
	cnt := map[int]int{}
	for _, x := range nums[:k-1] { // 先统计 k-1 个数
		sum += int64(x)
		cnt[x]++
	}
	for i, in := range nums[k-1:] {
		sum += int64(in) // 再添加一个数就是 k 个数了
		cnt[in]++
		if len(cnt) >= m && sum > ans {
			ans = sum
		}

		out := nums[i]
		sum -= int64(out) // 下一个子数组不包含 out，移出窗口
		cnt[out]--
		if cnt[out] == 0 {
			delete(cnt, out)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。哈希表的大小不会超过窗口长度，即 $k$。

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

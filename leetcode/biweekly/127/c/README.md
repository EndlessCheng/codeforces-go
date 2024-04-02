请看 [视频讲解](https://www.bilibili.com/video/BV19t421g7Pd/) 第三题，我结合例子给大家讲讲这个算法。

或者看我之前写的一篇教程：[题解方法二](https://leetcode.cn/problems/smallest-subarrays-with-maximum-bitwise-or/solution/by-endlesscheng-zai1/)。

本题只需在内层循环中添加一个 `if` 即可：如果子数组 OR 值大于等于 $k$，更新子数组长度的最小值。

```py [sol-Python3 字典]
class Solution:
    def minimumSubarrayLength(self, nums: List[int], k: int) -> int:
        ans = inf
        d = dict()  # key 是右端点为 i 的子数组 OR, value 是该子数组左端点的最大值
        for i, x in enumerate(nums):
            # 注意 key 是按照插入顺序排的，所以在相同 OR 时，会自动取到更大的 left 作为 value
            d = {or_ | x: left for or_, left in d.items()}
            d[x] = i  # 只包含 x 的子数组
            for or_, left in d.items():
                if or_ >= k:
                    ans = min(ans, i - left + 1)
        return ans if ans < inf else -1
```

```py [sol-Python3 列表]
class Solution:
    def minimumSubarrayLength(self, nums: List[int], k: int) -> int:
        ans = inf
        ors = []  # 保存 (右端点为 i 的子数组 OR, 该子数组左端点的最大值)
        for i, x in enumerate(nums):
            ors.append([0, i])
            j = 0
            for p in ors:
                p[0] |= x
                if p[0] >= k:
                    ans = min(ans, i - p[1] + 1)
                if ors[j][0] == p[0]:
                    ors[j][1] = p[1]  # 原地去重：合并相同值，左端点取靠右的
                else:
                    j += 1
                    ors[j] = p
            del ors[j + 1:]  # 去重：移除多余元素
        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public int minimumSubarrayLength(int[] nums, int k) {
        int ans = Integer.MAX_VALUE;
        List<int[]> ors = new ArrayList<>(); // 保存 (右端点为 i 的子数组 OR, 该子数组左端点的最大值)
        for (int i = 0; i < nums.length; i++) {
            ors.add(new int[]{0, i});
            int j = 0;
            for (int[] or : ors) {
                or[0] |= nums[i];
                if (or[0] >= k) {
                    ans = Math.min(ans, i - or[1] + 1);
                }
                if (ors.get(j)[0] == or[0]) {
                    ors.get(j)[1] = or[1]; // 原地去重：合并相同值，左端点取靠右的
                } else {
                    ors.set(++j, or);
                }
            }
            ors.subList(j + 1, ors.size()).clear(); // 去重：移除多余元素
        }
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

```java [sol-Java 数组]
class Solution {
    public int minimumSubarrayLength(int[] nums, int k) {
        int ans = Integer.MAX_VALUE;
        int[][] ors = new int[32][2]; // 保存 (右端点为 i 的子数组 OR, 该子数组左端点的最大值)
        int m = 0;
        for (int i = 0; i < nums.length; i++) {
            ors[m][0] = 0;
            ors[m++][1] = i;
            int j = 0;
            for (int idx = 0; idx < m; idx++) {
                ors[idx][0] |= nums[i];
                if (ors[idx][0] >= k) {
                    ans = Math.min(ans, i - ors[idx][1] + 1);
                }
                if (ors[j][0] != ors[idx][0]) {
                    ors[++j][0] = ors[idx][0];
                }
                ors[j][1] = ors[idx][1];
            }
            m = j + 1; // 去重：移除多余元素
        }
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumSubarrayLength(vector<int> &nums, int k) {
        int ans = INT_MAX;
        vector<pair<int, int>> ors; // 保存 (右端点为 i 的子数组 OR, 该子数组左端点的最大值)
        for (int i = 0; i < nums.size(); i++) {
            ors.emplace_back(0, i);
            int j = 0;
            for (auto &p : ors) {
                auto &[or_, left] = p;
                or_ |= nums[i];
                if (or_ >= k) {
                    ans = min(ans, i - left + 1);
                }
                if (ors[j].first == or_) {
                    ors[j].second = left; // 原地去重：合并相同值，左端点取靠右的
                } else {
                    ors[++j] = p;
                }
            }
            ors.resize(j + 1); // 去重：移除多余元素
        }
        return ans == INT_MAX ? -1 : ans;
    }
};
```

```go [sol-Go]
func minimumSubarrayLength(nums []int, k int) int {
	ans := math.MaxInt
	type pair struct{ or, left int }
	ors := []pair{} // 保存 (右端点为 i 的子数组 OR, 该子数组左端点的最大值)
	for i, x := range nums {
		ors = append(ors, pair{0, i})
		j := 0
		for idx := range ors {
			p := &ors[idx]
			p.or |= x
			if p.or >= k {
				ans = min(ans, i-p.left+1)
			}
			if ors[j].or == p.or {
				ors[j].left = p.left // 原地去重：合并相同值，左端点取靠右的
			} else {
				j++
				ors[j] = *p
			}
		}
		ors = ors[:j+1] // 去重：移除多余元素
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```

```go [sol-Go 写法二]
func minimumSubarrayLength(nums []int, k int) int {
	ans := math.MaxInt
	type pair struct{ or, left int }
	ors := []pair{} // 保存 (右端点为 i 的子数组 OR, 该子数组左端点的最大值)
	for i, x := range nums {
		ors = append(ors, pair{0, i})
		ors[0].or |= x
		j := 0
		for _, p := range ors[1:] {
			p.or |= x
			if ors[j].or == p.or {
				ors[j].left = p.left // 原地去重：合并相同值，左端点取靠右的
			} else {
				j++
				ors[j] = p
			}
		}
		ors = ors[:j+1] // 去重：移除多余元素
		for len(ors) > 0 && ors[0].or >= k {
			ans = min(ans, i-ors[0].left+1)
			ors = ors[1:]
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(\log U)$。

#### 思考题

如果改成**异或**，要怎么做？

#### 相似题目

见 [题解](https://leetcode.cn/problems/smallest-subarrays-with-maximum-bitwise-or/solution/by-endlesscheng-zai1/)。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

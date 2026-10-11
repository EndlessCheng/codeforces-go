注意本题只要求 $i\ne j$，而不是通常的 $i<j$。

所以 $\textit{nums}[i] > \textit{nums}[j]$ 可以视作 $\textit{nums}[i] \ne \textit{nums}[j]$。如果最后发现 $\textit{nums}[i] < \textit{nums}[j]$，交换 $i$ 和 $j$ 即可。

在遍历 $\textit{nums}$ 的过程中，维护左边遍历过的每个数的出现位置。遍历到 $x = \textit{nums}[i]$ 时。如果 $x\ne \textit{target}-x$，且乘积 $x(\textit{target}-x)$ 比之前的乘积更大，那么更新答案。

> **注**：由于 $x(\textit{target}-x)$ 的最大值在 $\dfrac{\textit{target}}{2}$ 处取到，我们也可以比较谁离 $\dfrac{\textit{target}}{2}$ 更近。

[本题视频讲解](https://www.bilibili.com/video/BV1rHpj6kE8u/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def maxProductPair(self, nums: list[int], target: int) -> list[int]:
        pos = {}
        best_i = best_j = -1
        for i, x in enumerate(nums):
            if x * 2 == target:  # x = target-x
                continue
            j = pos.get(target - x, -1)
            if j >= 0 and (best_i < 0 or x * (target - x) > nums[best_i] * (target - nums[best_i])):  # 也可以比谁离 target/2 更近
                best_i, best_j = i, j
            pos[x] = i

        if best_i >= 0 and nums[best_i] < nums[best_j]:
            best_i, best_j = best_j, best_i  # 保证 nums[best_i] > nums[best_j]
        return [best_i, best_j]
```

```java [sol-Java]
class Solution {
    public int[] maxProductPair(int[] nums, int target) {
        Map<Integer, Integer> pos = new HashMap<>();
        int bestI = -1;
        int bestJ = -1;
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (x * 2 == target) { // x = target-x
                continue;
            }
            Integer j = pos.get(target - x);
            if (j != null && (bestI < 0 || x * (target - x) > nums[bestI] * (target - nums[bestI]))) { // 也可以比谁离 target/2 更近
                bestI = i;
                bestJ = j;
            }
            pos.put(x, i);
        }

        if (bestI >= 0 && nums[bestI] < nums[bestJ]) {
            // 保证 nums[bestI] > nums[bestJ]
            int tmp = bestI;
            bestI = bestJ;
            bestJ = tmp;
        }
        return new int[]{bestI, bestJ};
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> maxProductPair(vector<int>& nums, int target) {
        unordered_map<int, int> pos;
        int best_i = -1, best_j = -1;
        for (int i = 0; i < nums.size(); i++) {
            int x = nums[i];
            if (x * 2 == target) { // x = target-x
                continue;
            }
            auto it = pos.find(target - x);
            if (it != pos.end() && (best_i < 0 || x * (target - x) > nums[best_i] * (target - nums[best_i]))) { // 也可以比谁离 target/2 更近
                best_i = i;
                best_j = it->second;
            }
            pos[x] = i;
        }

        if (best_i >= 0 && nums[best_i] < nums[best_j]) {
            swap(best_i, best_j); // 保证 nums[best_i] > nums[best_j]
        }
        return {best_i, best_j};
    }
};
```

```go [sol-Go]
func maxProductPair(nums []int, target int) []int {
	pos := map[int]int{}
	bestI, bestJ := -1, -1
	for i, x := range nums {
		if x*2 == target { // x = target-x
			continue
		}
		j, ok := pos[target-x]
		if ok && (bestI < 0 || x*(target-x) > nums[bestI]*(target-nums[bestI])) { // 也可以比谁离 target/2 更近
			bestI, bestJ = i, j
		}
		pos[x] = i
	}

	if bestI >= 0 && nums[bestI] < nums[bestJ] {
		bestI, bestJ = bestJ, bestI // 保证 nums[bestI] > nums[bestJ]
	}
	return []int{bestI, bestJ}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§0.1 枚举右，维护左**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

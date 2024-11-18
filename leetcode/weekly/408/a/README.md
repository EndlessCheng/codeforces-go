设 $s_1$ 为 $\textit{nums}$ 中的所有个位数之和，$s_2$ 为 $\textit{nums}$ 中的所有两位数之和。注意题目保证只有个位数和两位数。

小红若要获胜，必须满足 $s_1 > s_2$ 或者 $s_2 > s_1$，即

$$
s_1 \ne s_2
$$

代码实现时，可以令 $s = s_1 - s_2$，即累加 $\textit{nums}$ 的所有元素，把其中的两位数变成相反数累加。这样最后只需判断 $s\ne 0$ 即可。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Mi421a7cZ/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def canAliceWin(self, nums: List[int]) -> bool:
        return sum(x if x < 10 else -x for x in nums) != 0
```

```java [sol-Java]
class Solution {
    public boolean canAliceWin(int[] nums) {
        int s = 0;
        for (int x : nums) {
            s += x < 10 ? x : -x;
        }
        return s != 0;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    bool canAliceWin(vector<int>& nums) {
        int s = 0;
        for (int x : nums) {
            s += x < 10 ? x : -x;
        }
        return s != 0;
    }
};
```

```c [sol-C]
bool canAliceWin(int* nums, int numsSize) {
    int s = 0;
    for (int i = 0; i < numsSize; i++) {
        s += nums[i] < 10 ? nums[i] : -nums[i];
    }
    return s != 0;
}
```

```go [sol-Go]
func canAliceWin(nums []int) bool {
	s := 0
	for _, x := range nums {
		if x < 10 {
			s += x
		} else {
			s -= x
		}
	}
	return s != 0
}
```

```js [sol-JavaScript]
var canAliceWin = function(nums) {
    let s = 0;
    for (const x of nums) {
        s += x < 10 ? x : -x;
    }
    return s !== 0;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn can_alice_win(nums: Vec<i32>) -> bool {
        nums.iter().map(|&x| if x < 10 { x } else { -x }).sum::<i32>() != 0
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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

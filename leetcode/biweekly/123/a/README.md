把 $\textit{nums}$ 从小到大排序，可以简化判断逻辑。

设排序后 $\textit{nums}=[a,b,c]$，那么有 $1\le a\le b\le c$。

- 先判**是否合法**，即三角形任意两边之和必须大于第三边。由于排序后 $a+c > b$ 和 $b+c>a$ 自动成立（注意数组元素都是正数），所以只需判断 $a+b > c$ 是否成立。如果 $a+b\le c$，那么无法构成三角形。
- 然后判**等边**：只需判断 $a=c$。注意已经排序了，如果 $a=c$，那么必然有 $a=b=c$。
- 然后判**等腰**：判断 $a=b$ 或者 $b=c$。
- **其他情况**，三条边长度一定不相等，无需判断。

```py [sol-Python3]
class Solution:
    def triangleType(self, nums: List[int]) -> str:
        nums.sort()
        a, b, c = nums
        if a + b <= c:
            return "none"
        if a == c:
            return "equilateral"
        if a == b or b == c:
            return "isosceles"
        return "scalene"
```

```java [sol-Java]
class Solution {
    public String triangleType(int[] nums) {
        Arrays.sort(nums);
        int a = nums[0];
        int b = nums[1];
        int c = nums[2];
        if (a + b <= c) {
            return "none";
        }
        if (a == c) {
            return "equilateral";
        }
        if (a == b || b == c) {
            return "isosceles";
        }
        return "scalene";
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    string triangleType(vector<int>& nums) {
        ranges::sort(nums);
        int a = nums[0], b = nums[1], c = nums[2];
        if (a + b <= c) {
            return "none";
        }
        if (a == c) {
            return "equilateral";
        }
        if (a == b || b == c) {
            return "isosceles";
        }
        return "scalene";
    }
};
```

```c [sol-C]
int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

char* triangleType(int* nums, int numsSize) {
    qsort(nums, numsSize, sizeof(int), cmp);
    int a = nums[0], b = nums[1], c = nums[2];
    if (a + b <= c) {
        return "none";
    }
    if (a == c) {
        return "equilateral";
    }
    if (a == b || b == c) {
        return "isosceles";
    }
    return "scalene";
}
```

```go [sol-Go]
func triangleType(nums []int) string {
	slices.Sort(nums)
	a, b, c := nums[0], nums[1], nums[2]
	if a+b <= c {
		return "none"
	}
	if a == c {
		return "equilateral"
	}
	if a == b || b == c {
		return "isosceles"
	}
	return "scalene"
}
```

```js [sol-JavaScript]
var triangleType = function(nums) {
    nums.sort((a, b) => a - b);
    const [a, b, c] = nums;
    if (a + b <= c) {
        return "none";
    }
    if (a === c) {
        return "equilateral";
    }
    if (a === b || b === c) {
        return "isosceles";
    }
    return "scalene";
};
```

```rust [sol-Rust]
impl Solution {
    pub fn triangle_type(mut nums: Vec<i32>) -> String {
        nums.sort_unstable();
        let (a, b, c) = (nums[0], nums[1], nums[2]);
        if a + b <= c {
            return "none".to_string();
        }
        if a == c {
            return "equilateral".to_string();
        }
        if a == b || b == c {
            return "isosceles".to_string();
        }
        "scalene".to_string()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 附：哈希表做法

在三边长能构成三角形的情况下，用哈希表计算 $\textit{nums}$ 中有 $c$ 个不同元素，然后判断：

- $c=1$：答案是 $\texttt{equilateral}$。
- $c=2$：答案是 $\texttt{isosceles}$。
- $c=3$：答案是 $\texttt{scalene}$。

```py
class Solution:
    def triangleType(self, nums: List[int]) -> str:
        nums.sort()
        if nums[0] + nums[1] <= nums[2]:
            return "none"
        return ("equilateral", "isosceles", "scalene")[len(set(nums)) - 1]
```

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

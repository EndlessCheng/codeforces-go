从最低位开始枚举，取三个数在该数位的最小值加入答案。然后把三个数都除以 $10$，继续枚举数位。

循环直到其中一个数等于 $0$ 为止，因为后面的数位，最小值都是 $0$。

为方便写代码，下面把变量名改成 $x,y,z$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ajHYeoEG5/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def generateKey(self, x: int, y: int, z: int) -> int:
        ans = 0
        pow10 = 1
        while x and y and z:
            ans += min(x % 10, y % 10, z % 10) * pow10
            x //= 10
            y //= 10
            z //= 10
            pow10 *= 10
        return ans
```

```java [sol-Java]
class Solution {
    int generateKey(int x, int y, int z) {
        int ans = 0;
        for (int pow10 = 1; x > 0 && y > 0 && z > 0; pow10 *= 10) {
            ans += Math.min(Math.min(x % 10, y % 10), z % 10) * pow10;
            x /= 10;
            y /= 10;
            z /= 10;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int generateKey(int x, int y, int z) {
        int ans = 0;
        for (int pow10 = 1; x && y && z; pow10 *= 10) {
            ans += min({x % 10, y % 10, z % 10}) * pow10;
            x /= 10;
            y /= 10;
            z /= 10;
        }
        return ans;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

int generateKey(int x, int y, int z) {
    int ans = 0;
    for (int pow10 = 1; x && y && z; pow10 *= 10) {
        ans += MIN(MIN(x % 10, y % 10), z % 10) * pow10;
        x /= 10;
        y /= 10;
        z /= 10;
    }
    return ans;
}
```

```go [sol-Go]
func generateKey(x, y, z int) (ans int) {
	for pow10 := 1; x > 0 && y > 0 && z > 0; pow10 *= 10 {
		ans += min(x%10, y%10, z%10) * pow10
		x /= 10
		y /= 10
		z /= 10
	}
	return
}
```

```js [sol-JavaScript]
var generateKey = function(x, y, z) {
    let ans = 0;
    for (let pow10 = 1; x && y && z; pow10 *= 10) {
        ans += Math.min(x % 10, y % 10, z % 10) * pow10;
        x = Math.floor(x / 10);
        y = Math.floor(y / 10);
        z = Math.floor(z / 10);
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn generate_key(mut x: i32, mut y: i32, mut z: i32) -> i32 {
        let mut ans = 0;
        let mut pow10 = 1;
        while x > 0 && y > 0 && z > 0 {
            ans += (x % 10).min(y % 10).min(z % 10) * pow10;
            x /= 10;
            y /= 10;
            z /= 10;
            pow10 *= 10;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log \min(x,y,z))$。
- 空间复杂度：$\mathcal{O}(1)$。

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

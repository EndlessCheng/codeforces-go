思考：$4325$ 怎么分？

如果分成 $432$ 和 $5$ 这两组，不如分成 $32$ 和 $45$，因为 $4$ 在 $432$ 中是 $400$，在 $45$ 中是 $40$。这启发我们要尽量**均匀分**。如果有奇数个数，多出的那个数放在哪一组都可以。

如果均匀分成 $32$ 和 $45$ 这两组，那么 $32$ 这组调整一下顺序，$23$ 比 $32$ 更好。

进一步地，均匀分成 $24$ 和 $35$ 更好，也就是**把小的数字排在高位，大的数字排在低位**。

设 $s$ 是 $\textit{num}$ 的字符串形式，这等价于把 $s$ 排序后，按照奇偶下标分组。

[视频讲解](https://www.bilibili.com/video/BV1dY4y1C77x/)

```py [sol-Python3]
class Solution:
    def splitNum(self, num: int) -> int:
        s = ''.join(sorted(str(num)))
        return int(s[::2]) + int(s[1::2])
```

```java [sol-Java]
class Solution {
    public int splitNum(int num) {
        char[] s = Integer.toString(num).toCharArray();
        Arrays.sort(s);
        int[] a = new int[2];
        for (int i = 0; i < s.length; i++) {
            a[i % 2] = a[i % 2] * 10 + s[i] - '0'; // 按照奇偶下标分组
        }
        return a[0] + a[1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int splitNum(int num) {
        string s = to_string(num);
        ranges::sort(s);
        int a[2]{};
        for (int i = 0; i < s.length(); i++) {
            a[i % 2] = a[i % 2] * 10 + s[i] - '0'; // 按照奇偶下标分组
        }
        return a[0] + a[1];
    }
};
```

```go [sol-Go]
func splitNum(num int) int {
	s := []byte(strconv.Itoa(num))
	slices.Sort(s)
	a := [2]int{}
	for i, c := range s {
		a[i%2] = a[i%2]*10 + int(c-'0') // 按照奇偶下标分组
	}
	return a[0] + a[1]
}
```

```js [sol-JavaScript]
var splitNum = function(num) {
    const s = num.toString().split('').sort();
    const a = [0, 0];
    for (let i = 0; i < s.length; i++) {
        a[i % 2] = a[i % 2] * 10 + parseInt(s[i]);
    }
    return a[0] + a[1];
};
```

```rust [sol-Rust]
impl Solution {
    pub fn split_num(num: i32) -> i32 {
        let mut s = num.to_string();
        let mut s = unsafe { s.as_bytes_mut() };
        s.sort_unstable();
        let mut a = [0, 0];
        for (i, &c) in s.iter().enumerate() {
            a[i % 2] = a[i % 2] * 10 + (c - b'0') as i32;
        }
        a[0] + a[1]
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m)$，其中 $m$ 为 $\textit{num}$ 转成字符串后的长度。
- 空间复杂度：$\mathcal{O}(m)$。

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

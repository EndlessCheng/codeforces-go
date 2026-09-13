操作不改变 $\textit{nums}[i]$ 的奇偶性。

问题相当于：

- 找到与 $\textit{nums}[i]$ 同奇偶的最近的回文数。

我们可以预处理范围内的所有回文数（按奇偶分成两组），然后在回文数中 [二分查找](https://www.bilibili.com/video/BV1AP41137w7/) $\ge \textit{nums}[i]$ 的最小的数，以及 $< \textit{nums}[i]$ 的最大的数。

为简化二分逻辑，可以把预处理的范围上界置为 $2\times 10^9+2$，这是大于 $10^9$ 的最小的回文偶数。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
def gen_palindrome() -> Iterator[int]:
    base = 1
    while True:
        # 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
        for i in range(base, base * 10):
            s = str(i)
            x = int(s + s[::-1][1:])
            yield x

        # 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
        for i in range(base, base * 10):
            s = str(i)
            x = int(s + s[::-1])
            yield x

        base *= 10


MX = 2_000_000_002
palindromes = [[0], [0]]  # 哨兵
for x in gen_palindrome():
    if x > MX:
        break
    palindromes[x % 2].append(x)


class Solution:
    def minOperations(self, nums: list[int]) -> int:
        ans = 0
        for x in nums:
            p = palindromes[x % 2]
            i = bisect_left(p, x)
            ans += min(p[i] - x, x - p[i - 1])
        return ans // 2
```

```java [sol-Java]
class Solution {
    private static final int MX = 2_000_000_002;
    private static final List<Integer>[] palindromes = new ArrayList[2];
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        Arrays.setAll(palindromes, _ -> new ArrayList<>());
        palindromes[0].add(0);
        palindromes[1].add(0); // 哨兵

        // 预处理 [1, MX] 中的回文数
        for (int base = 1; ; base *= 10) {
            // 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
            for (int i = base; i < base * 10; i++) {
                int x = i;
                for (int t = i / 10; t > 0; t /= 10) {
                    x = x * 10 + t % 10;
                }
                if (x > MX) {
                    return;
                }
                palindromes[x % 2].add(x);
            }

            // 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
            for (int i = base; i < base * 10; i++) {
                int x = i;
                for (int t = i; t > 0; t /= 10) {
                    x = x * 10 + t % 10;
                }
                if (x > MX) {
                    return;
                }
                palindromes[x % 2].add(x);
            }
        }
    }

    public long minOperations(int[] nums) {
        long ans = 0;
        for (int x : nums) {
            List<Integer> p = palindromes[x % 2];
            int i = lowerBound(p, x);
            ans += Math.min(p.get(i) - x, x - p.get(i - 1));
        }
        return ans / 2;
    }

    // 开区间写法
    private int lowerBound(List<Integer> nums, int target) {
        int left = -1;
        int right = nums.size(); // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[right] >= target
            // nums[left] < target
            int mid = (left + right) >>> 1;
            if (nums.get(mid) >= target) {
                right = mid; // 范围缩小到 (left, mid)
            } else {
                left = mid; // 范围缩小到 (mid, right)
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
const int MX = 2'000'000'002;
vector<int> palindromes[2] = {{0}, {0}};

// 预处理 [1, MX] 中的回文数
auto init = []() {
    for (int base = 1; ; base *= 10) {
        // 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
        for (int i = base; i < base * 10; i++) {
            int x = i;
            for (int t = i / 10; t > 0; t /= 10) {
                x = x * 10 + t % 10;
            }
            if (x > MX) {
                return 0;
            }
            palindromes[x % 2].push_back(x);
        }

        // 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
        for (int i = base; i < base * 10; i++) {
            int x = i;
            for (int t = i; t > 0; t /= 10) {
                x = x * 10 + t % 10;
            }
            if (x > MX) {
                return 0;
            }
            palindromes[x % 2].push_back(x);
        }
    }
}();

class Solution {
public:
    long long minOperations(vector<int>& nums) {
        long long ans = 0;
        for (int x : nums) {
            auto it = ranges::lower_bound(palindromes[x % 2], x);
            ans += min(*it - x, x - *prev(it));
        }
        return ans / 2;
    }
};
```

```go [sol-Go]
const mx = 2_000_000_002
var palindromes = [2][]int{{0}, {0}} // 哨兵

// 预处理 [1, mx] 中的回文数
func init() {
	for base := 1; ; base *= 10 {
		// 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
		for i := base; i < base*10; i++ {
			x := i
			for t := i / 10; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			if x > mx {
				return
			}
			// 按照 x 的奇偶性分组
			palindromes[x%2] = append(palindromes[x%2], x)
		}

		// 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
		for i := base; i < base*10; i++ {
			x := i
			for t := i; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			if x > mx {
				return
			}
			palindromes[x%2] = append(palindromes[x%2], x)
		}
	}
}

func minOperations(nums []int) (ans int64) {
	for _, x := range nums {
		p := palindromes[x%2]
		i := sort.SearchInts(p, x)
		ans += int64(min(p[i]-x, x-p[i-1]))
	}
	return ans / 2
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(\log U)$，其中 $U = \sqrt{10^9}$。
- 空间复杂度：$\mathcal{O}(1)$。

## 值域范围更大的做法

做法类似 [564. 寻找最近的回文数](https://leetcode.cn/problems/find-the-closest-palindrome/)，[我的题解](https://leetcode.cn/problems/find-the-closest-palindrome/solutions/3855597/zhi-xu-kao-lu-5-ge-shu-zi-pythonjavacgo-3td25/)。

## 专题训练

见下面数学题单的「**§7.1 回文数**」。

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

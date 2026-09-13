## 方法一：预处理回文数 + 二分查找

操作不改变 $\textit{nums}[i]$ 的奇偶性。

问题相当于：

- 找到与 $\textit{nums}[i]$ 同奇偶的**最近**的回文数。

我们可以预处理范围内的所有回文数（按奇偶分成两组），然后在回文数中 [二分查找](https://www.bilibili.com/video/BV1AP41137w7/) $\ge \textit{nums}[i]$ 的最小的数，以及 $< \textit{nums}[i]$ 的最大的数。

为了简化边界情况的判断逻辑，可以把预处理的范围上界置为 $2\times 10^9+2$，这是大于 $10^9$ 的最小的回文偶数。

[本题视频讲解](https://www.bilibili.com/video/BV1k7Yv6WE3i/)，欢迎点赞关注~

```py [sol-Python3]
# 模板来自数学题单 https://leetcode.cn/discuss/post/3584388/
def gen_palindrome() -> Iterator[int]:
    base = 1
    while True:
        # 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
        for i in range(base, base * 10):
            s = str(i)
            x = int(s + s[::-1][1:])  # 去掉 i 的最低位，反转，拼在 i 的右边
            yield x

        # 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
        for i in range(base, base * 10):
            s = str(i)
            x = int(s + s[::-1])  # 反转 i，拼在 i 的右边
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
    // 模板来自数学题单 https://leetcode.cn/discuss/post/3584388/
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
                    x = x * 10 + t % 10; // 去掉 i 的最低位，反转，拼在 i 的右边
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
                    x = x * 10 + t % 10; // 反转 i，拼在 i 的右边
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
// 模板来自数学题单 https://leetcode.cn/discuss/post/3584388/
const int MX = 2'000'000'002;
vector<int> palindromes[2] = {{0}, {0}}; // 哨兵

// 预处理 [1, MX] 中的回文数
auto init = []() {
    for (int base = 1; ; base *= 10) {
        // 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
        for (int i = base; i < base * 10; i++) {
            int x = i;
            for (int t = i / 10; t > 0; t /= 10) {
                x = x * 10 + t % 10; // 去掉 i 的最低位，反转，拼在 i 的右边
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
                x = x * 10 + t % 10; // 反转 i，拼在 i 的右边
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
// 模板来自数学题单 https://leetcode.cn/discuss/post/3584388/
const mx = 2_000_000_002
var palindromes = [2][]int{{0}, {0}} // 哨兵

// 预处理 [1, mx] 中的回文数
func init() {
	for base := 1; ; base *= 10 {
		// 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
		for i := base; i < base*10; i++ {
			x := i
			for t := i / 10; t > 0; t /= 10 {
				x = x*10 + t%10 // 去掉 i 的最低位，反转，拼在 i 的右边
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
				x = x*10 + t%10 // 反转 i，拼在 i 的右边
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

- 时间复杂度：$\mathcal{O}(n\log \sqrt{U}) = \mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U = 10^9$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：至多考虑 5 个数

如果 $n\le 10^{18}$，预处理消耗的空间就太大了。我们需要直接找到最近的回文数。

做法类似 [564. 寻找最近的回文数](https://leetcode.cn/problems/find-the-closest-palindrome/)，下面接着 [我的题解](https://leetcode.cn/problems/find-the-closest-palindrome/solutions/3855597/zhi-xu-kao-lu-5-ge-shu-zi-pythonjavacgo-3td25/) 继续讲。

本题需要保证回文数的最高位与 $\textit{nums}[i]$ 的奇偶性相同。

设 $\textit{nums}[i]$ 的十进制长度为 $m$。如果 $\textit{nums}[i]$ 是偶数，额外考虑的回文数改为：

- 十进制长为 $m-1$ 的最大偶回文数 $9\cdot 10^{m-2}-2\ (m\ge 3)$。特别地，如果 $m=2$，则十进制长为 $m-1$ 的最大偶回文数为 $8$。
- 十进制长为 $m+1$ 的最小偶回文数 $2\cdot 10^m+2$。

> 注：代码实现时，无需考虑 $\textit{left}-1$ 生成的回文数的十进制长度小于 $m$ 的情况，按照我们的规则，这种回文数的十进制长度是 $m-2$，远远小于我们需要考虑的数。同理，无需考虑 $\textit{left}+1$ 生成的回文数的十进制长度大于 $m$ 的情况。

```py [sol-Python3]
class Solution:
    def nearestPalindromicDiff(self, num: int) -> int:
        if num <= 9:
            return 0  # num 已经是回文数

        min_d = inf

        def update(pal: int) -> None:
            nonlocal min_d
            min_d = min(min_d, abs(pal - num))

        s = str(num)
        m = len(s)

        if num % 2:
            update(10 ** (m - 1) - 1)  # 十进制长为 m-1 的最大奇回文数 999..999
            update(10 ** m + 1)  # 十进制长为 m+1 的最小奇回文数 100..001
        else:
            if m == 2:
                update(8)
            else:
                update(10 ** (m - 2) * 9 - 2)  # 十进制长为 m-1 的最大偶回文数 899..998
            update(10 ** m * 2 + 2)  # 十进制长为 m+1 的最小偶回文数 200..002

        high_digit = int(s[0])
        if high_digit % 2 == num % 2:
            left = int(s[:(m + 1) // 2])
            # 枚举十进制长为 m 的邻近回文数
            for l in range(left - 1, left + 2):
                # l 最高位的奇偶性必须与 num 的相同
                sl = str(l)
                if int(sl[0]) % 2 == num % 2:
                    update(int(sl + sl[::-1][m % 2:]))
        else:
            # 最高位为 high_digit - 1
            # 例如 num = 354..676，生成回文数 299..992
            if high_digit > 1:
                update((10 ** (m - 1) + 1) * high_digit - 11)

            # 最高位为 high_digit + 1
            # 例如 num = 354..676，生成回文数 400..004
            if high_digit < 9:
                update((10 ** (m - 1) + 1) * (high_digit + 1))

        return min_d

    def minOperations(self, nums: list[int]) -> int:
        return sum(self.nearestPalindromicDiff(x) for x in nums) // 2
```

```java [sol-Java]
class Solution {
    public long minOperations(int[] nums) {
        long ans = 0;
        for (int x : nums) {
            ans += nearestPalindromicDiff(x);
        }
        return ans / 2;
    }

    // 返回离 num 最近的与 num 同奇偶的正回文数与 num 的绝对差
    public long nearestPalindromicDiff(int num) {
        if (num <= 9) {
            return 0; // num 已经是回文数
        }

        minD = Long.MAX_VALUE;

        String s = String.valueOf(num);
        int m = s.length(); // num 的十进制长度

        if (num % 2 > 0) {
            update((long) Math.pow(10, m - 1) - 1, num); // 十进制长为 m-1 的最大奇回文数 999..999
            update((long) Math.pow(10, m) + 1, num); // 十进制长为 m+1 的最小奇回文数 100..001
        } else {
            if (m == 2) {
                update(8, num);
            } else {
                update((long) Math.pow(10, m - 2) * 9 - 2, num); // 十进制长为 m-1 的最大偶回文数 899..998
            }
            update((long) Math.pow(10, m) * 2 + 2, num); // 十进制长为 m+1 的最小偶回文数 200..002
        }

        int highDigit = s.charAt(0) - '0';
        if (highDigit % 2 == num % 2) {
            int left = Integer.parseInt(s.substring(0, (m + 1) / 2));
            // 枚举十进制长为 m 的邻近回文数
            for (int l = left - 1; l <= left + 1; l++) {
                // l 最高位奇偶性必须与 num 相同
                if (String.valueOf(l).charAt(0) % 2 != num % 2) {
                    continue;
                }

                long pal = l;
                for (int x = m % 2 > 0 ? l / 10 : l; x > 0; x /= 10) {
                    pal = pal * 10 + x % 10;
                }
                update(pal, num);
            }
        } else {
            // 最高位为 highDigit - 1
            // 例如 num = 354..676，生成回文数 299..992
            if (highDigit > 1) {
                update(((long) Math.pow(10, m - 1) + 1) * highDigit - 11, num);
            }

            // 最高位为 highDigit + 1
            // 例如 num = 354..676，生成回文数 400..004
            if (highDigit < 9) {
                update(((long) Math.pow(10, m - 1) + 1) * (highDigit + 1), num);
            }
        }

        return minD;
    }

    private long minD;

    private void update(long pal, int num) {
        minD = Math.min(minD, Math.abs(pal - num));
    }
}
```

```cpp [sol-C++]
class Solution {
    long long nearestPalindromicDiff(int num) {
        if (num <= 9) {
            return 0; // num 已经是回文数
        }

        long long min_d = LLONG_MAX;
        auto update = [&](long long pal) -> void {
            min_d = min(min_d, abs(pal - num));
        };

        string s = to_string(num);
        int m = s.size();

        if (num % 2) {
            update((long long) pow(10, m - 1) - 1); // 十进制长为 m-1 的最大奇回文数 999..999
            update((long long) pow(10, m) + 1); // 十进制长为 m+1 的最小奇回文数 100..001
        } else {
            if (m == 2) {
                update(8);
            } else {
                update((long long) pow(10, m - 2) * 9 - 2); // 十进制长为 m-1 的最大偶回文数 899..998
            }
            update((long long) pow(10, m) * 2 + 2); // 十进制长为 m+1 的最小偶回文数 200..002
        }

        int high_digit = s[0] - '0';
        if (high_digit % 2 == num % 2) {
            int left = stoi(s.substr(0, (m + 1) / 2));
            // 枚举十进制长为 m 的邻近回文数
            for (int l = left - 1; l <= left + 1; l++) {
                // l 最高位的奇偶性必须与 num 的相同
                if (to_string(l)[0] % 2 != num % 2) {
                    continue;
                }

                long long pal = l;
                for (int x = m % 2 ? l / 10 : l; x > 0; x /= 10) {
                    pal = pal * 10 + x % 10;
                }
                update(pal);
            }
        } else {
            // 最高位为 high_digit - 1
            // 例如 num = 354..676，生成回文数 299..992
            if (high_digit > 1) {
                update(((long long) pow(10, m - 1) + 1) * high_digit - 11);
            }

            // 最高位为 high_digit + 1
            // 例如 num = 354..676，生成回文数 400..004
            if (high_digit < 9) {
                update(((long long) pow(10, m - 1) + 1) * (high_digit + 1));
            }
        }

        return min_d;
    }

public:
    long long minOperations(vector<int>& nums) {
        long long ans = 0;
        for (int x : nums) {
            ans += nearestPalindromicDiff(x);
        }
        return ans / 2;
    }
};
```

```go [sol-Go]
// 返回离 num 最近的与 num 同奇偶的正回文数与 num 的绝对差
func nearestPalindromicDiff(num int) int {
	if num <= 9 {
		return 0 // num 已经是回文数
	}

	minD := math.MaxInt
	update := func(pal int) {
		minD = min(minD, abs(pal-num))
	}

	s := strconv.Itoa(num)
	m := len(s) // num 的十进制长度

	if num%2 > 0 {
		update(int(math.Pow10(m-1)) - 1) // 十进制长为 m-1 的最大奇回文数 999..999
		update(int(math.Pow10(m)) + 1)   // 十进制长为 m+1 的最小奇回文数 100..001
	} else {
		if m == 2 {
			update(8)
		} else {
			update(int(math.Pow10(m-2))*9 - 2) // 十进制长为 m-1 的最大偶回文数 899..998
		}
		update(int(math.Pow10(m))*2 + 2) // 十进制长为 m+1 的最小偶回文数 200..002
	}

	highDigit := int(s[0] - '0')
	if highDigit%2 == num%2 {
		left, _ := strconv.Atoi(s[:(m+1)/2])
		// 枚举十进制长为 m 的邻近回文数
		for l := left - 1; l <= left+1; l++ {
			// l 最高位的奇偶性必须与 num 的相同
			if int(strconv.Itoa(l)[0]%2) != num%2 {
				continue
			}

			pal := l
			x := l
			if m%2 > 0 {
				x /= 10
			}
			for ; x > 0; x /= 10 {
				pal = pal*10 + x%10
			}
			update(pal)
		}
	} else {
		// 最高位为 highDigit - 1
		// 例如 num = 354..676，生成回文数 299..992
		if highDigit > 1 {
			update((int(math.Pow10(m-1))+1)*highDigit - 11)
		}

		// 最高位为 highDigit + 1
		// 例如 num = 354..676，生成回文数 400..004
		if highDigit < 9 {
			update((int(math.Pow10(m-1)) + 1) * (highDigit + 1))
		}
	}

	return minD
}

func minOperations(nums []int) (ans int64) {
	for _, x := range nums {
		ans += int64(nearestPalindromicDiff(x))
	}
	return ans / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(\log U)$ 或 $\mathcal{O}(1)$，取决于是否用到字符串。

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

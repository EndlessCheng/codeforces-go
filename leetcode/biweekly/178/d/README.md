**前置知识**：见 [动态规划题单](https://leetcode.cn/circle/discuss/tXLS3i/)「**十、数位 DP**」的两个讲解视频。

对于本题，需要在模板的基础上，添加三个参数：

- $\textit{digitSum}$，表示已填数字之和（数位和）。
- $\textit{prev}$，表示上一个填的数字（$i-1$ 填的数字）。
- $\textit{state}$，表示已填数字的状态：
   - $0$，表示已经填了至多一个数（不含前导零）。
   - $1$，表示已填数字是严格递增的。
   - $2$，表示已填数字是严格递减的。
   - $3$，表示已填数字不是好数。

递归入口：$\textit{digitSum} = \textit{prev} = \textit{state} = 0$。

在递归的过程中，$\textit{state}$ 变化如下：

- 如果 $\textit{state} = 0$：
  - 在填了至多一个数的情况下，如果之前填过数，那么填的数一定非零。所以，我们可以根据 $\textit{prev}$ 是否为 $0$，判断之前是否填过数。
  - 如果 $\textit{prev} = 0$，说明之前没有填过数，那么状态仍然为 $0$。
  - 如果 $\textit{prev} > 0$，说明之前填过数，那么把当前填的数字 $d$ 和上一个填的数字 $\textit{prev}$ 比大小：
      - 如果 $d > \textit{prev}$，新状态为 $1$。
      - 如果 $d < \textit{prev}$，新状态为 $2$。
      - 如果 $d = \textit{prev}$，新状态为 $3$。
- 如果 $\textit{state} = 1$：
  - 如果 $d\le \textit{prev}$，新状态为 $3$，否则不变。
- 如果 $\textit{state} = 2$：
  - 如果 $d\ge \textit{prev}$，新状态为 $3$，否则不变。
- 如果 $\textit{state} = 3$：
  - 状态不变。

递归边界：如果 $\textit{state} \ne 3$，或者 $\textit{digitSum}$ 是好数，那么找到了一个合法方案（奇妙数），返回 $1$；否则返回 $0$。

下午两点 B站@灵茶山艾府 直播讲题，欢迎关注~

## 优化前

```py [sol-Python3]
class Solution:
    # 判断数位和 s 是否为好数
    def is_good(self, s: int) -> bool:
        if s < 100:  # s 是个位数或者两位数
            return s // 10 != s % 10  # 十位和个位不相等即为好数
        # s 是三位数，其百位一定是 1
        return 1 < s // 10 % 10 < s % 10  # 只能严格递增

    def countFancy(self, l: int, r: int) -> int:
        STATE_INIT = 0      # 已经填了至多一个数（不含前导零）
        STATE_INC = 1       # 已填数字是严格递增的
        STATE_DEC = 2       # 已填数字是严格递减的
        STATE_NOT_GOOD = 3  # 已填数字不是好数

        low_s = list(map(int, str(l)))  # 避免在 dfs 中频繁调用 int()
        high_s = list(map(int, str(r)))
        n = len(high_s)
        diff_lh = n - len(low_s)

        @cache
        def dfs(i: int, digit_sum: int, prev: int, state: int, limit_low: bool, limit_high: bool) -> int:
            if i == n:
                return 1 if state != STATE_NOT_GOOD or self.is_good(digit_sum) else 0

            lo = low_s[i - diff_lh] if limit_low and i >= diff_lh else 0
            hi = high_s[i] if limit_high else 9

            res = 0
            start = lo

            # 通过 limit_low 和 i 可以判断能否不填数字，无需 is_num 参数
            if limit_low and i < diff_lh:
                # 不填数字，上界不受约束
                res = dfs(i + 1, 0, 0, STATE_INIT, True, False)
                start = 1  # 下面填数字，从 1 开始填

            for d in range(start, hi + 1):
                new_state = state
                if state == STATE_INIT:
                    if prev > 0:  # 之前填过数
                        if d > prev:
                            new_state = STATE_INC
                        elif d < prev:
                            new_state = STATE_DEC
                        else:
                            new_state = STATE_NOT_GOOD
                elif state == STATE_INC:
                    if d <= prev:
                        new_state = STATE_NOT_GOOD
                elif state == STATE_DEC:
                    if d >= prev:
                        new_state = STATE_NOT_GOOD

                res += dfs(i + 1,
                           digit_sum + d,
                           d,
                           new_state,
                           limit_low and d == lo,
                           limit_high and d == hi)

            return res

        return dfs(0, 0, 0, STATE_INIT, True, True)
```

```java [sol-Java]
class Solution {
    private static final int STATE_INIT = 0;     // 已经填了至多一个数（不含前导零）
    private static final int STATE_INC = 1;      // 已填数字是严格递增的
    private static final int STATE_DEC = 2;      // 已填数字是严格递减的
    private static final int STATE_NOT_GOOD = 3; // 已填数字不是好数

    public long countFancy(long l, long r) {
        char[] lowS = String.valueOf(l).toCharArray();
        char[] highS = String.valueOf(r).toCharArray();
        int n = highS.length;
        long[][][][] memo = new long[n][n * 9 + 1][10][4]; // 数位和最大 9n
        return dfs(0, 0, 0, STATE_INIT, true, true, lowS, highS, memo);
    }

    private long dfs(int i, int digitSum, int prev, int state, boolean limitLow, boolean limitHigh,
                     char[] lowS, char[] highS, long[][][][] memo) {
        if (i == highS.length) {
            return state != STATE_NOT_GOOD || isGood(digitSum) ? 1 : 0;
        }

        if (!limitLow && !limitHigh && memo[i][digitSum][prev][state] > 0) {
            return memo[i][digitSum][prev][state] - 1;
        }

        int diffLh = highS.length - lowS.length;
        int lo = limitLow && i >= diffLh ? lowS[i - diffLh] - '0' : 0;
        int hi = limitHigh ? highS[i] - '0' : 9;

        long res = 0;
        int d = lo;

        // 通过 limitLow 和 i 可以判断能否不填数字，无需 isNum 参数
        if (limitLow && i < diffLh) {
            // 不填数字，上界不受约束
            res = dfs(i + 1, 0, 0, STATE_INIT, true, false, lowS, highS, memo);
            d = 1; // 下面填数字，从 1 开始填
        }

        for (; d <= hi; d++) {
            int newState = state;
            switch (state) {
                case STATE_INIT:
                    if (prev > 0) { // 之前填过数
                        if (d > prev) {
                            newState = STATE_INC;
                        } else if (d < prev) {
                            newState = STATE_DEC;
                        } else {
                            newState = STATE_NOT_GOOD;
                        }
                    }
                    break;
                case STATE_INC:
                    if (d <= prev) {
                        newState = STATE_NOT_GOOD;
                    }
                    break;
                case STATE_DEC:
                    if (d >= prev) {
                        newState = STATE_NOT_GOOD;
                    }
                    break;
            }
            res += dfs(i + 1, digitSum + d, d, newState,
                    limitLow && d == lo, limitHigh && d == hi, lowS, highS, memo);
        }

        if (!limitLow && !limitHigh) {
            memo[i][digitSum][prev][state] = res + 1; // 这样写无需初始化 memo 为 -1
        }
        return res;
    }

    // 判断数位和 s 是否为好数
    private boolean isGood(int s) {
        if (s < 100) { // s 是个位数或者两位数
            return s / 10 != s % 10; // 十位和个位不相等即为好数
        }
        // s 是三位数，其百位一定是 1
        return 1 < s / 10 % 10 && s / 10 % 10 < s % 10; // 只能严格递增
    }
}
```

```cpp [sol-C++]
class Solution {
    static constexpr int STATE_INIT = 0;     // 已经填了至多一个数（不含前导零）
    static constexpr int STATE_INC = 1;      // 已填数字是严格递增的
    static constexpr int STATE_DEC = 2;      // 已填数字是严格递减的
    static constexpr int STATE_NOT_GOOD = 3; // 已填数字不是好数

    // 判断数位和 s 是否为好数
    bool is_good(int s) {
        if (s < 100) { // s 是个位数或者两位数
            return s / 10 != s % 10; // 十位和个位不相等即为好数
        }
        // s 是三位数，其百位一定是 1
        return 1 < s / 10 % 10 && s / 10 % 10 < s % 10; // 只能严格递增
    }

public:
    long long countFancy(long long l, long long r) {
        string low_s = to_string(l);
        string high_s = to_string(r);
        int n = high_s.size();
        int diff_lh = n - low_s.size();
        vector memo(n, vector<array<array<long long, 4>, 10>>(n * 9 + 1)); // 数位和最大 9n

        auto dfs = [&](this auto&& dfs, int i, int digit_sum, int prev, int state, bool limit_low, bool limit_high) -> long long {
            if (i == n) {
                return state != STATE_NOT_GOOD || is_good(digit_sum);
            }

            auto& ref = memo[i][digit_sum][prev][state];
            if (!limit_low && !limit_high && ref) {
                return ref - 1; // 见后面的 ref = res + 1;
            }

            int lo = limit_low && i >= diff_lh ? low_s[i - diff_lh] - '0' : 0;
            int hi = limit_high ? high_s[i] - '0' : 9;

            long long res = 0;
            int d = lo;

            // 通过 limit_low 和 i 可以判断能否不填数字，无需 is_num 参数
            if (limit_low && i < diff_lh) {
                // 不填数字，上界不受约束
                res = dfs(i + 1, 0, 0, STATE_INIT, true, false);
                d = 1; // 下面填数字，从 1 开始填
            }

            for (; d <= hi; d++) {
                int new_state = state;
                switch (state) {
                    case STATE_INIT:
                        if (prev > 0) { // 之前填过数
                            if (d > prev) {
                                new_state = STATE_INC;
                            } else if (d < prev) {
                                new_state = STATE_DEC;
                            } else {
                                new_state = STATE_NOT_GOOD;
                            }
                        }
                        break;
                    case STATE_INC:
                        if (d <= prev) {
                            new_state = STATE_NOT_GOOD;
                        }
                        break;
                    case STATE_DEC:
                        if (d >= prev) {
                            new_state = STATE_NOT_GOOD;
                        }
                        break;
                }
                res += dfs(i + 1, digit_sum + d, d, new_state, limit_low && d == lo, limit_high && d == hi);
            }

            if (!limit_low && !limit_high) {
                ref = res + 1; // 这样写无需初始化 memo 为 -1
            }
            return res;
        };

        return dfs(0, 0, 0, STATE_INIT, true, true);
    }
};
```

```go [sol-Go]
// 判断数位和 s 是否为好数
func isGood(s int) bool {
	if s < 100 { // s 是个位数或者两位数
		return s/10 != s%10 // 十位和个位不相等即为好数
	}
	// s 是三位数，其百位一定是 1
	return 1 < s/10%10 && s/10%10 < s%10 // 只能严格递增
}

func countFancy(l, r int64) int64 {
	lowS := strconv.FormatInt(l, 10)
	highS := strconv.FormatInt(r, 10)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][][10][4]int64, n)
	for i := range memo {
		memo[i] = make([][10][4]int64, n*9+1) // 数位和最大 9n
	}

	const (
		stateInit    = iota // 已经填了至多一个数（不含前导零）
		stateInc            // 已填数字是严格递增的
		stateDec            // 已填数字是严格递减的
		stateNotGood        // 已填数字不是好数
	)

	var dfs func(int, int, int, int, bool, bool) int64
	dfs = func(i, digitSum, prev, state int, limitLow, limitHigh bool) (res int64) {
		if i == n {
			if state != stateNotGood || isGood(digitSum) {
				return 1 // 合法
			}
			return 0 // 不合法
		}

		if !limitLow && !limitHigh {
			dv := &memo[i][digitSum][prev][state]
			if *dv > 0 {
				return *dv - 1
			}
			defer func() { *dv = res + 1 }() // 这样写无需初始化 memo 为 -1
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		d := lo
		// 通过 limitLow 和 i 可以判断能否不填数字，无需 isNum 参数
		if limitLow && i < diffLH {
			// 不填数字，上界不受约束
			res = dfs(i+1, 0, 0, stateInit, true, false)
			d = 1 // 下面填数字，从 1 开始填
		}

		for ; d <= hi; d++ {
			newState := state
			switch state {
			case stateInit:
				if prev > 0 { // 之前填过数
					if d > prev {
						newState = stateInc
					} else if d < prev {
						newState = stateDec
					} else {
						newState = stateNotGood
					}
				}
			case stateInc:
				if d <= prev {
					newState = stateNotGood
				}
			case stateDec:
				if d >= prev {
					newState = stateNotGood
				}
			}
			res += dfs(i+1, digitSum+d, d, newState, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}

	return dfs(0, 0, 0, stateInit, true, true)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(D^3\log^2 r)$，其中 $D=10$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(D^2\log^2 r)$，单个状态的计算时间为 $\mathcal{O}(D)$，所以总的时间复杂度为 $\mathcal{O}(D^3\log^2 r)$。
- 空间复杂度：$\mathcal{O}(D^2\log^2 r)$。保存多少状态，就需要多少空间。

## 优化

原问题（奇妙数的个数）可以拆分成两个计数问题：

1. 数位和是好数的数（这个数是不是好数都可以）。
2. 数位和不是好数的数，那么这个数必须是好数。

第一个问题，是一个（相比优化前）更简单的数位 DP。我们只需 $\textit{digitSum}$ 这个参数，在递归边界判断 $\textit{digitSum}$ 是否为好数即可。

第二个问题，我们可以直接枚举所有的好数。怎么枚举？

由于好数是严格递增或者严格递减的，所以**各个数位互不相同**。我们可以枚举 $\{0,1,2,\ldots,9\}$ 的**子集**，把子集排成严格递减或者严格递增，就枚举了所有的好数。注意严格递增的好数不能包含 $0$（无前导零）。

如何枚举子集？见 [78. 子集](https://leetcode.cn/problems/subsets/)。

```py [sol-Python3]
# 判断数位和 s 是否为好数
def is_good(s: int) -> bool:
    if s < 100:  # s 是个位数或者两位数
        return s // 10 != s % 10  # 十位和个位不相等即为好数
    # s 是三位数，其百位一定是 1
    return 1 < s // 10 % 10 < s % 10  # 只能严格递增


# 预处理数位和不是好数的好数（只有 139 个）
good_nums = []

for mask in range(1, 1 << 10):
    # 构造严格递减好数
    x = s = 0
    for i in range(9, -1, -1):
        if mask >> i & 1:
            x = x * 10 + i
            s += i
    if not is_good(s):
        good_nums.append(x)

    # 构造严格递增好数
    if mask & 1:  # 不能包含 0
        continue
    x = s = 0
    for i in range(1, 10):
        if mask >> i & 1:
            x = x * 10 + i
            s += i
    if not is_good(s):
        good_nums.append(x)

good_nums.sort()  # 方便二分求个数


class Solution:
    def countFancy(self, l: int, r: int) -> int:
        # 计算 [l, r] 内的数位和不是好数的好数的个数
        ans = bisect_right(good_nums, r) - bisect_left(good_nums, l)

        # 计算 [l, r] 内的数位和是好数的数的个数（这个数是不是好数都可以）
        low_s = list(map(int, str(l)))  # 避免在 dfs 中频繁调用 int()
        high_s = list(map(int, str(r)))
        n = len(high_s)
        diff_lh = n - len(low_s)

        @cache
        def dfs(i: int, digit_sum: int, limit_low: bool, limit_high: bool) -> int:
            if i == n:
                return 1 if is_good(digit_sum) else 0

            lo = low_s[i - diff_lh] if limit_low and i >= diff_lh else 0
            hi = high_s[i] if limit_high else 9

            res = 0
            start = lo

            # 通过 limit_low 和 i 可以判断能否不填数字，无需 is_num 参数
            if limit_low and i < diff_lh:
                # 不填数字，上界不受约束
                res = dfs(i + 1, 0, True, False)
                start = 1  # 下面填数字，从 1 开始填

            for d in range(start, hi + 1):
                res += dfs(i + 1,
                           digit_sum + d,
                           limit_low and d == lo,
                           limit_high and d == hi)

            return res

        return ans + dfs(0, 0, True, True)
```

```java [sol-Java]
class Solution {
    // 数位和不是好数的好数（只有 139 个）
    private static final List<Long> goodNums = new ArrayList<>();
    private static boolean initialized = false;

    // 这样写比 static block 快
    public Solution() {
        if (initialized) {
            return;
        }
        initialized = true;

        for (int mask = 1; mask < (1 << 10); mask++) {
            // 构造严格递减好数
            long x = 0;
            int sum = 0;
            for (int i = 9; i >= 0; i--) {
                if ((mask >> i & 1) > 0) {
                    x = x * 10 + i;
                    sum += i;
                }
            }
            if (!isGood(sum)) {
                goodNums.add(x);
            }

            // 构造严格递增好数
            if ((mask & 1) > 0) { // 不能包含 0
                continue;
            }
            x = 0;
            sum = 0;
            for (int i = 1; i < 10; i++) {
                if ((mask >> i & 1) > 0) {
                    x = x * 10 + i;
                    sum += i;
                }
            }
            if (!isGood(sum)) {
                goodNums.add(x);
            }
        }

        Collections.sort(goodNums); // 方便二分求个数
    }

    public long countFancy(long l, long r) {
        // 计算 [l, r] 内的数位和不是好数的好数的个数
        long ans = lowerBound(r + 1) - lowerBound(l);

        // 计算 [l, r] 内的数位和是好数的数的个数（这个数是不是好数都可以）
        char[] lowS = String.valueOf(l).toCharArray();
        char[] highS = String.valueOf(r).toCharArray();
        int n = highS.length;
        long[][] memo = new long[n][n * 9 + 1]; // 数位和最大 9n
        return ans + dfs(0, 0, true, true, lowS, highS, memo);
    }

    private long dfs(int i, int digitSum, boolean limitLow, boolean limitHigh,
                     char[] lowS, char[] highS, long[][] memo) {
        if (i == highS.length) {
            return isGood(digitSum) ? 1 : 0;
        }

        if (!limitLow && !limitHigh && memo[i][digitSum] > 0) {
            return memo[i][digitSum] - 1;
        }

        int diffLh = highS.length - lowS.length;
        int lo = limitLow && i >= diffLh ? lowS[i - diffLh] - '0' : 0;
        int hi = limitHigh ? highS[i] - '0' : 9;

        long res = 0;
        int d = lo;

        // 通过 limitLow 和 i 可以判断能否不填数字，无需 isNum 参数
        if (limitLow && i < diffLh) {
            // 不填数字，上界不受约束
            res = dfs(i + 1, 0, true, false, lowS, highS, memo);
            d = 1; // 下面填数字，从 1 开始填
        }

        for (; d <= hi; d++) {
            res += dfs(i + 1, digitSum + d,
                    limitLow && d == lo, limitHigh && d == hi, lowS, highS, memo);
        }

        if (!limitLow && !limitHigh) {
            memo[i][digitSum] = res + 1; // 这样写无需初始化 memo 为 -1
        }
        return res;
    }

    // 判断数位和 s 是否为好数
    private boolean isGood(int s) {
        if (s < 100) { // s 是个位数或者两位数
            return s / 10 != s % 10; // 十位和个位不相等即为好数
        }
        // s 是三位数，其百位一定是 1
        return 1 < s / 10 % 10 && s / 10 % 10 < s % 10; // 只能严格递增
    }

    // 返回 goodNums 中第一个 >= x 的数的下标
    // 如果没有这样的数，返回 goodNums.size()
    private int lowerBound(long x) {
        // goodNums 没有重复元素，可以用库函数
        int i = Collections.binarySearch(goodNums, x);
        return i < 0 ? ~i : i;
    }
}
```

```cpp [sol-C++]
// 判断数位和 s 是否为好数
bool is_good(int s) {
    if (s < 100) { // s 是个位数或者两位数
        return s / 10 != s % 10; // 十位和个位不相等即为好数
    }
    // s 是三位数，其百位一定是 1
    return 1 < s / 10 % 10 && s / 10 % 10 < s % 10; // 只能严格递增
}

// 数位和不是好数的好数（只有 139 个）
vector<long long> good_nums;

int initGoodNums = []() {
    for (int mask = 1; mask < (1 << 10); mask++) {
        // 构造严格递减好数
        long long x = 0;
        int sum = 0;
        for (int i = 9; i >= 0; i--) {
            if (mask >> i & 1) {
                x = x * 10 + i;
                sum += i;
            }
        }
        if (!is_good(sum)) {
            good_nums.push_back(x);
        }

        // 构造严格递增好数
        if (mask & 1) { // 不能包含 0
            continue;
        }
        x = 0;
        sum = 0;
        for (int i = 1; i < 10; i++) {
            if (mask >> i & 1) {
                x = x * 10 + i;
                sum += i;
            }
        }
        if (!is_good(sum)) {
            good_nums.push_back(x);
        }
    }

    ranges::sort(good_nums); // 方便二分求个数
    return 0;
}();

class Solution {
public:
    long long countFancy(long long l, long long r) {
        // 计算 [l, r] 内的数位和不是好数的好数的个数
        long long ans = ranges::upper_bound(good_nums, r) - ranges::lower_bound(good_nums, l);

        // 计算 [l, r] 内的数位和是好数的数的个数（这个数是不是好数都可以）
        string low_s = to_string(l);
        string high_s = to_string(r);
        int n = high_s.size();
        int diff_lh = n - low_s.size();
        vector memo(n, vector<long long>(n * 9 + 1, -1)); // 数位和最大 9n

        auto dfs = [&](this auto&& dfs, int i, int digit_sum, bool limit_low, bool limit_high) -> long long {
            if (i == n) {
                return is_good(digit_sum);
            }

            auto& ref = memo[i][digit_sum];
            if (!limit_low && !limit_high && ref >= 0) {
                return ref;
            }

            int lo = limit_low && i >= diff_lh ? low_s[i - diff_lh] - '0' : 0;
            int hi = limit_high ? high_s[i] - '0' : 9;

            long long res = 0;
            int d = lo;

            // 通过 limit_low 和 i 可以判断能否不填数字，无需 is_num 参数
            if (limit_low && i < diff_lh) {
                // 不填数字，上界不受约束
                res = dfs(i + 1, 0, true, false);
                d = 1; // 下面填数字，从 1 开始填
            }

            for (; d <= hi; d++) {
                res += dfs(i + 1, digit_sum + d, limit_low && d == lo, limit_high && d == hi);
            }

            if (!limit_low && !limit_high) {
                ref = res;
            }
            return res;
        };

        return ans + dfs(0, 0, true, true);
    }
};
```

```go [sol-Go]
// 数位和不是好数的好数（只有 139 个）
var goodNums []int

func init() {
	for mask := 1; mask < 1<<10; mask++ {
		// 构造严格递减好数
		x := 0
		sum := 0
		for i := 9; i >= 0; i-- {
			if mask>>i&1 > 0 {
				x = x*10 + i
				sum += i
			}
		}
		if !isGood(sum) {
			goodNums = append(goodNums, x)
		}

		// 构造严格递增好数
		if mask&1 > 0 { // 不能包含 0
			continue
		}
		x = 0
		sum = 0
		for i := 1; i < 10; i++ {
			if mask>>i&1 > 0 {
				x = x*10 + i
				sum += i
			}
		}
		if !isGood(sum) {
			goodNums = append(goodNums, x)
		}
	}

	slices.Sort(goodNums) // 方便二分求个数
}

// 判断数位和 s 是否为好数
func isGood(s int) bool {
	if s < 100 { // s 是个位数或者两位数
		return s/10 != s%10 // 十位和个位不相等即为好数
	}
	// s 是三位数，其百位一定是 1
	return 1 < s/10%10 && s/10%10 < s%10 // 只能严格递增
}

func countFancy(l, r int64) int64 {
	// 计算 [l, r] 内的数位和不是好数的好数的个数
	ans := int64(sort.SearchInts(goodNums, int(r+1)) - sort.SearchInts(goodNums, int(l)))

	// 计算 [l, r] 内的数位和是好数的数的个数（这个数是不是好数都可以）
	lowS := strconv.FormatInt(l, 10)
	highS := strconv.FormatInt(r, 10)
	n := len(highS)
	diffLH := n - len(lowS)
	memo := make([][]int64, n)
	for i := range memo {
		memo[i] = make([]int64, n*9+1) // 数位和最大 9n
	}

	var dfs func(int, int, bool, bool) int64
	dfs = func(i, digitSum int, limitLow, limitHigh bool) (res int64) {
		if i == n {
			if isGood(digitSum) {
				return 1 // 合法
			}
			return 0 // 不合法
		}
		if !limitLow && !limitHigh {
			dv := &memo[i][digitSum]
			if *dv > 0 {
				return *dv - 1
			}
			defer func() { *dv = res + 1 }() // 这样写无需初始化 memo 为 -1
		}

		lo := 0
		if limitLow && i >= diffLH {
			lo = int(lowS[i-diffLH] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(highS[i] - '0')
		}

		d := lo
		// 通过 limitLow 和 i 可以判断能否不填数字，无需 isNum 参数
		if limitLow && i < diffLH {
			// 不填数字，上界不受约束
			res = dfs(i+1, 0, true, false)
			d = 1 // 下面填数字，从 1 开始填
		}

		for ; d <= hi; d++ {
			res += dfs(i+1, digitSum+d, limitLow && d == lo, limitHigh && d == hi)
		}
		return
	}

	return ans + dfs(0, 0, true, true)
}
```

#### 复杂度分析

不计入预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(D^2\log^2 r)$，其中 $D=10$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(D\log^2 r)$，单个状态的计算时间为 $\mathcal{O}(D)$，所以总的时间复杂度为 $\mathcal{O}(D^2\log^2 r)$。
- 空间复杂度：$\mathcal{O}(D\log^2 r)$。保存多少状态，就需要多少空间。

## 专题训练

见下面动态规划题单的「**十、数位 DP**」。

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
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

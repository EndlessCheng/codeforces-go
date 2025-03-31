## 方法一：正反两次扫描

为了计算酿造药水的时间，定义 $\textit{lastFinish}[i]$ 表示巫师 $i$ 完成上一瓶药水的时间。

示例 1 在处理完 $\textit{mana}[0]$ 后，有

$$
\textit{lastFinish} = [5,30,40,60]
$$

如果接着 $\textit{lastFinish}$ 继续酿造下一瓶药水 $\textit{mana}[1]=1$，完成时间是多少？注意开始酿造的时间不能早于 $\textit{lastFinish}[i]$。

| $i$  | $\textit{skill}[i]$  | $\textit{lastFinish}[i]$  | 完成时间  |
|---|---|---|---|
| $0$  | $1$  | $5$  |  $5+1=6$ |
| $1$  | $5$  | $30$  |  $\max(6,30)+5=35$ |
| $2$  | $2$  | $40$  |  $\max(35,40)+2=42$ |
| $3$  | $4$  | $60$  |  $\max(42,60)+4=64$ |

题目要求「药水在当前巫师完成工作后必须立即传递给下一个巫师并开始处理」，也就是说，酿造药水的过程中是**不能有停顿**的。

从 $64$ 开始**倒推**，可以得到每名巫师的**实际完成时间**。比如倒数第二位巫师的完成时间，就是 $64$ 减去最后一名巫师花费的时间 $4\cdot 1$，得到 $60$。

| $i$  | $\textit{skill}[i+1]$  | 实际完成时间  |
|---|---|---|
| $3$  |  - |  $64$ |
| $2$  | $4$  |  $64-4\cdot 1=60$ |
| $1$  | $2$  |  $60-2\cdot 1=58$ |
| $0$  | $5$  |  $58-5\cdot 1=53$ |

按照上述过程处理每瓶药水，最终答案为 $\textit{lastFinish}[n-1]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV12eXYYVE5H/?t=7m48s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minTime(self, skill: List[int], mana: List[int]) -> int:
        n = len(skill)
        last_finish = [0] * n  # 第 i 名巫师完成上一瓶药水的时间
        for m in mana:
            # 按题意模拟
            sum_t = 0
            for x, last in zip(skill, last_finish):
                if last > sum_t: sum_t = last  # 手写 max
                sum_t += x * m
            # 倒推：如果酿造药水的过程中没有停顿，那么 last_finish[i] 应该是多少
            last_finish[-1] = sum_t
            for i in range(n - 2, -1, -1):
                last_finish[i] = last_finish[i + 1] - skill[i + 1] * m
        return last_finish[-1]
```

```java [sol-Java]
class Solution {
    public long minTime(int[] skill, int[] mana) {
        int n = skill.length;
        long[] lastFinish = new long[n]; // 第 i 名巫师完成上一瓶药水的时间
        for (int m : mana) {
            // 按题意模拟
            long sumT = 0;
            for (int i = 0; i < n; i++) {
                sumT = Math.max(sumT, lastFinish[i]) + skill[i] * m;
            }
            // 倒推：如果酿造药水的过程中没有停顿，那么 lastFinish[i] 应该是多少
            lastFinish[n - 1] = sumT;
            for (int i = n - 2; i >= 0; i--) {
                lastFinish[i] = lastFinish[i + 1] - skill[i + 1] * m;
            }
        }
        return lastFinish[n - 1];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minTime(vector<int>& skill, vector<int>& mana) {
        int n = skill.size();
        vector<long long> last_finish(n); // 第 i 名巫师完成上一瓶药水的时间
        for (int m : mana) {
            // 按题意模拟
            long long sum_t = 0;
            for (int i = 0; i < n; i++) {
                sum_t = max(sum_t, last_finish[i]) + skill[i] * m;
            }
            // 倒推：如果酿造药水的过程中没有停顿，那么 lastFinish[i] 应该是多少
            last_finish[n - 1] = sum_t;
            for (int i = n - 2; i >= 0; i--) {
                last_finish[i] = last_finish[i + 1] - skill[i + 1] * m;
            }
        }
        return last_finish[n - 1];
    }
};
```

```go [sol-Go]
func minTime(skill, mana []int) int64 {
	n := len(skill)
	lastFinish := make([]int, n) // 第 i 名巫师完成上一瓶药水的时间
	for _, m := range mana {
		// 按题意模拟
		sumT := 0
		for i, x := range skill {
			sumT = max(sumT, lastFinish[i]) + x*m
		}
		// 倒推：如果酿造药水的过程中没有停顿，那么 lastFinish[i] 应该是多少
		lastFinish[n-1] = sumT
		for i := n - 2; i >= 0; i-- {
			lastFinish[i] = lastFinish[i+1] - skill[i+1]*m
		}
	}
	return int64(lastFinish[n-1])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $\textit{skill}$ 的长度，$m$ 是 $\textit{mana}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：递推开始时间

由于酿造药水的过程是连续的，所以知道了开始时间（或者完成时间）就能知道每个 $\textit{lastFinish}[i]$。所以 $\textit{lastFinish}$ 数组是多余的。

设开始酿造 $\textit{mana}[j]$ 的时间为 $\textit{start}_j$，那么有

$$
\textit{lastFinish}_j[i] = \textit{start}_j + \textit{mana}[j]\cdot \sum_{k=0}^{i} \textit{skill}[k]
$$

在已知 $\textit{start}_{j-1}$ 的前提下，能否递推算出 $\textit{start}_j$？

哪位巫师决定了开始时间？假设第 $i$ 位巫师决定了开始时间，那么这位巫师**完成** $\textit{mana}[j-1]$ 的时间，同时也是他**开始** $\textit{mana}[j]$ 的时间。

所以有

$$
\textit{lastFinish}_{j-1}[i] + \textit{mana}[j]\cdot \textit{skill}[i] = \textit{lastFinish}_j[i]
$$

两边代入 $\textit{lastFinish}_j[i]$ 的式子，得

$$
\textit{start}_{j-1} + \textit{mana}[j-1]\cdot \sum_{k=0}^{i} \textit{skill}[k] + \textit{mana}[j]\cdot \textit{skill}[i] = \textit{start}_j + \textit{mana}[j]\cdot \sum_{k=0}^{i} \textit{skill}[k]
$$

移项得

$$
\textit{start}_j = \textit{start}_{j-1} + \textit{mana}[j-1]\cdot \sum_{k=0}^{i} \textit{skill}[k] - \textit{mana}[j]\cdot \sum_{k=0}^{i-1} \textit{skill}[k]
$$

计算 $\textit{skill}$ 的 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 数组 $s$，上式为

$$
\textit{start}_j = \textit{start}_{j-1} + \textit{mana}[j-1]\cdot s[i+1] - \textit{mana}[j]\cdot s[i]
$$

枚举 $i$，取最大值，得

$$
\textit{start}_j = \textit{start}_{j-1} + \max_{i=0}^{n-1} \left\{\textit{mana}[j-1]\cdot s[i+1] - \textit{mana}[j]\cdot s[i]\right\}
$$

初始值 $\textit{start}_0 = 0$。

答案为 $\textit{lastFinish}_{m-1}[n-1] = \textit{start}_{m-1} + \textit{mana}[m-1]\cdot s[n]$。

```py [sol-Python3]
class Solution:
    def minTime(self, skill: List[int], mana: List[int]) -> int:
        n = len(skill)
        s = list(accumulate(skill, initial=0))  # skill 的前缀和
        start = 0
        for pre, cur in pairwise(mana):
            start += max(pre * s[i + 1] - cur * s[i] for i in range(n))
        return start + mana[-1] * s[-1]
```

```java [sol-Java]
class Solution {
    public long minTime(int[] skill, int[] mana) {
        int n = skill.length;
        int[] s = new int[n + 1]; // skill 的前缀和
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + skill[i];
        }

        int m = mana.length;
        long start = 0;
        for (int j = 1; j < m; j++) {
            long mx = 0;
            for (int i = 0; i < n; i++) {
                mx = Math.max(mx, (long) mana[j - 1] * s[i + 1] - (long) mana[j] * s[i]);
            }
            start += mx;
        }
        return start + (long) mana[m - 1] * s[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minTime(vector<int>& skill, vector<int>& mana) {
        int n = skill.size(), m = mana.size();
        vector<int> s(n + 1); // skill 的前缀和
        partial_sum(skill.begin(), skill.end(), s.begin() + 1);

        long long start = 0;
        for (int j = 1; j < m; j++) {
            long long mx = 0;
            for (int i = 0; i < n; i++) {
                mx = max(mx, 1LL * mana[j - 1] * s[i + 1] - 1LL * mana[j] * s[i]);
            }
            start += mx;
        }
        return start + 1LL * mana[m - 1] * s[n];
    }
};
```

```go [sol-Go]
func minTime(skill, mana []int) int64 {
	n, m := len(skill), len(mana)
	s := make([]int, n+1) // skill 的前缀和
	for i, x := range skill {
		s[i+1] = s[i] + x
	}

	start := 0
	for j := 1; j < m; j++ {
		mx := 0
		for i := range n {
			mx = max(mx, mana[j-1]*s[i+1]-mana[j]*s[i])
		}
		start += mx
	}
	return int64(start + mana[m-1]*s[n])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $\textit{skill}$ 的长度，$m$ 是 $\textit{mana}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。如果在遍历的同时计算前缀和，则可以做到 $\mathcal{O}(1)$ 空间。

## 方法三：record 优化

将递推式

$$
\textit{start}_j = \textit{start}_{j-1} + \max_{i=0}^{n-1} \left\{\textit{mana}[j-1]\cdot s[i+1] - \textit{mana}[j]\cdot s[i]\right\}
$$

变形为

$$
\textit{start}_j = \textit{start}_{j-1} + \max_{i=0}^{n-1} \left\{(\textit{mana}[j-1]-\textit{mana}[j])\cdot s[i]  + \textit{mana}[j-1]\cdot \textit{skill}[i] \right\}
$$

设 $d = \textit{mana}[j-1]-\textit{mana}[j]$。分类讨论：

- 如果 $d > 0$。由于 $s$ 是单调递增数组，如果 $\textit{skill}[3] < \textit{skill}[5]$，那么 $i=3$ 绝对不会算出最大值；但如果 $\textit{skill}[3] > \textit{skill}[5]$，谁会算出最大值就不一定了。所以我们只需要考虑 $\textit{skill}$ 的**逆序 record**，这才是可能成为最大值的数据。其中逆序 record 的意思是，倒序遍历 $\textit{skill}$，每次遍历到更大的数，就记录下标。
- 如果 $d < 0$。由于 $s$ 是单调递增数组，如果 $\textit{skill}[5] < \textit{skill}[3]$，那么 $i=5$ 绝对不会算出最大值；但如果 $\textit{skill}[5] > \textit{skill}[3]$，谁会算出最大值就不一定了。所以我们只需要考虑 $\textit{skill}$ 的**正序 record**，这才是可能成为最大值的数据。其中正序 record 的意思是，正序遍历 $\textit{skill}$，每次遍历到更大的数，就记录下标。
- $d = 0$ 的情况可以并入 $d>0$ 的情况。

```py [sol-Python3]
class Solution:
    def minTime(self, skill: List[int], mana: List[int]) -> int:
        n = len(skill)
        s = list(accumulate(skill, initial=0))

        suf_record = [n - 1]
        for i in range(n - 2, -1, -1):
            if skill[i] > skill[suf_record[-1]]:
                suf_record.append(i)

        pre_record = [0]
        for i in range(1, n):
            if skill[i] > skill[pre_record[-1]]:
                pre_record.append(i)

        start = 0
        for pre, cur in pairwise(mana):
            record = pre_record if pre < cur else suf_record
            start += max(pre * s[i + 1] - cur * s[i] for i in record)
        return start + mana[-1] * s[-1]
```

```java [sol-Java]
class Solution {
    public long minTime(int[] skill, int[] mana) {
        int n = skill.length;
        int[] s = new int[n + 1];
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + skill[i];
        }

        List<Integer> suf = new ArrayList<>();
        suf.add(n - 1);
        for (int i = n - 2; i >= 0; i--) {
            if (skill[i] > skill[suf.getLast()]) {
                suf.add(i);
            }
        }

        List<Integer> pre = new ArrayList<>();
        pre.add(0);
        for (int i = 1; i < n; i++) {
            if (skill[i] > skill[pre.getLast()]) {
                pre.add(i);
            }
        }

        int m = mana.length;
        long start = 0;
        for (int j = 1; j < m; j++) {
            List<Integer> record = mana[j - 1] < mana[j] ? pre : suf;
            long mx = 0;
            for (int i : record) {
                mx = Math.max(mx, (long) mana[j - 1] * s[i + 1] - (long) mana[j] * s[i]);
            }
            start += mx;
        }
        return start + (long) mana[m - 1] * s[n];
    }
}
```

```java [sol-Java 数组]
class Solution {
    public long minTime(int[] skill, int[] mana) {
        int n = skill.length;
        int[] s = new int[n + 1];
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + skill[i];
        }

        int[] suf = new int[n];
        int sufLen = 0;
        suf[sufLen++] = n - 1;
        for (int i = n - 2; i >= 0; i--) {
            if (skill[i] > skill[suf[sufLen - 1]]) {
                suf[sufLen++] = i;
            }
        }

        int[] pre = new int[n];
        int preLen = 0;
        pre[preLen++] = 0;
        for (int i = 1; i < n; i++) {
            if (skill[i] > skill[pre[preLen - 1]]) {
                pre[preLen++] = i;
            }
        }

        int m = mana.length;
        long start = 0;
        for (int j = 1; j < m; j++) {
            int[] record = mana[j - 1] < mana[j] ? pre : suf;
            int recordLen = mana[j - 1] < mana[j] ? preLen : sufLen;
            long mx = 0;
            for (int k = 0; k < recordLen; k++) {
                int i = record[k];
                mx = Math.max(mx, (long) mana[j - 1] * s[i + 1] - (long) mana[j] * s[i]);
            }
            start += mx;
        }
        return start + (long) mana[m - 1] * s[n];
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minTime(vector<int>& skill, vector<int>& mana) {
        int n = skill.size(), m = mana.size();
        vector<int> s(n + 1);
        partial_sum(skill.begin(), skill.end(), s.begin() + 1);

        vector<int> suf = {n - 1};
        for (int i = n - 2; i >= 0; i--) {
            if (skill[i] > skill[suf.back()]) {
                suf.push_back(i);
            }
        }

        vector<int> pre = {0};
        for (int i = 1; i < n; i++) {
            if (skill[i] > skill[pre.back()]) {
                pre.push_back(i);
            }
        }

        long long start = 0;
        for (int j = 1; j < m; j++) {
            auto& record = mana[j - 1] < mana[j] ? pre : suf;
            long long mx = 0;
            for (int i : record) {
                mx = max(mx, 1LL * mana[j - 1] * s[i + 1] - 1LL * mana[j] * s[i]);
            }
            start += mx;
        }
        return start + 1LL * mana[m - 1] * s[n];
    }
};
```

```go [sol-Go]
func minTime(skill, mana []int) int64 {
	n, m := len(skill), len(mana)
	s := make([]int, n+1)
	for i, x := range skill {
		s[i+1] = s[i] + x
	}

	suf := []int{n - 1}
	for i := n - 2; i >= 0; i-- {
		if skill[i] > skill[suf[len(suf)-1]] {
			suf = append(suf, i)
		}
	}

	pre := []int{0}
	for i := 1; i < n; i++ {
		if skill[i] > skill[pre[len(pre)-1]] {
			pre = append(pre, i)
		}
	}

	start := 0
	for j := 1; j < m; j++ {
		record := suf
		if mana[j-1] < mana[j] {
			record = pre
		}
		mx := 0
		for _, i := range record {
			mx = max(mx, mana[j-1]*s[i+1]-mana[j]*s[i])
		}
		start += mx
	}
	return int64(start + mana[m-1]*s[n])
}
```

#### 复杂度分析（最坏情况）

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $\textit{skill}$ 的长度，$m$ 是 $\textit{mana}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

#### 复杂度分析（随机情况）

力扣喜欢出随机数据，上述算法在随机数据下的性能如何？

换句话说，在随机数据下，record 的期望长度是多少？

为方便分析，假设 $\textit{skill}$ 是一个随机的 $[1,n]$ 的排列。

$\textit{skill}[i]$ 如果是一个新的最大值，那么它是 $[0,i]$ 中的最大值。在随机排列的情况下，$[0,i]$ 中的排列也是随机的，所以这等价于该排列的最后一个数是最大值的概率，即

$$
\dfrac{1}{i+1}
$$

record 的期望长度，等于「每个位置能否成为新的最大值」之和，能就贡献 $1$，不能就贡献 $0$。

所以 $\textit{skill}[i]$ 给期望的贡献是 $\dfrac{1}{i+1}$。所以 record 的期望长度为

$$
\sum_{i=0}^{n-1} \dfrac{1}{i+1}
$$

由调和级数可知，record 的期望长度为 $\Theta(\log n)$。

- 平均情况下的时间复杂度：$\Theta(n + m\log n)$，其中 $n$ 是 $\textit{skill}$ 的长度，$m$ 是 $\textit{mana}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法四：凸包 + 二分

**前置知识**：二维计算几何，凸包，Andrew 算法。

把递推式

$$
\textit{start}_j = \textit{start}_{j-1} + \max_{i=0}^{n-1} \left\{(\textit{mana}[j-1]-\textit{mana}[j])\cdot s[i]  + \textit{mana}[j-1]\cdot \textit{skill}[i] \right\}
$$

中的 

$$
(\textit{mana}[j-1]-\textit{mana}[j])\cdot s[i]  + \textit{mana}[j-1]\cdot \textit{skill}[i]
$$

改成点积的形式，这样我们能得到来自几何意义上的观察。

设向量 $\mathbf{v}_i = (s[i],\textit{skill}[i])$。

设向量 $\mathbf{p} = (\textit{mana}[j-1]-\textit{mana}[j], \textit{mana}[j-1])$。

那么我们求的是

$$
\max_{i=0}^{n-1} \mathbf{p}\cdot \mathbf{v}_i
$$

根据点积的几何意义，我们求的是 $\mathbf{v}_i$ 在 $\mathbf{p}$ 方向上的投影长度，再乘以 $\mathbf{p}$ 的模长 $||\mathbf{p}||$。由于 $||\mathbf{p}||$ 是个定值，所以要最大化投影长度。

考虑 $\mathbf{v}_i$ 的**上凸包**（用 Andrew 算法计算），在凸包内的点，就像是山坳，比凸包顶点的投影长度短。所以只需考虑凸包顶点。

这样有一个很好的性质：顺时针（或者逆时针）遍历凸包顶点，$\mathbf{p}\cdot \mathbf{v}_i$ 会先变大再变小（单峰函数）。那么要计算最大值，就类似 [852. 山脉数组的峰顶索引](https://leetcode.cn/problems/peak-index-in-a-mountain-array/)，**二分**首个「下坡」的位置，具体见 [我的题解](https://leetcode.cn/problems/peak-index-in-a-mountain-array/solutions/2984800/er-fen-gen-ju-shang-po-huan-shi-xia-po-p-uoev/)。

```py [sol-Python3]
class Vec:
    __slots__ = 'x', 'y'

    def __init__(self, x: int, y: int):
        self.x = x
        self.y = y

    def __sub__(self, b: "Vec") -> "Vec":
        return Vec(self.x - b.x, self.y - b.y)

    def det(self, b: "Vec") -> int:
        return self.x * b.y - self.y * b.x

    def dot(self, b: "Vec") -> int:
        return self.x * b.x + self.y * b.y

class Solution:
    # Andrew 算法，计算 points 的上凸包
    # 由于横坐标（前缀和）是严格递增的，所以无需排序
    def convex_hull(self, points: List[Vec]) -> List[Vec]:
        q = []
        for p in points:
            while len(q) > 1 and (q[-1] - q[-2]).det(p - q[-1]) >= 0:
                q.pop()
            q.append(p)
        return q

    def minTime(self, skill: List[int], mana: List[int]) -> int:
        s = list(accumulate(skill, initial=0))
        vs = [Vec(pre_sum, x) for pre_sum, x in zip(s, skill)]
        vs = self.convex_hull(vs)  # 去掉无用数据

        start = 0
        for pre, cur in pairwise(mana):
            p = Vec(pre - cur, pre)
            # p.dot(vs[i]) 是个单峰函数，二分找最大值
            check = lambda i: p.dot(vs[i]) > p.dot(vs[i + 1])
            i = bisect_left(range(len(vs) - 1), True, key=check)
            start += p.dot(vs[i])
        return start + mana[-1] * s[-1]
```

```java [sol-Java]
class Solution {
    private record Vec(int x, int y) {
        Vec sub(Vec b) {
            return new Vec(x - b.x, y - b.y);
        }

        long det(Vec b) {
            return (long) x * b.y - (long) y * b.x;
        }

        long dot(Vec b) {
            return (long) x * b.x + (long) y * b.y;
        }
    }

    // Andrew 算法，计算 points 的上凸包
    // 由于横坐标（前缀和）是严格递增的，所以无需排序
    private List<Vec> convexHull(Vec[] points) {
        List<Vec> q = new ArrayList<>();
        for (Vec p : points) {
            while (q.size() > 1 && q.getLast().sub(q.get(q.size() - 2)).det(p.sub(q.getLast())) >= 0) {
                q.removeLast();
            }
            q.add(p);
        }
        return q;
    }

    public long minTime(int[] skill, int[] mana) {
        int n = skill.length;
        int[] s = new int[n + 1];
        Vec[] points = new Vec[n];
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + skill[i];
            points[i] = new Vec(s[i], skill[i]);
        }
        List<Vec> vs = convexHull(points); // 去掉无用数据

        int m = mana.length;
        long start = 0;
        for (int j = 1; j < m; j++) {
            Vec p = new Vec(mana[j - 1] - mana[j], mana[j - 1]);
            // p.dot(vs[i]) 是个单峰函数，二分找最大值
            int l = -1, r = vs.size() - 1;
            while (l + 1 < r) {
                int mid = (l + r) >>> 1;
                if (p.dot(vs.get(mid)) > p.dot(vs.get(mid + 1))) {
                    r = mid;
                } else {
                    l = mid;
                }
            }
            start += p.dot(vs.get(r));
        }
        return start + (long) mana[m - 1] * s[n];
    }
}
```

```cpp [sol-C++]
struct Vec {
    int x, y;
    Vec operator-(const Vec& b) { return {x - b.x, y - b.y}; }
    long long det(const Vec& b) { return 1LL * x * b.y - 1LL * y * b.x; }
    long long dot(const Vec& b) { return 1LL * x * b.x + 1LL * y * b.y; }
};

class Solution {
    // Andrew 算法，计算 points 的上凸包
    // 由于横坐标（前缀和）是严格递增的，所以无需排序
    vector<Vec> convex_hull(vector<Vec>& points) {
        vector<Vec> q;
        for (auto& p : points) {
            while (q.size() > 1 && (q.back() - q[q.size() - 2]).det(p - q.back()) >= 0) {
                q.pop_back();
            }
            q.push_back(p);
        }
        return q;
    }

public:
    long long minTime(vector<int>& skill, vector<int>& mana) {
        int n = skill.size(), m = mana.size();
        vector<int> s(n + 1);
        vector<Vec> vs(n);
        for (int i = 0; i < n; i++) {
            s[i + 1] = s[i] + skill[i];
            vs[i] = {s[i], skill[i]};
        }
        vs = convex_hull(vs); // 去掉无用数据

        long long start = 0;
        for (int j = 1; j < m; j++) {
            Vec p = {mana[j - 1] - mana[j], mana[j - 1]};
            // p.dot(vs[i]) 是个单峰函数，二分找最大值
            int l = -1, r = vs.size() - 1;
            while (l + 1 < r) {
                int mid = l + (r - l) / 2;
                (p.dot(vs[mid]) > p.dot(vs[mid + 1]) ? r : l) = mid;
            }
            start += p.dot(vs[r]);
        }
        return start + 1LL * mana[m - 1] * s[n];
    }
};
```

```go [sol-Go]
type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }

// Andrew 算法，计算 points 的上凸包
// 由于横坐标（前缀和）是严格递增的，所以无需排序
func convexHull(points []vec) (q []vec) {
	for _, p := range points {
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) >= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, p)
	}
	return
}

func minTime(skill, mana []int) int64 {
	n, m := len(skill), len(mana)
	s := make([]int, n+1)
	vs := make([]vec, n)
	for i, x := range skill {
		s[i+1] = s[i] + x
		vs[i] = vec{s[i], x}
	}
	vs = convexHull(vs) // 去掉无用数据

	start := 0
	for j := 1; j < m; j++ {
		p := vec{mana[j-1] - mana[j], mana[j-1]}
		// p.dot(vs[i]) 是个单峰函数，二分找最大值
		i := sort.Search(len(vs)-1, func(i int) bool { return p.dot(vs[i]) > p.dot(vs[i+1]) })
		start += p.dot(vs[i])
	}
	return int64(start + mana[m-1]*s[n])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + m\log n)$，其中 $n$ 是 $\textit{skill}$ 的长度，$m$ 是 $\textit{mana}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

- [1840. 最高建筑高度](https://leetcode.cn/problems/maximum-building-height/)

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

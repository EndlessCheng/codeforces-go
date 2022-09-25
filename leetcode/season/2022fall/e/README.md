个人赛五道题目的 [视频讲解](https://www.bilibili.com/video/BV1zN4y1K762) 已出炉，欢迎点赞三连，在评论区分享你对这场比赛的看法~

---

数形结合更好理解，推荐先看视频哦，下面整理了视频中讲的重点。

1. 画折线图，问题转换成最小化折线图中最大值与最小值的差。
2. 定义 $f[i][j]$ 表示考虑 $\textit{operate}$ 的前 $i$ 个数，其中某些数字变成负数后，折线图最右端点到折线图最下端点的纵坐标距离为 $j$ 时，折线图中最大值与最小值的差的最小值。
3. 设 $x=\textit{operate}[i]$，分类讨论（**下面的等号表示左值和右值取 $\min$ 后赋给左值**）：
   - 取正号，折线图往上走：$f[i][j+x] = \max(f[i-1][j],j+x)$；
   - 取负号，折线图往下走，且纵坐标没有小于最下端点的纵坐标：$f[i][j-x] = f[i-1][j]$；
   - 取负号，折线图往下走，且纵坐标小于最下端点的纵坐标，那么产生了一个新的最下端点，按照定义：$f[i][0] = f[i-1][j]-j+x$。
4. 初始值 $f[0][0] = 0$，其余为 $+\infty$。
5. 答案为 $\min(f[n-1])$。
6. 代码实现时，用滚动数组优化空间。

```py [sol1-Python3]
class Solution:
    def unSuitability(self, operate: List[int]) -> int:
        mx = max(operate) * 2
        pre = [0] + [inf] * mx
        for x in operate:
            f = [inf] * (mx + 1)
            for j, dis in enumerate(pre):
                if dis == inf: continue  # 无效的长度（无法组成）
                if j + x <= mx: f[j + x] = min(f[j + x], max(dis, j + x))
                if j >= x: f[j - x] = min(f[j - x], dis)
                else: f[0] = min(f[0], dis - j + x)
            pre = f
        return min(pre)
```

```java [sol1-Java]
class Solution {
    public int unSuitability(int[] operate) {
        var mx = Arrays.stream(operate).max().orElseThrow() * 2 + 1;
        int[] pre = new int[mx], f = new int[mx];
        Arrays.fill(pre, Integer.MAX_VALUE);
        pre[0] = 0;
        for (var x : operate) {
            Arrays.fill(f, Integer.MAX_VALUE);
            for (var j = 0; j < mx; ++j) {
                var dis = pre[j];
                if (dis == Integer.MAX_VALUE) continue; // 无效的长度（无法组成）
                if (j + x < mx) f[j + x] = Math.min(f[j + x], Math.max(dis, j + x));
                if (j >= x) f[j - x] = Math.min(f[j - x], dis);
                else f[0] = Math.min(f[0], dis - j + x);
            }
            var tmp = pre;
            pre = f;
            f = tmp;
        }
        return Arrays.stream(pre).min().orElseThrow();
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int unSuitability(vector<int> &operate) {
        int mx = *max_element(operate.begin(), operate.end()) * 2 + 1;
        int pre[mx], f[mx];
        memset(pre, 0x3f, sizeof(pre));
        pre[0] = 0;
        for (int x : operate) {
            memset(f, 0x3f, sizeof(f));
            for (int j = 0; j < mx; ++j) {
                int dis = pre[j];
                if (dis == 0x3f3f3f3f) continue; // 无效的长度（无法组成）
                if (j + x < mx) f[j + x] = min(f[j + x], max(dis, j + x));
                if (j >= x) f[j - x] = min(f[j - x], dis);
                else f[0] = min(f[0], dis - j + x);
            }
            memcpy(pre, f, sizeof(f));
        }
        return *min_element(pre, pre + mx);
    }
};
```

```go [sol1-Go]
func unSuitability(operate []int) int {
	const inf = math.MaxInt32
	mx := 0
	for _, x := range operate {
		mx = max(mx, x)
	}
	mx *= 2
	pre := make([]int, mx+1)
	for i := range pre {
		pre[i] = inf
	}
	pre[0] = 0
	f := make([]int, mx+1)
	for _, x := range operate {
		for i := range f {
			f[i] = inf
		}
		for j, dis := range pre {
			if pre[j] == inf { // 无效的长度（无法组成）
				continue
			}
			if j+x <= mx {
				f[j+x] = min(f[j+x], max(dis, j+x))
			}
			if j >= x {
				f[j-x] = min(f[j-x], dis)
			} else {
				f[0] = min(f[0], dis-j+x)
			}
		}
		pre, f = f, pre
	}
	ans := inf
	for _, x := range pre {
		ans = min(ans, x)
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(nU)$，其中 $n$ 为 $\textit{operate}$ 的长度，$U=max(\textit{operate})$。
- 空间复杂度：$O(U)$。

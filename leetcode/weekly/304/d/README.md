本题 [视频讲解](https://www.bilibili.com/video/BV1Ba411N78j) 已出炉，包含**思考题**的讲解，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

对于内向基环树的概念和性质，我之前在一篇 [题解](https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/solution/nei-xiang-ji-huan-shu-tuo-bu-pai-xu-fen-c1i1b/) 中作了详细介绍，本文不再赘述（把我那篇题解的代码拿来改一改就能过）。

除了使用那篇题解中的通用做法（拓扑排序）外，我们还可以利用**时间戳**来实现找环的逻辑。

具体来说，初始时间戳 $\textit{clock}=1$，首次访问一个点 $x$ 时，记录访问这个点的时间 $\textit{time}[x]=\textit{clock}$，然后将 $\textit{clock}$ 加一。

如果首次访问一个点，则记录当前时间 $\textit{startTime}=\textit{clock}$，并尝试从这个点出发，看能否找到环。如果找到了一个之前访问过的点 $x$，且访问 $x$ 的时间不早于 $\textit{startTime}$，则说明我们找到了一个**新的**环，此时的环长就是前后两次访问 $x$ 的时间差，即 $\textit{clock}-\textit{time}[x]$。

取所有环长的最大值作为答案。若没有找到环，则返回 $-1$。

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{edges}$ 的长度。
- 空间复杂度：$O(n)$。

#### 思考题

如果题目要你返回最长环上的所有节点呢？（见 [视频讲解](https://www.bilibili.com/video/BV1Ba411N78j)）

```py [sol1-Python3]
class Solution:
    def longestCycle(self, edges: List[int]) -> int:
        time = [0] * len(edges)
        clock, ans = 1, -1
        for x, t in enumerate(time):
            if t: continue
            start_time = clock
            while x >= 0:
                if time[x]:  # 重复访问
                    if time[x] >= start_time:  # 找到了一个新的环
                        ans = max(ans, clock - time[x])
                    break
                time[x] = clock
                clock += 1
                x = edges[x]
        return ans
```

```java [sol1-Java]
class Solution {
    public int longestCycle(int[] edges) {
        int n = edges.length, ans = -1;
        var time = new int[n];
        for (int i = 0, clock = 1; i < n; ++i) {
            if (time[i] > 0) continue;
            for (int x = i, start_time = clock; x >= 0; x = edges[x]) {
                if (time[x] > 0) { // 重复访问
                    if (time[x] >= start_time) // 找到了一个新的环
                        ans = Math.max(ans, clock - time[x]);
                    break;
                }
                time[x] = clock++;
            }
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int longestCycle(vector<int> &edges) {
        int n = edges.size(), time[n], ans = -1;
        memset(time, 0, sizeof(time));
        for (int i = 0, clock = 1; i < n; ++i) {
            if (time[i]) continue;
            for (int x = i, start_time = clock; x >= 0; x = edges[x]) {
                if (time[x]) { // 重复访问
                    if (time[x] >= start_time) // 找到了一个新的环
                        ans = max(ans, clock - time[x]);
                    break;
                }
                time[x] = clock++;
            }
        }
        return ans;
    }
};
```

```go [sol1-Go]
func longestCycle(edges []int) int {
	time := make([]int, len(edges))
	clock, ans := 1, -1
	for x, t := range time {
		if t > 0 {
			continue
		}
		for startTime := clock; x >= 0; x = edges[x] {
			if time[x] > 0 { // 重复访问
				if time[x] >= startTime { // 找到了一个新的环
					ans = max(ans, clock-time[x])
				}
				break
			}
			time[x] = clock
			clock++
		}
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
```

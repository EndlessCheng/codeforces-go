本题 [视频讲解](https://www.bilibili.com/video/BV1Ba411N78j) 已出炉，包含**思考题**的讲解，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

对于内向基环树的概念和性质，我之前在一篇 [题解](https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/solution/nei-xiang-ji-huan-shu-tuo-bu-pai-xu-fen-c1i1b/) 中作了详细介绍，本文不再赘述。

求出 $\textit{node}_1$ 到每个点的距离 $d_1$ 和 $\textit{node}_2$ 到每个点的距离 $d_2$（无法到达时设为一个比较大的数），然后遍历 $d_1$ 和 $d_2$，输出 $\max(d_1[i],d_2[i])$ 的最小值对应的节点编号。若没有这样的节点，输出 $-1$。

由于题目的输入是内向基环树（森林），每个连通块至多有一个环，我们可以用一个简单的循环来求出距离数组。

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{edges}$ 的长度。
- 空间复杂度：$O(n)$。

#### 思考题

1. 如果输入的不止两个节点 $\textit{node}_1$ 和 $\textit{node}_2$，而是一个很长的 $\textit{nodes}$ 列表，要怎么做呢？
2. 如果输入的是 $\textit{queries}$ 询问数组，每个询问包含两个节点 $\textit{node}_1$ 和 $\textit{node}_2$，要你回答 `closestMeetingNode(edges, node1, node2)` 的答案，要怎么做呢？

解答见 [视频讲解](https://www.bilibili.com/video/BV1Ba411N78j)。

```py [sol1-Python3]
class Solution:
    def closestMeetingNode(self, edges: List[int], node1: int, node2: int) -> int:
        n, min_dis, ans = len(edges), len(edges), -1
        def calc_dis(x: int) -> List[int]:
            dis = [n] * n
            d = 0
            while x >= 0 and dis[x] == n:
                dis[x] = d
                d += 1
                x = edges[x]
            return dis
        for i, d1, d2 in zip(range(n), calc_dis(node1), calc_dis(node2)):
            d = max(d1, d2)
            if d < min_dis:
                min_dis, ans = d, i
        return ans
```

```java [sol1-Java]
class Solution {
    public int closestMeetingNode(int[] edges, int node1, int node2) {
        int[] d1 = calcDis(edges, node1), d2 = calcDis(edges, node2);
        int ans = -1, n = edges.length;
        for (int i = 0, minDis = n; i < n; ++i) {
            var d = Math.max(d1[i], d2[i]);
            if (d < minDis) {
                minDis = d;
                ans = i;
            }
        }
        return ans;
    }

    int[] calcDis(int[] edges, int x) {
        var n = edges.length;
        var dis = new int[n];
        Arrays.fill(dis, n);
        for (var d = 0; x >= 0 && dis[x] == n; x = edges[x])
            dis[x] = d++;
        return dis;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    int closestMeetingNode(vector<int> &edges, int node1, int node2) {
        int n = edges.size(), min_dis = n, ans = -1;
        auto calc_dis = [&](int x) -> vector<int> {
            vector<int> dis(n, n);
            for (int d = 0; x >= 0 && dis[x] == n; x = edges[x])
                dis[x] = d++;
            return dis;
        };
        auto d1 = calc_dis(node1), d2 = calc_dis(node2);
        for (int i = 0; i < n; ++i) {
            int d = max(d1[i], d2[i]);
            if (d < min_dis) {
                min_dis = d;
                ans = i;
            }
        }
        return ans;
    }
};
```

```go [sol1-Go]
func closestMeetingNode(edges []int, node1, node2 int) int {
	n := len(edges)
	calcDis := func(x int) []int {
		dis := make([]int, n)
		for i := range dis {
			dis[i] = n
		}
		for d := 0; x >= 0 && dis[x] == n; x = edges[x] {
			dis[x] = d
			d++
		}
		return dis
	}

	d1 := calcDis(node1)
	d2 := calcDis(node2)
	minDis, ans := n, -1
	for i, d := range d1 {
		if d2[i] > d {
			d = d2[i]
		}
		if d < minDis {
			minDis, ans = d, i
		}
	}
	return ans
}
```

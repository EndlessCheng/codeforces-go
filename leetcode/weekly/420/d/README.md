## 前置知识

1. **DFS 时间戳**，见 [DFS 时间戳——处理树上问题的有力工具](https://leetcode.cn/problems/minimum-score-after-removals-on-a-tree/solution/dfs-shi-jian-chuo-chu-li-shu-shang-wen-t-x1kk/)。文章讲的是先序遍历，而本题是后序遍历，调整一下顺序即可。
2. **Manacher 算法**，具体请看 [视频讲解](https://www.bilibili.com/video/BV1UcyYY4EnQ/)，欢迎点赞关注~

## 核心思路

构造 $\textit{dfsStr}$ 的过程是**后序遍历**。

子树的后序遍历字符串，是整棵树的后序遍历字符串的**子串**。

## 算法

1. 后序遍历这棵树，得到从根节点 $0$ 开始的后序遍历的字符串 $\textit{dfsStr}$。
2. 后序遍历的同时，计算每个节点 $i$ 在后序遍历中的开始时间戳和结束时间戳，这也是子树 $i$ 的后序遍历字符串在 $\textit{dfsStr}$ 上的开始下标和结束下标（代码用的左闭右开区间）。
3. 在 $\textit{dfsStr}$ 上跑 Manacher 算法，这样就可以 $\mathcal{O}(1)$ 判断任意子串是否回文了。

**细节**：建图时，由于我们是从左到右遍历 $\textit{parent}$ 数组的，下标 $i$ 是递增的，所以子节点列表一定是升序，所以无需排序。

```py [sol-Python3]
class Solution:
    def findAnswer(self, parent: List[int], s: str) -> List[bool]:
        n = len(parent)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            p = parent[i]
            # 由于 i 是递增的，所以 g[p] 必然是有序的，下面无需排序
            g[p].append(i)

        # dfsStr 是后序遍历整棵树得到的字符串
        dfsStr = [''] * n
        # nodes[i] 表示子树 i 的后序遍历的开始时间戳和结束时间戳+1（左闭右开区间）
        nodes = [[0, 0] for _ in range(n)]
        time = 0

        def dfs(x: int) -> None:
            nonlocal time
            nodes[x][0] = time
            for y in g[x]:
                dfs(y)
            dfsStr[time] = s[x]  # 后序遍历
            time += 1
            nodes[x][1] = time
        dfs(0)

        # Manacher 模板
        # 将 dfsStr 改造为 t，这样就不需要讨论 n 的奇偶性，因为新串 t 的每个回文子串都是奇回文串（都有回文中心）
        # dfsStr 和 t 的下标转换关系：
        # (dfsStr_i+1)*2 = ti
        # ti/2-1 = dfsStr_i
        # ti 为偶数，对应奇回文串（从 2 开始）
        # ti 为奇数，对应偶回文串（从 3 开始）
        t = '#'.join(['^'] + dfsStr + ['$'])

        # 定义一个奇回文串的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余字符串的长度
        # halfLen[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子串的回文半径
        # 即 [i-halfLen[i]+1,i+halfLen[i]-1] 是 t 上的一个回文子串
        halfLen = [0] * (len(t) - 2)
        halfLen[1] = 1
        # boxR 表示当前右边界下标最大的回文子串的右边界下标+1
        # boxM 为该回文子串的中心位置，二者的关系为 r=mid+halfLen[mid]
        boxM = boxR = 0
        for i in range(2, len(halfLen)):
            hl = 1
            if i < boxR:
                # 记 i 关于 boxM 的对称位置 i'=boxM*2-i
                # 若以 i' 为中心的最长回文子串范围超出了以 boxM 为中心的回文串的范围（即 i+halfLen[i'] >= boxR）
                # 则 halfLen[i] 应先初始化为已知的回文半径 boxR-i，然后再继续暴力匹配
                # 否则 halfLen[i] 与 halfLen[i'] 相等
                hl = min(halfLen[boxM * 2 - i], boxR - i)
            # 暴力扩展
            # 算法的复杂度取决于这部分执行的次数
            # 由于扩展之后 boxR 必然会更新（右移），且扩展的的次数就是 boxR 右移的次数
            # 因此算法的复杂度 = O(len(t)) = O(n)
            while t[i - hl] == t[i + hl]:
                hl += 1
                boxM, boxR = i, i + hl
            halfLen[i] = hl

        # t 中回文子串的长度为 hl*2-1
        # 由于其中 # 的数量总是比字母的数量多 1
        # 因此其在 dfsStr 中对应的回文子串的长度为 hl-1
        # 这一结论可用在 isPalindrome 中

        # 判断左闭右开区间 [l,r) 是否为回文串  0<=l<r<=n
        # 根据下标转换关系得到 dfsStr 的 [l,r) 子串在 t 中对应的回文中心下标为 l+r+1
        # 需要满足 halfLen[l + r + 1] - 1 >= r - l，即 halfLen[l + r + 1] > r - l
        def isPalindrome(l: int, r: int) -> bool:
            return halfLen[l + r + 1] > r - l

        return [isPalindrome(l, r) for l, r in nodes]
```

```java [sol-Java]
class Solution {
    private int time = 0;

    public boolean[] findAnswer(int[] parent, String s) {
        int n = parent.length;
        List<Integer>[] g = new ArrayList[n];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (int i = 1; i < n; i++) {
            int p = parent[i];
            // 由于 i 是递增的，所以 g[p] 必然是有序的，下面无需排序
            g[p].add(i);
        }

        // dfsStr 是后序遍历整棵树得到的字符串
        char[] dfsStr = new char[n];
        // nodes[i] 表示子树 i 的后序遍历的开始时间戳和结束时间戳+1（左闭右开区间）
        int[][] nodes = new int[n][2];
        dfs(0, g, s.toCharArray(), dfsStr, nodes);

        // Manacher 模板
        // 将 dfsStr 改造为 t，这样就不需要讨论 n 的奇偶性，因为新串 t 的每个回文子串都是奇回文串（都有回文中心）
        // dfsStr 和 t 的下标转换关系：
        // (dfsStr_i+1)*2 = ti
        // ti/2-1 = dfsStr_i
        // ti 为偶数，对应奇回文串（从 2 开始）
        // ti 为奇数，对应偶回文串（从 3 开始）
        char[] t = new char[n * 2 + 3];
        Arrays.fill(t, '#');
        t[0] = '^';
        for (int i = 0; i < n; i++) {
            t[i * 2 + 2] = dfsStr[i];
        }
        t[n * 2 + 2] = '$';

        // 定义一个奇回文串的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余字符串的长度
        // halfLen[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子串的回文半径
        // 即 [i-halfLen[i]+1,i+halfLen[i]-1] 是 t 上的一个回文子串
        int[] halfLen = new int[t.length - 2];
        halfLen[1] = 1;
        // boxR 表示当前右边界下标最大的回文子串的右边界下标+1
        // boxM 为该回文子串的中心位置，二者的关系为 r=mid+halfLen[mid]
        int boxM = 0;
        int boxR = 0;
        for (int i = 2; i < halfLen.length; i++) { // 循环的起止位置对应着原串的首尾字符
            int hl = 1;
            if (i < boxR) {
                // 记 i 关于 boxM 的对称位置 i'=boxM*2-i
                // 若以 i' 为中心的最长回文子串范围超出了以 boxM 为中心的回文串的范围（即 i+halfLen[i'] >= boxR）
                // 则 halfLen[i] 应先初始化为已知的回文半径 boxR-i，然后再继续暴力匹配
                // 否则 halfLen[i] 与 halfLen[i'] 相等
                hl = Math.min(halfLen[boxM * 2 - i], boxR - i);
            }
            // 暴力扩展
            // 算法的复杂度取决于这部分执行的次数
            // 由于扩展之后 boxR 必然会更新（右移），且扩展的的次数就是 boxR 右移的次数
            // 因此算法的复杂度 = O(len(t)) = O(n)
            while (t[i - hl] == t[i + hl]) {
                hl++;
                boxM = i;
                boxR = i + hl;
            }
            halfLen[i] = hl;
        }

        // t 中回文子串的长度为 hl*2-1
        // 由于其中 # 的数量总是比字母的数量多 1
        // 因此其在 dfsStr 中对应的回文子串的长度为 hl-1
        // 这一结论可用在下面的判断回文串中

        boolean[] ans = new boolean[n];
        for (int i = 0; i < n; i++) {
            // 判断左闭右开区间 [l,r) 是否为回文串  0<=l<r<=n
            // 根据下标转换关系得到 dfsStr 的 [l,r) 子串在 t 中对应的回文中心下标为 l+r+1
            // 需要满足 halfLen[l + r + 1] - 1 >= r - l，即 halfLen[l + r + 1] > r - l
            int l = nodes[i][0];
            int r = nodes[i][1];
            ans[i] = halfLen[l + r + 1] > r - l;
        }
        return ans;
    }

    private void dfs(int x, List<Integer>[] g, char[] s, char[] dfsStr, int[][] nodes) {
        nodes[x][0] = time;
        for (int y : g[x]) {
            dfs(y, g, s, dfsStr, nodes);
        }
        dfsStr[time++] = s[x]; // 后序遍历
        nodes[x][1] = time;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<bool> findAnswer(vector<int>& parent, string s) {
        int n = parent.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; i++) {
            int p = parent[i];
            // 由于 i 是递增的，所以 g[p] 必然是有序的，下面无需排序
            g[p].push_back(i);
        }

        // dfsStr 是后序遍历整棵树得到的字符串
        string dfsStr(n, 0);
        // nodes[i] 表示子树 i 的后序遍历的开始时间戳和结束时间戳+1（左闭右开区间）
        vector<pair<int, int>> nodes(n);
        int time = 0;
        auto dfs = [&](auto&& dfs, int x) -> void {
            nodes[x].first = time;
            for (int y : g[x]) {
                dfs(dfs, y);
            }
            dfsStr[time++] = s[x]; // 后序遍历
            nodes[x].second = time;
        };
        dfs(dfs, 0);

        // Manacher 模板
        // 将 dfsStr 改造为 t，这样就不需要讨论 n 的奇偶性，因为新串 t 的每个回文子串都是奇回文串（都有回文中心）
        // dfsStr 和 t 的下标转换关系：
        // (dfsStr_i+1)*2 = ti
        // ti/2-1 = dfsStr_i
        // ti 为偶数，对应奇回文串（从 2 开始）
        // ti 为奇数，对应偶回文串（从 3 开始）
        string t = "^";
        for (char c : dfsStr) {
            t += '#';
            t += c;
        }
        t += "#$";

        // 定义一个奇回文串的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余字符串的长度
        // halfLen[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子串的回文半径
        // 即 [i-halfLen[i]+1,i+halfLen[i]-1] 是 t 上的一个回文子串
        vector<int> halfLen(t.length() - 2);
        halfLen[1] = 1;
        // boxR 表示当前右边界下标最大的回文子串的右边界下标+1
        // boxM 为该回文子串的中心位置，二者的关系为 r=mid+halfLen[mid]
        int boxM = 0, boxR = 0;
        for (int i = 2; i < halfLen.size(); i++) { // 循环的起止位置对应着原串的首尾字符
            int hl = 1;
            if (i < boxR) {
                // 记 i 关于 boxM 的对称位置 i'=boxM*2-i
                // 若以 i' 为中心的最长回文子串范围超出了以 boxM 为中心的回文串的范围（即 i+halfLen[i'] >= boxR）
                // 则 halfLen[i] 应先初始化为已知的回文半径 boxR-i，然后再继续暴力匹配
                // 否则 halfLen[i] 与 halfLen[i'] 相等
                hl = min(halfLen[boxM * 2 - i], boxR - i);
            }
            // 暴力扩展
            // 算法的复杂度取决于这部分执行的次数
            // 由于扩展之后 boxR 必然会更新（右移），且扩展的的次数就是 boxR 右移的次数
            // 因此算法的复杂度 = O(len(t)) = O(n)
            while (t[i - hl] == t[i + hl]) {
                hl++;
                boxM = i;
                boxR = i + hl;
            }
            halfLen[i] = hl;
        }

        // t 中回文子串的长度为 hl*2-1
        // 由于其中 # 的数量总是比字母的数量多 1
        // 因此其在 dfsStr 中对应的回文子串的长度为 hl-1
        // 这一结论可用在 isPalindrome 中

        // 判断左闭右开区间 [l,r) 是否为回文串  0<=l<r<=n
        // 根据下标转换关系得到 dfsStr 的 [l,r) 子串在 t 中对应的回文中心下标为 l+r+1
        // 需要满足 halfLen[l + r + 1] - 1 >= r - l，即 halfLen[l + r + 1] > r - l
        auto isPalindrome = [&](int l, int r) -> bool {
            return halfLen[l + r + 1] > r - l;
        };

        vector<bool> ans(n);
        for (int i = 0; i < n; i++) {
            ans[i] = isPalindrome(nodes[i].first, nodes[i].second);
        }
        return ans;
    }
};
```

```go [sol-Go]
func findAnswer(parent []int, s string) []bool {
    n := len(parent)
    g := make([][]int, n)
    for i := 1; i < n; i++ {
        p := parent[i]
        // 由于 i 是递增的，所以 g[p] 必然是有序的，下面无需排序
        g[p] = append(g[p], i)
    }

    // dfsStr 是后序遍历整棵树得到的字符串
    dfsStr := make([]byte, n)
    // nodes[i] 表示子树 i 的后序遍历的开始时间戳和结束时间戳+1（左闭右开区间）
    nodes := make([]struct{ begin, end int }, n)
    time := 0
    var dfs func(int)
    dfs = func(x int) {
        nodes[x].begin = time
        for _, y := range g[x] {
            dfs(y)
        }
        dfsStr[time] = s[x] // 后序遍历
        time++
        nodes[x].end = time
    }
    dfs(0)

    // Manacher 模板
    // 将 dfsStr 改造为 t，这样就不需要讨论 n 的奇偶性，因为新串 t 的每个回文子串都是奇回文串（都有回文中心）
    // dfsStr 和 t 的下标转换关系：
    // (dfsStr_i+1)*2 = ti
    // ti/2-1 = dfsStr_i
    // ti 为偶数，对应奇回文串（从 2 开始）
    // ti 为奇数，对应偶回文串（从 3 开始）
    t := append(make([]byte, 0, n*2+3), '^')
    for _, c := range dfsStr {
        t = append(t, '#', c)
    }
    t = append(t, '#', '$')

    // 定义一个奇回文串的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余字符串的长度
    // halfLen[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子串的回文半径
    // 即 [i-halfLen[i]+1,i+halfLen[i]-1] 是 t 上的一个回文子串
    halfLen := make([]int, len(t)-2)
    halfLen[1] = 1
    // boxR 表示当前右边界下标最大的回文子串的右边界下标+1
    // boxM 为该回文子串的中心位置，二者的关系为 r=mid+halfLen[mid]
    boxM, boxR := 0, 0
    for i := 2; i < len(halfLen); i++ { // 循环的起止位置对应着原串的首尾字符
        hl := 1
        if i < boxR {
            // 记 i 关于 boxM 的对称位置 i'=boxM*2-i
            // 若以 i' 为中心的最长回文子串范围超出了以 boxM 为中心的回文串的范围（即 i+halfLen[i'] >= boxR）
            // 则 halfLen[i] 应先初始化为已知的回文半径 boxR-i，然后再继续暴力匹配
            // 否则 halfLen[i] 与 halfLen[i'] 相等
            hl = min(halfLen[boxM*2-i], boxR-i)
        }
        // 暴力扩展
        // 算法的复杂度取决于这部分执行的次数
        // 由于扩展之后 boxR 必然会更新（右移），且扩展的的次数就是 boxR 右移的次数
        // 因此算法的复杂度 = O(len(t)) = O(n)
        for t[i-hl] == t[i+hl] {
            hl++
            boxM, boxR = i, i+hl
        }
        halfLen[i] = hl
    }

    // t 中回文子串的长度为 hl*2-1
    // 由于其中 # 的数量总是比字母的数量多 1
    // 因此其在 dfsStr 中对应的回文子串的长度为 hl-1
    // 这一结论可用在 isPalindrome 中

    // 判断左闭右开区间 [l,r) 是否为回文串  0<=l<r<=n
    // 根据下标转换关系得到 dfsStr 的 [l,r) 子串在 t 中对应的回文中心下标为 l+r+1
    // 需要满足 halfLen[l+r+1]-1 >= r-l，即 halfLen[l+r+1] > r-l
    isPalindrome := func(l, r int) bool { return halfLen[l+r+1] > r-l }

    ans := make([]bool, n)
    for i, p := range nodes {
        ans[i] = isPalindrome(p.begin, p.end)
    }
    return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面字符串题单中的「**三、Manacher 算法**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

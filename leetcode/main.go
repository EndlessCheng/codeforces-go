package main

import (
    . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
    "math/bits"
    "sort"
    "strconv"
    "strings"
)

func main() {
    toBytes := func(g [][]string) [][]byte {
        n, m := len(g), len(g[0])
        bytes := make([][]byte, n)
        for i := range bytes {
            bytes[i] = make([]byte, m)
            for j := range bytes[i] {
                bytes[i][j] = g[i][j][0]
            }
        }
        return bytes
    }
    _ = toBytes
    _ = MustBuildTreeNode
}

// LC 18
func fourSum(a []int, target int) (ans [][]int) {
    sort.Ints(a)
    n := len(a)
    for i := 0; i < n-3 && a[i]+a[i+1]+a[i+2]+a[i+3] <= target; i++ {
        if i > 0 && a[i] == a[i-1] || a[i]+a[n-3]+a[n-2]+a[n-1] < target {
            continue
        }
        for j := i + 1; j < n-2 && a[i]+a[j]+a[j+1]+a[j+2] <= target; j++ {
            if j > i+1 && a[j] == a[j-1] || a[i]+a[j]+a[n-2]+a[n-1] < target {
                continue
            }
            for l, r := j+1, n-1; l < r; {
                if s := a[i] + a[j] + a[l] + a[r]; s == target {
                    ans = append(ans, []int{a[i], a[j], a[l], a[r]})
                    for l++; l < r && a[l] == a[l-1]; l++ {
                    }
                    for r--; l < r && a[r] == a[r+1]; r-- {
                    }
                } else if s < target {
                    l++
                } else {
                    r--
                }
            }
        }
    }
    return
}

// LC 19
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    l, r := dummy, head
    // r 先走 n 步，这样当 r 到末尾时，l 恰好在目标位置之前
    for i := 0; i < n; i++ {
        r = r.Next
    }
    for ; r != nil; r = r.Next {
        l = l.Next
    }
    l.Next = l.Next.Next
    return dummy.Next
}

// LC 31
func nextPermutation(nums []int) {
    n := len(nums)
    i := n - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }
    if i >= 0 {
        j := n - 1
        for j >= 0 && nums[i] >= nums[j] {
            j--
        }
        nums[i], nums[j] = nums[j], nums[i]
    }
    reverse := func(a []int) {
        for i, n := 0, len(a); i < n/2; i++ {
            a[i], a[n-1-i] = a[n-1-i], a[i]
        }
    }
    reverse(nums[i+1:])
}

// LC 37
func solveSudoku(board [][]byte) {
    var line, column [9]uint
    var block [3][3]uint
    var spaces [][2]int

    flip := func(i, j int, digit byte) {
        line[i] ^= 1 << digit
        column[j] ^= 1 << digit
        block[i/3][j/3] ^= 1 << digit
    }

    for i, row := range board {
        for j, b := range row {
            if b != '.' {
                digit := b - '1'
                flip(i, j, digit)
            }
        }
    }

    for {
        modified := false
        for i, row := range board {
            for j, b := range row {
                if b != '.' {
                    continue
                }
                mask := 0x1ff &^ (line[i] | column[j] | block[i/3][j/3])
                if mask&(mask-1) == 0 {
                    digit := byte(bits.TrailingZeros(mask))
                    flip(i, j, digit)
                    board[i][j] = digit + '1'
                    modified = true
                }
            }
        }
        if !modified {
            break
        }
    }

    for i, row := range board {
        for j, b := range row {
            if b == '.' {
                spaces = append(spaces, [2]int{i, j})
            }
        }
    }

    var dfs func(int) bool
    dfs = func(pos int) bool {
        if pos == len(spaces) {
            return true
        }
        i, j := spaces[pos][0], spaces[pos][1]
        mask := 0x1ff &^ (line[i] | column[j] | block[i/3][j/3])
        for ; mask > 0; mask &= mask - 1 {
            digit := byte(bits.TrailingZeros(mask))
            flip(i, j, digit)
            board[i][j] = digit + '1'
            if dfs(pos + 1) {
                return true
            }
            flip(i, j, digit)
        }
        return false
    }
    dfs(0)
}

// LC 39
func combinationSum(a []int, target int) (ans [][]int) {
    b := []int{}
    var f func(p, rest int)
    f = func(p, rest int) {
        if p == len(a) {
            return
        }
        if rest == 0 {
            ans = append(ans, append([]int(nil), b...))
            return
        }
        f(p+1, rest)
        if rest-a[p] >= 0 {
            b = append(b, a[p])
            f(p, rest-a[p])
            b = b[:len(b)-1]
        }
    }
    f(0, target)
    return
}

// LC 40
func combinationSum2(a []int, target int) (ans [][]int) {
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }

    sort.Ints(a)
    var freq [][2]int
    for _, v := range a {
        if freq == nil || v != freq[len(freq)-1][0] {
            freq = append(freq, [2]int{v, 1})
        } else {
            freq[len(freq)-1][1]++
        }
    }

    var b []int
    var f func(p, rest int)
    f = func(p, rest int) {
        if rest == 0 {
            ans = append(ans, append([]int(nil), b...))
            return
        }
        if p == len(freq) || rest < freq[p][0] {
            return
        }
        f(p+1, rest)
        most := min(rest/freq[p][0], freq[p][1])
        for i := 1; i <= most; i++ {
            b = append(b, freq[p][0])
            f(p+1, rest-i*freq[p][0])
        }
        b = b[:len(b)-most]
    }
    f(0, target)
    return
}

// LC 41
func firstMissingPositive(a []int) int {
    n := len(a)
    for i, v := range a {
        for 0 < v && v <= n && v != a[v-1] {
            a[i], a[v-1] = a[v-1], a[i]
            v = a[i]
        }
    }
    for i, v := range a {
        if i+1 != v {
            return i + 1
        }
    }
    return n + 1
}

// LC 47
// 给定一个可包含重复数字的序列，返回所有不重复的全排列
func permuteUnique(nums []int) (ans [][]int) {
    n := len(nums)
    sort.Ints(nums)
    perm := []int{}
    vis := make([]bool, n)
    var f func(int)
    f = func(p int) {
        if p == n {
            ans = append(ans, append([]int(nil), perm...))
            return
        }
        for i, v := range nums {
            if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
                continue
            }
            perm = append(perm, v)
            vis[i] = true
            f(p + 1)
            vis[i] = false
            perm = perm[:len(perm)-1]
        }
    }
    f(0)
    return
}

// LC 52
func totalNQueens(n int) (ans int) {
    var solve func(row, columns, diagonals1, diagonals2 int)
    solve = func(row, columns, diagonals1, diagonals2 int) {
        if row == 1 {
            ans++
            return
        }
        availablePositions := (1<<n - 1) &^ (columns | diagonals1 | diagonals2)
        for availablePositions > 0 {
            position := availablePositions & -availablePositions
            solve(row+1, columns|position, (diagonals1|position)<<1, (diagonals2|position)>>1)
            availablePositions &^= position // 移除该比特位
        }
    }
    solve(0, 0, 0, 0)
    return
}

// LC 75
// 荷兰国旗问题
func sortColors(nums []int) {
    p0, p2 := 0, len(nums)-1
    for i := 0; i <= p2; i++ {
        for ; i <= p2 && nums[i] == 2; p2-- {
            nums[i], nums[p2] = nums[p2], nums[i]
        }
        if nums[i] == 0 {
            nums[i], nums[p0] = nums[p0], nums[i]
            p0++
        }
    }
}

// LC 79
func exist(board [][]byte, word string) bool {
    type pair struct{ x, y int }
    var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    h, w := len(board), len(board[0])
    vis := make([][]bool, h)
    for i := range vis {
        vis[i] = make([]bool, w)
    }
    var f func(i, j, k int) bool
    f = func(i, j, k int) bool {
        if board[i][j] != word[k] {
            return false
        }
        if k == len(word)-1 {
            return true
        }
        vis[i][j] = true
        defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
        for _, d := range dir4 {
            if x, y := i+d.x, j+d.y; 0 <= x && x < h && 0 <= y && y < w && !vis[x][y] {
                if f(x, y, k+1) {
                    return true
                }
            }
        }
        return false
    }
    for i, r := range board {
        for j := range r {
            if f(i, j, 0) {
                return true
            }
        }
    }
    return false
}

// LC 94
// Morris 中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
    for root != nil {
        if root.Left != nil {
            // predecessor 节点表示当前 root 节点向左走一步，然后一直向右走至无法走为止的节点
            predecessor := root.Left
            for predecessor.Right != nil && predecessor.Right != root {
                // 有右子树且没有设置过指向 root，则继续向右走
                predecessor = predecessor.Right
            }
            if predecessor.Right == nil {
                // 将 predecessor 的右指针指向 root，这样后面遍历完左子树 root.Left 后，就能通过这个指向回到 root
                predecessor.Right = root
                // 遍历左子树
                root = root.Left
            } else { // predecessor 的右指针已经指向了 root，则表示左子树 root.Left 已经访问完了
                res = append(res, root.Val)
                // 恢复原样
                predecessor.Right = nil
                // 遍历右子树
                root = root.Right
            }
        } else { // 没有左子树
            res = append(res, root.Val)
            // 若有右子树，则遍历右子树
            // 若没有右子树，则整颗左子树已遍历完，root 会通过之前设置的指向回到这颗子树的父节点
            root = root.Right
        }
    }
    return
}

// LC 99
func recoverTree(root *TreeNode) {
    nodes := []*TreeNode{}
    var f func(o *TreeNode)
    f = func(o *TreeNode) {
        if o == nil {
            return
        }
        f(o.Left)
        nodes = append(nodes, o)
        f(o.Right)
    }
    f(root)
    so := append([]*TreeNode(nil), nodes...)
    sort.Slice(so, func(i, j int) bool { return so[i].Val < so[j].Val })
    do := []*TreeNode{}
    for i, o := range nodes {
        if o.Val != so[i].Val {
            do = append(do, o)
        }
    }
    do[0].Val, do[1].Val = do[1].Val, do[0].Val
}

// LC 106
// 从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 {
        return nil
    }
    rootVal := postorder[len(postorder)-1]
    for i, v := range inorder {
        if v == rootVal {
            return &TreeNode{
                rootVal,
                buildTree(inorder[:i], postorder[:i]),
                buildTree(inorder[i+1:], postorder[i:len(postorder)-1]),
            }
        }
    }
    panic(1)
}

// LC 117, O(1) 空间复杂度
func connect(root *Node) *Node {
    start := root
    for start != nil {
        var nextStart, last *Node
        do := func(cur *Node) {
            if cur == nil {
                return
            }
            if nextStart == nil {
                nextStart = cur
            }
            if last != nil {
                last.Next = cur
            }
            last = cur
        }
        for o := start; o != nil; o = o.Next {
            do(o.Left)
            do(o.Right)
        }
        start = nextStart
    }
    return root
}

// LC 124
func maxPathSum(root *TreeNode) int {
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    ans := int(-1e18)
    var f func(*TreeNode) int
    f = func(o *TreeNode) int {
        if o == nil {
            return -1e18
        }
        l := max(f(o.Left), 0)
        r := max(f(o.Right), 0)
        ans = max(ans, o.Val+l+r)
        return o.Val + max(l, r)
    }
    f(root)
    return ans
}

// LC 127 双向 BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
    wid := map[string]int{}
    g := [][]int{}
    addWord := func(w string) int {
        id, has := wid[w]
        if !has {
            id = len(wid)
            wid[w] = id
            g = append(g, []int{})
        }
        return id
    }
    addEdge := func(w string) int {
        id1 := addWord(w)
        s := []byte(w)
        for i, b := range s {
            s[i] = '*'
            id2 := addWord(string(s))
            g[id1] = append(g[id1], id2)
            g[id2] = append(g[id2], id1)
            s[i] = b
        }
        return id1
    }

    for _, w := range wordList {
        addEdge(w)
    }
    st := addEdge(beginWord)
    end, has := wid[endWord]
    if !has {
        return 0
    }

    const inf int = 1e9
    dst := make([]int, len(wid))
    for i := range dst {
        dst[i] = inf
    }
    dst[st] = 0
    qst := []int{st}

    dend := make([]int, len(wid))
    for i := range dend {
        dend[i] = inf
    }
    dend[end] = 0
    qend := []int{end}

    for len(qst) > 0 && len(qend) > 0 {
        q := qst
        qst = nil
        for _, v := range q {
            if dend[v] < inf {
                return (dst[v]+dend[v])/2 + 1
            }
            for _, w := range g[v] {
                if dst[w] == inf {
                    dst[w] = dst[v] + 1
                    qst = append(qst, w)
                }
            }
        }

        q = qend
        qend = nil
        for _, v := range q {
            if dst[v] < inf {
                return (dst[v]+dend[v])/2 + 1
            }
            for _, w := range g[v] {
                if dend[w] == inf {
                    dend[w] = dend[v] + 1
                    qend = append(qend, w)
                }
            }
        }
    }
    return 0
}

// LC 140
func wordBreak(s string, wordDict []string) (sentences []string) {
    wordSet := map[string]bool{}
    for _, w := range wordDict {
        wordSet[w] = true
    }
    n := len(s)
    dp := make([][][]string, n)
    var f func(int) [][]string
    f = func(p int) [][]string {
        if dp[p] != nil {
            return dp[p]
        }
        res := [][]string{}
        for r := p + 1; r < n; r++ {
            if w := s[p:r]; wordSet[w] {
                for _, words := range f(r) {
                    res = append(res, append([]string{w}, words...))
                }
            }
        }
        if w := s[p:]; wordSet[w] {
            res = append(res, []string{w})
        }
        dp[p] = res
        return res
    }
    for _, words := range f(0) {
        sentences = append(sentences, strings.Join(words, " "))
    }
    return
}

// LC 141 O(1) 判环
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    slow, fast := head, head.Next
    for fast != slow {
        if fast == nil || fast.Next == nil {
            return false
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    return true
}

// LC 142
func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil {
        slow = slow.Next
        if fast.Next == nil {
            return nil
        }
        fast = fast.Next.Next
        if fast == slow {
            p := head
            for p != slow {
                p = p.Next
                slow = slow.Next
            }
            return p
        }
    }
    return nil
}

// LC 148
func mergeList(head1, head2 *ListNode) *ListNode {
    dummyHead := &ListNode{}
    temp, temp1, temp2 := dummyHead, head1, head2
    for temp1 != nil && temp2 != nil {
        if temp1.Val <= temp2.Val {
            temp.Next = temp1
            temp1 = temp1.Next
        } else {
            temp.Next = temp2
            temp2 = temp2.Next
        }
        temp = temp.Next
    }
    if temp1 != nil {
        temp.Next = temp1
    } else if temp2 != nil {
        temp.Next = temp2
    }
    return dummyHead.Next
}

func sort2(head, tail *ListNode) *ListNode {
    if head == nil {
        return head
    }

    if head.Next == tail {
        head.Next = nil
        return head
    }

    slow, fast := head, head
    for fast != tail {
        slow = slow.Next
        fast = fast.Next
        if fast != tail {
            fast = fast.Next
        }
    }

    mid := slow
    return mergeList(sort2(head, mid), sort2(mid, tail))
}

func sortList(head *ListNode) *ListNode {
    return sort2(head, nil)
}

// LC 152
func maxProduct(a []int) int {
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    mi, mx, ans := a[0], a[0], a[0]
    for _, v := range a[1:] {
        mi, mx = min(v, min(v*mi, v*mx)), max(v, max(v*mi, v*mx))
        ans = max(ans, mx)
    }
    return ans
}

// LC 216
func combinationSum3(k int, n int) (ans [][]int) {
    var temp []int
    var dfs func(cur, rest int)
    dfs = func(cur, rest int) {
        // 找到一个答案
        if len(temp) == k && rest == 0 {
            ans = append(ans, append([]int(nil), temp...))
            return
        }
        // 剪枝：跳过的数字过多，后面已经无法选到 k 个数字
        if len(temp)+10-cur < k || rest < 0 {
            return
        }
        // 跳过当前数字
        dfs(cur+1, rest)
        // 选当前数字
        temp = append(temp, cur)
        dfs(cur+1, rest-cur)
        temp = temp[:len(temp)-1]
    }
    dfs(1, n)
    return
}

// LC 222
// O(logn) 求节点个数
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    level := 0
    for node := root; node.Left != nil; node = node.Left {
        level++
    }
    return sort.Search(1<<(level+1), func(k int) bool {
        if k <= 1<<level {
            return false
        }
        bits := 1 << (level - 1)
        node := root
        for node != nil && bits > 0 {
            if bits&k == 0 {
                node = node.Left
            } else {
                node = node.Right
            }
            bits >>= 1
        }
        return node == nil
    }) - 1
}

// LC 226
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    left := invertTree(root.Left)
    right := invertTree(root.Right)
    root.Left = right
    root.Right = left
    return root
}

// LC 235
func lowestCommonAncestor(root, p, q *TreeNode) (ancestor *TreeNode) {
    ancestor = root
    for {
        if p.Val < ancestor.Val && q.Val < ancestor.Val {
            ancestor = ancestor.Left
        } else if p.Val > ancestor.Val && q.Val > ancestor.Val {
            ancestor = ancestor.Right
        } else {
            return
        }
    }
}

// LC 321
func maxSubsequence(a []int, k int) (s []int) {
    for i, v := range a {
        for len(s) > 0 && len(s)+len(a)-1-i >= k && v > s[len(s)-1] {
            s = s[:len(s)-1]
        }
        if len(s) < k {
            s = append(s, v)
        }
    }
    return
}

func lexicographicalLess(a, b []int) bool {
    for i := 0; i < len(a) && i < len(b); i++ {
        if a[i] != b[i] {
            return a[i] < b[i]
        }
    }
    return len(a) < len(b)
}

func greatFirstMerge(a, b []int) []int {
    merged := make([]int, len(a)+len(b))
    for i := range merged {
        if lexicographicalLess(a, b) {
            merged[i], b = b[0], b[1:]
        } else {
            merged[i], a = a[0], a[1:]
        }
    }
    return merged
}

func maxNumber(nums1, nums2 []int, k int) (res []int) {
    start := 0
    if k > len(nums2) {
        start = k - len(nums2)
    }
    for i := start; i <= k && i <= len(nums1); i++ {
        s1 := maxSubsequence(nums1, i)
        s2 := maxSubsequence(nums2, k-i)
        merged := greatFirstMerge(s1, s2)
        if lexicographicalLess(res, merged) {
            res = merged
        }
    }
    return
}

// LC 327 基于求逆序对的思路
func countRangeSum(nums []int, lower, upper int) int {
    var mergeCount func([]int) int
    mergeCount = func(a []int) int {
        n := len(a)
        if n <= 1 {
            return 0
        }

        n1 := append([]int(nil), a[:n/2]...)
        n2 := append([]int(nil), a[n/2:]...)
        cnt := mergeCount(n1) + mergeCount(n2)

        // 统计下标对的数量
        l, r := 0, 0
        for _, v := range n1 {
            for l < len(n2) && n2[l]-v < lower {
                l++
            }
            for r < len(n2) && n2[r]-v <= upper {
                r++
            }
            cnt += r - l
        }

        // n1 和 n2 归并填入 a
        p1, p2 := 0, 0
        for i := range a {
            if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
                a[i] = n1[p1]
                p1++
            } else {
                a[i] = n2[p2]
                p2++
            }
        }
        return cnt
    }

    sum := make([]int, len(nums)+1)
    for i, v := range nums {
        sum[i+1] = sum[i] + v
    }
    return mergeCount(sum)
}

// LC 332
func findItinerary(tickets [][]string) []string {
    g := map[string][]string{}
    for _, p := range tickets {
        g[p[0]] = append(g[p[0]], p[1])
    }
    for _, vs := range g {
        sort.Strings(vs)
    }

    path := make([]string, 0, len(tickets)+1)
    var f func(string)
    f = func(v string) {
        for len(g[v]) > 0 {
            w := g[v][0]
            g[v] = g[v][1:]
            f(w)
        }
        path = append(path, v)
    }
    f("JFK")

    for i, j := 0, len(path)-1; i < j; i++ {
        path[i], path[j] = path[j], path[i]
        j--
    }
    return path
}

// LC 501
func findMode(root *TreeNode) (ans []int) {
    var base, cnt, maxCnt int

    update := func(x int) {
        if x == base {
            cnt++
        } else {
            base, cnt = x, 1
        }
        if cnt == maxCnt {
            ans = append(ans, base)
        } else if cnt > maxCnt {
            maxCnt = cnt
            ans = []int{base}
        }
    }

    var f func(*TreeNode)
    f = func(o *TreeNode) {
        if o == nil {
            return
        }
        f(o.Left)
        update(o.Val)
        f(o.Right)
    }
    f(root)
    return
}

// LC 538 1038
// 反序中序遍历
func convertBST(root *TreeNode) *TreeNode {
    sum := 0
    var f func(*TreeNode)
    f = func(o *TreeNode) {
        if o != nil {
            f(o.Right)
            sum += o.Val
            o.Val = sum
            f(o.Left)
        }
    }
    f(root)
    return root
}

// LC 600 不含连续1的非负整数
func findIntegers(N int) int {
    s := strconv.FormatInt(int64(N), 2)
    n := len(s)
    dp := make([][2]int, n)
    for i := range dp {
        dp[i] = [2]int{-1, -1}
    }
    var f func(p, prevIsOne int, isUpper bool) int
    f = func(p, prevIsOne int, isUpper bool) (res int) {
        if p == n {
            return 1
        }
        if !isUpper {
            dv := &dp[p][prevIsOne]
            if *dv >= 0 {
                return *dv
            }
            defer func() { *dv = res }()
        }
        up := 1
        if isUpper {
            up = int(s[p] & 1)
        }
        res = f(p+1, 0, isUpper && 0 == up)
        if prevIsOne == 0 && up == 1 {
            res += f(p+1, 1, isUpper)
        }
        return
    }
    return f(0, 0, true)
}

// LC 621 任务调度器
func leastInterval(tasks []byte, n int) int {
    cnt := map[byte]int{}
    for _, t := range tasks {
        cnt[t]++
    }

    maxExec, maxExecCnt := 0, 0
    for _, c := range cnt {
        if c > maxExec {
            maxExec, maxExecCnt = c, 1
        } else if c == maxExec {
            maxExecCnt++
        }
    }

    if time := (maxExec-1)*(n+1) + maxExecCnt; time > len(tasks) {
        return time
    }
    return len(tasks)
}

// LC 659 分割数组为长度至少为三且元素连续的子序列
func isPossible(nums []int) bool {
    left := map[int]int{} // 每个数字的剩余个数
    for _, v := range nums {
        left[v]++
    }
    endCnt := map[int]int{} // 以某个数字结尾的连续子序列的个数
    for _, v := range nums {
        if left[v] == 0 {
            continue
        }
        if endCnt[v-1] > 0 { // 若存在以 v-1 结尾的连续子序列，则将 v 加到其末尾
            left[v]--
            endCnt[v-1]--
            endCnt[v]++
        } else if left[v+1] > 0 && left[v+2] > 0 { // 否则，生成一个长度为 3 的连续子序列
            left[v]--
            left[v+1]--
            left[v+2]--
            endCnt[v+2]++
        } else { // 无法生成
            return false
        }
    }
    return true
}

// LC 834
// 返回一个表示节点 i 与其他所有节点距离之和的列表 ans
func sumOfDistancesInTree(n int, edges [][]int) []int {
    g := make([][]int, n)
    for _, e := range edges {
        v, w := e[0], e[1]
        g[v] = append(g[v], w)
        g[w] = append(g[w], v)
    }

    size := make([]int, n)
    sum := make([]int, n)
    var f func(v, fa int) int
    f = func(v, fa int) int {
        sz := 1
        for _, w := range g[v] {
            if w == fa {
                continue
            }
            s := f(w, v)
            sum[v] += sum[w] + s
            sz += s
        }
        size[v] = sz
        return sz
    }
    f(0, -1)

    ans := make([]int, n)
    var f2 func(u, f int)
    f2 = func(v, fa int) {
        ans[v] = sum[v]
        for _, w := range g[v] {
            if w == fa {
                continue
            }
            sumV, sumW := sum[v], sum[w]
            sizeV, sizeW := size[v], size[w]

            sum[v] -= sum[w] + size[w]
            size[v] -= size[w]
            sum[w] += sum[v] + size[v]
            size[w] += size[v]

            f2(w, v)

            sum[v], sum[w] = sumV, sumW
            size[v], size[w] = sizeV, sizeW
        }
    }
    f2(0, -1)
    return ans
}

// LC 968
func minCameraCover(root *TreeNode) int {
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }
    var f func(*TreeNode) (a, b, c int)
    f = func(o *TreeNode) (a, b, c int) {
        if o == nil {
            return 1e9, 0, 0
        }
        la, lb, lc := f(o.Left)
        ra, rb, rc := f(o.Right)
        a = lc + rc + 1
        b = min(a, min(la+rb, ra+lb))
        c = min(a, lb+rb)
        return
    }
    _, ans, _ := f(root)
    return ans
}

// LC 1190
func reverse(a []byte) string {
    for i, n := 0, len(a); i < n/2; i++ {
        a[i], a[n-1-i] = a[n-1-i], a[i]
    }
    return string(a)
}
func reverseParentheses(s string) string {
    l := strings.IndexByte(s, '(')
    if l == -1 {
        return s
    }
    for r, c := l+1, 1; ; r++ {
        if s[r] == '(' {
            c++
        } else if s[r] == ')' {
            if c--; c == 0 {
                return s[:l] + reverse([]byte(reverseParentheses(s[l+1:r]))) + reverseParentheses(s[r+1:])
            }
        }
    }
}

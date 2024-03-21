package copypasta

/* 

## 二叉树练习题

#### 自顶向下

- [112. 路径总和](https://leetcode.cn/problems/path-sum/)
- [129. 求根节点到叶节点数字之和](https://leetcode.cn/problems/sum-root-to-leaf-numbers/)
- [257. 二叉树的所有路径](https://leetcode.cn/problems/binary-tree-paths/)
- [113. 路径总和 II](https://leetcode.cn/problems/path-sum-ii/)
- [437. 路径总和 III](https://leetcode.cn/problems/path-sum-iii/) *前缀和
- [104. 二叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)
- [111. 二叉树的最小深度](https://leetcode.cn/problems/minimum-depth-of-binary-tree/)

#### 自底向上

- [100. 相同的树](https://leetcode.cn/problems/same-tree/)
- [101. 对称二叉树](https://leetcode.cn/problems/symmetric-tree/)
- [110. 平衡二叉树](https://leetcode.cn/problems/balanced-binary-tree/)
- [226. 翻转二叉树](https://leetcode.cn/problems/invert-binary-tree/)
- [104. 二叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)
- [111. 二叉树的最小深度](https://leetcode.cn/problems/minimum-depth-of-binary-tree/)

#### 前序中序后序

- [105. 从前序与中序遍历序列构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)
- [106. 从中序与后序遍历序列构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
- [889. 根据前序和后序遍历构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/)
- 只有每个节点度为 2 或者 0 的时候，前序和后序才能唯一确定一颗二叉树，只有一个子节点是无法确定的，因为无法判断它是左儿子还是右儿子。

#### BFS

- [102. 二叉树的层序遍历](https://leetcode.cn/problems/binary-tree-level-order-traversal/)
- [103. 二叉树的锯齿形层序遍历](https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/)
- [199. 二叉树的右视图](https://leetcode.cn/problems/binary-tree-right-side-view/)
- [513. 找树左下角的值](https://leetcode.cn/problems/find-bottom-left-tree-value/)
- [116. 填充每个节点的下一个右侧节点指针](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/)
- [117. 填充每个节点的下一个右侧节点指针 II](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/)

#### 练习 B（右边数字为难度分）

- [1161. 最大层内元素和](https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/) 1250
- [1448. 统计二叉树中好节点的数目](https://leetcode.cn/problems/count-good-nodes-in-binary-tree/) 1360
- [1302. 层数最深叶子节点的和](https://leetcode.cn/problems/deepest-leaves-sum/) 1388
- [2415. 反转二叉树的奇数层](https://leetcode.cn/problems/reverse-odd-levels-of-binary-tree/) 1431
- [1609. 奇偶树](https://leetcode.cn/problems/even-odd-tree/) 1438
- [1026. 节点与其祖先之间的最大差值](https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/) 1446
- [1110. 删点成林](https://leetcode.cn/problems/delete-nodes-and-return-forest/) 1511
- [865. 具有所有最深节点的最小子树](https://leetcode.cn/problems/smallest-subtree-with-all-the-deepest-nodes/) 1534
- [1367. 二叉树中的链表](https://leetcode.cn/problems/linked-list-in-binary-tree/) 1650
- [863. 二叉树中所有距离为 K 的结点](https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/) 1663
- [987. 二叉树的垂序遍历](https://leetcode.cn/problems/vertical-order-traversal-of-a-binary-tree/) 1676
- [2641. 二叉树的堂兄弟节点 II](https://leetcode.cn/problems/cousins-in-binary-tree-ii/) 1677
- [1372. 二叉树中的最长交错路径](https://leetcode.cn/problems/longest-zigzag-path-in-a-binary-tree/) 1713
- [1080. 根到叶路径上的不足节点](https://leetcode.cn/problems/insufficient-nodes-in-root-to-leaf-paths/) 1805
- [2096. 从二叉树一个节点到另一个节点每一步的方向](https://leetcode.cn/problems/step-by-step-directions-from-a-binary-tree-node-to-another/) 1805

*/

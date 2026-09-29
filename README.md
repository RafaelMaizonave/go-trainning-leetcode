# Treino LeetCode em Go

Repositório pessoal para resolver problemas e acompanhar o progresso. O catálogo (título, dificuldade e tags) segue o índice de [keep-practicing/leetcode-go](https://github.com/keep-practicing/leetcode-go). As soluções ficam nas pastas criadas por você.

Este arquivo é gerado pelo comando `go run ./cmd/leetcode readme`. Os comandos de progresso também o atualizam.

## Comandos

```bash
go run ./cmd/leetcode list
go run ./cmd/leetcode list --difficulty easy --status todo
go run ./cmd/leetcode stats
go run ./cmd/leetcode start 1
go run ./cmd/leetcode done 1
go run ./cmd/leetcode reset 1
go run ./cmd/leetcode review
go run ./cmd/leetcode review 1
```

Pelo Makefile: `make list`, `make start ID=1`, `make done ID=1`, `make reset ID=1`, `make stats` e `make review`.

- **A fazer**: o problema ainda não foi iniciado.
- **Tentando**: `start` criou a pasta da solução.
- **Resolvido**: `done` passou em `go test` e o README da solução tem abordagem e complexidade.
- `reset` volta o status para a fazer e mantém os arquivos da solução.
- `review` lista o que já esfriou. `review <id>` grava a data da revisão.

## Progresso

```
Resolvidos: 1/114

Easy    1/60     █░░░░░░░░░
Medium  0/50     ░░░░░░░░░░
Hard    0/4      ░░░░░░░░░░

Tentando: 0
A fazer: 113
```

## Para revisar

Nenhum problema para revisar.

## Problemas

| ID | Título | Dificuldade | Tags | Status | Solução |
| --- | --- | --- | --- | --- | --- |
| 0001 | [Two Sum](https://leetcode.com/problems/two-sum/) | Easy | array, lookup table | Resolvido | [solução](solutions/0001_two_sum/solution.go) |
| 0002 | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) | Medium | linked list | A fazer | — |
| 0003 | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | Medium | sliding window | A fazer | — |
| 0004 | [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) | Hard | binary search | A fazer | — |
| 0007 | [Reverse Integer](https://leetcode.com/problems/reverse-integer/) | Easy | math | A fazer | — |
| 0009 | [Palindrome Number](https://leetcode.com/problems/palindrome-number/) | Easy | math | A fazer | — |
| 0011 | [Container With Most Water](https://leetcode.com/problems/container-with-most-water/) | Medium | array, double index | A fazer | — |
| 0013 | [Roman to Integer](https://leetcode.com/problems/roman-to-integer/) | Easy | math | A fazer | — |
| 0014 | [Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/) | Easy | — | A fazer | — |
| 0015 | [3Sum](https://leetcode.com/problems/3sum/) | Medium | — | A fazer | — |
| 0017 | [Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) | Medium | tree | A fazer | — |
| 0019 | [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | Medium | linked list | A fazer | — |
| 0020 | [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) | Easy | string, stack | A fazer | — |
| 0021 | [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) | Easy | linked list | A fazer | — |
| 0023 | [Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/) | Hard | linked list, heap | A fazer | — |
| 0024 | [Swap Nodes in Pairs](https://leetcode.com/problems/swap-nodes-in-pairs/) | Medium | linked list | A fazer | — |
| 0025 | [Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/) | Hard | linked list | A fazer | — |
| 0026 | [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) | Easy | array, double index | A fazer | — |
| 0027 | [Remove Element](https://leetcode.com/problems/remove-element/) | Easy | array | A fazer | — |
| 0028 | [Implement strStr()](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/) | Easy | double index | A fazer | — |
| 0033 | [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/) | Medium | binary search | A fazer | — |
| 0034 | [Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/) | Medium | binary search | A fazer | — |
| 0035 | [Search Insert Position](https://leetcode.com/problems/search-insert-position/) | Easy | binary search | A fazer | — |
| 0048 | [Rotate Image](https://leetcode.com/problems/rotate-image/) | Medium | array | A fazer | — |
| 0053 | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) | Easy | dynamic programming | A fazer | — |
| 0058 | [Length of Last Word](https://leetcode.com/problems/length-of-last-word/) | Easy | — | A fazer | — |
| 0061 | [Rotate List](https://leetcode.com/problems/rotate-list/) | Medium | linked list | A fazer | — |
| 0062 | [Unique Paths](https://leetcode.com/problems/unique-paths/) | Medium | recursion, memory search, dynamic programming | A fazer | — |
| 0063 | [Unique Paths II](https://leetcode.com/problems/unique-paths-ii/) | Medium | recursion, memory search, dynamic programming | A fazer | — |
| 0064 | [Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum/) | Medium | dynamic programming, dfs | A fazer | — |
| 0066 | [Plus One](https://leetcode.com/problems/plus-one/) | Easy | math | A fazer | — |
| 0067 | [Add Binary](https://leetcode.com/problems/add-binary/) | Easy | — | A fazer | — |
| 0069 | [Sqrt(x)](https://leetcode.com/problems/sqrtx/) | Easy | binary search | A fazer | — |
| 0070 | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) | Easy | dynamic programming | A fazer | — |
| 0075 | [Sort Colors](https://leetcode.com/problems/sort-colors/) | Medium | sort | A fazer | — |
| 0076 | [Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/) | Hard | sliding window | A fazer | — |
| 0077 | [Combinations](https://leetcode.com/problems/combinations/) | Medium | backtracking, combine | A fazer | — |
| 0079 | [Word Search](https://leetcode.com/problems/word-search/) | Medium | backtracking, array | A fazer | — |
| 0080 | [Remove Duplicates from Sorted Array II](https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/) | Medium | double index | A fazer | — |
| 0082 | [Remove Duplicates from Sorted List II](https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/) | Medium | linked list | A fazer | — |
| 0083 | [Remove Duplicates from Sorted List](https://leetcode.com/problems/remove-duplicates-from-sorted-list/) | Easy | linked list | A fazer | — |
| 0086 | [Partition List](https://leetcode.com/problems/partition-list/) | Medium | linked list | A fazer | — |
| 0088 | [Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/) | Easy | sort | A fazer | — |
| 0092 | [Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/) | Medium | linked list | A fazer | — |
| 0094 | [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/) | Medium | binary tree | A fazer | — |
| 0100 | [Same Tree](https://leetcode.com/problems/same-tree/) | Easy | binary tree | A fazer | — |
| 0101 | [Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) | Easy | stack, recursion, iterative | A fazer | — |
| 0102 | [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) | Medium | binary tree, bfs | A fazer | — |
| 0104 | [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | Easy | binary tree | A fazer | — |
| 0107 | [Binary Tree Level Order Traversal II](https://leetcode.com/problems/binary-tree-level-order-traversal-ii/) | Easy | binary tree | A fazer | — |
| 0111 | [Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/) | Easy | binary tree | A fazer | — |
| 0112 | [Path Sum](https://leetcode.com/problems/path-sum/) | Easy | binary tree | A fazer | — |
| 0120 | [Triangle](https://leetcode.com/problems/triangle/) | Medium | dynamic programming, dfs | A fazer | — |
| 0121 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | Easy | — | A fazer | — |
| 0122 | [Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) | Easy | greedy | A fazer | — |
| 0125 | [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/) | Easy | — | A fazer | — |
| 0136 | [Single Number](https://leetcode.com/problems/single-number/) | Easy | hash table, bit manipulation | A fazer | — |
| 0144 | [Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/) | Medium | binary tree | A fazer | — |
| 0148 | [Sort List](https://leetcode.com/problems/sort-list/) | Medium | sort, linked list | A fazer | — |
| 0150 | [Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/) | Medium | stack | A fazer | — |
| 0153 | [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/) | Medium | binary search | A fazer | — |
| 0155 | [Min Stack](https://leetcode.com/problems/min-stack/) | Easy | stack | A fazer | — |
| 0165 | [Compare Version Numbers](https://leetcode.com/problems/compare-version-numbers/) | Medium | string | A fazer | — |
| 0167 | [Two Sum II - Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) | Easy | two pointers | A fazer | — |
| 0179 | [Largest Number](https://leetcode.com/problems/largest-number/) | Medium | sort | A fazer | — |
| 0198 | [House Robber](https://leetcode.com/problems/house-robber/) | Easy | memory search, dynamic programming | A fazer | — |
| 0200 | [Number of Islands](https://leetcode.com/problems/number-of-islands/) | Medium | dfs, bfs | A fazer | — |
| 0203 | [Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/) | Easy | linked list | A fazer | — |
| 0206 | [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/) | Easy | linked list | A fazer | — |
| 0208 | [Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/) | Medium | trie | A fazer | — |
| 0209 | [Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/) | Medium | sliding window | A fazer | — |
| 0211 | [Add and Search Word - Data Structure Design](https://leetcode.com/problems/design-add-and-search-words-data-structure/) | Medium | trie | A fazer | — |
| 0215 | [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) | Medium | sort | A fazer | — |
| 0217 | [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | Easy | map | A fazer | — |
| 0219 | [Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii/) | Easy | map | A fazer | — |
| 0226 | [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/) | Easy | recursion, binary tree | A fazer | — |
| 0235 | [Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) | Easy | recursion, binary tree | A fazer | — |
| 0236 | [Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/) | Medium | recursion, binary tree | A fazer | — |
| 0237 | [Delete Node in a Linked List](https://leetcode.com/problems/delete-node-in-a-linked-list/) | Easy | linked list | A fazer | — |
| 0257 | [Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/) | Easy | binary tree | A fazer | — |
| 0258 | [Add Digits](https://leetcode.com/problems/add-digits/) | Easy | math | A fazer | — |
| 0283 | [Move Zeroes](https://leetcode.com/problems/move-zeroes/) | Easy | array | A fazer | — |
| 0300 | [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/) | Medium | dp | A fazer | — |
| 0303 | [Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/) | Easy | — | A fazer | — |
| 0304 | [Range Sum Query 2D - Immutable](https://leetcode.com/problems/range-sum-query-2d-immutable/) | Medium | dp | A fazer | — |
| 0307 | [Range Sum Query - Mutable](https://leetcode.com/problems/range-sum-query-mutable/) | Medium | segment tree | A fazer | — |
| 0328 | [Odd Even Linked List](https://leetcode.com/problems/odd-even-linked-list/) | Medium | singly linked list | A fazer | — |
| 0343 | [Integer Break](https://leetcode.com/problems/integer-break/) | Medium | recursion, memory search, dynamic programming | A fazer | — |
| 0344 | [Reverse String](https://leetcode.com/problems/reverse-string/) | Easy | double index | A fazer | — |
| 0345 | [Reverse Vowels of a String](https://leetcode.com/problems/reverse-vowels-of-a-string/) | Easy | double index | A fazer | — |
| 0347 | [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) | Medium | map, heap, array | A fazer | — |
| 0349 | [Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/) | Easy | set | A fazer | — |
| 0350 | [Intersection of Two Arrays II](https://leetcode.com/problems/intersection-of-two-arrays-ii/) | Easy | map | A fazer | — |
| 0376 | [Wiggle Subsequence](https://leetcode.com/problems/wiggle-subsequence/) | Medium | dp | A fazer | — |
| 0392 | [Is Subsequence](https://leetcode.com/problems/is-subsequence/) | Medium | greedy algorithm | A fazer | — |
| 0404 | [Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves/) | Easy | binary tree | A fazer | — |
| 0416 | [Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/) | Medium | dp, 0-1 knapsack problem | A fazer | — |
| 0435 | [Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/) | Medium | dp, 0-1 knapsack problem | A fazer | — |
| 0437 | [Path Sum III](https://leetcode.com/problems/path-sum-iii/) | Easy | binary tree | A fazer | — |
| 0438 | [Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/) | Easy | sliding window | A fazer | — |
| 0447 | [Number of Boomerangs](https://leetcode.com/problems/number-of-boomerangs/) | Easy | — | A fazer | — |
| 0454 | [4Sum II](https://leetcode.com/problems/4sum-ii/) | Medium | — | A fazer | — |
| 0455 | [Assign Cookies](https://leetcode.com/problems/assign-cookies/) | Easy | greedy algorithm | A fazer | — |
| 0557 | [Reverse Words in a String III](https://leetcode.com/problems/reverse-words-in-a-string-iii/) | Easy | string | A fazer | — |
| 0674 | [Longest Continuous Increasing Subsequence](https://leetcode.com/problems/longest-continuous-increasing-subsequence/) | Easy | — | A fazer | — |
| 0677 | [Map Sum Pairs](https://leetcode.com/problems/map-sum-pairs/) | Medium | trie | A fazer | — |
| 0704 | [Binary Search](https://leetcode.com/problems/binary-search/) | Easy | binary search | A fazer | — |
| 0713 | [Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/) | Medium | sliding window | A fazer | — |
| 0717 | [1-bit and 2-bit Characters](https://leetcode.com/problems/1-bit-and-2-bit-characters/) | Easy | — | A fazer | — |
| 0728 | [Self Dividing Numbers](https://leetcode.com/problems/self-dividing-numbers/) | Easy | — | A fazer | — |
| 0735 | [Asteroid Collision](https://leetcode.com/problems/asteroid-collision/) | Medium | stack | A fazer | — |
| 0747 | [Largest Number At Least Twice of Others](https://leetcode.com/problems/largest-number-at-least-twice-of-others/) | Easy | — | A fazer | — |
| 0872 | [Leaf-Similar Trees](https://leetcode.com/problems/leaf-similar-trees/) | Easy | binary tree | A fazer | — |
| 1021 | [Remove Outermost Parentheses](https://leetcode.com/problems/remove-outermost-parentheses/) | Easy | stack | A fazer | — |


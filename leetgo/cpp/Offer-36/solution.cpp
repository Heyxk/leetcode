// Created by k at 2023/06/16 10:44
// https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/

#include <bits/stdc++.h>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;

    Node() {}

    Node(int _val) {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node* _left, Node* _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/
class Solution {
public:
    Node* treeToDoublyList(Node* root) {
        if (!root) return nullptr;
        dfs(root);
        head->left = pre;
        pre->right = head;
        return head;
    }
private:
    Node *pre;
    Node *head;
    void dfs(Node* node){
        if (node == nullptr) return;
        dfs(node->left);
        if (!pre) {
            // 此时为最小的节点, 也是头节点
            head = node;
            pre = node;
        }else{
            node->left = pre;
            pre->right = node;
            pre = node;
        }
        dfs(node->right);
    }
};

// @lc code=end

int main() {
	ios_base::sync_with_stdio(false);
	stringstream out_stream;

	 root; LeetCodeIO::scan(cin, root);

	Solution *obj = new Solution();

	auto res = obj->treeToDoublyList(root);

	LeetCodeIO::print(out_stream, res);
	cout << "
output: " << out_stream.rdbuf() << endl;

	delete obj;
	return 0;
}

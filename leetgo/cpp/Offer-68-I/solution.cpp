// Created by k at 2023/06/20 17:41
// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/

#include <bits/stdc++.h>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
		while(root != nullptr){
			// p q都在root的左子树
			if (root->val > p->val && root->val > q->val){
				root = root->left;
			}else if (root->val < p->val && root->val < q->val){
				// 都在节点的右子树
				root = root->right;
			}else{
				break;
			}
		}
		return root;
    }
};

// @lc code=end

int main() {
	ios_base::sync_with_stdio(false);
	stringstream out_stream;

	TreeNode* root; LeetCodeIO::scan(cin, root);
	int p; LeetCodeIO::scan(cin, p);
	int q; LeetCodeIO::scan(cin, q);

	Solution *obj = new Solution();

	auto res = obj->lowestCommonAncestor(root, p, q);

	LeetCodeIO::print(out_stream, res);
	cout << "
output: " << out_stream.rdbuf() << endl;

	delete obj;
	return 0;
}

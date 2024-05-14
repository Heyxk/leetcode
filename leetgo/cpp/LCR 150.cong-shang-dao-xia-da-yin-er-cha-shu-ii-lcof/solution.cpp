// Created by k at 2024/05/14 15:05
// leetgo: dev
// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/

#include <bits/stdc++.h>
#include <vector>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Solution {
  public:
    vector<vector<int>> decorateRecord(TreeNode *root)
    {
        vector<vector<int>> ans;
        if (root == nullptr) {
            return ans;
        }
        vector<TreeNode *> queue = {root};
        while (!queue.empty()) {
            int l = queue.size();
            vector<int> lans;
            for (int i = 0; i < l; i++) {
                lans.push_back(queue[i]->val);
                if (queue[i]->left != nullptr) {
                    queue.push_back(queue[i]->left);
                }
                if (queue[i]->right != nullptr) {
                    queue.push_back(queue[i]->right);
                }
            }
            ans.push_back(lans);
            queue.erase(queue.cbegin(), queue.cbegin() + l);
        }
        return ans;
    }
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    TreeNode *root;
    LeetCodeIO::scan(cin, root);

    Solution *obj = new Solution();
    auto res      = obj->decorateRecord(root);
    LeetCodeIO::print(out_stream, res);
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}

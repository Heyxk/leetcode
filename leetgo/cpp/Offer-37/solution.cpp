// Created by k at 2023/06/29 14:29
// https://leetcode.cn/problems/xu-lie-hua-er-cha-shu-lcof/

#include <bits/stdc++.h>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class Codec {
public:

   // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        // bfs
        string res;
        if (!root) return res;
        queue<TreeNode*> q;
        q.push(root); 

        while (!q.empty()) {
            TreeNode* front = q.front();
            q.pop();
            if (front) {
                res += to_string(front->val) + ",";
                q.push(front->left);
                q.push(front->right);
            }
            else res += "null,";
        }
        return res.substr(0, res.size() - 1);
    }

    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        if (data.empty()) return nullptr;
        vector<string> s = split(data);
        
        queue<TreeNode*> q;
        TreeNode* root = new TreeNode(stoi(s[0]));
        q.push(root);
        int i = 1;

        while (!q.empty()) {
            TreeNode* front = q.front();
            q.pop();
            if (s[i] != "null") {
                front->left = new TreeNode(stoi(s[i]));
                q.push(front->left);
            }
            ++i;
            if (s[i] != "null") {
                front->right = new TreeNode(stoi(s[i]));
                q.push(front->right);
            }
            ++i;
        }
        return root;
    }

    vector<string> split(string& s) {
        vector<string> res;
        int n = s.size();
        int i = 0;

        while (i < n) {
            int j = i + 1;
            while (j < n && s[j] != ',') ++j;
            res.push_back(s.substr(i, j - i));
            i = j + 1;
        }
        return res;
    }
};

// Your Codec object will be instantiated and called as such:
// Codec codec;
// codec.deserialize(codec.serialize(root));

// @lc code=end

int main() {
	ios_base::sync_with_stdio(false);
	stringstream out_stream;

	TreeNode* root; LeetCodeIO::scan(cin, root);

	Solution *obj = new Solution();

	auto res = obj->Codec(root);

	LeetCodeIO::print(out_stream, res);
	cout << "
output: " << out_stream.rdbuf() << endl;

	delete obj;
	return 0;
}

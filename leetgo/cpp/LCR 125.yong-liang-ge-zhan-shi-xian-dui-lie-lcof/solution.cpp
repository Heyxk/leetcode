// Created by k at 2024/05/13 15:08
// leetgo: dev
// https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

#include <bits/stdc++.h>
#include <vector>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class CQueue {
  public:
    CQueue() {}

    void appendTail(int value) { tail.push_back(value); }

    int deleteHead()
    {
        if (head.empty() && tail.empty()) {
            return -1;
        }
        if (head.empty()) {
            head.insert(head.cbegin(), tail.cbegin(), tail.cend());
            tail.clear();
        }
        int ans = head.front();
        head.erase(head.begin());
        return ans;
    }
    vector<int> head, tail;
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    vector<string> method_names;
    LeetCodeIO::scan(cin, method_names);

    CQueue *obj;
    const unordered_map<string, function<void()>> methods = {
        {"CQueue",
         [&]() {
             cin.ignore();
             obj = new CQueue();
             out_stream << "null,";
         }},
        {"appendTail",
         [&]() {
             int value;
             LeetCodeIO::scan(cin, value);
             cin.ignore();
             obj->appendTail(value);
             out_stream << "null,";
         }},
        {"deleteHead",
         [&]() {
             cin.ignore();
             LeetCodeIO::print(out_stream, obj->deleteHead());
             out_stream << ',';
         }},
    };
    cin >> ws;
    out_stream << '[';
    for (auto &&method_name : method_names) {
        cin.ignore(2);
        methods.at(method_name)();
    }
    cin.ignore();
    out_stream.seekp(-1, ios_base::end);
    out_stream << ']';
    cout << "\noutput: " << out_stream.rdbuf() << endl;

    delete obj;
    return 0;
}

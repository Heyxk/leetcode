// Created by k at 2024/05/13 14:10
// leetgo: dev
// https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/

#include <bits/stdc++.h>
#include <vector>
#include "LC_IO.h"
using namespace std;

// @lc code=begin

class MinStack {
  public:
    /** initialize your data structure here. */
    MinStack()
    {
        stack.reserve(0);
        minStack.reserve(0);
    }

    void push(int x)
    {
        stack.push_back(x);
        if (minStack.empty()) {
            minStack.push_back(x);
        }
        else if (minStack.back() >= x) {
            minStack.push_back(x);
        }
    }

    void pop()
    {
        int x = stack.back();
        stack.pop_back();
        if (minStack.back() == x) {
            minStack.pop_back();
        }
    }

    int top() { return stack.back(); }

    int getMin() { return minStack.back(); }
    vector<int> stack;
    vector<int> minStack;
};

// @lc code=end

int main()
{
    ios_base::sync_with_stdio(false);
    stringstream out_stream;

    vector<string> method_names;
    LeetCodeIO::scan(cin, method_names);

    MinStack *obj;
    const unordered_map<string, function<void()>> methods = {
        {"MinStack",
         [&]() {
             int maxSize;
             LeetCodeIO::scan(cin, maxSize);
             cin.ignore();
             obj = new MinStack();
             out_stream << "null,";
         }},
        {"push",
         [&]() {
             int x;
             LeetCodeIO::scan(cin, x);
             cin.ignore();
             obj->push(x);
             out_stream << "null,";
         }},
        {"pop",
         [&]() {
             cin.ignore();
             obj->pop();
             out_stream << "null,";
         }},
        {"top",
         [&]() {
             cin.ignore();
             LeetCodeIO::print(out_stream, obj->top());
             out_stream << ',';
         }},
        {"getMin",
         [&]() {
             cin.ignore();
             LeetCodeIO::print(out_stream, obj->getMin());
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

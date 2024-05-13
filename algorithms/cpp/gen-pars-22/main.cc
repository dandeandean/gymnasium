// Solution from:
// https://github.com/neetcode-gh/leetcode/blob/main/cpp/0022-generate-parentheses.cpp
#include <iostream>
#include <string>
using namespace std;
class Solution {
public:
  vector<string> generateParenthesis(int n) {
    vector<string> result;
    generate(n, 0, 0, "", result);
    return result;
  }

private:
  void generate(int n, int open, int close, string str,
                vector<string> &result) {
    if (open == n && close == n) {
      result.push_back(str);
      return;
    }
    if (open < n) {
      generate(n, open + 1, close, str + '(', result);
    }
    if (open > close) {
      generate(n, open, close + 1, str + ')', result);
    }
  }
};

int main() {
  Solution *s = new Solution;
  vector<string> answer = s->generateParenthesis(8);
  int count = 0;
  while (!answer.empty()) {
    cout << answer.back();
    answer.pop_back();
    count++;
  }
  cout << count;
  return 0;
}

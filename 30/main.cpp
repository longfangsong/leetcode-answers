#include <algorithm>
#include <iostream>
#include <map>
#include <string>
#include <vector>
using namespace std;
class Solution {
public:
  vector<int> findSubstring(string s, vector<string> &words) {
    if (s == "" || words.size() == 0)
      return {};
    const size_t sliding_window_length = words[0].size() * words.size();
    map<string, size_t> initial_word_require;
    vector<int> result;
    for (auto &&word : words) {
      ++initial_word_require[word];
    }
    for (size_t i = 0; i < s.size() - sliding_window_length + 1; ++i) {
      auto word_require = initial_word_require;
      auto this_window = s.substr(i, sliding_window_length);
      for (size_t j = 0; j < this_window.size(); j += words[0].size()) {
        auto this_word = this_window.substr(j, words[0].size());
        if (word_require[this_word] == 0) {
          break;
        } else {
          --word_require[this_word];
        }
      }
      if (all_of(word_require.begin(), word_require.end(),
                 [](const pair<string, size_t> &item) {
                   return item.second == 0;
                 })) {
        result.push_back(i);
      }
    }
    return result;
  }
};

int main() {
  Solution s;
  vector<string> v({"foo", "bar"});
  auto result = s.findSubstring("barfoothefoobarman", v);
  for (auto &&i : result) {
    cout << i << ',';
  }
  return 0;
}
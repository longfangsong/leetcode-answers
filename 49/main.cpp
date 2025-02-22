#include <algorithm>
#include <string>
#include <unordered_map>
#include <vector>
using namespace std;

class Solution {
public:
  vector<vector<string>> groupAnagrams(vector<string> &strs) {
    unordered_map<string, vector<string>> groups;
    for (auto &str : strs) {
      string s = str;
      sort(s.begin(), s.end());
      groups[s].push_back(str);
    }
    vector<vector<string>> result;
    for (auto &[key, value] : groups) {
      result.push_back(value);
    }
    return result;
  }
};

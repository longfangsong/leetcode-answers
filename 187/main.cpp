#include <bitset>
#include <string>
#include <unordered_set>
#include <vector>

using namespace std;

class Solution {
public:
  void append_char(bitset<20> &bits, char ch) {
    bits <<= 2;
    switch (ch) {
    case 'A':
      break;
    case 'C':
      bits.set(0, 1);
      break;
    case 'G':
      bits.set(1, 1);
      break;
    case 'T':
      bits.set(0, 1);
      bits.set(1, 1);
      break;
    default:
      break;
    }
  }
  vector<string> findRepeatedDnaSequences(string s) {
    unordered_set<bitset<20>> pattern_exists;
    unordered_set<bitset<20>> pattern_consumed;
    vector<string> result;
    if (s.length() <= 10)
      return vector<string>();
    bitset<20> current;
    size_t i;
    for (i = 0; i < 10; ++i) {
      append_char(current, s[i]);
    }
    for (; i < s.length(); ++i) {
      if (pattern_exists.find(current) != pattern_exists.end() &&
          pattern_consumed.find(current) == pattern_consumed.end()) {
        result.push_back(s.substr(i - 10, 10));
        pattern_consumed.insert(current);
      } else {
        pattern_exists.insert(current);
      }
      append_char(current, s[i]);
    }
    if (pattern_exists.find(current) != pattern_exists.end() &&
        pattern_consumed.find(current) == pattern_consumed.end()) {
      result.push_back(s.substr(i - 10, 10));
    }
    return vector<string>(result.begin(), result.end());
  }
};
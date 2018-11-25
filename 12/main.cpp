#include <functional>
#include <iostream>
#include <map>
#include <string>
using namespace std;

class Solution {
 public:
  string intToRoman(int num) {
    static const map<int, string, std::greater<int>> dictionary{
        {1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"},
        {90, "XC"},  {50, "L"},   {40, "XL"}, {10, "X"},   {9, "IX"},
        {5, "V"},    {4, "IV"},   {1, "I"}};
    string result;
    for (auto&& values : dictionary) {
      while (num >= values.first) {
        result += values.second;
        num -= values.first;
      }
    }
    return result;
  }
};

int main(int argc, char const* argv[]) {
  Solution s;
  cout << s.intToRoman(1994) << endl;
  return 0;
}

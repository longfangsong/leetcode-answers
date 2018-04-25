#include <string>
#include <map>
#include <algorithm>

using std::string;
using std::map;
using std::for_each;

class Solution {
public:
    char findTheDifference(string s, string t) {
        map<char, size_t> used_map_s;
        map<char, size_t> used_map_t;
        for_each(s.begin(),s.end(),[&](const char& ch) {
            ++used_map_s[ch];
        });
        for_each(t.begin(),t.end(),[&](const char& ch) {
            ++used_map_t[ch];
        });
        for(char ch = 'a';ch<='z';++ch) {
            if(used_map_s[ch] != used_map_t[ch]) {
                return ch;
            }
        }
        return ' ';
    }
};
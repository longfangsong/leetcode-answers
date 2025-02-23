#include <vector>
using namespace std;

class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        if (nums.empty()) { return 0; }
        sort(nums.begin(), nums.end());
        size_t max_size = 1;
        size_t current_size = 1;
        for(size_t i = 1; i<nums.size(); ++i) {
            if(nums[i-1] + 1 == nums[i]) {
                ++current_size;
            } else if (nums[i-1] != nums[i]) {
                if (current_size > max_size) {
                    max_size = current_size;
                }
                current_size = 1;
            }
        }
        if (current_size > max_size) {
            max_size = current_size;
        }
        return max_size;
    }
};

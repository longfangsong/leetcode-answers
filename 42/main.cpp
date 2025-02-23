#include <vector>
using namespace std;

class Solution {
public:
    int trap(vector<int>& height) {
        int current_left_highest = 0;
        vector<int> left_highest(height.size());
        int current_right_highest = 0;
        vector<int> right_highest(height.size());
        int i;
        for (i = 0; i < height.size() && height[i] == 0; ++i) {}
        for (/* reuse i */; i < height.size(); ++i) {
            if(current_left_highest < height[i]) {
                current_left_highest = height[i];
            }
            left_highest[i] = current_left_highest;
        }
        for (i = height.size()-1; i >= 0 && height[i] == 0; --i) {}
        for (/* reuse i */; i >= 0; --i) {
            if(current_right_highest < height[i]) {
                current_right_highest = height[i];
            }
            right_highest[i] = current_right_highest;
        }

        int result = 0;
        for (i = 0; i < left_highest.size(); ++i) {
            int left_height= left_highest[i];
            int right_hight = right_highest[i];
            if (left_height != 0 && right_hight != 0) {
                int water_height = min(left_height, right_hight);
                int water_amount = water_height - height[i];
                result += water_amount;
            }
        }
        return result;
    }
};

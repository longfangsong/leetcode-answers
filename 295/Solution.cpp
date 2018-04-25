#include <queue>
#include <vector>
#include <functional>
class MedianFinder {
private:
    std::priority_queue<int,std::vector<int>,std::greater<int>> big;
    std::priority_queue<int> small;
public:
    /** initialize your data structure here. */
    MedianFinder() {
        
    }
    
    void addNum(int num) {
        if(small.empty()) {
            small.push(num);
        } else if(big.empty()) {
            if(num < small.top()) {
                big.push(small.top());
                small.pop();
                small.push(num);
            } else {
                big.push(num);
            }
        } else {
            if(num > findMedian()) {
                big.push(num);
                if(big.size()>small.size() + 1) {
                    small.push(big.top());
                    big.pop();
                }
            } else {
                small.push(num);
                if(small.size()>big.size() + 1) {
                    big.push(small.top());
                    small.pop();
                }
            }
        }
    }
    
    double findMedian() {
        if(small.size() < big.size()) {
            return big.top();
        } else if(small.size() > big.size()) {
            return small.top();
        } else {
            return (small.top() + big.top()) / 2.0;
        }
    }
};
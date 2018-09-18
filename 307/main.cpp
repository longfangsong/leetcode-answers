#include <algorithm>
#include <vector>
#include <iostream>

using namespace std;

template<typename T>
class segment_tree {
    static_assert(is_integral<T>::value,
                  "Only integral types can be stored in a segment_tree!");

private:
    vector<T> tree;
    vector<T> origin;

    static inline size_t left_child_index(size_t parent_index) {
        return parent_index * 2 + 1;
    }

    static inline size_t right_child_index(size_t parent_index) {
        return parent_index * 2 + 2;
    }

    static inline size_t parent_index(size_t child_index) {
        return (child_index - 1) / 2;
    }

    void construct(size_t node_index, size_t start_index, size_t end_index) {
        if (start_index == end_index) {
            tree[node_index] = origin[start_index];
        } else {
            size_t mid_index = (start_index + end_index) / 2;
            construct(left_child_index(node_index), start_index, mid_index);
            construct(right_child_index(node_index), mid_index + 1, end_index);
            tree[node_index] = tree[left_child_index(node_index)] +
                               tree[right_child_index(node_index)];
        }
    }

    T getSumForRange(size_t current_start, size_t current_end,
                     size_t target_start, size_t target_end,
                     size_t current_node_index) {
        if (current_start > target_end || current_end < target_start) {
            return T(0);
        } else if (target_start <= current_start && current_end <= target_end) {
            return tree[current_node_index];
        } else {
            size_t current_mid = (current_start + current_end) / 2;
            return getSumForRange(current_start, current_mid, target_start,
                                  target_end, left_child_index(current_node_index)) +
                   getSumForRange(current_mid + 1, current_end, target_start,
                                  target_end, right_child_index(current_node_index));
        }
    }

public:
    template<typename InputIter>
    segment_tree(InputIter it1, InputIter it2)
            : origin(it1, it2), tree(distance(it1, it2) * 4) {
        if (!origin.empty())
            construct(0, 0, origin.size() - 1);
    }

    explicit segment_tree(vector<int> &&nums) : origin(nums), tree(nums.size() * 4) {
        if (!origin.empty())
            construct(0, 0, origin.size() - 1);
    }

    void update(size_t index_in_origin, T new_num) {
        assert(index_in_origin < origin.size());
        T delta = new_num - origin[index_in_origin];
        origin[index_in_origin] = new_num;
        size_t current_index = 0;
        size_t current_range_from = 0;
        size_t current_range_to = origin.size() - 1;
        while (current_range_from != current_range_to && left_child_index(current_index) < tree.size()) {
            tree[current_index] += delta;
            size_t current_range_mid = (current_range_from + current_range_to) / 2;
            if (index_in_origin <= current_range_mid) {
                current_range_to = current_range_mid;
                current_index = left_child_index(current_index);
            } else {
                current_range_from = current_range_mid + 1;
                current_index = right_child_index(current_index);
            }
        }
        tree[current_index] += delta;
    }

    T getSum(size_t target_start, size_t target_end) {
        return getSumForRange(0, origin.size() - 1, target_start, target_end, 0);
    }
};

class NumArray {
private:
    segment_tree<int> st;

public:
    NumArray(vector<int> nums) : st(move(nums)) {}

    void update(int i, int val) { st.update(i, val); }

    int sumRange(int i, int j) { return st.getSum(i, j); }
};

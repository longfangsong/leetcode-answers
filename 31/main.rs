fn next_permutation(nums: &mut Vec<i32>) {
    let mut i=0;
    for i_ in (0..nums.len()).rev() {
        i = i_;
        if i_ != 0 && nums[i] > nums[i-1] {
            break;
        }
    }
    if i == 0 {
        nums.reverse();
        return;
    }
    let mut j=0;
    for j_ in (i..nums.len()).rev() {
        j = j_;
        if nums[j_] > nums[i-1] {
            break;
        }
    }
    nums.swap(i-1,j);
    let (mut useless,rest) = nums.split_at_mut(i);
    rest.reverse();
}
    
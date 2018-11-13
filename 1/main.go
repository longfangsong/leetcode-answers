package main

func twoSum(nums []int, target int) []int {
	elements := map[int]int{}
	for index, num := range nums {
		if priviousIndex, has := elements[target-num]; has {
			return []int{priviousIndex, index}
		}
		elements[num] = index
	}
	return []int{}
}

func main() {

}

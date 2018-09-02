import "math"

func majorityElement(nums []int) []int {
	counter1 := 0
	counter2 := 0
	num1 := math.MaxInt32
	num2 := math.MaxInt32
	for _, number := range nums {
		if number == num1 {
			counter1++
		} else if number == num2 {
			counter2++
		} else if counter1 == 0 {
			counter1++
			num1 = number
		} else if counter2 == 0 {
			counter2++
			num2 = number
		} else {
			counter1--
			counter2--
		}
	}
	counter1 = 0
	counter2 = 0
	for _, number := range nums {
		if number == num1 {
			counter1++
		} else if number == num2 {
			counter2++
		}
	}
	var result []int
	if counter1 > len(nums)/3 {
		result = append(result, num1)
	}
	if counter2 > len(nums)/3 {
		result = append(result, num2)
	}
	return result
}

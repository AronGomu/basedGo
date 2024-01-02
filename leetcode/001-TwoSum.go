import "fmt"

func main() {
	twoSums([]int{3, 2, 4}, 6)
}

func twoSum(nums []int, target int) []int {
	a := 0
	b := 1
	found := false

	for found != true {
		for i := b; i < len(nums); i++ {
			// fmt.Printf("a = %v, i = %v\n", a, i)
			// fmt.Printf("%v + %v == %v => %v \n\n", nums[a], nums[i], target, nums[a]+nums[i] == target)
			if nums[a]+nums[i] == target {
				// fmt.Printf("is true \n")
				return []int{a, i}
			}
		}
		a += 1
		b += 1
		// fmt.Printf("a = %v, b = %v\n", a, b)

		if b >= len(nums) {
			found = true
		}
	}

	return []int{0, 0}
}

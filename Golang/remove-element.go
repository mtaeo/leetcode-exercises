package main

func removeElement(nums []int, val int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}

func main() {
	print(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

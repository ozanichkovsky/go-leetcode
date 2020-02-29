package p1twosum

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i := 0; i < len(nums) - 1; i++ {
		complement := target - nums[i]

		if v, ok := hash[complement]; ok {
			return []int{i, v}
		}

		hash[nums[i]] = i
	}

	return nil
}

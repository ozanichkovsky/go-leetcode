package p1twosum

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == target - nums[j] {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, n := range nums {
		hash[n] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if v, ok := hash[complement]; ok && v != i {
			return []int{i, v}
		}
	}

	return nil
}

func twoSum3(nums []int, target int) []int {
	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if v, ok := hash[complement]; ok {
			return []int{i, v}
		}

		hash[nums[i]] = i
	}

	return nil
}

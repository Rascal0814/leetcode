package _83_move_zeroes

func moveZeroes(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == 0 && nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

}

func moveZeroesPoint(nums []int) {
	f, s := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[f] != 0 {
			nums[s] = nums[f]
			s++
		}
		f++
	}

	for s < len(nums) {
		nums[s] = 0
		s++
	}
}

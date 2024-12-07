package main

import (
	"strconv"
	"strings"
)

func sol2(lines []string) int64 {

	var ret int64
	ret = 0

	for _, line := range lines {

		var sum int64
		var nums []int64

		tmp_str := strings.Split(line, ": ")

		sum, _ = strconv.ParseInt(tmp_str[0], 10, 64)

		for _, val := range strings.Split(tmp_str[1], " ") {
			v, _ := strconv.ParseInt(val, 10, 64)
			nums = append(nums, v)
		}

		plus_possible := is_possible2(sum, nums, nums[0], 1, 0)
		mul_possible := is_possible2(sum, nums, nums[0], 1, 1)
		concat_possible := is_possible2(sum, nums, nums[0], 1, 2)

		if plus_possible || mul_possible || concat_possible {

			ret += sum
		}
	}

	return ret

}

func is_possible2(target_sum int64, nums []int64, running_total int64, nums_pos, operation int) bool {

	switch operation {
	case 0:
		running_total += nums[nums_pos]
	case 1:
		running_total *= nums[nums_pos]
	case 2:
		str_num := strconv.Itoa(int(running_total)) + strconv.Itoa(int(nums[nums_pos]))
		num, _ := strconv.Atoi(str_num)
		running_total = int64(num)

	}

	if running_total > target_sum {
		return false
	}

	if nums_pos == len(nums)-1 {
		if target_sum == running_total {
			return true
		} else {
			return false
		}
	}

	return is_possible2(target_sum, nums, running_total, nums_pos+1, 0) || is_possible2(target_sum, nums, running_total, nums_pos+1, 1) || is_possible2(target_sum, nums, running_total, nums_pos+1, 2)

}

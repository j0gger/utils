// Contains mathematical utils
package math

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}


func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Gcd(a, b int) int {
	mx := Max(a, b)
	mn := Min(a, b)

	for mn != 0 {
		tmp := mn
		mn = mx % mn
		mx = tmp
	}

	return mx
}


func Lcm(nums []int) int {
	var lcm int = nums[0]

	for i := 1; i < len(nums); i += 1 {
		lcm = lcm * nums[i] / Gcd(lcm, nums[i])
	}

	return lcm
}

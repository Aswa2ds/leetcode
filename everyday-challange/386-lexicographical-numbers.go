package everydaychallange

func lexicalOrder(n int) []int {
	result := make([]int, n)
	num := 1

	for i := range result {
		result[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return result
}

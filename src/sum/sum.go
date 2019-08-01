package sum

func Sum(numbers []int) (sum int) {
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func SumAllTails(slicesToSum ...[]int) []int {
	var sumSlice []int
	for _, numbers := range slicesToSum {
		if len(numbers) > 0 {
			tail := numbers[1:]
			sumSlice = append(sumSlice, Sum(tail))
		} else {
			sumSlice = append(sumSlice, 0)
		}
	}
	return sumSlice
}

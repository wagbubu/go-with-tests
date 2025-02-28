package main

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	for i, numbers := range numbersToSum {
		if numbers != nil {
			sums[i] = Sum(numbers[1:])
		} else {
			sums[i] = 0
		}

	}
	return sums
}

func main() {

}

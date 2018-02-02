package basic

func getNumbers(numbers []int, fn func(num int) bool) []int{
	newNumbers := make([]int, 0)
	for _, number := range numbers{
		if fn(number) {
			newNumbers = append(newNumbers, number)
		}
	}
	return newNumbers
}

func isEvenNumber(number int) bool{
	return number%2==0
}

func isOddNumber(number int) bool{
	return !isEvenNumber(number)
}

func GetOddNumbers(numbers []int) []int {
	return getNumbers(numbers, isOddNumber)
}

func GetEvenNumbers(numbers []int) []int {
	return getNumbers(numbers, isEvenNumber)
}
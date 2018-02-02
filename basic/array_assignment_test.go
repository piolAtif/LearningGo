package basic

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetOddNumbers(t *testing.T){
	numbers := []int{1,2,3,4,5,6}
	oddNumbers := GetOddNumbers(numbers)
	expectedOddNumbers := []int{1,3,5}
	assert.Equal(t, expectedOddNumbers, oddNumbers)

}

func TestGetEvenNumbers(t *testing.T){
	numbers := []int{1,2,3,4,5,6}
	evenNumbers := GetEvenNumbers(numbers)
	expectedOddNumbers := []int{2,4,6}
	assert.Equal(t, expectedOddNumbers, evenNumbers)

}


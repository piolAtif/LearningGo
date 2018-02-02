package basic

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestMyCopy(t *testing.T){
	var greeting []string = make([]string, 2)
	greeting[0] = "Hello"
	greeting[1] = "world"

	var emptyGreeting []string = make([]string, (cap(greeting)+1)*2)

	Mycopy(emptyGreeting, greeting)
	expected := []string{"Hello", "world"}

	assert.Equal(t,expected, greeting,"expected and actual are not same")
}

func TestAppendByte(t *testing.T){
	primerNumbers := []byte{1, 2, 3}

	assert.Equal(t, 3, len(primerNumbers),"length of actual greeting")
	assert.Equal(t, 3, cap(primerNumbers),"capacity of actual greeting")

	primerNumbers = AppendByte(primerNumbers, 5, 7)

	assert.Equal(t, 5, len(primerNumbers),"length of actual greeting")
	assert.Equal(t, 12, cap(primerNumbers),"capacity of actual greeting")
	expectedPrimeNumbers := []byte{1,2,3,5,7}

	assert.Equal(t, expectedPrimeNumbers, primerNumbers, "prime numbers are not same")
}

func TestIsPalindrome(t *testing.T) {
	var tests = []struct{
		input string
		want bool
	}{
		{"",true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak",true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama",true},
		{"Evil I did dwell; lewd did I live", true},
		{"palindrome", false},
		{"desserts",false},
	}

	for _, test := range tests{
		if got := IsPalindrome(test.input); got != test.want{
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i<b.N; i++{
		IsPalindrome("A man, a plan, a canal: PanamaR")
	}
}


func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}


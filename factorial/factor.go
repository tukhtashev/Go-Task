package factorial

func Factorial(number int)(result int){
	result = 1
	for i := 1; i <= number; i++{
		result *= i
	}
	return
}
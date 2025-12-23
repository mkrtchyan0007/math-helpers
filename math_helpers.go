package mathhelpers

// Factorial возвращает факториал числа n
func Factorial(n int) int {
	if n < 0 {
		return -1 // ошибочное значение
	}
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// IsPrime проверяет, является ли число простым
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Sum возвращает сумму всех элементов слайса
func Sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

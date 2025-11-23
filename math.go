package main

//Сложение двух чисел
func Add(a, b int) int {
	return a + b
}

//Вычитание двух чисел
func Sub(a, b int) int {
	return a - b
}

//Умножение двух чисел
func Mul(a, b int) int {
	return a * b
}

//Деление двух чисел
func Div(a, b int) float64 {
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}

//Возведение числа в степень
func Pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

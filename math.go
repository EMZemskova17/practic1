package main

/// Add возвращает сумму двух чисел a и b.
///
/// Пример:
/// ```go
/// sum := Add(2, 3) // sum = 5
/// ```
func Add(a, b int) int {
	return a + b
}

/// Sub возвращает разность двух чисел a и b.
///
/// Пример:
/// ```go
/// diff := Sub(5, 2) // diff = 3
/// ```
func Sub(a, b int) int {
	return a - b
}

/// Mul возвращает произведение двух чисел a и b.
///
/// Пример:
/// ```go
/// product := Mul(2, 3) // product = 6
/// ```
func Mul(a, b int) int {
	return a * b
}

/// Div возвращает результат деления a на b.
/// Если b = 0, функция возвращает 0.
///
/// Пример:
/// ```go
/// quotient := Div(6, 3) // quotient = 2.0
/// ```
func Div(a, b int) float64 {
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}

/// Pow возводит число base в степень exp и возвращает результат.
///
/// Пример:
/// ```go
/// result := Pow(2, 3) // result = 8
/// ```
func Pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

package main

// Add возвращает сумму
func Add(a, b int) int {
	return a + b
}

// Sub возвращает разность
func Sub(a, b int) int {
	return a - b
}

// Mul возвращает произведение
func Mul(a, b int) int {
	return a * b
}

// Div возвращает частное (если b != 0)
func Div(a, b int) float64 {
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}

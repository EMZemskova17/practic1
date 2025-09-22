package main

func Add(a, b int) int {
	return a + b
}
func Sub(a, b int) int {
	return a - b
}
func Mul(a, b int) int {
	return a * b
}
func Div(a, b int) float64 {
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}

package main

import (
	"fmt"
)

// GetInput запрашивает два числа
// Два числа
func GetInput() (int, int) {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	return a, b
}

// PrintResults выводит результаты операций
func PrintResults(a, b, sum, diff, prod int, quot float64) {
	fmt.Printf("\nРезультаты для %d и %d:\n", a, b)
	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Разность: %d\n", diff)
	fmt.Printf("Произведение: %d\n", prod)
	fmt.Printf("Частное: %.2f\n", quot)
}

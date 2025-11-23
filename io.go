package main

import (
	"fmt"
)

// / GetInput запрашивает у пользователя два целых числа и возвращает их.
// /
// / Возвращаемые значения:
// / - int: первое число
// / - int: второе число
// /
// / Пример:
// / ```
// / a, b := GetInput()
// / ```
func GetInput() (int, int) {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	return a, b
}

// / PrintResults выводит результаты арифметических операций для двух чисел.
// /
// / Параметры:
// / - a, b: исходные числа
// / - sum: сумма a и b
// / - diff: разность a и b
// / - prod: произведение a и b
// / - quot: частное a / b
// /
// / Пример:
// / ```
// / PrintResults(2, 3, 5, -1, 6, 0.6667)
// / ```
func PrintResults(a, b, sum, diff, prod int, quot float64) {
	fmt.Printf("\nРезультаты для %d и %d:\n", a, b)
	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Разность: %d\n", diff)
	fmt.Printf("Произведение: %d\n", prod)
	fmt.Printf("Частное: %.2f\n", quot)
}

package main

func main() {
	//Получение значений, введенных пользователем
	a, b := GetInput()
	//Получение результатов операций
	sum := Add(a, b)
	prod := Mul(a, b)
	diff := Sub(a, b)
	quot := Div(a, b)

	PrintResults(a, b, sum, diff, prod, quot)
}

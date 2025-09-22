package main

func main() {
	a, b := GetInput()

	sum := Add(a, b)
	diff := Sub(a, b)
	prod := Mul(a, b)
	quot := Div(a, b)

	PrintResults(a, b, sum, diff, prod, quot)
}

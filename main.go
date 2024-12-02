package main

import "fmt"

func main() {
	// создать три константы конвертации
	const usdToEur = 0.94
	const usdToRub = 108.34

	// расчитать EUR в RUB на основании первых двух констант
	eurToRub := usdToRub / usdToEur

	// вывести результат на экран
	fmt.Printf("1 EUR = %.2f RUB\n", eurToRub)

}

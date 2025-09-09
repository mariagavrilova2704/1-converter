package main

import "fmt"

func main() {
	const UEur = 0.86
	const URub = 81.85
	ERub := URub / UEur
	fmt.Print(ERub)
}

func currencyInput(float64, string, string) {
	var moneyAmount float64
	var currentCurrency string
	var targetCurrency string
	fmt.Print("Введите количество денег, которое мы хотим сконвертировать:")
	fmt.Scan(&moneyAmount)
	fmt.Print("Введите исходную валюту: ")
	fmt.Scan(&currentCurrency)
	fmt.Print("Введите таргетированную валюту: ")
	fmt.Scan(&targetCurrency)
	return moneyAmount, currentCurrency, targetCurrency
}

func currencyCalculate(float64, string, string) {
	moneyAmount, currentCurrency, targetCurrency := currencyInput()

	return result
}

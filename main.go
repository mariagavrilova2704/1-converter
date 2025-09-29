package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== Перевод валют ===")
	var initial string
	var target string
	var err error
	var result float64

	for {
		initial, err = initialCurrency()
		if err != nil {
			fmt.Println("Ошибка:", err)
			fmt.Println("Попробуйте еще раз..")
			continue
		}
		break
	}

	amount := getNumberFromUser()
	//fmt.Printf("Вы ввели: %.2f\n", amount)
	// Дальше можно использовать amount для перевода валют

	for {
		target, err = targetCurrency(initial)
		if err != nil {
			fmt.Println("Ошибка:", err)
			fmt.Println("Попробуйте еще раз..")
			continue
		}
		break
	}

	fmt.Printf("Конвертируем из %s в %s\n", initial, target)

	if initial == "USD" && target == "RUB" {
		result = amount * 82.87
		fmt.Printf("Результат: %.2f\n", result)
	}

	if initial == "RUB" && target == "USD" {
		result = amount / 82.87
		fmt.Printf("Результат: %.2f\n", result)
	}

	if initial == "USD" && target == "EUR" {
		result = amount * 0.85
		fmt.Printf("Результат: %.2f\n", result)
	}

	if initial == "EUR" && target == "USD" {
		result = amount / 0.85
		fmt.Printf("Результат: %.2f\n", result)
	}

	if initial == "EUR" && target == "RUB" {
		result = amount * 97.53
		fmt.Printf("Результат: %.2f\n", result)
	}
	if initial == "RUB" && target == "EUR" {
		result = amount / 97.53
		fmt.Printf("Результат: %.2f\n", result)
	}
}

func initialCurrency() (string, error) {
	var initialCurrency string
	fmt.Print("Введите исходную валюту (USD/RUB/EUR): ")
	fmt.Scan(&initialCurrency)
	if initialCurrency == "USD" || initialCurrency == "RUB" || initialCurrency == "EUR" {
		return initialCurrency, nil
	}
	return "", errors.New("неверная валюта")
}

func getNumberFromUser() float64 {
	var input string

	for {
		fmt.Print("Введите сумму для перевода: ")
		fmt.Scan(&input)

		// Пытаемся преобразовать строку в число
		num, err := stringToInt(input)

		// Если ошибки нет (err == nil), значит это число
		if err == nil {
			return num
		}

		fmt.Printf("Ошибка: %v. Пожалуйста, введите число.\n", err)
	}
}

func stringToInt(s string) (float64, error) {
	// strconv.Atoi возвращает число и ошибку
	// Если ошибка nil - преобразование успешно
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("'%s' не является числом", s)
	}
	return num, nil
}

func targetCurrency(initial string) (string, error) {
	var targetCurrency string
	fmt.Print("Введите конечную валюту (USD/RUB/EUR): ")
	fmt.Scan(&targetCurrency)
	if targetCurrency == "USD" || targetCurrency == "RUB" || targetCurrency == "EUR" {
		if targetCurrency == initial {
			return "", errors.New("конечная валюта не должна совпадать с исходной")
		}
		return targetCurrency, nil
	}
	return "", errors.New("неверная валюта")
}

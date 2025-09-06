package main

import "fmt"

func main() {
	const UEur = 0.86
	const URub = 81.85
	ERub := URub / UEur
	fmt.Print(ERub)
}

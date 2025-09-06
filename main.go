package main

import "fmt"

func main() {
	const UEur = 0.85
	const URub = 81.85
	ERub := URub / UEur
	fmt.Print(ERub)
}

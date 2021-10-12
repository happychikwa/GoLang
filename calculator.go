package main

import "fmt"

//basic calculator with addition, multiplication, subtraction and division
func addition(a int, b int) int {
	return a + b
}
func subtraction(a int, b int) int {
	return a - b
}
func multiplication(a int, b int) int {
	return a * b
}
func division(a int, b int) int {
	return a / b
}

func main() {
	fmt.Println("Welcome to simple calc")
	var option int
	var a, b int
	for {
		fmt.Println("==MENU==")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("Enter Option :")
		fmt.Scanln(&option)
		switch option {
		case 1:
			fmt.Println("Enter First Number:")
			fmt.Scanln(&a)
			fmt.Println("Enter second Number:")
			fmt.Scanln(&b)
			fmt.Println(addition(a, b))
			continue
		case 2:
			fmt.Println("Enter First Number:")
			fmt.Scanln(&a)
			fmt.Println("Enter second Number:")
			fmt.Scanln(&b)
			fmt.Println(subtraction(a, b))
			continue
		case 3:
			fmt.Println("Enter First Number:")
			fmt.Scanln(&a)
			fmt.Println("Enter second Number:")
			fmt.Scanln(&b)
			fmt.Println(multiplication(a, b))
			continue
		case 4:
			fmt.Println("Enter First Number:")
			fmt.Scanln(&a)
			fmt.Println("Enter second Number:")
			fmt.Scanln(&b)
			fmt.Println(division(a, b))
			continue

		}
	}
}

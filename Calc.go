package main

import "fmt"

func main() {
	var num1, num2 float64 
	var choice string
	fmt.Println("Enter two numbers")
	num1, num2 = scanNumber()
	fmt.Printf("Enter your choice\n1.Add\n2.Subtract\n3.Multiply\n4.Divide\n5.Mod\nAny other key to exit...")
	fmt.Scanf("%s", &choice)
	result := calcResult(num1, num2, choice)
	fmt.Printf("Calculation result: %.2f", result)
}


func scanNumber() (float64, float64) {
	var x, y float64
	fmt.Println("Enter 1st number")
	fmt.Scanf("%f\n", &x)
	fmt.Println("Enter 2nd number")
	fmt.Scanf("%f\n", &y)
	return x, y
}

func calcResult(a, b float64, z string) float64 {
	var result float64
	switch z {
	case "1": //Case 1 would be Addition
			result = float64(a + b)  
			break

	case "2": //Case 2 would be Subtraction
			result = float64(a - b)
			break

	case "3": //Case 3 would be Multiplication
			result = float64(a * b)
			break

	case "4": //Case 4 would be Divide
			result = float64(a / b)
			break

	default: //Default case
			fmt.Println("See ya!!")
			break
	}

	return result
}

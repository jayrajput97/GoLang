package main

import "fmt"

func main() {
	var choice , num1, num2 int 
	fmt.Println("Enter two numbers")
	num1, num2 = scanNumber()
	fmt.Println("Enter your choice\n1.Add\n2.Subtract\n3.Multiply\n4.Divide\n5.Mod\nAny other key to exit...")
	fmt.Scanf("%d", &choice)
	result := calcResult(num1, num2, choice)
	fmt.Println("Calculation result: ", result)
}

func scanNumber() (int, int) { 
	var x, y int
	//fmt.Println("Enter 1st number")
	fmt.Scanf("%f %f", &x, &y)
	//fmt.Println("Enter 2nd number")
	//fmt.Scanf("%f", &y)
	return x, y
}

func calcResult(a, b int, z int) float64 {
	var result float64
	switch z {
	case 1: //Case 1 would be Addition
			result = float64(a + b)  
			break

	case 2: //Case 2 would be Subtraction
			result = float64(a - b)
			break

	case 3: //Case 3 would be Multiplication
			result = float64(a * b)
			break

	case 4: //Case 4 would be Divide
			result = float64(a / b)
			break

	case 5: //Case 5 would be Modulus
			result = float64(a % b)
			break
	
	default: //Default case
			fmt.Println("See ya!!")
			break
	}

	return result
}
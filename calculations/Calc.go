//Program for Simple Calculator in GoLang
package main

import "fmt"

//Main function
func main() {
	var num1, num2 float64 
	var choice string
	fmt.Println("Enter two numbers")
	num1, num2 = scanNumber()	//Function call to scan numbers
	fmt.Printf("Enter your choice\n1.Add\n2.Subtract\n3.Multiply\n4.Divide\n5.Mod\nAny other key to exit...")
	fmt.Scanf("%s", &choice)
	result := calcResult(num1, num2, choice)		//Function call to calculate result
	fmt.Printf("Calculation result: %.2f", result)
}

//Function to scan numbers
func scanNumber() (float64, float64) {
	var x, y float64
	fmt.Println("Enter 1st number")
	fmt.Scanf("%f\n", &x)		//\n is used to end the read after reading the first number
	fmt.Println("Enter 2nd number")
	fmt.Scanf("%f\n", &y)
	return x, y
}

//Calculator fucntion
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
	
	
			//NOTE: The modulus operator (%) is not in the switch case because it is not able to perform Modulus on Float values
	
	
	default: //Default case
			fmt.Println("See ya!!")
			break
	}

	return result
}

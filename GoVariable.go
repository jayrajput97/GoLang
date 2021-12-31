/*NOTE:
--> GoLang does not require semicolons (;) at the end of the statement.
--> GoLang produces errors if the declared variables or libraries are not used after declaring them.
	Unlike other programming languages such as C, C++, Java does not produce errors if the declared variables or libraries are not used.
*/

package main

import "fmt" //fmt stands for Format

func main() {

/*
VARIABLE DATA TYPES ARE AS FOLLOWS:
--> Integer - int
--> Unsigned Integer - uint
--> Floating Point - float
--> Boolean - bool
--> String - string
--> Byte - byte
*/



	//BASIC VARIABLE DECLARATION METHOD
	var x = 1	//VAR keyword automatically asigns the data-type to X i.e., 1
	var y int = 2	//You can expicitly assign the data-type to variables by mentioning the same after the variable-name

	fmt.Print(x, " & ", y)	//Printing variables with space or text in between 
	
	
	//ASSIGNING MULTIPLE VARIABLES HAVEING THE SAME DATA-TYPE IN ONE LINE
	var a, b int
	a = 10
	b = 20

	fmt.Print(a+b)	//Operations performed on the variables such as Multiplication (*), Division (/), Subtraction (-), Modulus (%), etc

	//ADVANCED METHOD TO DECLARE AND ASSIGN VALUE TO VARIABLES
	number := 1000	//The colon (:) and the equal to (=) sign declare and assign the value

	fmt.Print(number)
}
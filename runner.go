//Author: Conor Raftery

package main

import (
	"bufio"
	"fmt"
	"os"

	//Import folder for using shunt.go & rega.go
	assets "./src/assets"
)

func main() {

	option := 0
	exit := true

	//"While" loop to keep app running
	for exit {

		//Initial selection
		fmt.Print("\nPlease select an option:\n1)Infix Expression to Postfix Expresssion\n2)Postix Regular Expresssion to NFA\n3)Exit\n")
		fmt.Scanln(&option)

		//Enter inifix to postfix
		if option == 1 {
			fmt.Print("Please enter the expression to be converted:")
			//Read in string and use created function to trim string
			reader := bufio.NewReader(os.Stdin)
			expression, _ := reader.ReadString('\n')
			expression = assets.TrimEndString(expression)

			//Use for testing
			//a.b.c* -> ab.c*.	(a.(b|d))* -> abd|.*	a.(b|d).c* -> abd|.c*.		a.(b.b)+.c -> abb.+.c.

			//Output expression
			fmt.Println("Infix:  ", expression)
			//Pass expression into Intopost function and output result
			fmt.Println("Postfix: ", assets.Intopost(expression))

			//Enter Postfix to NFA
		} else if option == 2 {

			fmt.Print("Please enter the expression to be converted:")
			////Read in string and use created function to trim string
			reader := bufio.NewReader(os.Stdin)
			expression, _ := reader.ReadString('\n')
			expression = assets.TrimEndString(expression)

			//Output expression
			fmt.Println("Postfix:  ", expression)
			//Pass expression into Poregtonfa function and output result
			fmt.Println("NFA: ", assets.Poregtonfa(expression))
			fmt.Print("\n")

			//Prompt for basic string
			fmt.Print("Please enter a regular string to see if it matches the NFA:")
			////Read in string and use created function to trim string
			regString, _ := reader.ReadString('\n')
			regString = assets.TrimEndString(regString)
			regString = assets.Intopost(regString)

			//Use for testing
			//"ab.c*|", "ccc"

			//Pass expression and basic string into Pomatch

			//If returns true, strings match
			if assets.Pomatch(expression, regString) == true {
				//Output result
				fmt.Println("Regular string, ", regString, " matches the expression: ", expression)

				//If returns false, strings do not match
			} else if assets.Pomatch(expression, regString) == false {
				fmt.Print("String does not match")
				//Catch any errors
			} else {
				fmt.Print("An error occured")
			}

			//Exit program
		} else if option == 3 {
			fmt.Print("\nProgram Exited\n")
			//Stop loop
			exit = false
		} else {
			//Catch any input errors
			fmt.Print("\nPlease enter a valid option.\n")
		}
	}
} //main

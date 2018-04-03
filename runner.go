package main

import (
	"bufio"
	"fmt"
	"os"

	assets "./src/assets"
)

func main() {

	option := 0
	exit := true

	for exit {

		fmt.Print("Please select an option:\n1)Infix Expression to Postfix Expresssion\n2)Postix Regular Expresssion to NFA\n3)Exit\n")
		fmt.Scanln(&option)

		if option == 1 {
			fmt.Print("Please enter the expression to be converted:")
			reader := bufio.NewReader(os.Stdin)
			expression, _ := reader.ReadString('\n')
			expression = assets.TrimEndString(expression)

			//a.b.c* -> ab.c*.	(a.(b|d))* -> abd|.*	a.(b|d).c* -> abd|.c*.		a.(b.b)+.c -> abb.+.c.
			fmt.Println("Infix:  ", expression)
			fmt.Println("Postfix: ", assets.Intopost(expression))

		} else if option == 2 {

			fmt.Print("Please enter the expression to be converted:")
			reader := bufio.NewReader(os.Stdin)
			expression, _ := reader.ReadString('\n')

			expression = assets.TrimEndString(expression)

			fmt.Println("Postfix:  ", expression)
			fmt.Println("NFA: ", assets.Poregtonfa(expression))
			fmt.Print("\n")

			fmt.Print("Please enter a regular string to see if it matches the NFA:")
			//reader := bufio.NewReader(os.Stdin)
			regString, _ := reader.ReadString('\n')
			regString = assets.TrimEndString(regString)
			regString = assets.Intopost(regString) // Remove ending of string

			//"ab.c*|", "ccc"
			if assets.Pomatch(expression, regString) == true {
				fmt.Println("Regular string, ", regString, " matches the expression: ", expression)

			} else if assets.Pomatch(expression, regString) == false {
				fmt.Print("String does not match")
			} else {
				fmt.Print("An error occured")
			}

		} else if option == 3 {
			fmt.Print("\nProgram Exited\n")
			exit = false
		} else {
			fmt.Print("\nPlease enter a valid option.\n")
		}
	}
} //main

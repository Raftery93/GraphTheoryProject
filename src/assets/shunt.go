//Author: Conor Raftery
//Adapted from: https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e

package assets

func Intopost(infix string) string {
	//Creates a map with special characters and maps them to ints
	specials := map[rune]int{'*': 10, '.': 9, '|': 8, '?': 7, '+': 6}

	//Arrray of runes, initializes an array of empty runes
	pofix := []rune{} //rune-character on the screen diplayed in UTF-8

	//Stack which stores operators from the infix regular expression
	s := []rune{}

	//Loop over the input
	//range(convert into array of runes using UTF)
	for _, r := range infix {
		switch {

		case r == '(':
			//puts open bracket at the end of the stack
			s = append(s, r)
		case r == ')':
			//Pop of the stack until an opening bracket is found
			//len(s)-1 = the last element on the stack
			for s[len(s)-1] != '(' {
				pofix = append(pofix, s[len(s)-1]) //last element of 's'
				s = s[:len(s)-1]                   //everything in 's' except the last character
			}

			s = s[:len(s)-1]
		//if a special character is found
		case specials[r] > 0:
			//while the stack still has an element on it and the precedence of the current character that reads is <= the precedence of the top element of the stack
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				//Takes character of the stack and sticks into pofix
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			//Less precendence then append current character onto the stack
			s = append(s, r)
		//Not a bracket or a special character
		default:
			//Takes the default characters and sticks it at the end of pofix
			pofix = append(pofix, r)
		}
	}

	//If there is anything left on the stack, append to pofix
	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(pofix)
} //intopost

// Remove ending of string, user input adds a carriage return and a line feed
// to each string, so these characters needs to be trimmed back
func TrimEndString(s string) string {
	//Make sure string is not blank
	if len(s) > 0 {
		//Puts everything except the last 2 items (ASCII 12 & ASCII 10) into 's'
		s = s[:len(s)-2]
	}
	return s
}

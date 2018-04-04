//Author: Conor Raftery
//Adapted from: https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
//Adapted from: https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
//Adapted from: https://www.cs.york.ac.uk/fp/lsa/lectures/REToC.pdf

package assets

import (
	"fmt"
)

type state struct {
	//letter to determine epsilon
	symbol rune
	//2 arrows that come from a state which points at other states
	edge1 *state
	edge2 *state
}

//Helper struct
type nfa struct {
	//Keeps track of fragment of NFA
	initial *state
	accept  *state
}

//Post Reg Expression to NFA, returns a pointer to an NFA struct
func Poregtonfa(pofix string) *nfa {
	//Keeps track of fragments on a stack (pointer to nfastack)
	nfastack := []*nfa{}

	//Loop through pofix expression 1 char(rune) at a time
	for _, r := range pofix {
		switch r {
		//Concatenate character
		case '.':
			//Pop frag2 of nfastack
			frag2 := nfastack[len(nfastack)-1] //Index of last item on nfastack
			//Give me everything (up to but not including the last item) in nfastack
			nfastack = nfastack[:len(nfastack)-1]

			//Pop frag1 of nfastack
			frag1 := nfastack[len(nfastack)-1] //Index of last item on nfastack
			//Give me everything (up to but not including the last item) in nfastack
			nfastack = nfastack[:len(nfastack)-1]

			//frag1 and frag2 are pointers to 2 nfa fragments

			//(Concatenate)Join frag1 and frag2 - accept state of frag1's edge1
			// points to frag2's initial state
			frag1.accept.edge1 = frag2.initial

			//Append new pointer to nfa struct that represents new bigger nfa fragment which
			//is constructed from frag1 and frag2
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

		//Union character
		case '|':
			//Same as in case above
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			//New state which just points at the initial state of the 2 fragments that were popped off the stack
			//New initial state where edge1 points at frag1.initial and edge2 points at frag2.initial
			//Order does not matter because its an 'or'
			initial := state{edge1: frag1.initial, edge2: frag2.initial}

			//create a normal new state
			accept := state{}

			//The fragments edges has to point at the accept state
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			//Still a pointer to an NFA but we want the initial state of the fragment we are pushing to be
			//the new initial state we created above and we want the new accept state to be the accept state
			//of the fragment that we are pushing to the nfa stack nfa stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		//Kleene star
		case '*':
			//Pop 1 fragment of nfastack (Kleene star only works on one fragment of the nfa)
			frag := nfastack[len(nfastack)-1]
			//Give me everything (up to but not including the last item) in nfastack
			nfastack = nfastack[:len(nfastack)-1]
			//create a normal new state
			accept := state{}
			//Edge1 needs to be in iintial  state of fragment we popped off and edge2 needs to point at the new accept state
			initial := state{edge1: frag.initial, edge2: &accept}
			//Join accept state (edge1) for the fragment to initial state of fragment that was popped off
			frag.accept.edge1 = frag.initial
			//Join edge2 of old fragment to the accept state
			frag.accept.edge2 = &accept

			//Push new fragment to nfastack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '+':
			//Pop 1 fragment of nfastack (Kleene star only works on one fragment of the nfa)
			frag := nfastack[len(nfastack)-1]

			//create a normal new state
			accept := state{}

			//Edge1 needs to be in iintial  state of fragment we popped off and edge2 needs to point at the new accept state
			initial := state{edge1: frag.initial, edge2: &accept}

			//The fragment edge has to point at the initial state
			frag.accept.edge1 = &initial

			//Push new fragment to nfastack
			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})

		case '?':
			//Pop 1 fragment of nfastack (Kleene star only works on one fragment of the nfa)
			frag := nfastack[len(nfastack)-1]

			initial := state{edge1: frag.initial, edge2: frag.accept}

			//
			accept := state{edge1: frag.initial, edge2: frag.accept}
			frag.accept.edge1 = &accept
			//

			//The fragment edge has to point at the initial state
			frag.accept.edge1 = &initial

			//Push new fragment to nfastack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})

		//Any other character
		default:
			//Create a new accept state
			accept := state{}
			//Create a new initial state and set symbol to r and its only edge points to the accept state
			initial := state{symbol: r, edge1: &accept}

			//Push new fragment to nfastack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}

	//If the length of the nfastack is anything but 1
	if len(nfastack) != 1 {
		fmt.Println("Uh oh, there was ", len(nfastack), " NFA's found, they are: ", nfastack)
	}

	//Return only item(actual nfa) that is left on the stack
	return nfastack[0]
}

//Helper function, works recursively
//accepts a list of pointers to states, a single pointer to a state, and the accept state
//Returns a list of pointers to state
func addState(l []*state, s *state, a *state) []*state {
	//Append the state that has been passed in to the list
	l = append(l, s)

	//If 'a' does not, and the rune(s) equals 0 then it contains 'e' arrows
	if s != a && s.symbol == 0 {

		l = addState(l, s.edge1, a)
		//if edge2 is not null, do the same
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}

	//Return the list of pointers to a state
	return l
}

//Takes in postfix reg expression and a regular string, returns true if they match, false if they dont
func Pomatch(po string, s string) bool {

	//Initialise boolean
	ismatch := false

	//Variable which passes in the postfix expression, which is ran through poregtonfa
	ponfa := Poregtonfa(po)

	//Creates an array of pointers to state (basically a linked list of states)
	current := []*state{}

	//An array of pointers to state that I can get through from current
	next := []*state{}

	//Pass current, initial and accept state to addState
	current = addState(current[:], ponfa.initial, ponfa.accept)

	//Loop through one s, one char at a time
	for _, r := range s {

		//Everytime a char is read, loop through current array
		for _, c := range current {

			//If current symbol is the same is the symbol I am currently reading from s
			if c.symbol == r {

				//Returns the list of states and puts them into next
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		//Swap current with next, and then create a new array for next
		current, next = next, []*state{}
	}

	//Loop through the current array
	for _, c := range current {

		//If the state in the current array is equal to the accept state of ponfa
		if c == ponfa.accept {

			//Set ismatch to true
			ismatch = true
			break
		}
	}

	//Return either true or false
	return ismatch
}

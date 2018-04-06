package main

import (
	"fmt"
)

//Holds the state parameters, up to two arrows coming from state and at least one will have a letter value on an arrow
type state struct {
	//Letter stored with one of arrows (If arrow has Epsilon then value will be left at default)
	symbol rune
	//Two possible arrows from current state, which will point to other states
	edge1 *state
	edge2 *state
}

//Keeps track of inital state and accept state of nfa fragment
type nfa struct {
	initial *state
	accept  *state
}

//Function which changes infix expression to postfix expression
func inToPost(infix string) string {
	//Mapping special charaters to int values to easily determin order of precedence
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	//Arrays to hold final postfix and the working stack
	postfix := []rune{}
	stack := []rune{}

	//Loops through infix input until its end ('_' holds position in string, 'r' holds value at position)
	for _, r := range infix {
		switch {
		//When character is '('
		case r == '(':
			//Add '(' to stack
			stack = append(stack, r)
		//When character is ')'
		case r == ')':
			//Pop characters off of stack and append onto postfix until '(' is found
			//stack[len(stack)-1] finds the last element of the stack
			for stack[len(stack)-1] != '(' {
				//Puts element onto postfix
				postfix = append(postfix, stack[len(stack)-1])
				//Sets stack to everything up to the last element on stack (basically removes the last element)
				stack = stack[:len(stack)-1]
			}
			//Pops the '(' off the stack
			stack = stack[:len(stack)-1]

		//When character exists in specials array (any character not in array returns a value of int 0)
		case specials[r] > 0:
			//While the stack still contains elements && the precedence of the current character is <= the precedence of teh last element on the stack
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] {
				//Puts element onto postfix
				postfix = append(postfix, stack[len(stack)-1])
				//Sets stack to everything up to the last element on stack (basically removes the last element)
				stack = stack[:len(stack)-1]
			}
			//When current character has more precedence than the last element on the stack, add current character onto the stack
			stack = append(stack, r)
		//When character doesn't match any case, simply add the character to the end of postfix
		default:
			postfix = append(postfix, r)
		}
	}

	//If anything is still on the stack, add it to output and remove it from the stack
	for len(stack) > 0 {
		//Puts element onto postfix
		postfix = append(postfix, stack[len(stack)-1])
		//Sets stack to everything up to the last element on stack (basically removes the last element)
		stack = stack[:len(stack)-1]
	}

	//Return the postfix expression
	return string(postfix)
}

//Function which changes postfix expression to a non-deterministic finite automaton (NFA)
func regexToNFA(postfix string) *nfa {

}

func main() {
	//inToPost test cases
	//Answer: ab.c*
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: ", inToPost("a.b.c*"))

	//Answer: abd|.*
	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("postFix: ", inToPost("(a.(b|d))*"))

	//Answer: abd|.c*
	fmt.Println("Infix: ", "a.(b|d).c*")
	fmt.Println("postFix: ", inToPost("a.(b|d).c*"))

	//Answer: abb.+.c.
	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("postFix: ", inToPost("a.(b.b)+.c"))

	//regexToNFA test case
	nfa := regexToNFA("ab.c*|")
	fmt.Println(nfa)
}

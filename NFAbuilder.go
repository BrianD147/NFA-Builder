package main

import (
	"fmt"
)

func inToPost(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	postfix := []rune{}
	stack := []rune{}

	return string(postfix)
}

func main() {
	//Answer: ab.c*
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: ", inToPost("a.b.c*"))

	//Answer: abd|.*
	fmt.Println("Infix: ", "(a.b|d))*")
	fmt.Println("postFix: ", inToPost("(a.b|d))*"))

	//Answer: abd|.c*
	fmt.Println("Infix: ", "a.(b|d).c*")
	fmt.Println("postFix: ", inToPost("a.(b|d).c*"))

	//Answer: abb.+.c.
	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("postFix: ", inToPost("a.(b.b)+.c"))
}

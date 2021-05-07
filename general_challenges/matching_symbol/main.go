package main

import "fmt"

func main() {
	str := "a(bcd)e[f(ghij)]"

	fmt.Println(process(str))
}

// A checking for symmetry using a stack
func process(str string) string {
	stack := []string{} // arrays are natural stacks

	var matchingSymbol = map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, r := range []rune(str) { // In go we iterate over a string as runes
		char := string(r) // we immediately cast the rune to a string and work with that
		fmt.Println("->", char)

		if char == "(" || char == "[" || char == "{" { // we are going into a new context -- push to stack
			stack = append(stack, char)
			fmt.Println("pushed to stack", char)
		}

		if char == ")" || char == "]" || char == "}" { // end the current context
			if lnStk := len(stack); lnStk > 0 { // we will be referring to the array via indexes, so be safe!
				lastItem := stack[lnStk-1] // access the last element
				stack = stack[:lnStk-1]    // actual "pop" (re-slice)
				fmt.Println("popped from stack", lastItem)

				match, ok := matchingSymbol[char]
				if !ok {
					fmt.Println("Unknown match for symbol - check the map:", char)
				} else if lastItem != match {
					return "invalid character - no matching " + match + " found"
				}
			} else {
				return "invalid - no matching " + char
			}
		}
	}

	return "valid"
}

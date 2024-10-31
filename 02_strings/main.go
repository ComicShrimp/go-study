package main

import "fmt"

func main() {
	myString := []rune("résumé")
	indexed := myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nThe length of 'myString' is %v", len(myString))
	myRune := 'a'
	fmt.Printf("\nmyRune = %v", myRune)
	strSlice := []string{"s", "u", "b", "s", "c", "p", "i", "b", "e"}
	catStr := ""

	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)
}

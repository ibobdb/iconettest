package main

import (
	"fmt"
	"unicode"
)

func countDigits(slice []string) int {
	count := 0

	for _, item := range slice {
		if len(item) == 1 && unicode.IsDigit(rune(item[0])) {
			count++
		}
	}
	return count
}
func soal5() {
	fmt.Println("SOAL 5")
	list1 := []string{"b", "7", "h", "6", "h", "k", "i", "5", "g", "7", "8"}
	list2 := []string{"7", "b", "8", "5", "6", "9", "n", "f", "y", "6", "9"}
	list3 := []string{"u", "h", "b", "n", "7", "6", "5", "1", "g", "7", "9"}

	fmt.Println(list1,"=",countDigits(list1))
	fmt.Println(list2,"=",countDigits(list2)) 
	fmt.Println(list3,"=",countDigits(list3)) 
}
package main

import (
	"fmt"
	"strings"
)

func revString(s string)string{
	var solve string
	for i :=0 ; i < len(s);i++{
		solve = string(s[i])+ solve 
	}
	return solve
}

func revWord (s string)string{
		words := strings.Fields(s)
		for i, word := range words {
			words[i] = revString(word)
		}
		return strings.Join(words, " ")
}
func soal1(){
	fmt.Println("SOAL 1")
	var arr [3]string
	arr[0]="italem irad irigayaj"
	arr[1]="iadab itsap ulalreb"
	arr[2]="nalub kusutret gnalali"
	for i :=0;i < len(arr);i++{
		fmt.Println(revWord(arr[i]))
	}
}

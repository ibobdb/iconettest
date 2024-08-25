package main

import "fmt"

func soal2(){
	fmt.Println("SOAL 2")
	var ket string
	for i :=1;i <= 100;i++{
		ket=""
		if(i %3==0){
		ket = "FIZZ"
		}
		if(i%5==0){
		ket="BUZZ"
		}
		if(i %3==0 &&i%5==0){
			ket="FIZZBUZZ"
		}
		
		fmt.Println(i,ket)
	}
}
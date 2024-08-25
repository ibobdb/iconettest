package main

import "fmt"

func deret(n int) []int {
	num := make([]int, n)
	if n > 0 {
			num[0] = 0
	}
	if n > 1 {
			num[1] = 1
	}
	for i := 2; i < n; i++ {
			num[i] = num[i-1] + num[i-2]
	}
	return num
}

func soal3(){
	fmt.Println("SOAL 3")
	n :=9
	hasil := deret(n)
	fmt.Println(hasil)

}
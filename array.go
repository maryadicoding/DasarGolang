package main

import "fmt"

func main(){
	//tanpa slot
	var numbers = [...]int{2,3,4,5,6}
	// dengan slot 
	var arraystr = [2]string{"aku", "kamu"}
	//array multi
	var arraymulti = [2][3]int{[3]int{2,3,4}, [3]int{1,2,3}}
	var arraymultilain = [2][3]int{{21,3,4}, {1,2,3}}

	fmt.Println("data array :", numbers)
	fmt.Println("Jumlah array :", len(numbers))
	fmt.Println("Array Multi :", arraymulti)
	fmt.Println("Array String :", arraystr)
	fmt.Println("Array Multi Lain :", arraymultilain)
}

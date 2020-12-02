package main

import "fmt"

func main(){
	var buah = [3]string{"Nanas", "Apel", "Nangka"}

	for i := 0; i < len(buah); i++ {
		fmt.Println("Elemen", i ,":", buah[i])
	}
}
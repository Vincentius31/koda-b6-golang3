package main

import (
	"fmt"
	"koda-b6-golang3/lib"
)

func main(){
	var users []string = []string{
		"John",
		"Dave",
		"Robert",
		"Johanna",
		"Johnny",
		"Darius",
	}

	var keyword string

	fmt.Print("Masukan nama yang ingin dicari: ")
	fmt.Scanln(&keyword)

	var results[] string = lib.SearchPerson(users, keyword)

	if len(results) == 0 {
		fmt.Println("[]")
	}else {
		fmt.Println("Data ditemukan:")
		for i := range results {
			fmt.Println("-", results[i])
		}
	}
}
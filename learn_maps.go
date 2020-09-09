package main

import "fmt"

func main() {

birthdays := map[string]string {
	"keith": "01/01/2200",
	"test": "02/02/2020",
	}

	fmt.Println(birthdays)
	delete(birthdays,"test")
	fmt.Println(birthdays)
	ages:= map[string]int {}
	ages["keuth"] = 25
	ages["test"] = 21
	ages["Map"] = 10
	ages["keuth"] = 9
fmt.Println(ages)




}


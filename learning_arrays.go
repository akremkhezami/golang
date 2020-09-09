package main

import "fmt"

func main(){
	var items [3]string	
names := [3]string{"Akrem","Marwen","Amin"}
	fmt.Println(names)
	items[0] = "test"
	items[1] = "Kevin"
	items[2] = "golang"

fmt.Println(items)



	// Working with slices
	array := []string{}
	array = append(array,"my name")
	array = append(array,"is Akrem", "Marwen", "golang")
	fmt.Println(array)

	// Using makes function

tab := make([]string,4)
tab [0] = "test"
tab [1] = "test"
tab [2] = "test"
tab [3] = "test"

fmt.Println(tab)
}


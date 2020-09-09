package main

import "fmt"

func main() {


	ages := map[string]int {}
	ages["keuth"] = 77
	if ages["keuth"] < 66 {

fmt.Println(ages)
fmt.Println("Kevin is not of retirment age")

	} else { 
	fmt.Println("Kevin has retired")
	}


switch {

	case ages["Kevin"] <18:
	fmt.Println("Kevin cant vote")
	case ages["Kevin"] <67:
	fmt.Println("Kevin is not of retirment age")
	default:
	fmt.Println("Kevin has retired")
}






}

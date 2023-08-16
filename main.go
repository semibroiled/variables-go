package main

import "fmt"

func main(){

	penniesPerCoin := 2.0

	if penniesPerCoin > 0.0 {
	fmt.Printf("Type penniesPerCoin is %T\n", penniesPerCoin)
	} else {
		fmt.Println("No Pennies")
	}

	//functions
	fmt.Printf("The factorial is %d", factorial(10))

	var my_list = []string {"sth", "sth", "oh"}

	fmt.Printf("\n%v", concat(my_list))
}

func factorial(n uint) uint {
	if n == 1 {
		return 1
	} else{
	return n * factorial(n-1)
	}
}

func concat(list []string) string {
	combo := ""
	for i:= 0; i<len(list); i++ {
		combo+=list[i]
	}
	return combo
}
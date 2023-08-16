package main

import "fmt"

func main(){

	penniesPerCoin := 2.0

	if penniesPerCoin > 0 {
	fmt.Printf("Type penniesPerCoin is %T", penniesPerCoin)
	} else {
		fmt.Println("No Pennies")
	}
}
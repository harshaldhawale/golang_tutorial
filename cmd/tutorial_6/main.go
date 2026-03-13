package main

import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
	owner
}

type owner struct{
	name string
}

func main(){
	var myEngine gasEngine
	fmt.Println(myEngine.mpg, myEngine.gallons)

	var myEngine2 gasEngine = gasEngine{25, 15, owner{"Alex"}}
	myEngine2.mpg = 20
	fmt.Println(myEngine2)
}
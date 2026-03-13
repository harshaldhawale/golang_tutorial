package main

import "fmt"

func main(){
	//string, runes, bytes

	var myString = "myString"
	var indexed = myString[0]
	fmt.Println(myString)
	fmt.Println(indexed)
	fmt.Printf("%v,%T \n",indexed, indexed)

	for i,v := range myString{
		fmt.Println(i,v)
	}

	fmt.Printf("\n the length of my string is %v \n",len(myString))

	var myRune = 'a'
	fmt.Printf("\n my rune = %v \n", myRune)

	var strSlice = []string{"a","b","c","d"}
	var catStr = ""
	for i := range strSlice{
		catStr += strSlice[i]
	}

	fmt.Printf("\n%v\n",catStr)

}
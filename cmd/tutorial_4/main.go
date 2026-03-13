package main

import "fmt"

func main(){
	var intArr[3] int32 // arrays are fixed length, same type, indexable, contiguous in memory

	intArr[1] = 123

	fmt.Println(intArr)
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])


	var intArr2 [3]int32 = [3]int32{1,2,3} 
	fmt.Println(intArr2)

	intArr3 := [3]int32{3,2,1}
	fmt.Println(intArr3)

	intArr4 := [...]int32{4,5,6}
	fmt.Println(intArr4)

///////////////////////////////////////////////////////////

	var intSlice []int32 = []int32{1,2,3}
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	intSlice = append(intSlice,7)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 10)
	fmt.Println(intSlice3)

///////////////////////////////////////////////////////////

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":23, "Sarah":22}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Json"])

	var age, ok = myMap2["Json"]
	if ok{
		fmt.Printf("The age is %v", age)
	}else{
		fmt.Println("Invalid Name")
	}

	for name, age := range myMap2{
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	for i,v := range intArr2{
		fmt.Printf("Index: %v, Value: %v \n",i,v)
	}

	// var i int = 0

	// for i<10{
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for{
	// 	if i > 10{
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	for i:=0; i < 15; i++{
		fmt.Println(i)
	}
}
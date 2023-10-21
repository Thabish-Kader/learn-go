package learngo

import (
	"fmt"
)



func changeValueByPointer (initalNum *int) int {
	*initalNum = 5
	return *initalNum
}

func changeValue (initalNum int) int {
	initalNum = 5
	return initalNum
}

func level4(){

	num := 4;
	// ref := &num
	// num = 5


	// fmt.Println(num)
	// fmt.Println(ref)

	changedNumber := changeValue(&num)
	fmt.Printf("Orginal number %d \n", num)
	fmt.Printf("Changed number %d \n", changedNumber)


}
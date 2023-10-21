package learngo

import "fmt"

func level2() {
	var numbers [5]int

	numbers[0] = 5
	numbers[1] = 10
	numbers[2] = 15

	for i :=0 ; i < len(numbers) ; i++ {
		fmt.Println(numbers[i])
	}

	numbers2 := []int{1,2,3,4,5}
	// var numbers = [...]int{10, 20, 30, 40, 50} - note if you use ... the elipssis operator go will determine the size

	for i :=0 ; i < len(numbers2) ; i++ {
		fmt.Println(numbers2[i])
	}

	for {
		fmt.Println("Hello")
	}

	for {
        fmt.Println("first statment")
        return // This will exit the main function
        fmt.Println("Second statement")
    }
	fmt.Println("third statement")

	for {
        fmt.Println("first statment")
        break // This will exit the main function
        fmt.Println("Second statement")
    }
	fmt.Println("third statement")
}
package learngo

import "fmt"


func level3() {
	ages := map[string]int{
        "Alice": 25,
        "Bob":   30,
        "Carol": 35,
    }

	if age, exists := ages["Carol"]; exists {
		fmt.Println(age)
	} else {
		fmt.Println("Carol's age is not available.")
	}

	ages["Eve"] = 28
	fmt.Println("Eve's age:", ages["Eve"])

	delete(ages, "Eve")
	fmt.Println(ages)

	for key, value := range ages {
        fmt.Printf("Key: %s , Value: %d \n", key, value)
    }
}
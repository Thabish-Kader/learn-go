package learngo

import "fmt"

type Person struct {
    Name string
    Age  int
}

func modifyPerson(p Person) {
    p.Name = "John Doe"
    p.Age = 30
}

func modifyPersonByPointer(p *Person) {
    p.Name = "Jane Smith"
    p.Age = 25
}

func level5() {
    // Direct Initialization
    person1 := Person{}
    modifyPerson(person1)
    fmt.Printf("Direct Initialization - Name: %s, Age: %d\n", person1.Name, person1.Age)

    // Pointer Initialization
    person2 := &Person{}
    modifyPersonByPointer(person2)
    fmt.Printf("Pointer Initialization - Name: %s, Age: %d\n", person2.Name, person2.Age)
}

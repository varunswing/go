package main
import "fmt"

type person struct {
    name string
	age int
}
// Method with pointer receiver
func (p *person) updateName(newName string, age int) {
    p.name = newName
	p.age = age
}

// Method with value receiver
func (p person) showName() {
    fmt.Println("Name:", p.name) // Name: b
}

func methods() {
    a := person{name: "a"}
    
    // Calling pointer method with value
    a.updateName("b", 23)
    fmt.Println("After pointer method:", a.name, a.age) // b 23
    
    // Calling value method with pointer
    (&a).showName()
}
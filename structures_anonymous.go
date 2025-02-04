package main
import "fmt"

// Student struct with an anonymous structure and fields
type Student struct {
    Info struct {   // Named anonymous structure
        Name       string
        Enrollment int
    }
    GPA float64  // Standard field
}

func structures_anonymous() {
    student := Student{
        Info: struct {
            Name       string
            Enrollment int
        }{
            Name:       "A",
            Enrollment: 12345,
        },
        GPA: 3.8,
    }
    fmt.Println("Name:", student.Info.Name)
    fmt.Println("Enrollment:", student.Info.Enrollment)
    fmt.Println("GPA:", student.GPA)
}
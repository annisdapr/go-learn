package library

import "fmt"

var Student1 = struct {
    Name  string
    Grade int
}{}

func init() {
    Student1.Name = "John Wick"
    Student1.Grade = 2

    fmt.Println("--> library/library.go imported")
}
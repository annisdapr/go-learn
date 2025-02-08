package main

import "go-level-akses/library"
import "fmt"


func main() {
    library.SayHello("ethan")
	var s1 = library.Student{"ethan", 21}
    fmt.Println("name ", s1.Name)
    fmt.Println("grade", s1.Grade)

	//partial
	sayHello("ethan")

	//init
	fmt.Printf("Name  : %s\n", library.Student1.Name)
    fmt.Printf("Grade : %d\n", library.Student1.Grade)
}
package library

import "fmt"

func SayHello(name string) {
    fmt.Println("hello")
    introduce(name)
}
// func SayHello() {
//     fmt.Println("hello")
// }

func introduce(name string) {
    fmt.Println("nama saya", name)
}


type Student struct {
    Name  string
    Grade int
}
package main

import "fmt"
import "strings"


type person struct {
    name string
    age  int
}

func main() {
    var secret interface{}

    secret = "ethan hunt"
    fmt.Println(secret)

    secret = []string{"apple", "manggo", "banana"}
    fmt.Println(secret)

    secret = 12.4
    fmt.Println(secret)

    fmt.Printf("\n")

    //Casting variabel to anyInterface
    var secrets interface{}

    secrets = 2
    var number = secrets.(int) * 10
    fmt.Println(secrets, "multiplied by 10 is :", number)

    secrets = []string{"apple", "manggo", "banana"}
    var gruits = strings.Join(secrets.([]string), ", ")
    fmt.Println(gruits, "is my favorite fruits")

    fmt.Printf("\n")

    //Casting to Pointer
    var secretss interface{} = &person{name: "wick", age: 27}
    var name = secretss.(*person).name
    fmt.Println(name)

    fmt.Printf("\n")
    //slice interface
    var fruits = []interface{}{
        map[string]interface{}{"name": "strawberry", "total": 10},
        []string{"manggo", "pineapple", "papaya"},
        "orange",
    }
    
    for _, each := range fruits {
        fmt.Println(each)
    }
}
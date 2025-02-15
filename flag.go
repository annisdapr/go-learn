package main

import "flag"
import "fmt"

func main() {
    var name = flag.String("name", "anonymous", "type your name")
    var age = flag.Int64("age", 25, "type your age")

    flag.Parse()
    fmt.Printf("name\t: %s\n", *name)
    fmt.Printf("age\t: %d\n", *age)


	//passing flag to variabel
	// cara ke-1
	var data1 = flag.String("name", "anonymous", "type your name")
	fmt.Println(*data1)

	// cara ke-2
	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender")
	fmt.Println(data2)
}
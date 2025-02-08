package main

import "fmt"
import "reflect"


type Student struct {
    Name  string
    Grade int
}

func main() {
    var number = 23
    var reflectValue = reflect.ValueOf(number)

    fmt.Println("tipe  variabel :", reflectValue.Type())

    if reflectValue.Kind() == reflect.Int {
        fmt.Println("nilai variabel :", reflectValue.Int())
    }
	fmt.Println("tipe  variabel :", reflectValue.Type())
	fmt.Println("nilai variabel :", reflectValue.Interface())


	var s1 = &Student{Name: "wick", Grade: 2}
    s1.getPropertyInfo()

    fmt.Println("nama :", s1.Name)
    var method = reflectValue.MethodByName("SetName")
    method.Call([]reflect.Value{
        reflect.ValueOf("wick"),
    })

    fmt.Println("nama :", s1.Name)

    
}

func (s *Student) getPropertyInfo() {
    var reflectValue = reflect.ValueOf(s)

    if reflectValue.Kind() == reflect.Ptr {
        reflectValue = reflectValue.Elem()
    }

    var reflectType = reflectValue.Type()

    for i := 0; i < reflectValue.NumField(); i++ {
        fmt.Println("nama      :", reflectType.Field(i).Name)
        fmt.Println("tipe data :", reflectType.Field(i).Type)
        fmt.Println("nilai     :", reflectValue.Field(i).Interface())
        fmt.Println("")
    }
}
func (s *Student) SetName(name string) {
    s.Name = name
}
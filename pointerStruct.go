package main
import "fmt"

func main(){
	var s1 = student{name: "wick", grade: 2}
	var s2 *student = &s1

	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)

	s2.name = "ethan"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)
  	var s11 = students{}
    s11.name = "wick"
    s11.age = 21
    s11.grade = 2

    fmt.Println("name  :", s11.name)
    fmt.Println("age   :", s11.age)
    fmt.Println("age   :", s11.person.age)
    fmt.Println("grade :", s11.grade, "\n")

	/*anonymous struct

	//without properties
	 var s1 = struct {
        person
        grade int
    }{}
	
	// with properties
    var s2 = struct {
    person
    grade int
	}{
    person: person{"wick", 21},
    grade:  2,
	}
	*/

	//slice struct

	fmt.Println("**Slice struct**\n")
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}
	
	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
	//anonymous slice struct
	var allStudents = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}
	//anonymous var struct
	var student struct {
		person
		grade int
	}

	// declaration only
	var student struct {
    grade int
	}

	// deklaration + initialize
	var student = struct {
		grade int
	} {
		12,
	}
}

type student struct {
    name string
    grade int
}

type students struct {
    grade int
    person
}

type person struct {
    name string
    age  int
}

/*nested struct
type student struct {
    person struct {
        name string
        age  int
    }
    grade   int
    hobbies []string
}
	*/

//Type alias
/*
type Person struct {
    name string
    age  int
}
type People = Person

type People2 = struct {
    name string
    age  int
}

type Number = int
var num Number = 12
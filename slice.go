package main
import "fmt"

func main(){

	// inisialisasi slice
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // "apple"

	// perbedaan slice dan array
	/*
	var fruitsA = []string{"apple", "grape"}      // slice
	var fruitsB = [2]string{"banana", "melon"}    // array
	var fruitsC = [...]string{"papaya", "grape"}  // array
	*/

	//operasi slice
	var newFruits = fruits[0:2]
	var sliFruits = fruits[2:]
	fmt.Println(newFruits) // ["apple", "grape"]
	fmt.Println(sliFruits) 

	//fungsi slice
	fmt.Println(len(fruits))

	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	var bFruits = fruits[1:4]
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	var zFruits = fruits[0:2:2]
	fmt.Println(len(zFruits)) // len: 2
	fmt.Println(cap(zFruits)) // cap: 2

	var cFruits = append(fruits, "papaya")
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]

	dst := make([]string, 3)

	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3
}
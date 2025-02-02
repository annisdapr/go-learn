package main
import "fmt"

func main(){
	var fruits [4]string

	/*cara horizontal
	fruits  = [4]string{"apple", "grape", "banana", "melon"}
	*/

	// cara vertikal
	fruits  = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}
	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	var buah = make([]string, 2)
	buah[0] = "apple"
	buah[1] = "manggo"

	fmt.Println(buah)  
	}
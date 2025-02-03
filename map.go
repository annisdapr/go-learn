package main
import "fmt"

func main(){
	// inisialisasi map
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["jan"] = 50
	chicken["feb"] = 40

	fmt.Println("jan", chicken["jan"]) // januari 50
	fmt.Println("mei",     chicken["mei"])   

	var data map[string]int
	/*
	data["one"] = 1
	// akan muncul error!
*/
	data = map[string]int{}
	data["one"] = 1
	fmt.Println("one",     data["one"])  
	// tidak ada error

	/*
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)
	*/


	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	// hapus
	delete(chicken, "januari")
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// slice map
	var data1 = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	for _, item := range data1 {
		fmt.Println("Name:", item["name"], "Gender:", item["gender"])
	}
}
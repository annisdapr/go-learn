package main

import "encoding/json"
import "fmt"

type User struct {
    FullName string `json:"Name"`
    Age      int
}


func main() {
    var jsonString = `{"Name": "john wick", "Age": 27}`
    var jsonData = []byte(jsonString)

    var data User

    var err = json.Unmarshal(jsonData, &data)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("user :", data.FullName)
    fmt.Println("age  :", data.Age)

	//JSON to Map
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])

	//Interface
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])

	//JSON to array
	var jsonString1 = `[
    {"Name": "john wick", "Age": 27},
    {"Name": "ethan hunt", "Age": 32}
	]`
	var data3 []User

	var err1 = json.Unmarshal([]byte(jsonString1), &data3)
	if err1 != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1:", data3[0].FullName)
	fmt.Println("user 2:", data3[1].FullName)

	//JSON to String
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData1, err2 = json.Marshal(object)
	if err != nil {
		fmt.Println(err2.Error())
		return
	}

	var jsonString2 = string(jsonData1)
	fmt.Println(jsonString2)

}


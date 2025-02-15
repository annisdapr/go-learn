package main

import "encoding/base64"
import "fmt"

func main() {
    var data = "john wick"

    var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println("encoded:", encodedString)

    var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
    var decodedString = string(decodedByte)
    fmt.Println("decoded:", decodedString)

	//another one
	var encoded1 = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded1, []byte(data))
	var encodedString1 = string(encoded1)
	fmt.Println(encodedString1)

	var decoded1 = make([]byte, base64.StdEncoding.DecodedLen(len(encoded1)))
	var _, err = base64.StdEncoding.Decode(decoded1, encoded1)
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodedString1 = string(decoded1)
	fmt.Println(decodedString1)

	//Encode Decode URL
	var dataURL = "https://kalipare.com/"

	var encodedStringURL = base64.URLEncoding.EncodeToString([]byte(dataURL))
	fmt.Println(encodedStringURL)

	var decodedByteURL, _ = base64.URLEncoding.DecodeString(encodedStringURL)
	var decodedStringURL = string(decodedByteURL)
	fmt.Println(decodedStringURL)
}
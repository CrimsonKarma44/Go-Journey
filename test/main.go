package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func encrypt(plainText, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, plainText, nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}
func decrypt(cipherText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, cipherTextBytes := data[:nonceSize], data[nonceSize:]

	plainText, err := aesGCM.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func viewUsers() []string {
	var profileNames []string

	file, _ := os.ReadDir("../form/profiles")
	for _, f := range file {
		profileNames = append(profileNames, f.Name()[:len(f.Name())-5])
	}
	return profileNames
}

// Define a struct type
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

func main() {
	//nameList := viewUsers()
	//fmt.Println(nameList[0])
	//fmt.Println(nameList[1])
	//comment := "65"
	//
	//value, _ := strconv.Atoi(comment)
	//fmt.Printf("%s\n", comment)
	//fmt.Printf("%v\n", value)
	//file, _ := os.ReadDir(".")
	//fmt.Println(file)

	// Create instances of Person
	//person1 := Person{Name: "Alice", Age: 30, Country: "USA"}
	//person2 := Person{Name: "Bob", Age: 35, Country: "Canada"}

	// Append the struct instances to an array
	//people := []Person{person1, person2}
	//
	//// Marshal the array into JSON
	//jsonData, err := json.Marshal(people)
	//if err != nil {
	//	fmt.Println("Error marshaling JSON:", err)
	//	return
	//}

	//Print the JSON data
	//fmt.Println(string(jsonData))
	encrypted, err := encrypt([]byte("Darksummer44."), []byte("a very very very very secret key"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted")
	fmt.Println(encrypted)
	recoveredText, err := decrypt(encrypted, []byte("a very very very very secret key"))
	if err != nil {
		panic(err)
	}
	fmt.Println(recoveredText)
}

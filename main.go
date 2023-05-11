package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	var length int
	
	fmt.Println("Enter the length of password: ")
	fmt.Scanln(&length)

	rand.Seed(time.Now().UnixNano())

	charset := "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm01234567890@#$%^&*()_+-=[]{}|;':\",./<>?"

	for {
		password := make([]byte, length)
		for i := range password {
			password[i] = charset[rand.Intn(len(charset))]
		}
		fmt.Printf("%s\n", password)

		fmt.Println("Do you want to generate another password? [Y/N]")
		var input string
		fmt.Scanln(&input)

		if input != "y" && input != "Y" {
			return
		} 

		fmt.Println("Enther the length of password: ")
		fmt.Scanln(&length)

	}
}
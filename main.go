package main

import (
	"fmt"
	"math/rand"
	"strings"
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
		fmt.Println(string(password))
		
		fmt.Println("Do you want to generate another password? [Y/N]")
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "y" {
			fmt.Println("Enther the length of password: ")
			fmt.Scanln(&length)

		} else {
			return
		}
	}
}
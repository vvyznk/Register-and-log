package main

import (
	"fmt"
	"os"
)

type dados struct {
	username string
	password string
	food     string
	nickname string
	age      int
}
type Type struct {
	username string
	password string
}

var choice string
var users []dados

func main() {
	for {
		menu()
	}
}

func register() {
	var user dados

	fmt.Println("Choose an avaliable username")
	fmt.Scanln(&user.username)
	fmt.Println("Choose your password")
	fmt.Scanln(&user.password)
	fmt.Println("Write your favorite food")
	fmt.Scanln(&user.food)
	fmt.Println("Write your childhood nickname")
	fmt.Scanln(&user.nickname)
	fmt.Println("Write your age")
	_, err := fmt.Scanln(&user.age)

	if err != nil {
		fmt.Println("Just type numbers ")
		fmt.Println("Please try again.")
		fmt.Scanln()
		return
	} else {
		switch {
		case user.age >= 18:
			users = append(users, user)
			fmt.Println("Registred profile!")
			return
		default:
			fmt.Println("Sorry, your age is not allowed in our system.")
			return
		}
	}

}

func log() {
	var validation Type
	if len(users) == 0 {
		fmt.Println("No registred users")
		return
	}

	fmt.Println("Choose an user registered in the system")
	fmt.Scanln(&validation.username)
	fmt.Println("Enter your password")
	fmt.Scanln(&validation.password)

	for _, u := range users {
		if u.username == validation.username && u.password == validation.password {
			fmt.Println("Enjoy!")
			system()
			return
		}
	}
	fmt.Println("Wrong user or password.")
}

func menu() {
	var choice string
	fmt.Println("----------------------/")
	fmt.Println("What do you want? ")
	fmt.Println("-------------------/")
	fmt.Println("Type 1 to register")
	fmt.Println("-------------------/")
	fmt.Println("Type 2 to log")
	fmt.Println("--------------/")
	fmt.Println("Type 3 to exit")
	fmt.Println("---------------/")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		register()
	case "2":
		log()
	case "3":
		os.Exit(0)
	default:
		fmt.Println("That was not a valid option. ")
		return
	}
}

func system() {
	fmt.Println("")
	fmt.Println("-------- ENTROU DENTRO DO SISTEMA --------")
	fmt.Println("")
	return
}

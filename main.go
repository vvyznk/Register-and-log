package main

import (
	"fmt"
	"os"
	"strings"
)

type user struct {
	username string
	password string
	food     string
	nickname string
	validAge bool
}
type login struct {
	username string
	password string
}

var choice string
var users []user

func main() {
	for {
		menu()
	}
}
func userExists(username string) bool {
	for _, u := range users {
		if u.username == username {
			return true
		}
	}
	return false
}
func createPassword() string {
	for {
		var password string
		fmt.Println("- Choose your password")
		fmt.Print(":")
		fmt.Scanln(&password)
		if len(password) < 8 {
			fmt.Println("-- Your password need have 8 caracteres.")
			continue
		}
		return password
	}
}

func createUsername() string {
	for {
		var username string
		fmt.Println("- Choose an avaliable username")
		fmt.Print(":")
		fmt.Scanln(&username)

		username = strings.ToLower(username)

		if len(username) < 4 {
			fmt.Println("-- Username too short, try another username bigger at 4 caracteres.")
			continue
		}

		if userExists(username) {
			fmt.Println("-- Username already registred, Please try another username")
			continue
		}
		return username
	}
}
func justnumber() int {
	var number int
	for {
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("- Just type numbers, please try again.")
			fmt.Print(": ")

			var garbage string
			fmt.Scanln(&garbage)

		} else {
			return number
		}
	}
}
func validAge() bool {
	var age int
	for {
		fmt.Println("- Write your age")
		fmt.Print(":")
		age = justnumber()
		if age >= 18 {
			return true
		}
		return false
	}
}
func register() {
	var newUser user

	newUser.username = createUsername()
	newUser.password = createPassword()

	fmt.Println("- Write your favorite food")
	fmt.Print(":")
	fmt.Scanln(&newUser.food)
	fmt.Println("- Write your childhood nickname")
	fmt.Print(":")
	fmt.Scanln(&newUser.nickname)

	newUser.validAge = validAge()

	switch {
	case newUser.validAge == true:
		users = append(users, newUser)
		fmt.Println("- Registred profile!")
		fmt.Println()
		return
	default:
		fmt.Println("- Sorry, your age is not allowed in our system.")
		fmt.Println()
		return
	}
}

func log() {
	var validation login
	if len(users) == 0 {
		fmt.Println("- No registred users")
		fmt.Println()
		return
	}

	fmt.Println("- Choose an user registered in the system")
	fmt.Print(":")
	fmt.Scanln(&validation.username)
	fmt.Println("- Enter your password")
	fmt.Print(":")
	fmt.Scanln(&validation.password)
	validation.username = strings.ToLower(validation.username)

	for _, u := range users {
		if u.username == validation.username && u.password == validation.password {
			fmt.Println("- Enjoy!")
			fmt.Println()
			system()
			return
		}
	}
	fmt.Println("Wrong user or password.")
	fmt.Println()
}

func menu() {
	var choice string
	fmt.Println("======= MENU =======")
	fmt.Println("1 - Register")
	fmt.Println("2 - Log")
	fmt.Println("3 - Exit")
	fmt.Println("=====================")
	fmt.Print("Choose: ")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		fmt.Println("-----------------")
		register()
	case "2":
		fmt.Println("-----------------")
		log()
	case "3":
		fmt.Println("-----------------")
		fmt.Println("- Leaving the system...")
		os.Exit(0)
	default:
		fmt.Println("That was not a valid option. ")
		fmt.Println()
		return
	}
}

func system() {
	var answer string
	fmt.Println("-------- SYSTEM --------")
	for {
		fmt.Println("")
		fmt.Println("1 - Back to menu")
		fmt.Println("2 - Exit")
		fmt.Print("Choose: ")
		fmt.Scanln(&answer)
		switch answer {
		case "1":
			return
		case "2":
			os.Exit(0)
		default:
			fmt.Println("Just choose 1 or 2")
		}
	}

}
